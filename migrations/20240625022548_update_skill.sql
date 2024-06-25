ALTER TABLE skills
ADD COLUMN skill_type_id UUID REFERENCES skill_type(id);
