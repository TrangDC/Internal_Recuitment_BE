CREATE TABLE IF NOT EXISTS hiring_team_approvers (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users (id),
  hiring_team_id UUID NOT NULL REFERENCES hiring_teams (id),
  order_id INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
