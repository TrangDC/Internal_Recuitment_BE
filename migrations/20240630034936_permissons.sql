CREATE TABLE permissions (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  group_id UUID REFERENCES permission_groups(id) ON DELETE CASCADE,
  for_owner BOOLEAN DEFAULT FALSE,
  for_team BOOLEAN DEFAULT FALSE,
  for_all BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);