// Code generated by ent, DO NOT EDIT.

package attachment

import (
	"time"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DocumentID applies equality check predicate on the "document_id" field. It's identical to DocumentIDEQ.
func DocumentID(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDocumentID), v))
	})
}

// DocumentName applies equality check predicate on the "document_name" field. It's identical to DocumentNameEQ.
func DocumentName(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDocumentName), v))
	})
}

// RelationID applies equality check predicate on the "relation_id" field. It's identical to RelationIDEQ.
func RelationID(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRelationID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// DocumentIDEQ applies the EQ predicate on the "document_id" field.
func DocumentIDEQ(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDocumentID), v))
	})
}

// DocumentIDNEQ applies the NEQ predicate on the "document_id" field.
func DocumentIDNEQ(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDocumentID), v))
	})
}

// DocumentIDIn applies the In predicate on the "document_id" field.
func DocumentIDIn(vs ...uuid.UUID) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDocumentID), v...))
	})
}

// DocumentIDNotIn applies the NotIn predicate on the "document_id" field.
func DocumentIDNotIn(vs ...uuid.UUID) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDocumentID), v...))
	})
}

// DocumentIDGT applies the GT predicate on the "document_id" field.
func DocumentIDGT(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDocumentID), v))
	})
}

// DocumentIDGTE applies the GTE predicate on the "document_id" field.
func DocumentIDGTE(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDocumentID), v))
	})
}

// DocumentIDLT applies the LT predicate on the "document_id" field.
func DocumentIDLT(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDocumentID), v))
	})
}

// DocumentIDLTE applies the LTE predicate on the "document_id" field.
func DocumentIDLTE(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDocumentID), v))
	})
}

// DocumentNameEQ applies the EQ predicate on the "document_name" field.
func DocumentNameEQ(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDocumentName), v))
	})
}

// DocumentNameNEQ applies the NEQ predicate on the "document_name" field.
func DocumentNameNEQ(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDocumentName), v))
	})
}

// DocumentNameIn applies the In predicate on the "document_name" field.
func DocumentNameIn(vs ...string) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDocumentName), v...))
	})
}

// DocumentNameNotIn applies the NotIn predicate on the "document_name" field.
func DocumentNameNotIn(vs ...string) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDocumentName), v...))
	})
}

// DocumentNameGT applies the GT predicate on the "document_name" field.
func DocumentNameGT(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDocumentName), v))
	})
}

// DocumentNameGTE applies the GTE predicate on the "document_name" field.
func DocumentNameGTE(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDocumentName), v))
	})
}

// DocumentNameLT applies the LT predicate on the "document_name" field.
func DocumentNameLT(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDocumentName), v))
	})
}

// DocumentNameLTE applies the LTE predicate on the "document_name" field.
func DocumentNameLTE(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDocumentName), v))
	})
}

// DocumentNameContains applies the Contains predicate on the "document_name" field.
func DocumentNameContains(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDocumentName), v))
	})
}

// DocumentNameHasPrefix applies the HasPrefix predicate on the "document_name" field.
func DocumentNameHasPrefix(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDocumentName), v))
	})
}

// DocumentNameHasSuffix applies the HasSuffix predicate on the "document_name" field.
func DocumentNameHasSuffix(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDocumentName), v))
	})
}

// DocumentNameEqualFold applies the EqualFold predicate on the "document_name" field.
func DocumentNameEqualFold(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDocumentName), v))
	})
}

// DocumentNameContainsFold applies the ContainsFold predicate on the "document_name" field.
func DocumentNameContainsFold(v string) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDocumentName), v))
	})
}

// RelationTypeEQ applies the EQ predicate on the "relation_type" field.
func RelationTypeEQ(v RelationType) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRelationType), v))
	})
}

// RelationTypeNEQ applies the NEQ predicate on the "relation_type" field.
func RelationTypeNEQ(v RelationType) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRelationType), v))
	})
}

// RelationTypeIn applies the In predicate on the "relation_type" field.
func RelationTypeIn(vs ...RelationType) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRelationType), v...))
	})
}

// RelationTypeNotIn applies the NotIn predicate on the "relation_type" field.
func RelationTypeNotIn(vs ...RelationType) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRelationType), v...))
	})
}

// RelationIDEQ applies the EQ predicate on the "relation_id" field.
func RelationIDEQ(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRelationID), v))
	})
}

