CREATE TABLE IF NOT EXISTS scores (
                                      id SERIAL PRIMARY KEY,
                                      user_id INT REFERENCES users(id),
    game_id VARCHAR(50),
    score INT,
    created_at TIMESTAMPTZ DEFAULT NOW()
    );
