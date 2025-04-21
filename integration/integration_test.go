package integration

import (
	"context"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
)

func TestIntegration_FullFlow(t *testing.T) {
	db := mustStartPostgresAndMigrate(t)

	addr := mustStartGRPCServer(t, db)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()
	client := desc.NewPVZServiceClient(conn)

	ctx := context.Background()

	_, err = client.Register(ctx, &desc.RegisterRequest{
		Login:    "moderator",
		Password: "moderator",
		Role:     "модератор",
	})
	require.NoError(t, err)
	loginResp, err := client.Login(ctx, &desc.LoginRequest{
		Login:    "moderator",
		Password: "moderator",
	})
	require.NoError(t, err)
	tokenModerator := loginResp.GetToken()

	createResp, err := client.CreatePVZ(ctx, &desc.CreatePVZRequest{
		Token: tokenModerator,
		City:  "Москва",
	})
	require.NoError(t, err)
	pvzID := createResp.GetId()

	_, err = client.Register(ctx, &desc.RegisterRequest{
		Login:    "employee",
		Password: "employee",
		Role:     "сотрудник ПВЗ",
	})
	require.NoError(t, err)
	loginResp, err = client.Login(ctx, &desc.LoginRequest{
		Login:    "employee",
		Password: "employee",
	})
	require.NoError(t, err)
	tokenEmployee := loginResp.GetToken()

	_, err = client.OpenReception(ctx, &desc.OpenReceptionRequest{
		Token: tokenEmployee,
		Id:    pvzID,
	})
	require.NoError(t, err)

	for i := 0; i < 50; i++ {
		_, err = client.AddProduct(ctx, &desc.AddProductRequest{
			Token: tokenEmployee,
			Id:    pvzID,
			Type:  "электроника",
		})
		require.NoErrorf(t, err, "AddProduct #%d failed", i)
	}

	_, err = client.CloseReception(ctx, &desc.CloseReceptionRequest{
		Token: tokenEmployee,
		Id:    pvzID,
	})
	require.NoError(t, err)
}
