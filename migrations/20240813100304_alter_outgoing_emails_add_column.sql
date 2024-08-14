CREATE TYPE recipient_type_enum AS ENUM ('interviewer', 'job_request', 'hiring_team_manager', 'hiring_team_member', 'role', 'candidate');

ALTER TABLE outgoing_emails
  ADD COLUMN candidate_id   UUID REFERENCES candidates (id),
  ADD COLUMN recipient_type recipient_type_enum;