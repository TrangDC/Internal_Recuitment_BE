CREATE TABLE candidate_exps (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "candidate_id" UUID NOT NULL REFERENCES candidates(id),
  "position" VARCHAR(256) NOT NULL,
  "company" VARCHAR(256) NOT NULL,
  "location" VARCHAR(256) NULL,
  "start_date" TIMESTAMP WITHOUT TIME ZONE,
  "end_date" TIMESTAMP WITHOUT TIME ZONE,
  "description" VARCHAR(512) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
