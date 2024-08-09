ALTER TYPE candidate_job_enum RENAME TO candidate_job_enum_old;
CREATE TYPE candidate_job_enum AS ENUM('applied', 'interviewing', 'offering', 'hired', 'failed_cv', 'kiv','failed_interview', 'offer_lost', 'ex_staff');

ALTER TABLE public.candidate_jobs ALTER COLUMN status DROP DEFAULT;
ALTER TABLE public.candidate_job_feedbacks ALTER COLUMN candidate_job_status DROP DEFAULT;
ALTER TABLE public.candidate_interviews ALTER COLUMN candidate_job_status DROP DEFAULT;
ALTER TABLE public.candidate_job_steps ALTER COLUMN candidate_job_status DROP DEFAULT;

ALTER TABLE candidate_jobs ALTER COLUMN status TYPE candidate_job_enum USING "status"::text::candidate_job_enum;
ALTER TABLE candidate_job_feedbacks ALTER COLUMN candidate_job_status TYPE candidate_job_enum USING "candidate_job_status"::text::candidate_job_enum;
ALTER TABLE candidate_interviews ALTER COLUMN candidate_job_status TYPE candidate_job_enum USING "candidate_job_status"::text::candidate_job_enum;
ALTER TABLE candidate_job_steps ALTER COLUMN candidate_job_status TYPE candidate_job_enum USING "candidate_job_status"::text::candidate_job_enum;

DROP TYPE candidate_job_enum_old;

-- change all value of status column is kiv to failed_cv
DO $$
BEGIN
  UPDATE candidate_jobs
  SET status = 'failed_cv'
  WHERE status = 'kiv';
END $$;
DO $$
BEGIN
  UPDATE candidate_job_feedbacks
  SET candidate_job_status = 'failed_cv'
  WHERE candidate_job_status = 'kiv';
END $$;
DO $$
BEGIN
  UPDATE candidate_interviews
  SET candidate_job_status = 'failed_cv'
  WHERE candidate_job_status = 'kiv';
END $$;
DO $$
BEGIN
  UPDATE candidate_job_steps
  SET candidate_job_status = 'failed_cv'
  WHERE candidate_job_status = 'kiv';
END $$;

-- remove enum kiv out of candidate_job_enum type
ALTER TYPE candidate_job_enum RENAME TO candidate_job_enum_old;
CREATE TYPE candidate_job_enum AS ENUM('applied', 'interviewing', 'offering', 'hired', 'failed_cv', 'failed_interview', 'offer_lost', 'ex_staff');

ALTER TABLE candidate_jobs ALTER COLUMN status TYPE candidate_job_enum USING "status"::text::candidate_job_enum;
ALTER TABLE candidate_job_feedbacks ALTER COLUMN candidate_job_status TYPE candidate_job_enum USING "candidate_job_status"::text::candidate_job_enum;
ALTER TABLE candidate_interviews ALTER COLUMN candidate_job_status TYPE candidate_job_enum USING "candidate_job_status"::text::candidate_job_enum;
ALTER TABLE candidate_job_steps ALTER COLUMN candidate_job_status TYPE candidate_job_enum USING "candidate_job_status"::text::candidate_job_enum;

DROP TYPE candidate_job_enum_old;

ALTER TABLE public.candidate_jobs ALTER COLUMN status SET DEFAULT 'applied';
ALTER TABLE public.candidate_job_feedbacks ALTER COLUMN candidate_job_status SET DEFAULT 'applied';
ALTER TABLE public.candidate_interviews ALTER COLUMN candidate_job_status SET DEFAULT 'applied';
ALTER TABLE public.candidate_job_steps ALTER COLUMN candidate_job_status SET DEFAULT 'applied';
