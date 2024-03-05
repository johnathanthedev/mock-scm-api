CREATE TABLE
  IF NOT EXISTS vehicles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    operation_id UUID DEFAULT NULL,
    vehicle_type VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    carry_volume FLOAT NOT NULL,
    max_weight INT NOT NULL,
    last_location GEOGRAPHY (Point, 4326),
    departure_location GEOGRAPHY (Point, 4326),
    arrival_location GEOGRAPHY (Point, 4326),
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (operation_id) REFERENCES operations (id)
  );