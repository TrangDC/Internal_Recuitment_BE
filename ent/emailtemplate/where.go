// Code generated by ent, DO NOT EDIT.

package emailtemplate

import (
	"time"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubject), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Signature applies equality check predicate on the "signature" field. It's identical to SignatureEQ.
func Signature(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSignature), v))
	})
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// EventEQ applies the EQ predicate on the "event" field.
func EventEQ(v Event) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEvent), v))
	})
}

// EventNEQ applies the NEQ predicate on the "event" field.
func EventNEQ(v Event) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEvent), v))
	})
}

// EventIn applies the In predicate on the "event" field.
func EventIn(vs ...Event) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEvent), v...))
	})
}

// EventNotIn applies the NotIn predicate on the "event" field.
func EventNotIn(vs ...Event) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEvent), v...))
	})
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubject), v))
	})
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubject), v))
	})
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSubject), v...))
	})
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSubject), v...))
	})
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubject), v))
	})
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubject), v))
	})
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubject), v))
	})
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubject), v))
	})
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSubject), v))
	})
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSubject), v))
	})
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSubject), v))
	})
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSubject), v))
	})
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSubject), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// SignatureEQ applies the EQ predicate on the "signature" field.
func SignatureEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSignature), v))
	})
}

// SignatureNEQ applies the NEQ predicate on the "signature" field.
func SignatureNEQ(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSignature), v))
	})
}

// SignatureIn applies the In predicate on the "signature" field.
func SignatureIn(vs ...string) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSignature), v...))
	})
}

// SignatureNotIn applies the NotIn predicate on the "signature" field.
func SignatureNotIn(vs ...string) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSignature), v...))
	})
}

// SignatureGT applies the GT predicate on the "signature" field.
func SignatureGT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSignature), v))
	})
}

// SignatureGTE applies the GTE predicate on the "signature" field.
func SignatureGTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSignature), v))
	})
}

// SignatureLT applies the LT predicate on the "signature" field.
func SignatureLT(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSignature), v))
	})
}

// SignatureLTE applies the LTE predicate on the "signature" field.
func SignatureLTE(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSignature), v))
	})
}

// SignatureContains applies the Contains predicate on the "signature" field.
func SignatureContains(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSignature), v))
	})
}

// SignatureHasPrefix applies the HasPrefix predicate on the "signature" field.
func SignatureHasPrefix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSignature), v))
	})
}

// SignatureHasSuffix applies the HasSuffix predicate on the "signature" field.
func SignatureHasSuffix(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSignature), v))
	})
}

// SignatureIsNil applies the IsNil predicate on the "signature" field.
func SignatureIsNil() predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSignature)))
	})
}

// SignatureNotNil applies the NotNil predicate on the "signature" field.
func SignatureNotNil() predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSignature)))
	})
}

// SignatureEqualFold applies the EqualFold predicate on the "signature" field.
func SignatureEqualFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSignature), v))
	})
}

// SignatureContainsFold applies the ContainsFold predicate on the "signature" field.
func SignatureContainsFold(v string) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSignature), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventID), v))
	})
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v uuid.UUID) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventID), v))
	})
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...uuid.UUID) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEventID), v...))
	})
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...uuid.UUID) predicate.EmailTemplate {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEventID), v...))
	})
}

// HasRoleEdges applies the HasEdge predicate on the "role_edges" edge.
func HasRoleEdges() predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoleEdgesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RoleEdgesTable, RoleEdgesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleEdgesWith applies the HasEdge predicate on the "role_edges" edge with a given conditions (other predicates).
func HasRoleEdgesWith(preds ...predicate.Role) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoleEdgesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RoleEdgesTable, RoleEdgesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEventEdge applies the HasEdge predicate on the "event_edge" edge.
func HasEventEdge() predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EventEdgeTable, EventEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventEdgeWith applies the HasEdge predicate on the "event_edge" edge with a given conditions (other predicates).
func HasEventEdgeWith(preds ...predicate.EmailEvent) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EventEdgeTable, EventEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoleEmailTemplates applies the HasEdge predicate on the "role_email_templates" edge.
func HasRoleEmailTemplates() predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoleEmailTemplatesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RoleEmailTemplatesTable, RoleEmailTemplatesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleEmailTemplatesWith applies the HasEdge predicate on the "role_email_templates" edge with a given conditions (other predicates).
func HasRoleEmailTemplatesWith(preds ...predicate.EmailRoleAttribute) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoleEmailTemplatesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RoleEmailTemplatesTable, RoleEmailTemplatesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmailTemplate) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmailTemplate) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
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
func Not(p predicate.EmailTemplate) predicate.EmailTemplate {
	return predicate.EmailTemplate(func(s *sql.Selector) {
		p(s.Not())
	})
}
