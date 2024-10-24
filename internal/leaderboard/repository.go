package leaderboard

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// LeaderboardRepository handles all interactions with Redis
type LeaderboardRepository struct {
	redisClient *redis.Client
}

// NewLeaderboardRepository creates a new repository instance
func NewLeaderboardRepository(redisClient *redis.Client) *LeaderboardRepository {
	return &LeaderboardRepository{redisClient: redisClient}
}

// AddScore adds a score to the leaderboard in Redis
func (r *LeaderboardRepository) AddScore(ctx context.Context, userID string, score int) error {
	return r.redisClient.ZAdd(ctx, "leaderboard", &redis.Z{
		Score:  float64(score),
		Member: userID,
	}).Err()
}

// GetTopPlayers retrieves the top players from the leaderboard
func (r *LeaderboardRepository) GetTopPlayers(ctx context.Context, limit int64) ([]redis.Z, error) {
	return r.redisClient.ZRevRangeWithScores(ctx, "leaderboard", 0, limit-1).Result()
}

// GetUserRank retrieves the rank of a specific user
func (r *LeaderboardRepository) GetUserRank(ctx context.Context, userID string) (int64, error) {
	rank, err := r.redisClient.ZRevRank(ctx, "leaderboard", userID).Result()
	if err != nil {
		return 0, err
	}
	return rank + 1, nil // Redis ranks start at 0, so we add 1
}
