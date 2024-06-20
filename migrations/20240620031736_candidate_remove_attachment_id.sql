ALTER TABLE candidates
DROP COLUMN attachment_id;

ALTER TYPE attachment_model_enum ADD VALUE 'candidates';
