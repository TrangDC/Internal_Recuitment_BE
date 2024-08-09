// Code generated by ent, DO NOT EDIT.

package candidatejob

import (
	"time"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// HiringJobID applies equality check predicate on the "hiring_job_id" field. It's identical to HiringJobIDEQ.
func HiringJobID(v uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHiringJobID), v))
	})
}

// CandidateID applies equality check predicate on the "candidate_id" field. It's identical to CandidateIDEQ.
func CandidateID(v uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCandidateID), v))
	})
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// OnboardDate applies equality check predicate on the "onboard_date" field. It's identical to OnboardDateEQ.
func OnboardDate(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnboardDate), v))
	})
}

// OfferExpirationDate applies equality check predicate on the "offer_expiration_date" field. It's identical to OfferExpirationDateEQ.
func OfferExpirationDate(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfferExpirationDate), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// HiringJobIDEQ applies the EQ predicate on the "hiring_job_id" field.
func HiringJobIDEQ(v uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHiringJobID), v))
	})
}

// HiringJobIDNEQ applies the NEQ predicate on the "hiring_job_id" field.
func HiringJobIDNEQ(v uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHiringJobID), v))
	})
}

// HiringJobIDIn applies the In predicate on the "hiring_job_id" field.
func HiringJobIDIn(vs ...uuid.UUID) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHiringJobID), v...))
	})
}

// HiringJobIDNotIn applies the NotIn predicate on the "hiring_job_id" field.
func HiringJobIDNotIn(vs ...uuid.UUID) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHiringJobID), v...))
	})
}

// HiringJobIDIsNil applies the IsNil predicate on the "hiring_job_id" field.
func HiringJobIDIsNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHiringJobID)))
	})
}

// HiringJobIDNotNil applies the NotNil predicate on the "hiring_job_id" field.
func HiringJobIDNotNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHiringJobID)))
	})
}

// CandidateIDEQ applies the EQ predicate on the "candidate_id" field.
func CandidateIDEQ(v uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCandidateID), v))
	})
}

// CandidateIDNEQ applies the NEQ predicate on the "candidate_id" field.
func CandidateIDNEQ(v uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCandidateID), v))
	})
}

// CandidateIDIn applies the In predicate on the "candidate_id" field.
func CandidateIDIn(vs ...uuid.UUID) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCandidateID), v...))
	})
}

// CandidateIDNotIn applies the NotIn predicate on the "candidate_id" field.
func CandidateIDNotIn(vs ...uuid.UUID) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCandidateID), v...))
	})
}

// CandidateIDIsNil applies the IsNil predicate on the "candidate_id" field.
func CandidateIDIsNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCandidateID)))
	})
}

// CandidateIDNotNil applies the NotNil predicate on the "candidate_id" field.
func CandidateIDNotNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCandidateID)))
	})
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uuid.UUID) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uuid.UUID) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uuid.UUID) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedBy)))
	})
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedBy)))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// FailedReasonIsNil applies the IsNil predicate on the "failed_reason" field.
func FailedReasonIsNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFailedReason)))
	})
}

// FailedReasonNotNil applies the NotNil predicate on the "failed_reason" field.
func FailedReasonNotNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFailedReason)))
	})
}

// OnboardDateEQ applies the EQ predicate on the "onboard_date" field.
func OnboardDateEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnboardDate), v))
	})
}

// OnboardDateNEQ applies the NEQ predicate on the "onboard_date" field.
func OnboardDateNEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOnboardDate), v))
	})
}

// OnboardDateIn applies the In predicate on the "onboard_date" field.
func OnboardDateIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOnboardDate), v...))
	})
}

// OnboardDateNotIn applies the NotIn predicate on the "onboard_date" field.
func OnboardDateNotIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOnboardDate), v...))
	})
}

// OnboardDateGT applies the GT predicate on the "onboard_date" field.
func OnboardDateGT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOnboardDate), v))
	})
}

// OnboardDateGTE applies the GTE predicate on the "onboard_date" field.
func OnboardDateGTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOnboardDate), v))
	})
}

// OnboardDateLT applies the LT predicate on the "onboard_date" field.
func OnboardDateLT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOnboardDate), v))
	})
}

// OnboardDateLTE applies the LTE predicate on the "onboard_date" field.
func OnboardDateLTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOnboardDate), v))
	})
}

// OnboardDateIsNil applies the IsNil predicate on the "onboard_date" field.
func OnboardDateIsNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOnboardDate)))
	})
}

// OnboardDateNotNil applies the NotNil predicate on the "onboard_date" field.
func OnboardDateNotNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOnboardDate)))
	})
}

// OfferExpirationDateEQ applies the EQ predicate on the "offer_expiration_date" field.
func OfferExpirationDateEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOfferExpirationDate), v))
	})
}

// OfferExpirationDateNEQ applies the NEQ predicate on the "offer_expiration_date" field.
func OfferExpirationDateNEQ(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOfferExpirationDate), v))
	})
}

