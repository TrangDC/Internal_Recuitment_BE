CREATE TYPE candidate_job_level_enum AS ENUM ('intern', 'fresher', 'junior', 'middle', 'senior', 'manager', 'director');

ALTER TABLE candidate_jobs
  ADD COLUMN level candidate_job_level_enum
