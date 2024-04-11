CREATE TYPE audit_trail_action_type_enum AS ENUM ('create', 'update', 'delete');
CREATE TYPE audit_trail_module_enum AS ENUM ('teams');

CREATE TABLE public.audit_trails (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	created_by uuid NULL REFERENCES users(id),
	record_id uuid NULL,
	"module" public."audit_trail_module_enum" NOT NULL,
	action_type public."audit_trail_action_type_enum" DEFAULT 'create'::audit_trail_action_type_enum NOT NULL,
	note varchar(500) NULL,
	record_changes text NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);
