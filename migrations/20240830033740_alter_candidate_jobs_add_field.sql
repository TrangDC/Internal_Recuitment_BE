ALTER TABLE candidate_jobs
ADD COLUMN rec_in_charge_id uuid REFERENCES users(id);
