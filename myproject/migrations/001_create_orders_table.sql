
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INT NOT NULL,
                        items TEXT[] NOT NULL,
                        created_at TIMESTAMP DEFAULT NOW()
);
