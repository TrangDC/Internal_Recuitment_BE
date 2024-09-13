ALTER TYPE email_send_to_enum ADD VALUE IF NOT EXISTS 'job_rec_in_charge';
ALTER TYPE email_send_to_enum ADD VALUE IF NOT EXISTS 'cd_job_rec_in_charge';
ALTER TYPE email_send_to_enum ADD VALUE IF NOT EXISTS 'hiring_approver';
ALTER TYPE email_send_to_enum ADD VALUE IF NOT EXISTS 'rec_leader';
ALTER TYPE email_send_to_enum ADD VALUE IF NOT EXISTS 'rec_member';

ALTER TYPE recipient_type_enum ADD VALUE IF NOT EXISTS 'job_rec_in_charge';
ALTER TYPE recipient_type_enum ADD VALUE IF NOT EXISTS 'cd_job_rec_in_charge';
ALTER TYPE recipient_type_enum ADD VALUE IF NOT EXISTS 'hiring_approver';
ALTER TYPE recipient_type_enum ADD VALUE IF NOT EXISTS 'rec_leader';
ALTER TYPE recipient_type_enum ADD VALUE IF NOT EXISTS 'rec_member';
