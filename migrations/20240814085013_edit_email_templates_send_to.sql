DROP TRIGGER IF EXISTS validate_email_send_to_trigger ON email_templates;
DROP FUNCTION IF EXISTS validate_email_send_to;

ALTER TYPE email_send_to_enum RENAME TO email_send_to_enum_old;
Create Type email_send_to_enum as enum ('interviewer', 'job_request', 'hiring_team_manager', 'hiring_team_member', 'role', 'candidate');
ALTER TABLE email_templates
  ALTER COLUMN "send_to" TYPE jsonb;

CREATE OR REPLACE FUNCTION validate_email_send_to()
  RETURNS TRIGGER AS
$$
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
  BEFORE INSERT OR UPDATE
  ON email_templates
  FOR EACH ROW
EXECUTE FUNCTION validate_email_send_to();

DROP TYPE email_send_to_enum_old;
