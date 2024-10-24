package routes

import (
	"github.com/gorilla/mux"

	"real-time-leaderboard/internal/auth"
	"real-time-leaderboard/internal/leaderboard"
	"real-time-leaderboard/internal/middleware"
)

// SetupRoutes defines all the application routes
func SetupRoutes(
	authService *auth.AuthService,
	leaderboardService *leaderboard.LeaderboardService,
) *mux.Router {
	r := mux.NewRouter()

	// Middleware: Add logging middleware for all routes
	r.Use(middleware.LoggingMiddleware)

	// Auth routes
	r.HandleFunc("/register", auth.RegisterHandler(authService)).Methods("POST")
	r.HandleFunc("/login", auth.LoginHandler(authService)).Methods("POST")

	// Leaderboard routes
	r.HandleFunc("/submit-score", leaderboard.SubmitScoreHandler(leaderboardService)).Methods("POST")
	r.HandleFunc("/leaderboard", leaderboard.GetLeaderboardHandler(leaderboardService)).Methods("GET")
	r.HandleFunc("/user-ranking", leaderboard.GetUserRankingHandler(leaderboardService)).Methods("GET")

	return r
}
