
# Real-Time Leaderboard System

## Description

This is a real-time leaderboard system built in Go using Redis and PostgreSQL. The application allows users to register, submit scores, and view leaderboards in real-time, with scores being stored in Redis sorted sets.

## Features

- User Authentication (using JWT)
- Submit Scores for various games/activities
- Real-time Leaderboard Updates
- View Global Leaderboard
- View User's Rank on Leaderboard
- Automatic Database Migrations on Startup

## Requirements

- Go (1.18 or newer)
- PostgreSQL
- Redis
- `go-redis` package for Redis connection
- `golang-migrate` package for migrations

## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/Mehran-tr/Real-time-Leaderboard
cd real-time-leaderboard
```

### 2. Set up environment variables:

Create a `.env` file in the root of the project with the following content:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=leaderboard_db
DATABASE_URL=postgres://postgres:yourpassword@localhost:5432/leaderboard_db?sslmode=disable

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=""
REDIS_DB=0
JWT_SECRET=my_secret_key
```

### 3. Install Dependencies:

Run the following command to install required Go modules:

```bash
go mod download
```

### 4. Run Database Migrations:

The application will automatically check for and apply new migrations when it starts. However, you can manually run migrations using the `golang-migrate` tool:

```bash
# Install the migrate tool if you don't have it installed
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run the migrations manually
migrate -path ./migrations -database postgres://postgres:yourpassword@localhost:5432/leaderboard_db?sslmode=disable up
```

### 5. Run the Application:

Start the server:

```bash
go run cmd/main.go
```

The server will start on port 8080 by default. You can access the API at `http://localhost:8080`.

## API Endpoints

### 1. Register User

**POST** `/register`

Request:

```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "secretpassword"
}
```

### 2. Login User

**POST** `/login`

Request:

```json
{
  "email": "john@example.com",
  "password": "secretpassword"
}
```

Response (Success):

```json
{
  "token": "your_jwt_token_here"
}
```

### 3. Submit Score

**POST** `/submit-score`

Request (with JWT token in Authorization header):

```json
{
  "game_id": "game_123",
  "score": 1500
}
```

### 4. Get Global Leaderboard

**GET** `/leaderboard?limit=5`

Response:

```json
[
  {"username": "john_doe", "score": 1500, "rank": 1},
  {"username": "jane_smith", "score": 1400, "rank": 2}
]
```

### 5. Get User Ranking

**GET** `/user-ranking?user_id=1`

Response:

```json
{
  "username": "john_doe",
  "score": 1500,
  "rank": 1
}
```

## Docker Setup (Optional)

You can also run the application using Docker. First, create a `Dockerfile` and build the image:

```bash
docker build -t leaderboard-app .
```

Then run the Docker container:

```bash
docker run -p 8080:8080 --env-file .env leaderboard-app
```

## Testing

To run unit tests for the application:

```bash
go test ./...
```

## License

This project is licensed under the MIT License.
