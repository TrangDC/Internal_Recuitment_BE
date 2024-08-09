CREATE TABLE IF NOT EXISTS candidate_notes (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "candidate_id" UUID NOT NULL REFERENCES candidates(id),
  "created_by_id" UUID NOT NULL REFERENCES users(id),
  "name" VARCHAR(256) NOT NULL,
  "description" TEXT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

ALTER TYPE attachment_model_enum ADD VALUE IF NOT EXISTS 'candidate_notes';
