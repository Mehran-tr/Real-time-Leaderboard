package leaderboard

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type LeaderboardService struct {
	repo *LeaderboardRepository
}

func NewLeaderboardService(repo *LeaderboardRepository) *LeaderboardService {
	return &LeaderboardService{repo: repo}
}

func (s *LeaderboardService) SubmitScore(ctx context.Context, userID string, score int) error {
	return s.repo.AddScore(ctx, userID, score)
}

func (s *LeaderboardService) GetTopPlayers(ctx context.Context, limit int64) ([]redis.Z, error) {
	return s.repo.GetTopPlayers(ctx, limit)
}

func (s *LeaderboardService) GetUserRank(ctx context.Context, userID string) (int64, error) {
	return s.repo.GetUserRank(ctx, userID)
}
