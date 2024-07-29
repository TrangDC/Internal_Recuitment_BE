-- Delete foreign keys
ALTER TABLE IF EXISTS team_managers DROP CONSTRAINT IF EXISTS team_managers_team_id_fkey;
ALTER TABLE IF EXISTS hiring_jobs DROP CONSTRAINT IF EXISTS hiring_jobs_team_id_fkey;
ALTER TABLE IF EXISTS users DROP COLUMN IF EXISTS team_id;

-- Rename teams to hiring_teams
ALTER TABLE IF EXISTS teams RENAME TO hiring_teams;
ALTER TABLE IF EXISTS team_managers RENAME TO hiring_team_managers;

-- Rename columns
ALTER TABLE IF EXISTS hiring_team_managers RENAME COLUMN team_id TO hiring_team_id;
ALTER TABLE IF EXISTS hiring_jobs RENAME COLUMN team_id TO hiring_team_id;

-- Add foreign keys
ALTER TABLE IF EXISTS hiring_team_managers ADD CONSTRAINT hiring_team_managers_hiring_team_id_fkey FOREIGN KEY (hiring_team_id) REFERENCES hiring_teams(id);
ALTER TABLE IF EXISTS hiring_jobs ADD CONSTRAINT hiring_jobs_hiring_team_id_fkey FOREIGN KEY (hiring_team_id) REFERENCES hiring_teams(id);
