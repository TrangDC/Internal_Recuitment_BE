Create type candidate_history_call_enum as ENUM ('candidate', 'others');

CREATE TABLE candidate_history_calls (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "candidate_id" UUID NOT NULL REFERENCES candidates(id),
  "type" candidate_history_call_enum NOT NULL,
  "contact_to" VARCHAR(256) NULL,
  "date" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  "start_time" TIMESTAMP WITHOUT TIME ZONE,
  "end_time" TIMESTAMP WITHOUT TIME ZONE,
  "description" VARCHAR(512) NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
