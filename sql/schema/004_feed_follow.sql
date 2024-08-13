-- +goose Up
CREATE TABLE feed_follow (
  id UUID PRIMARY KEY,
  feed_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL REFERENCES users
                  ON DELETE CASCADE,

  FOREIGN KEY(user_id) REFERENCES users (id)
);

-- +goose Down
DROP TABLE feed_follow;