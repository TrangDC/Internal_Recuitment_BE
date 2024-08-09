ALTER TABLE hiring_jobs
  ADD COLUMN job_position_id uuid REFERENCES job_positions (id);
