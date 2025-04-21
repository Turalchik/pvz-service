package pvz_service

import (
	"github.com/Turalchik/pvz-service/internal/app/tokenizer"
	desc "github.com/Turalchik/pvz-service/pkg/pvz_service"
)

type PVZServiceAPI struct {
	desc.UnimplementedPVZServiceServer
	repo               RepoInterface
	uuidInterface      UUIDInterface
	timerInterface     TimerInterface
	tokenizerInterface tokenizer.TokenizerInterface
}

func NewPVZServiceServer(
	repo RepoInterface,
	uuidInterface UUIDInterface,
	timerInterface TimerInterface,
	tokenizerInterface tokenizer.TokenizerInterface,
) (desc.PVZServiceServer, error) {
	return &PVZServiceAPI{
		repo:               repo,
		uuidInterface:      uuidInterface,
		timerInterface:     timerInterface,
		tokenizerInterface: tokenizerInterface,
	}, nil
}
