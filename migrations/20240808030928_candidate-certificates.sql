CREATE TABLE candidate_certificates (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "candidate_id" UUID NOT NULL REFERENCES candidates(id),
  "name" VARCHAR(256) NOT NULL,
  "achieved_date" TIMESTAMP WITHOUT TIME ZONE,
  "score" VARCHAR(256) NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
