ALTER TABLE public.candidates ALTER COLUMN description TYPE varchar(512) USING description::varchar(512);
ALTER TABLE public.skill_types ALTER COLUMN description TYPE varchar(512) USING description::varchar(512);
ALTER TABLE public.skills ALTER COLUMN description TYPE varchar(512) USING description::varchar(512);
