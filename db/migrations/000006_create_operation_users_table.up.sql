CREATE TABLE
  IF NOT EXISTS operation_users (
    operation_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_operation_users PRIMARY KEY (operation_id, user_id),
    CONSTRAINT fk_operation_users_operations FOREIGN KEY (operation_id) REFERENCES operations (id) ON DELETE CASCADE,
    CONSTRAINT fk_operation_users_users FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
  );