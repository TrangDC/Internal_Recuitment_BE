CREATE TYPE reference_type_enum AS ENUM ('eb', 'rec', 'hiring_platform', 'reference', 'headhunt');

ALTER TABLE candidates 
ADD COLUMN reference_uid uuid REFERENCES users,
ADD COLUMN reference_type reference_type_enum NULL,
ADD COLUMN reference_value VARCHAR NULL,
ADD COLUMN country VARCHAR NULL,
ADD COLUMN recruit_time TIMESTAMP NULL,
ADD COLUMN description VARCHAR(255) NULL,
ADD COLUMN attachment_id uuid NULL;