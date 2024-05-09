ALTER TYPE attachment_model_enum ADD VALUE 'candidate_job_feedbacks';
ALTER TABLE attachments ALTER COLUMN relation_type TYPE attachment_model_enum USING relation_type::text::attachment_model_enum;

CREATE TABLE candidate_feedbacks (
  id UUID PRIMARY KEY NOT NULL,
  created_by UUID NOT NULL REFERENCES users (id),
  candidate_job_id UUID NOT NULL REFERENCES candidate_jobs (id),
  feedback TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
