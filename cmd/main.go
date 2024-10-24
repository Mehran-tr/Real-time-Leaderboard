package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import for reading migration files
	"github.com/joho/godotenv"
	"real-time-leaderboard/internal/auth"
	"real-time-leaderboard/internal/leaderboard"
	"real-time-leaderboard/internal/routes"
	"real-time-leaderboard/pkg/database"
	"real-time-leaderboard/pkg/redis" // Import the Redis client setup
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Run database migrations before starting the server
	runMigrations(db)

	// Initialize services
	authService := setupAuthService(db)
	leaderboardService := setupLeaderboardService()

	// Setup the router with defined routes
	router := routes.SetupRoutes(authService, leaderboardService)

	// Start the HTTP server
	log.Println("Server running on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// runMigrations checks for and applies any new database migrations
func runMigrations(db *sql.DB) {
	// Create a migration driver from the PostgreSQL database connection
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create PostgreSQL driver: %v", err)
	}

	// Path to the migrations folder
	migrationsDir := "file://migrations"

	// Initialize the migrate instance
	m, err := migrate.NewWithDatabaseInstance(migrationsDir, "postgres", driver)
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}

	// Run the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

// Initialize AuthService (assuming it takes a db connection)
func setupAuthService(db *sql.DB) *auth.AuthService {
	userRepo := auth.NewUserRepository(db)
	authService := auth.NewAuthService(userRepo)
	return authService
}

// Initialize LeaderboardService (using Redis)
func setupLeaderboardService() *leaderboard.LeaderboardService {
	redisClient := redis.NewRedisClient() // Initialize Redis client
	leaderboardRepo := leaderboard.NewLeaderboardRepository(redisClient)
	leaderboardService := leaderboard.NewLeaderboardService(leaderboardRepo)
	return leaderboardService
}
