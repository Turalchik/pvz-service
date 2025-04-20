package pvz_service

import "github.com/google/uuid"

type UUIDInterface interface {
	NewString() string
}

func NewUUID() UUIDInterface {
	return &UUID{}
}

func (*UUID) NewString() string {
	return uuid.NewString()
}

type UUID struct{}
