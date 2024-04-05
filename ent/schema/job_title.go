package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// JobTitle holds the schema definition for the jobTitle entity.
type JobTitle struct {
	ent.Schema
}

// Fields of the JobTitle. NOTE : Part of the public API ( ultimately exposed to end users
func (JobTitle) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID")),
		field.String("code").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("CODE")),
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("description").Optional().MaxLen(255).Annotations(entgql.OrderField("DESCRIPTION")),
		field.String("specification").Optional().MaxLen(255).Annotations(entgql.OrderField("SPECIFICATION")),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Optional().Annotations(entgql.OrderField("UPDATED_AT")),
		field.Time("deleted_at").Optional().Annotations(entgql.OrderField("DELETED_AT")),
	}
}
