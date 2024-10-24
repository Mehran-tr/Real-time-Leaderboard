package leaderboard

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func SubmitScoreHandler(service *LeaderboardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID string `json:"user_id"`
			Score  int    `json:"score"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		err := service.SubmitScore(context.Background(), req.UserID, req.Score)
		if err != nil {
			http.Error(w, "Failed to submit score", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Score submitted"})
	}
}

func GetLeaderboardHandler(service *LeaderboardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitStr := r.URL.Query().Get("limit")
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil || limit <= 0 {
			limit = 10 // Default limit
		}

		players, err := service.GetTopPlayers(context.Background(), limit)
		if err != nil {
			http.Error(w, "Failed to fetch leaderboard", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(players)
	}
}

func GetUserRankingHandler(service *LeaderboardService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "User ID required", http.StatusBadRequest)
			return
		}

		rank, err := service.GetUserRank(context.Background(), userID)
		if err == redis.Nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Failed to fetch ranking", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]int64{"rank": rank})
	}
}
