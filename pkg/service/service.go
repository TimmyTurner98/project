package service

import "github.com/TimmyTurner98/project/pkg/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
