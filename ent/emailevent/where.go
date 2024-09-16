// Code generated by ent, DO NOT EDIT.

package emailevent

import (
	"time"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ModuleEQ applies the EQ predicate on the "module" field.
func ModuleEQ(v Module) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModule), v))
	})
}

// ModuleNEQ applies the NEQ predicate on the "module" field.
func ModuleNEQ(v Module) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModule), v))
	})
}

// ModuleIn applies the In predicate on the "module" field.
func ModuleIn(vs ...Module) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldModule), v...))
	})
}

// ModuleNotIn applies the NotIn predicate on the "module" field.
func ModuleNotIn(vs ...Module) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldModule), v...))
	})
}

// ActionEQ applies the EQ predicate on the "action" field.
func ActionEQ(v Action) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAction), v))
	})
}

// ActionNEQ applies the NEQ predicate on the "action" field.
func ActionNEQ(v Action) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAction), v))
	})
}

// ActionIn applies the In predicate on the "action" field.
func ActionIn(vs ...Action) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAction), v...))
	})
}

// ActionNotIn applies the NotIn predicate on the "action" field.
func ActionNotIn(vs ...Action) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAction), v...))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.EmailEvent {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// HasTemplateEdges applies the HasEdge predicate on the "template_edges" edge.
func HasTemplateEdges() predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TemplateEdgesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TemplateEdgesTable, TemplateEdgesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTemplateEdgesWith applies the HasEdge predicate on the "template_edges" edge with a given conditions (other predicates).
func HasTemplateEdgesWith(preds ...predicate.EmailTemplate) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TemplateEdgesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TemplateEdgesTable, TemplateEdgesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOutgoingEmailEdges applies the HasEdge predicate on the "outgoing_email_edges" edge.
func HasOutgoingEmailEdges() predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutgoingEmailEdgesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutgoingEmailEdgesTable, OutgoingEmailEdgesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutgoingEmailEdgesWith applies the HasEdge predicate on the "outgoing_email_edges" edge with a given conditions (other predicates).
func HasOutgoingEmailEdgesWith(preds ...predicate.OutgoingEmail) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OutgoingEmailEdgesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OutgoingEmailEdgesTable, OutgoingEmailEdgesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmailEvent) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmailEvent) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EmailEvent) predicate.EmailEvent {
	return predicate.EmailEvent(func(s *sql.Selector) {
		p(s.Not())
	})
}