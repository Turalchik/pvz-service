package integration

import (
	"context"
	"fmt"
	"github.com/Turalchik/pvz-service/internal/app/repo"
	"github.com/Turalchik/pvz-service/internal/app/tokenizer"
	"github.com/Turalchik/pvz-service/internal/service/pvz_service"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"google.golang.org/grpc"
	"net"
	"os"
	"path/filepath"
	"testing"
)

func mustStartPostgresAndMigrate(t *testing.T) *sqlx.DB {
	ctx := context.Background()

	// 1) Запускаем контейнер
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		Env:          map[string]string{"POSTGRES_PASSWORD": "secret", "POSTGRES_DB": "testdb"},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
	}
	pgC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req, Started: true,
	})
	require.NoError(t, err)

	// 2) Берём Host и Port из контейнера
	host, _ := pgC.Host(ctx)
	port, _ := pgC.MappedPort(ctx, "5432")
	dsn := fmt.Sprintf("postgres://postgres:secret@%s:%s/testdb?sslmode=disable", host, port.Port())

	// 3) Открываем соединение
	db := sqlx.MustConnect("pgx", dsn)

	// 4) Прогоняем миграции
	files, _ := filepath.Glob("../migrations/*.sql")
	for _, f := range files {
		sqlBytes, _ := os.ReadFile(f)
		_, err := db.Exec(string(sqlBytes))
		require.NoError(t, err, "migration %s failed", f)
	}

	return db
}

func mustStartGRPCServer(t *testing.T, db *sqlx.DB) string {
	lis, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	server := grpc.NewServer()
	// создаём зависимости так же, как в main.go
	repo := repo.NewRepo(db)
	uuidIfc := pvz_service.NewUUID()
	timerIfc := pvz_service.NewTimer()
	tokenizerIfc := tokenizer.NewTokenizer([]byte("fdjsfdsafkljsdfj"))
	svc, err := pvz_service.NewPVZServiceServer(repo, uuidIfc, timerIfc, tokenizerIfc)
	require.NoError(t, err)
	desc.RegisterPVZServiceServer(server, svc)

	go func() {
		if err := server.Serve(lis); err != nil {
			t.Logf("gRPC server exited: %v", err)
		}
	}()

	return lis.Addr().String()
}
