CREATE TABLE candidate_interviews (
  id UUID PRIMARY KEY NOT NULL,
  title VARCHAR(255) NOT NULL,
  candidate_job_id UUID NOT NULL REFERENCES candidate_jobs (id),
  candidate_job_status candidate_job_enum DEFAULT 'applied',
  interview_date TIMESTAMP NOT NULL,
  start_from TIMESTAMP,
  end_at TIMESTAMP,
  description TEXT
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
