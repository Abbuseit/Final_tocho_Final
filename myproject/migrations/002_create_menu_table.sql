
CREATE TABLE menu (
                      id SERIAL PRIMARY KEY,
                      name TEXT NOT NULL,
                      price DECIMAL(10, 2) NOT NULL,
                      description TEXT,
                      created_at TIMESTAMP DEFAULT NOW()
);
