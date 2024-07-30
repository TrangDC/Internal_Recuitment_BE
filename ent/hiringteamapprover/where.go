// Code generated by ent, DO NOT EDIT.

package hiringteamapprover

import (
	"time"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// HiringTeamID applies equality check predicate on the "hiring_team_id" field. It's identical to HiringTeamIDEQ.
func HiringTeamID(v uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHiringTeamID), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// HiringTeamIDEQ applies the EQ predicate on the "hiring_team_id" field.
func HiringTeamIDEQ(v uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHiringTeamID), v))
	})
}

// HiringTeamIDNEQ applies the NEQ predicate on the "hiring_team_id" field.
func HiringTeamIDNEQ(v uuid.UUID) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHiringTeamID), v))
	})
}

// HiringTeamIDIn applies the In predicate on the "hiring_team_id" field.
func HiringTeamIDIn(vs ...uuid.UUID) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHiringTeamID), v...))
	})
}

// HiringTeamIDNotIn applies the NotIn predicate on the "hiring_team_id" field.
func HiringTeamIDNotIn(vs ...uuid.UUID) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHiringTeamID), v...))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int) predicate.HiringTeamApprover {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v int) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v int) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v int) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v int) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHiringTeam applies the HasEdge predicate on the "hiring_team" edge.
func HasHiringTeam() predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HiringTeamTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HiringTeamTable, HiringTeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHiringTeamWith applies the HasEdge predicate on the "hiring_team" edge with a given conditions (other predicates).
func HasHiringTeamWith(preds ...predicate.HiringTeam) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HiringTeamInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HiringTeamTable, HiringTeamColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HiringTeamApprover) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HiringTeamApprover) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
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
func Not(p predicate.HiringTeamApprover) predicate.HiringTeamApprover {
	return predicate.HiringTeamApprover(func(s *sql.Selector) {
		p(s.Not())
	})
}
