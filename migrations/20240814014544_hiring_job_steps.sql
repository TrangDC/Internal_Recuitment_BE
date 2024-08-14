CREATE TYPE hiring_job_step_enum AS ENUM ('created', 'opened', 'closed');

CREATE TABLE hiring_job_steps (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "hiring_job_id" UUID NOT NULL REFERENCES hiring_jobs (id),
  "type" hiring_job_step_enum DEFAULT 'created',
  "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
  "updated_at" TIMESTAMP WITHOUT TIME ZONE
);