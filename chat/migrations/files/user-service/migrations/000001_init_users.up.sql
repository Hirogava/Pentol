CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  auth_user_id INTEGER NOT NULL UNIQUE,
  name VARCHAR(255)
);

CREATE TABLE users_desc (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  description TEXT,
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);