ALTER TABLE candidate_interviews
ADD COLUMN created_by uuid REFERENCES users(id);
