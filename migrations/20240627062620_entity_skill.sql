CREATE TYPE entity_skill_type_enum AS ENUM ('candidate', 'hiring_job');

CREATE TABLE entity_skills (
  id UUID PRIMARY KEY NOT NULL,
  entity_type entity_skill_type_enum NOT NULL,
  entity_id uuid NOT NULL,
  skill_id uuid NOT NULL,
  order_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (skill_id) REFERENCES skills(id)
);

ALTER TABLE public.entity_skills ALTER COLUMN id SET DEFAULT uuid_generate_v4();
