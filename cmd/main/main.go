package main

import (
	"github.com/Turalchik/pvz-service/internal/database"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"

	"github.com/Turalchik/pvz-service/internal/app/repo"
	"github.com/Turalchik/pvz-service/internal/app/tokenizer"
	"github.com/Turalchik/pvz-service/internal/service/pvz_service"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	dsn := database.NewPostgresDSN()
	db, err := database.NewDatabase(dsn, "pgx")
	if err != nil {
		log.Fatalf("Can't create database: %v", err)
	}

	repo := repo.NewRepo(db)
	uuidIfc := pvz_service.NewUUID()
	timerIfc := pvz_service.NewTimer()
	tokenizerIfc := tokenizer.NewTokenizer([]byte(os.Getenv("SECRET_KEY_FOR_JWT")))

	svc, err := pvz_service.NewPVZServiceServer(repo, uuidIfc, timerIfc, tokenizerIfc)
	if err != nil {
		log.Fatalf("failed to create PVZService: %v", err)
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen on :8080: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	desc.RegisterPVZServiceServer(grpcServer, svc)

	reflection.Register(grpcServer)

	log.Println("gRPC server listening on :8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
