ALTER TABLE IF EXISTS hiring_job_steps ALTER COLUMN type DROP DEFAULT;
ALTER TABLE IF EXISTS hiring_job_steps RENAME COLUMN type TO status;
ALTER TABLE IF EXISTS hiring_job_steps ADD COLUMN IF NOT EXISTS order_id INTEGER NOT NULL;
ALTER TABLE IF EXISTS hiring_job_steps RENAME COLUMN created_by_id TO user_id;
ALTER TYPE hiring_job_step_enum RENAME VALUE 'created' TO 'pending';
ALTER TYPE hiring_job_step_enum RENAME VALUE 'opened' TO 'accepted';
ALTER TYPE hiring_job_step_enum RENAME VALUE 'closed' TO 'rejected';
ALTER TYPE hiring_job_step_enum ADD VALUE IF NOT EXISTS 'waiting';
