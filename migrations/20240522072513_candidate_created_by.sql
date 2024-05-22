ALTER TABLE candidate_jobs
ADD COLUMN created_by uuid REFERENCES users(id);