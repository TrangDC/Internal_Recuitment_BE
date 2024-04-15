-- Add index on id column
CREATE INDEX idx_users_id ON users (id);
CREATE INDEX idx_users_oid ON users (oid);
CREATE INDEX idx_teams_id ON teams (id);
CREATE INDEX idx_audit_trails_id ON audit_trails (id);
CREATE INDEX idx_hiring_jobs_id ON hiring_jobs (id);

-- Add slug column
ALTER TABLE teams ADD COLUMN slug VARCHAR(255) NOT NULL;
ALTER TABLE hiring_jobs ADD COLUMN slug VARCHAR(255) NOT NULL;
