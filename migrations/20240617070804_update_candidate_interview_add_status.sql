CREATE TYPE candidate_interview_status_enum AS ENUM ('invited_to_interview', 'interviewing', 'done', 'cancelled');
ALTER TABLE candidate_interviews
ADD COLUMN candidate_interview_status candidate_interview_status_enum NOT NULL DEFAULT 'invited_to_interview';
