CREATE TABLE groups (
  id SERIAL PRIMARY KEY,
  owner_id INTEGER
);

CREATE TABLE group_desc (
  id SERIAL PRIMARY KEY,
  group_id INTEGER NOT NULL,
  name VARCHAR(255),
  description TEXT,
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

CREATE TABLE group_users (
  id SERIAL PRIMARY KEY,
  group_id INTEGER NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

CREATE TABLE groups_messages (
  id SERIAL PRIMARY KEY,
  message TEXT,
  group_id INTEGER NOT NULL,
  user_id INTEGER,
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);