CREATE TABLE candidate_job_steps (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  candidate_job_id UUID NOT NULL REFERENCES candidate_jobs (id),
  candidate_job_status candidate_job_enum not null,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
