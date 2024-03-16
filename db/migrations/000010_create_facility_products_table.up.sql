CREATE TABLE
  IF NOT EXISTS facility_products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    facility_id UUID NOT NULL,
    product_id UUID NOT NULL,
    demand_per_day integer DEFAULT NULL,
    production_per_day integer DEFAULT NULL,
    quantity_on_hand integer DEFAULT NULL,
    CONSTRAINT fk_facility_products_facilities FOREIGN KEY (facility_id) REFERENCES facilities (id) ON DELETE CASCADE,
    CONSTRAINT fk_facility_products_products FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
  );