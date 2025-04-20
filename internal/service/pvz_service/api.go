package pvz_service

import (
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

type PVZServiceAPI struct {
	desc.UnimplementedPVZServiceServer
	repo            RepoInterface
	uuidInterface   UUIDInterface
	secretKeyForJWT []byte
}

func NewPVZServiceServer(repo RepoInterface, uuidInterface UUIDInterface) (desc.PVZServiceServer, error) {
	if err := godotenv.Load(); err != nil {
		return nil, status.Error(codes.Internal, "Cannot load env vars")
	}
	return &PVZServiceAPI{
		repo:            repo,
		uuidInterface:   uuidInterface,
		secretKeyForJWT: []byte(os.Getenv("JWT_SECRET_KEY")),
	}, nil
}
