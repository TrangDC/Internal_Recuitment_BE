ALTER TABLE teams
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE users
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE candidates
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE hiring_jobs
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE team_managers
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE candidate_jobs
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE candidate_interviews
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE candidate_interviewers
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE candidate_job_feedbacks
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE attachments
ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE audit_trails
ALTER COLUMN id SET DEFAULT uuid_generate_v4();