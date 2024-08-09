create type candidate_gender_enum as enum ('male', 'female', 'others');

alter table candidates add column gender candidate_gender_enum not null default 'others';