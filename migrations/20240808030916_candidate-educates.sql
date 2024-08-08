CREATE TABLE candidate_educates (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "candidate_id" UUID NOT NULL REFERENCES candidates(id),
  "school_name" VARCHAR(256) NOT NULL,
  "major" VARCHAR(256) NULL,
  "gpa" VARCHAR(256) NULL,
  "start_date" TIMESTAMP WITHOUT TIME ZONE,
  "end_date" TIMESTAMP WITHOUT TIME ZONE,
  "location" VARCHAR(256) NULL,
  "description" VARCHAR(512) NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