// OfferExpirationDateIn applies the In predicate on the "offer_expiration_date" field.
func OfferExpirationDateIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOfferExpirationDate), v...))
	})
}

// OfferExpirationDateNotIn applies the NotIn predicate on the "offer_expiration_date" field.
func OfferExpirationDateNotIn(vs ...time.Time) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOfferExpirationDate), v...))
	})
}

// OfferExpirationDateGT applies the GT predicate on the "offer_expiration_date" field.
func OfferExpirationDateGT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOfferExpirationDate), v))
	})
}

// OfferExpirationDateGTE applies the GTE predicate on the "offer_expiration_date" field.
func OfferExpirationDateGTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOfferExpirationDate), v))
	})
}

// OfferExpirationDateLT applies the LT predicate on the "offer_expiration_date" field.
func OfferExpirationDateLT(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOfferExpirationDate), v))
	})
}

// OfferExpirationDateLTE applies the LTE predicate on the "offer_expiration_date" field.
func OfferExpirationDateLTE(v time.Time) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOfferExpirationDate), v))
	})
}

// OfferExpirationDateIsNil applies the IsNil predicate on the "offer_expiration_date" field.
func OfferExpirationDateIsNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOfferExpirationDate)))
	})
}

// OfferExpirationDateNotNil applies the NotNil predicate on the "offer_expiration_date" field.
func OfferExpirationDateNotNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOfferExpirationDate)))
	})
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v Level) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v Level) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLevel), v))
	})
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...Level) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLevel), v...))
	})
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...Level) predicate.CandidateJob {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLevel), v...))
	})
}

// LevelIsNil applies the IsNil predicate on the "level" field.
func LevelIsNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLevel)))
	})
}

// LevelNotNil applies the NotNil predicate on the "level" field.
func LevelNotNil() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLevel)))
	})
}

// HasAttachmentEdges applies the HasEdge predicate on the "attachment_edges" edge.
func HasAttachmentEdges() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttachmentEdgesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentEdgesTable, AttachmentEdgesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttachmentEdgesWith applies the HasEdge predicate on the "attachment_edges" edge with a given conditions (other predicates).
func HasAttachmentEdgesWith(preds ...predicate.Attachment) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
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

// HasHiringJobEdge applies the HasEdge predicate on the "hiring_job_edge" edge.
func HasHiringJobEdge() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HiringJobEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HiringJobEdgeTable, HiringJobEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHiringJobEdgeWith applies the HasEdge predicate on the "hiring_job_edge" edge with a given conditions (other predicates).
func HasHiringJobEdgeWith(preds ...predicate.HiringJob) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HiringJobEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HiringJobEdgeTable, HiringJobEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCandidateJobFeedback applies the HasEdge predicate on the "candidate_job_feedback" edge.
func HasCandidateJobFeedback() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobFeedbackTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CandidateJobFeedbackTable, CandidateJobFeedbackColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateJobFeedbackWith applies the HasEdge predicate on the "candidate_job_feedback" edge with a given conditions (other predicates).
func HasCandidateJobFeedbackWith(preds ...predicate.CandidateJobFeedback) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobFeedbackInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CandidateJobFeedbackTable, CandidateJobFeedbackColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCandidateEdge applies the HasEdge predicate on the "candidate_edge" edge.
func HasCandidateEdge() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateEdgeTable, CandidateEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateEdgeWith applies the HasEdge predicate on the "candidate_edge" edge with a given conditions (other predicates).
func HasCandidateEdgeWith(preds ...predicate.Candidate) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
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

// HasCandidateJobInterview applies the HasEdge predicate on the "candidate_job_interview" edge.
func HasCandidateJobInterview() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobInterviewTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CandidateJobInterviewTable, CandidateJobInterviewColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateJobInterviewWith applies the HasEdge predicate on the "candidate_job_interview" edge with a given conditions (other predicates).
func HasCandidateJobInterviewWith(preds ...predicate.CandidateInterview) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobInterviewInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CandidateJobInterviewTable, CandidateJobInterviewColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatedByEdge applies the HasEdge predicate on the "created_by_edge" edge.
func HasCreatedByEdge() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatedByEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatedByEdgeTable, CreatedByEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedByEdgeWith applies the HasEdge predicate on the "created_by_edge" edge with a given conditions (other predicates).
func HasCreatedByEdgeWith(preds ...predicate.User) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatedByEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatedByEdgeTable, CreatedByEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCandidateJobStep applies the HasEdge predicate on the "candidate_job_step" edge.
func HasCandidateJobStep() predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobStepTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CandidateJobStepTable, CandidateJobStepColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateJobStepWith applies the HasEdge predicate on the "candidate_job_step" edge with a given conditions (other predicates).
func HasCandidateJobStepWith(preds ...predicate.CandidateJobStep) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobStepInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CandidateJobStepTable, CandidateJobStepColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CandidateJob) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CandidateJob) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
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
func Not(p predicate.CandidateJob) predicate.CandidateJob {
	return predicate.CandidateJob(func(s *sql.Selector) {
		p(s.Not())
	})
}
