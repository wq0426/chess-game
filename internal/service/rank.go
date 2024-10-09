package service

import (
	"chess/internal/model"
	"chess/internal/repository"
	"context"
)

type RankService interface {
	GetRank(ctx context.Context) ([]*model.RankList, error)
	AddUserRank(ctx context.Context, param *model.UserRankParam) (bool, error)
}

func NewRankService(
	service *Service,
	rankRepository repository.RankRepository,
) RankService {
	return &rankService{
		Service:        service,
		rankRepository: rankRepository,
	}
}

type rankService struct {
	*Service
	rankRepository repository.RankRepository
}

func (s *rankService) GetRank(ctx context.Context) ([]*model.RankList, error) {
	return s.rankRepository.GetRank(ctx)
}

func (s *rankService) AddUserRank(ctx context.Context, param *model.UserRankParam) (bool, error) {
	return s.rankRepository.PostRank(ctx, param)
}
