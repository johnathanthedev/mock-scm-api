CREATE TABLE
  IF NOT EXISTS route_stops (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    route_id UUID NOT NULL,
    facility_id UUID NOT NULL,
    sequence integer DEFAULT 1 NOT NULL,
    CONSTRAINT fk_route_stops_routes FOREIGN KEY (route_id) REFERENCES routes (id) ON DELETE CASCADE,
    CONSTRAINT fk_route_stops_facilities FOREIGN KEY (facility_id) REFERENCES facilities (id) ON DELETE CASCADE
  );