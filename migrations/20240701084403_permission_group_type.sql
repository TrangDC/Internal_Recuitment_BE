CREATE TYPE permission_group_type AS ENUM ('function', 'system');

ALTER table permission_groups add column group_type permission_group_type not null default 'function';
