CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sender TEXT,
    message TEXT, 
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
