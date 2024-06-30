Drop Table user_permissions;
Drop Table role_permissions;

CREATE TYPE entity_permission_enum AS ENUM ('user', 'role');

CREATE TABLE entity_permissions (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  entity_id UUID,
  permission_id UUID REFERENCES permissions (id) ON DELETE CASCADE,
  for_owner BOOLEAN DEFAULT FALSE,
  for_team BOOLEAN DEFAULT FALSE,
  for_all BOOLEAN DEFAULT FALSE,
  entity_type entity_permission_enum,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);