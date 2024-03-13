CREATE TABLE
  IF NOT EXISTS routes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name VARCHAR(255) NOT NULL,
    pickup GEOGRAPHY (Point, 4326) DEFAULT NULL,
    destination GEOGRAPHY (Point, 4326) NOT NULL,
    estimated_time BIGINT DEFAULT NULL, -- ms
    distance FLOAT DEFAULT NULL, -- km
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );