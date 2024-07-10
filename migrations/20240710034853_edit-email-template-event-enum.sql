ALTER TYPE email_event_enum RENAME TO email_event_enum_old;
CREATE TYPE email_event_enum AS ENUM('candidate_applied_to_kiv', 'candidate_interviewing_to_kiv', 'candidate_interviewing_to_offering', 'created_interview', 'updating_interview', 'cancel_interview');
ALTER TABLE email_templates ALTER COLUMN event TYPE email_event_enum USING "event"::text::email_event_enum;
DROP TYPE email_event_enum_old;
