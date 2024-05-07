// Code generated by ent, DO NOT EDIT.

package candidateinterview

import (
	"time"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// CandidateJobID applies equality check predicate on the "candidate_job_id" field. It's identical to CandidateJobIDEQ.
func CandidateJobID(v uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCandidateJobID), v))
	})
}

// InterviewDate applies equality check predicate on the "interview_date" field. It's identical to InterviewDateEQ.
func InterviewDate(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterviewDate), v))
	})
}

// StartFrom applies equality check predicate on the "start_from" field. It's identical to StartFromEQ.
func StartFrom(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartFrom), v))
	})
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// CandidateJobStatusEQ applies the EQ predicate on the "candidate_job_status" field.
func CandidateJobStatusEQ(v CandidateJobStatus) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCandidateJobStatus), v))
	})
}

// CandidateJobStatusNEQ applies the NEQ predicate on the "candidate_job_status" field.
func CandidateJobStatusNEQ(v CandidateJobStatus) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCandidateJobStatus), v))
	})
}

// CandidateJobStatusIn applies the In predicate on the "candidate_job_status" field.
func CandidateJobStatusIn(vs ...CandidateJobStatus) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCandidateJobStatus), v...))
	})
}

// CandidateJobStatusNotIn applies the NotIn predicate on the "candidate_job_status" field.
func CandidateJobStatusNotIn(vs ...CandidateJobStatus) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCandidateJobStatus), v...))
	})
}

// CandidateJobIDEQ applies the EQ predicate on the "candidate_job_id" field.
func CandidateJobIDEQ(v uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCandidateJobID), v))
	})
}

// CandidateJobIDNEQ applies the NEQ predicate on the "candidate_job_id" field.
func CandidateJobIDNEQ(v uuid.UUID) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCandidateJobID), v))
	})
}

// CandidateJobIDIn applies the In predicate on the "candidate_job_id" field.
func CandidateJobIDIn(vs ...uuid.UUID) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCandidateJobID), v...))
	})
}

// CandidateJobIDNotIn applies the NotIn predicate on the "candidate_job_id" field.
func CandidateJobIDNotIn(vs ...uuid.UUID) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCandidateJobID), v...))
	})
}

// CandidateJobIDIsNil applies the IsNil predicate on the "candidate_job_id" field.
func CandidateJobIDIsNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCandidateJobID)))
	})
}

// CandidateJobIDNotNil applies the NotNil predicate on the "candidate_job_id" field.
func CandidateJobIDNotNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCandidateJobID)))
	})
}

// InterviewDateEQ applies the EQ predicate on the "interview_date" field.
func InterviewDateEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterviewDate), v))
	})
}

// InterviewDateNEQ applies the NEQ predicate on the "interview_date" field.
func InterviewDateNEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInterviewDate), v))
	})
}

// InterviewDateIn applies the In predicate on the "interview_date" field.
func InterviewDateIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInterviewDate), v...))
	})
}

// InterviewDateNotIn applies the NotIn predicate on the "interview_date" field.
func InterviewDateNotIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInterviewDate), v...))
	})
}

// InterviewDateGT applies the GT predicate on the "interview_date" field.
func InterviewDateGT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInterviewDate), v))
	})
}

// InterviewDateGTE applies the GTE predicate on the "interview_date" field.
func InterviewDateGTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInterviewDate), v))
	})
}

// InterviewDateLT applies the LT predicate on the "interview_date" field.
func InterviewDateLT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInterviewDate), v))
	})
}

// InterviewDateLTE applies the LTE predicate on the "interview_date" field.
func InterviewDateLTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInterviewDate), v))
	})
}

// InterviewDateIsNil applies the IsNil predicate on the "interview_date" field.
func InterviewDateIsNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInterviewDate)))
	})
}

// InterviewDateNotNil applies the NotNil predicate on the "interview_date" field.
func InterviewDateNotNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInterviewDate)))
	})
}

// StartFromEQ applies the EQ predicate on the "start_from" field.
func StartFromEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartFrom), v))
	})
}

// StartFromNEQ applies the NEQ predicate on the "start_from" field.
func StartFromNEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartFrom), v))
	})
}

// StartFromIn applies the In predicate on the "start_from" field.
func StartFromIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartFrom), v...))
	})
}

// StartFromNotIn applies the NotIn predicate on the "start_from" field.
func StartFromNotIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartFrom), v...))
	})
}

// StartFromGT applies the GT predicate on the "start_from" field.
func StartFromGT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartFrom), v))
	})
}

// StartFromGTE applies the GTE predicate on the "start_from" field.
func StartFromGTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartFrom), v))
	})
}

// StartFromLT applies the LT predicate on the "start_from" field.
func StartFromLT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartFrom), v))
	})
}

// StartFromLTE applies the LTE predicate on the "start_from" field.
func StartFromLTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartFrom), v))
	})
}

// StartFromIsNil applies the IsNil predicate on the "start_from" field.
func StartFromIsNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartFrom)))
	})
}

// StartFromNotNil applies the NotNil predicate on the "start_from" field.
func StartFromNotNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartFrom)))
	})
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndAt), v))
	})
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEndAt), v...))
	})
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEndAt), v...))
	})
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndAt), v))
	})
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndAt), v))
	})
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndAt), v))
	})
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndAt), v))
	})
}

// EndAtIsNil applies the IsNil predicate on the "end_at" field.
func EndAtIsNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndAt)))
	})
}

// EndAtNotNil applies the NotNil predicate on the "end_at" field.
func EndAtNotNil() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndAt)))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.CandidateInterview {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// HasCandidateJobEdge applies the HasEdge predicate on the "candidate_job_edge" edge.
func HasCandidateJobEdge() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateJobEdgeTable, CandidateJobEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateJobEdgeWith applies the HasEdge predicate on the "candidate_job_edge" edge with a given conditions (other predicates).
func HasCandidateJobEdgeWith(preds ...predicate.CandidateJob) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateJobEdgeTable, CandidateJobEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttachmentEdges applies the HasEdge predicate on the "attachment_edges" edge.
func HasAttachmentEdges() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttachmentEdgesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentEdgesTable, AttachmentEdgesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttachmentEdgesWith applies the HasEdge predicate on the "attachment_edges" edge with a given conditions (other predicates).
func HasAttachmentEdgesWith(preds ...predicate.Attachment) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
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

// HasInterviewerEdges applies the HasEdge predicate on the "interviewer_edges" edge.
func HasInterviewerEdges() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InterviewerEdgesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, InterviewerEdgesTable, InterviewerEdgesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInterviewerEdgesWith applies the HasEdge predicate on the "interviewer_edges" edge with a given conditions (other predicates).
func HasInterviewerEdgesWith(preds ...predicate.User) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InterviewerEdgesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, InterviewerEdgesTable, InterviewerEdgesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserInterviewers applies the HasEdge predicate on the "user_interviewers" edge.
func HasUserInterviewers() predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInterviewersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserInterviewersTable, UserInterviewersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserInterviewersWith applies the HasEdge predicate on the "user_interviewers" edge with a given conditions (other predicates).
func HasUserInterviewersWith(preds ...predicate.CandidateInterviewer) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInterviewersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserInterviewersTable, UserInterviewersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CandidateInterview) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CandidateInterview) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
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
func Not(p predicate.CandidateInterview) predicate.CandidateInterview {
	return predicate.CandidateInterview(func(s *sql.Selector) {
		p(s.Not())
	})
}
