create type outgoing_email_status as enum ('pending', 'sent', 'failed');

CREATE TABLE outgoing_emails (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "to" jsonb NOT NULL,
  "cc" jsonb,
  "bcc" jsonb,
  "subject" TEXT NOT NULL,
  "content" TEXT NOT NULL,
  "signature" TEXT,
  "email_template_id" UUID REFERENCES email_templates(id),
  "status" outgoing_email_status NOT NULL DEFAULT 'pending',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
