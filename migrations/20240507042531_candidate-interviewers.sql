CREATE TABLE candidate_interviewers (
  id UUID PRIMARY KEY NOT NULL,
  candidate_interview_id UUID REFERENCES candidate_interviews NOT NULL,
  user_id UUID REFERENCES users NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
