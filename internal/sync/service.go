package sync

import (
	"context"
	"fmt"
)

type Service struct {
	repo    *Repository
	rasRepo *RASRepository
}

func NewService(repo *Repository, rasRepo *RASRepository) *Service {
	return &Service{
		repo:    repo,
		rasRepo: rasRepo,
	}
}

func (s *Service) SyncProformas(ctx context.Context) error {
	return nil
}

func (s *Service) SyncStudents(ctx context.Context) error {
	rc, err := s.rasRepo.GetActiveRecuritmentCycle(ctx)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", rc)

	return nil
}
