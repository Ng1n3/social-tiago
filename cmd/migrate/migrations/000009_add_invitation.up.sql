CREATE TABLE IF NOT EXISTS users_invitations (
  token bytea PRIMARY KEY,
  user_id BIGINT NOT NULL 
)

