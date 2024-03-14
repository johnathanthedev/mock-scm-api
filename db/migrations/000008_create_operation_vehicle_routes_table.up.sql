CREATE TABLE
  operation_vehicle_routes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    operation_id UUID NOT NULL,
    vehicle_id UUID NOT NULL,
    route_id UUID NOT NULL,
    sequence INTEGER DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_operation FOREIGN KEY (operation_id) REFERENCES operations (id),
    CONSTRAINT fk_vehicle FOREIGN KEY (vehicle_id) REFERENCES vehicles (id),
    CONSTRAINT fk_route FOREIGN KEY (route_id) REFERENCES routes (id)
  );