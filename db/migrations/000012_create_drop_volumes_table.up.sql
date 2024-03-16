CREATE TABLE
  IF NOT EXISTS drop_volumes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    product_id UUID NOT NULL,
    route_stop_id UUID NOT NULL,
    deliver_quantity integer NOT NULL,
    pickup_quantity integer NOT NULL,
    CONSTRAINT fk_drop_volumes_products FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    CONSTRAINT fk_drop_volumes_route_stops FOREIGN KEY (route_stop_id) REFERENCES route_stops (id) ON DELETE CASCADE
  );