// RelationIDNEQ applies the NEQ predicate on the "relation_id" field.
func RelationIDNEQ(v uuid.UUID) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRelationID), v))
	})
}

// RelationIDIn applies the In predicate on the "relation_id" field.
func RelationIDIn(vs ...uuid.UUID) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRelationID), v...))
	})
}

// RelationIDNotIn applies the NotIn predicate on the "relation_id" field.
func RelationIDNotIn(vs ...uuid.UUID) predicate.Attachment {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRelationID), v...))
	})
}

// RelationIDIsNil applies the IsNil predicate on the "relation_id" field.
func RelationIDIsNil() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRelationID)))
	})
}

// RelationIDNotNil applies the NotNil predicate on the "relation_id" field.
func RelationIDNotNil() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRelationID)))
	})
}

// HasCandidateJobEdge applies the HasEdge predicate on the "candidate_job_edge" edge.
func HasCandidateJobEdge() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateJobEdgeTable, CandidateJobEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateJobEdgeWith applies the HasEdge predicate on the "candidate_job_edge" edge with a given conditions (other predicates).
func HasCandidateJobEdgeWith(preds ...predicate.CandidateJob) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
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

// HasCandidateJobFeedbackEdge applies the HasEdge predicate on the "candidate_job_feedback_edge" edge.
func HasCandidateJobFeedbackEdge() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobFeedbackEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateJobFeedbackEdgeTable, CandidateJobFeedbackEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateJobFeedbackEdgeWith applies the HasEdge predicate on the "candidate_job_feedback_edge" edge with a given conditions (other predicates).
func HasCandidateJobFeedbackEdgeWith(preds ...predicate.CandidateJobFeedback) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateJobFeedbackEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateJobFeedbackEdgeTable, CandidateJobFeedbackEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCandidateInterviewEdge applies the HasEdge predicate on the "candidate_interview_edge" edge.
func HasCandidateInterviewEdge() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateInterviewEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateInterviewEdgeTable, CandidateInterviewEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateInterviewEdgeWith applies the HasEdge predicate on the "candidate_interview_edge" edge with a given conditions (other predicates).
func HasCandidateInterviewEdgeWith(preds ...predicate.CandidateInterview) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateInterviewEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateInterviewEdgeTable, CandidateInterviewEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCandidateEdge applies the HasEdge predicate on the "candidate_edge" edge.
func HasCandidateEdge() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateEdgeTable, CandidateEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateEdgeWith applies the HasEdge predicate on the "candidate_edge" edge with a given conditions (other predicates).
func HasCandidateEdgeWith(preds ...predicate.Candidate) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
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

// HasCandidateEducateEdge applies the HasEdge predicate on the "candidate_educate_edge" edge.
func HasCandidateEducateEdge() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateEducateEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateEducateEdgeTable, CandidateEducateEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateEducateEdgeWith applies the HasEdge predicate on the "candidate_educate_edge" edge with a given conditions (other predicates).
func HasCandidateEducateEdgeWith(preds ...predicate.CandidateEducate) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateEducateEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateEducateEdgeTable, CandidateEducateEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCandidateAwardEdge applies the HasEdge predicate on the "candidate_award_edge" edge.
func HasCandidateAwardEdge() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateAwardEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateAwardEdgeTable, CandidateAwardEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateAwardEdgeWith applies the HasEdge predicate on the "candidate_award_edge" edge with a given conditions (other predicates).
func HasCandidateAwardEdgeWith(preds ...predicate.CandidateAward) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateAwardEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateAwardEdgeTable, CandidateAwardEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCandidateCertificateEdge applies the HasEdge predicate on the "candidate_certificate_edge" edge.
func HasCandidateCertificateEdge() predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateCertificateEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateCertificateEdgeTable, CandidateCertificateEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCandidateCertificateEdgeWith applies the HasEdge predicate on the "candidate_certificate_edge" edge with a given conditions (other predicates).
func HasCandidateCertificateEdgeWith(preds ...predicate.CandidateCertificate) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CandidateCertificateEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CandidateCertificateEdgeTable, CandidateCertificateEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
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
func Not(p predicate.Attachment) predicate.Attachment {
	return predicate.Attachment(func(s *sql.Selector) {
		p(s.Not())
	})
}
