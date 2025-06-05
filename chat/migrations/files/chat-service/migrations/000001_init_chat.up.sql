CREATE TABLE chats (
  id SERIAL PRIMARY KEY,
  user_1_id INTEGER,
  user_2_id INTEGER
);

CREATE TABLE chat_desc (
  id SERIAL PRIMARY KEY,
  chat_id INTEGER NOT NULL,
  name VARCHAR(255),
  description TEXT,
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_chat FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE
);

CREATE TABLE chat_messages (
  id SERIAL PRIMARY KEY,
  chat_id INTEGER NOT NULL,
  sender_id INTEGER,
  message TEXT,
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_chat FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE
);