CREATE TYPE user_status_enum AS ENUM ('active', 'inactive');
ALTER TABLE users
ADD COLUMN status user_status_enum NOT NULL DEFAULT 'active';
