Create Type email_event_enum as enum ('candidate_applied_to_kiv', 'candidate_interviewing_to_kiv', 'candidate_interviewing_to_offering', 'created_candidate', 'updating_interview');
Create Type email_send_to_enum as enum ('interviewer', 'job_request', 'team_manager', 'team_member', 'role', 'candidate');
CREATE TYPE email_template_status_enum as enum ('active', 'inactive');

CREATE TABLE email_templates (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "event" email_event_enum not null,
  "send_to" jsonb not null,
  "subject" TEXT NOT NULL,
  "content" TEXT NOT NULL,
  "signature" TEXT,
  "status" email_template_status_enum not null default 'active',
  "cc" jsonb,
  "bcc" jsonb,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE OR REPLACE FUNCTION validate_email_send_to()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.send_to IS NOT NULL THEN
        PERFORM 1
        FROM jsonb_array_elements_text(NEW.send_to) AS elem
        WHERE elem::email_send_to_enum IS NULL;

        IF FOUND THEN
            RAISE EXCEPTION 'Invalid value in send_to';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER validate_email_send_to_trigger
BEFORE INSERT OR UPDATE ON email_templates
FOR EACH ROW EXECUTE FUNCTION validate_email_send_to();
