package leaderboard

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockRedisClient struct{}

func (m *mockRedisClient) ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	return redis.NewIntResult(1, nil)
}

func (m *mockRedisClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	players := []redis.Z{
		{Score: 1500, Member: "user1"},
		{Score: 1200, Member: "user2"},
	}
	return redis.NewZSliceResult(players, nil)
}

func (m *mockRedisClient) ZRevRank(ctx context.Context, key string, member string) *redis.IntCmd {
	if member == "user1" {
		return redis.NewIntResult(0, nil) // rank 1 (0-indexed)
	}
	return redis.NewIntResult(1, nil) // rank 2 (0-indexed)
}

func TestSubmitScore(t *testing.T) {
	redisClient := &mockRedisClient{}
	service := NewLeaderboardService(redisClient)

	err := service.SubmitScore(context.Background(), "user1", 1500)
	assert.NoError(t, err)
}

func TestGetTopPlayers(t *testing.T) {
	redisClient := &mockRedisClient{}
	service := NewLeaderboardService(redisClient)

	players, err := service.GetTopPlayers(context.Background(), 2)
	assert.NoError(t, err)
	assert.Len(t, players, 2)
	assert.Equal(t, players[0].Member, "user1")
	assert.Equal(t, players[0].Score, 1500.0)
	assert.Equal(t, players[1].Member, "user2")
	assert.Equal(t, players[1].Score, 1200.0)
}

func TestGetUserRank(t *testing.T) {
	redisClient := &mockRedisClient{}
	service := NewLeaderboardService(redisClient)

	rank, err := service.GetUserRank(context.Background(), "user1")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rank) // rank 1 (1-indexed)
}
