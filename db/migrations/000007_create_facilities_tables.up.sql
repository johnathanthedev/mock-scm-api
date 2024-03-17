CREATE TABLE
  IF NOT EXISTS facilities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    max_storage_capacity INT NOT NULL,
    daily_operating_cost INT NOT NULL,
    daily_rent_cost INT NOT NULL,
    daily_carbon_output INT NOT NULL,
    location GEOGRAPHY (Point, 4326) NOT NULL,
    operation_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_facilities_operations FOREIGN KEY (operation_id) REFERENCES operations (id) ON DELETE CASCADE
  )