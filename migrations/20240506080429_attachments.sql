CREATE TYPE attachment_model_enum AS ENUM ('candidate_jobs');

CREATE TABLE attachments (
  id UUID PRIMARY KEY,
  document_name VARCHAR(255),
  document_id UUID UNIQUE,
  relation_type attachment_model_enum NOT NULL,
  relation_id UUID NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
