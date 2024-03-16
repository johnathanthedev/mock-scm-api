CREATE TABLE
  IF NOT EXISTS routes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name VARCHAR(255) NOT NULL,
    operation_id UUID NOT NULL,
    origin_facility_id UUID DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_routes_operations FOREIGN KEY (operation_id) REFERENCES operations (id) ON DELETE CASCADE,
    CONSTRAINT fk_routes_origin_facilities FOREIGN KEY (origin_facility_id) REFERENCES facilities (id) ON DELETE SET NULL
  );