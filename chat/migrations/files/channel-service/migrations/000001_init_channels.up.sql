CREATE TABLE channels (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  description TEXT,
  owner_id INTEGER
);

CREATE TABLE channel_posts (
  id SERIAL PRIMARY KEY,
  channel_id INTEGER NOT NULL,
  message TEXT,
  time TIMESTAMP,
  CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES channels(id)
);

CREATE TABLE channel_users (
  id SERIAL PRIMARY KEY,
  channel_id INTEGER NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES channels(id)
);