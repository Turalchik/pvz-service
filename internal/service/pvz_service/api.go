package pvz_service

import (
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
)

type PVZServiceAPI struct {
	desc.UnimplementedPVZServiceServer
	repo RepoInterface
}

func NewPVZServiceServer(repo RepoInterface) desc.PVZServiceServer {
	return &PVZServiceAPI{
		repo: repo,
	}
}
