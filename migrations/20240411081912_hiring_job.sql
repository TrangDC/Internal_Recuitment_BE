CREATE TYPE hiring_job_currencies_enum AS ENUM ('vnd', 'usd', 'jpy');
CREATE TYPE hiring_job_locations_enum AS ENUM ('ha_noi', 'ho_chi_minh', 'da_nang', 'japan');
CREATE TYPE hiring_job_salary_type_enum AS ENUM ('range', 'up_to', 'negotiate', 'minimum');
CREATE TYPE hiring_job_status_enum AS ENUM ('draft', 'opened', 'closed');

CREATE TABLE hiring_jobs (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(255) NOT NULL,
  "description" TEXT NOT NULL,
  created_by UUID NOT NULL REFERENCES users(id),
  team_id UUID NOT NULL REFERENCES teams(id),
  amount INT NOT NULL DEFAULT 0 CHECK (amount >= 0),
  "location" public."hiring_job_locations_enum" NOT NULL,
  "status" public."hiring_job_status_enum" NOT NULL DEFAULT 'draft',
  salary_type public."hiring_job_salary_type_enum" NOT NULL,
  salary_from INT DEFAULT 0 CHECK (salary_from >= 0),
  salary_to INT DEFAULT 0 CHECK (salary_to >= 0),
  currency public."hiring_job_currencies_enum",
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);