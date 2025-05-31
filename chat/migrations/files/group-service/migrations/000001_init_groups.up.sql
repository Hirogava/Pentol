CREATE TABLE groups (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  description TEXT,
  owner_id INTEGER
);

CREATE TABLE group_users (
  id SERIAL PRIMARY KEY,
  group_id INTEGER NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE groups_messages (
  id SERIAL PRIMARY KEY,
  message TEXT,
  group_id INTEGER NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_group FOREIGN KEY (group_id) REFERENCES groups(id)
);