package service

import "cmd/main.go/pkg/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
