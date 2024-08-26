ALTER TABLE IF EXISTS hiring_jobs
ADD COLUMN rec_in_charge_id uuid REFERENCES users(id),
ADD COLUMN rec_team_id uuid REFERENCES rec_teams(id);

ALTER TYPE hiring_job_status_enum RENAME VALUE 'draft' TO 'pending_approvals';
ALTER TYPE hiring_job_status_enum ADD VALUE 'cancelled';
ALTER TABLE hiring_jobs ALTER COLUMN status SET DEFAULT 'pending_approvals';

CREATE TYPE hiring_job_level_enum AS ENUM ('intern', 'fresher', 'junior', 'middle', 'senior', 'manager', 'director');

ALTER TABLE hiring_jobs
  ADD COLUMN level hiring_job_level_enum
