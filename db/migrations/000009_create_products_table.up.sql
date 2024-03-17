CREATE TABLE
  IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    weight_kg INT NOT NULL,
    volume_m3 NUMERIC(10, 3) NOT NULL,
    operation_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_products_operations FOREIGN KEY (operation_id) REFERENCES operations (id) ON DELETE CASCADE
  )