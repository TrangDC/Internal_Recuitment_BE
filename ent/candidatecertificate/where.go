// Code generated by ent, DO NOT EDIT.

package candidatecertificate

import (
	"time"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// CandidateID applies equality check predicate on the "candidate_id" field. It's identical to CandidateIDEQ.
func CandidateID(v uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCandidateID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// AchievedDate applies equality check predicate on the "achieved_date" field. It's identical to AchievedDateEQ.
func AchievedDate(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAchievedDate), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// CandidateIDEQ applies the EQ predicate on the "candidate_id" field.
func CandidateIDEQ(v uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCandidateID), v))
	})
}

// CandidateIDNEQ applies the NEQ predicate on the "candidate_id" field.
func CandidateIDNEQ(v uuid.UUID) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCandidateID), v))
	})
}

// CandidateIDIn applies the In predicate on the "candidate_id" field.
func CandidateIDIn(vs ...uuid.UUID) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCandidateID), v...))
	})
}

// CandidateIDNotIn applies the NotIn predicate on the "candidate_id" field.
func CandidateIDNotIn(vs ...uuid.UUID) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCandidateID), v...))
	})
}

// CandidateIDIsNil applies the IsNil predicate on the "candidate_id" field.
func CandidateIDIsNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCandidateID)))
	})
}

// CandidateIDNotNil applies the NotNil predicate on the "candidate_id" field.
func CandidateIDNotNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCandidateID)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScore), v))
	})
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...string) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldScore), v...))
	})
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...string) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldScore), v...))
	})
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScore), v))
	})
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScore), v))
	})
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScore), v))
	})
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScore), v))
	})
}

// ScoreContains applies the Contains predicate on the "score" field.
func ScoreContains(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldScore), v))
	})
}

// ScoreHasPrefix applies the HasPrefix predicate on the "score" field.
func ScoreHasPrefix(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldScore), v))
	})
}

// ScoreHasSuffix applies the HasSuffix predicate on the "score" field.
func ScoreHasSuffix(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldScore), v))
	})
}

// ScoreIsNil applies the IsNil predicate on the "score" field.
func ScoreIsNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldScore)))
	})
}

// ScoreNotNil applies the NotNil predicate on the "score" field.
func ScoreNotNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldScore)))
	})
}

// ScoreEqualFold applies the EqualFold predicate on the "score" field.
func ScoreEqualFold(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldScore), v))
	})
}

// ScoreContainsFold applies the ContainsFold predicate on the "score" field.
func ScoreContainsFold(v string) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldScore), v))
	})
}

// AchievedDateEQ applies the EQ predicate on the "achieved_date" field.
func AchievedDateEQ(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAchievedDate), v))
	})
}

// AchievedDateNEQ applies the NEQ predicate on the "achieved_date" field.
func AchievedDateNEQ(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAchievedDate), v))
	})
}

// AchievedDateIn applies the In predicate on the "achieved_date" field.
func AchievedDateIn(vs ...time.Time) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAchievedDate), v...))
	})
}

// AchievedDateNotIn applies the NotIn predicate on the "achieved_date" field.
func AchievedDateNotIn(vs ...time.Time) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAchievedDate), v...))
	})
}

// AchievedDateGT applies the GT predicate on the "achieved_date" field.
func AchievedDateGT(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAchievedDate), v))
	})
}

// AchievedDateGTE applies the GTE predicate on the "achieved_date" field.
func AchievedDateGTE(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAchievedDate), v))
	})
}

// AchievedDateLT applies the LT predicate on the "achieved_date" field.
func AchievedDateLT(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAchievedDate), v))
	})
}

// AchievedDateLTE applies the LTE predicate on the "achieved_date" field.
func AchievedDateLTE(v time.Time) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAchievedDate), v))
	})
}

// AchievedDateIsNil applies the IsNil predicate on the "achieved_date" field.
func AchievedDateIsNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAchievedDate)))
	})
}

// AchievedDateNotNil applies the NotNil predicate on the "achieved_date" field.
func AchievedDateNotNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAchievedDate)))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int) predicate.CandidateCertificate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v int) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v int) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v int) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v int) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// HasAttachmentEdges applies the HasEdge predicate on the "attachment_edges" edge.
func HasAttachmentEdges() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttachmentEdgesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentEdgesTable, AttachmentEdgesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttachmentEdgesWith applies the HasEdge predicate on the "attachment_edges" edge with a given conditions (other predicates).
func HasAttachmentEdgesWith(preds ...predicate.Attachment) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttachmentEdgesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentEdgesTable, AttachmentEdgesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCandidateEdge applies the HasEdge predicate on the "candidate_edge" edge.
func HasCandidateEdge() predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateEdgeTable, CandidateEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateEdgeWith applies the HasEdge predicate on the "candidate_edge" edge with a given conditions (other predicates).
func HasCandidateEdgeWith(preds ...predicate.Candidate) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateEdgeTable, CandidateEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CandidateCertificate) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CandidateCertificate) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
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
func Not(p predicate.CandidateCertificate) predicate.CandidateCertificate {
	return predicate.CandidateCertificate(func(s *sql.Selector) {
		p(s.Not())
	})
}
