ALTER TABLE users ADD COLUMN rec_team_id UUID REFERENCES rec_teams(id);