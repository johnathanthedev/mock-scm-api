CREATE TABLE
  IF NOT EXISTS vehicles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name VARCHAR(255) UNIQUE NOT NULL,
    make VARCHAR(50) DEFAULT NULL,
    model VARCHAR(50) DEFAULT NULL,
    status VARCHAR(20) NOT NULL,
    crew_capacity INT NOT NULL,
    attributes JSONB DEFAULT NULL,
    preferred_speed FLOAT DEFAULT NULL, -- km
    vehicle_type VARCHAR(50) NOT NULL,
    carry_volume FLOAT NOT NULL,
    max_weight INT NOT NULL,
    last_location GEOGRAPHY (Point, 4326),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );