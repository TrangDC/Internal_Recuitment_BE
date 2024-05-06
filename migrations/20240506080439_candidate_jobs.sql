-- Create candidate_job_enum values
CREATE TYPE candidate_job_enum AS ENUM ('applied', 'interviewing', 'offering', 'hired', 'kiv', 'offer_lost', 'ex_staff');

-- Create candidate_job table
CREATE TABLE candidate_jobs (
  id UUID PRIMARY KEY,
  candidate_id UUID REFERENCES candidates (id) NOT NULL,
  hiring_job_id UUID REFERENCES hiring_jobs (id),
  status candidate_job_enum DEFAULT 'applied',
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);
