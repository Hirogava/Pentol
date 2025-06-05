CREATE TABLE channels (
  id SERIAL PRIMARY KEY,
  owner_id INTEGER
);

CREATE TABLE channel_desc (
  id SERIAL PRIMARY KEY,
  channel_id INTEGER NOT NULL,
  name VARCHAR(255),
  description TEXT,
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE
);

CREATE TABLE channel_posts (
  id SERIAL PRIMARY KEY,
  channel_id INTEGER NOT NULL,
  message TEXT,
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE
);

CREATE TABLE channel_users (
  id SERIAL PRIMARY KEY,
  channel_id INTEGER NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE
);