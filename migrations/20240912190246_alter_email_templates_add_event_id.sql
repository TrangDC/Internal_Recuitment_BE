ALTER TABLE IF EXISTS email_templates ADD COLUMN IF NOT EXISTS event_id uuid REFERENCES email_events(id);

UPDATE email_templates
SET event_id = CASE
    WHEN event = 'candidate_applied_to_kiv' THEN '1f1120bd-32af-4a74-ae96-e544efdce755'::UUID
    WHEN event = 'candidate_interviewing_to_kiv' THEN 'd43349e1-8062-4b97-99de-1fbfa23183dd'::UUID
    WHEN event = 'candidate_interviewing_to_offering' THEN '097520f8-c92c-4753-9562-e40e83d1de5a'::UUID
    WHEN event = 'created_interview' THEN '2e25bedd-a87d-497c-aa39-1e9ee5879642'::UUID
    WHEN event = 'updating_interview' THEN '466f2006-d51e-4a16-98cf-6ff9b73838c7'::UUID
    WHEN event = 'cancel_interview' THEN 'ae4055be-72b2-46f3-9de6-6ee9789cc6b2'::UUID
  END
WHERE deleted_at IS NULL;

ALTER TABLE IF EXISTS outgoing_emails ADD COLUMN IF NOT EXISTS event_id uuid REFERENCES email_events(id);

UPDATE outgoing_emails
SET event_id = CASE
    WHEN event = 'candidate_applied_to_kiv' THEN '1f1120bd-32af-4a74-ae96-e544efdce755'::UUID
    WHEN event = 'candidate_interviewing_to_kiv' THEN 'd43349e1-8062-4b97-99de-1fbfa23183dd'::UUID
    WHEN event = 'candidate_interviewing_to_offering' THEN '097520f8-c92c-4753-9562-e40e83d1de5a'::UUID
    WHEN event = 'created_interview' THEN '2e25bedd-a87d-497c-aa39-1e9ee5879642'::UUID
    WHEN event = 'updating_interview' THEN '466f2006-d51e-4a16-98cf-6ff9b73838c7'::UUID
    WHEN event = 'cancel_interview' THEN 'ae4055be-72b2-46f3-9de6-6ee9789cc6b2'::UUID
  END
WHERE deleted_at IS NULL;
