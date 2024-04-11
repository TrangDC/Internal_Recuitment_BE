package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team. NOTE : Part of the public API ( ultimately exposed to end team
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID")),
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).Optional().Annotations(entgql.OrderField("UPDATED_AT")),
		field.Time("deleted_at").Optional().Annotations(entgql.OrderField("DELETED_AT")),
	}
}

// Edges of the Team
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_edges", User.Type).Ref("team_edges").Through("user_teams", TeamManager.Type).
			Comment("The uniqueness of the user is enforced on the edge schema"),
	}
}
