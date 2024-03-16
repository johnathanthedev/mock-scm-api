CREATE TABLE
  IF NOT EXISTS operation_users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    operation_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_operations_users_operations FOREIGN KEY (operation_id) REFERENCES operations (id) ON DELETE CASCADE,
    CONSTRAINT fk_operations_users_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
  );