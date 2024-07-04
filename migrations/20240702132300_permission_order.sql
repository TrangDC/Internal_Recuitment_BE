ALTER TABLE public.permission_groups ADD COLUMN order_id integer;

ALTER TABLE public.permissions ADD COLUMN order_id integer;
ALTER TABLE public.permissions ADD COLUMN parent_id UUID;