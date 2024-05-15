// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"trec/ent/attachment"
	"trec/ent/audittrail"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidateinterviewer"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/hiringjob"
	"trec/ent/team"
	"trec/ent/teammanager"
	"trec/ent/user"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     uuid.UUID `json:"id,omitempty"`     // node id.
	Type   string    `json:"type,omitempty"`   // node type.
	Fields []*Field  `json:"fields,omitempty"` // node fields.
	Edges  []*Edge   `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string      `json:"type,omitempty"` // edge type.
	Name string      `json:"name,omitempty"` // edge name.
	IDs  []uuid.UUID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (a *Attachment) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Attachment",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(a.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.DocumentID); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "uuid.UUID",
		Name:  "document_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.DocumentName); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "document_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.RelationType); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "attachment.RelationType",
		Name:  "relation_type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.RelationID); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "uuid.UUID",
		Name:  "relation_id",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "CandidateJob",
		Name: "candidate_job",
	}
	err = a.QueryCandidateJob().
		Select(candidatejob.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "CandidateJobFeedback",
		Name: "candidate_job_feedback",
	}
	err = a.QueryCandidateJobFeedback().
		Select(candidatejobfeedback.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "CandidateInterview",
		Name: "candidate_interview",
	}
	err = a.QueryCandidateInterview().
		Select(candidateinterview.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (at *AuditTrail) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     at.ID,
		Type:   "AuditTrail",
		Fields: make([]*Field, 9),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(at.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(at.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(at.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(at.CreatedBy); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "uuid.UUID",
		Name:  "created_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(at.RecordId); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "uuid.UUID",
		Name:  "recordId",
		Value: string(buf),
	}
	if buf, err = json.Marshal(at.Module); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "audittrail.Module",
		Name:  "module",
		Value: string(buf),
	}
	if buf, err = json.Marshal(at.ActionType); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "audittrail.ActionType",
		Name:  "actionType",
		Value: string(buf),
	}
	if buf, err = json.Marshal(at.Note); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "note",
		Value: string(buf),
	}
	if buf, err = json.Marshal(at.RecordChanges); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "string",
		Name:  "record_changes",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "user_edge",
	}
	err = at.QueryUserEdge().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Candidate) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.ID,
		Type:   "Candidate",
		Fields: make([]*Field, 9),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(c.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Name); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Email); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Phone); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "phone",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Dob); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "time.Time",
		Name:  "dob",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.IsBlacklist); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "bool",
		Name:  "is_blacklist",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.LastApplyDate); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "time.Time",
		Name:  "last_apply_date",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "CandidateJob",
		Name: "candidate_job_edges",
	}
	err = c.QueryCandidateJobEdges().
		Select(candidatejob.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ci *CandidateInterview) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ci.ID,
		Type:   "CandidateInterview",
		Fields: make([]*Field, 10),
		Edges:  make([]*Edge, 4),
	}
	var buf []byte
	if buf, err = json.Marshal(ci.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.Title); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "title",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.CandidateJobStatus); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "candidateinterview.CandidateJobStatus",
		Name:  "candidate_job_status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.CandidateJobID); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "uuid.UUID",
		Name:  "candidate_job_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.InterviewDate); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "time.Time",
		Name:  "interview_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.StartFrom); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "time.Time",
		Name:  "start_from",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.EndAt); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "time.Time",
		Name:  "end_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.Description); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "CandidateJob",
		Name: "candidate_job_edge",
	}
	err = ci.QueryCandidateJobEdge().
		Select(candidatejob.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Attachment",
		Name: "attachment_edges",
	}
	err = ci.QueryAttachmentEdges().
		Select(attachment.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "User",
		Name: "interviewer_edges",
	}
	err = ci.QueryInterviewerEdges().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "CandidateInterviewer",
		Name: "user_interviewers",
	}
	err = ci.QueryUserInterviewers().
		Select(candidateinterviewer.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ci *CandidateInterviewer) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ci.ID,
		Type:   "CandidateInterviewer",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(ci.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.CandidateInterviewID); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "uuid.UUID",
		Name:  "candidate_interview_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ci.UserID); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "uuid.UUID",
		Name:  "user_id",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "user_edge",
	}
	err = ci.QueryUserEdge().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "CandidateInterview",
		Name: "interview_edge",
	}
	err = ci.QueryInterviewEdge().
		Select(candidateinterview.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (cj *CandidateJob) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     cj.ID,
		Type:   "CandidateJob",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 5),
	}
	var buf []byte
	if buf, err = json.Marshal(cj.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cj.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cj.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cj.HiringJobID); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "uuid.UUID",
		Name:  "hiring_job_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cj.CandidateID); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "uuid.UUID",
		Name:  "candidate_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cj.Status); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "candidatejob.Status",
		Name:  "status",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Attachment",
		Name: "attachment_edges",
	}
	err = cj.QueryAttachmentEdges().
		Select(attachment.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "HiringJob",
		Name: "hiring_job_edge",
	}
	err = cj.QueryHiringJobEdge().
		Select(hiringjob.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "CandidateJobFeedback",
		Name: "candidate_job_feedback",
	}
	err = cj.QueryCandidateJobFeedback().
		Select(candidatejobfeedback.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "Candidate",
		Name: "candidate_edge",
	}
	err = cj.QueryCandidateEdge().
		Select(candidate.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "CandidateInterview",
		Name: "candidate_job_interview",
	}
	err = cj.QueryCandidateJobInterview().
		Select(candidateinterview.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (cjf *CandidateJobFeedback) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     cjf.ID,
		Type:   "CandidateJobFeedback",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(cjf.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cjf.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cjf.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cjf.CandidateJobID); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "uuid.UUID",
		Name:  "candidate_job_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cjf.CreatedBy); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "uuid.UUID",
		Name:  "created_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(cjf.Feedback); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "feedback",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "created_by_edge",
	}
	err = cjf.QueryCreatedByEdge().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "CandidateJob",
		Name: "candidate_job_edge",
	}
	err = cjf.QueryCandidateJobEdge().
		Select(candidatejob.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Attachment",
		Name: "attachment_edges",
	}
	err = cjf.QueryAttachmentEdges().
		Select(attachment.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (hj *HiringJob) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     hj.ID,
		Type:   "HiringJob",
		Fields: make([]*Field, 16),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(hj.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.Slug); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "slug",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.Name); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.Description); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.Amount); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "int",
		Name:  "amount",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.Status); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "hiringjob.Status",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.CreatedBy); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "uuid.UUID",
		Name:  "created_by",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.TeamID); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "uuid.UUID",
		Name:  "team_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.Location); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "hiringjob.Location",
		Name:  "location",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.SalaryType); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "hiringjob.SalaryType",
		Name:  "salary_type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.SalaryFrom); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "int",
		Name:  "salary_from",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.SalaryTo); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "int",
		Name:  "salary_to",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.Currency); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "hiringjob.Currency",
		Name:  "currency",
		Value: string(buf),
	}
	if buf, err = json.Marshal(hj.LastApplyDate); err != nil {
		return nil, err
	}
	node.Fields[15] = &Field{
		Type:  "time.Time",
		Name:  "last_apply_date",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "owner_edge",
	}
	err = hj.QueryOwnerEdge().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Team",
		Name: "team_edge",
	}
	err = hj.QueryTeamEdge().
		Select(team.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "CandidateJob",
		Name: "candidate_job_edges",
	}
	err = hj.QueryCandidateJobEdges().
		Select(candidatejob.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (t *Team) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     t.ID,
		Type:   "Team",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(t.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Slug); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "slug",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Name); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "user_edges",
	}
	err = t.QueryUserEdges().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "HiringJob",
		Name: "team_job_edges",
	}
	err = t.QueryTeamJobEdges().
		Select(hiringjob.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "TeamManager",
		Name: "user_teams",
	}
	err = t.QueryUserTeams().
		Select(teammanager.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (tm *TeamManager) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     tm.ID,
		Type:   "TeamManager",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(tm.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tm.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tm.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tm.TeamID); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "uuid.UUID",
		Name:  "team_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(tm.UserID); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "uuid.UUID",
		Name:  "user_id",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "user_edge",
	}
	err = tm.QueryUserEdge().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Team",
		Name: "team_edge",
	}
	err = tm.QueryTeamEdge().
		Select(team.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 7),
	}
	var buf []byte
	if buf, err = json.Marshal(u.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.DeletedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "deleted_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Name); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.WorkEmail); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "work_email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Status); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "user.Status",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Oid); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "oid",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "AuditTrail",
		Name: "audit_edge",
	}
	err = u.QueryAuditEdge().
		Select(audittrail.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "HiringJob",
		Name: "hiring_owner",
	}
	err = u.QueryHiringOwner().
		Select(hiringjob.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Team",
		Name: "team_edges",
	}
	err = u.QueryTeamEdges().
		Select(team.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "CandidateJobFeedback",
		Name: "candidate_job_feedback",
	}
	err = u.QueryCandidateJobFeedback().
		Select(candidatejobfeedback.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "CandidateInterview",
		Name: "interview_edges",
	}
	err = u.QueryInterviewEdges().
		Select(candidateinterview.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[5] = &Edge{
		Type: "TeamManager",
		Name: "team_users",
	}
	err = u.QueryTeamUsers().
		Select(teammanager.FieldID).
		Scan(ctx, &node.Edges[5].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[6] = &Edge{
		Type: "CandidateInterviewer",
		Name: "interview_users",
	}
	err = u.QueryInterviewUsers().
		Select(candidateinterviewer.FieldID).
		Scan(ctx, &node.Edges[6].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id uuid.UUID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, uuid.UUID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, uuid.UUID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, uuid.UUID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id uuid.UUID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id uuid.UUID) (Noder, error) {
	switch table {
	case attachment.Table:
		query := c.Attachment.Query().
			Where(attachment.ID(id))
		query, err := query.CollectFields(ctx, "Attachment")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case audittrail.Table:
		query := c.AuditTrail.Query().
			Where(audittrail.ID(id))
		query, err := query.CollectFields(ctx, "AuditTrail")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case candidate.Table:
		query := c.Candidate.Query().
			Where(candidate.ID(id))
		query, err := query.CollectFields(ctx, "Candidate")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case candidateinterview.Table:
		query := c.CandidateInterview.Query().
			Where(candidateinterview.ID(id))
		query, err := query.CollectFields(ctx, "CandidateInterview")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case candidateinterviewer.Table:
		query := c.CandidateInterviewer.Query().
			Where(candidateinterviewer.ID(id))
		query, err := query.CollectFields(ctx, "CandidateInterviewer")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case candidatejob.Table:
		query := c.CandidateJob.Query().
			Where(candidatejob.ID(id))
		query, err := query.CollectFields(ctx, "CandidateJob")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case candidatejobfeedback.Table:
		query := c.CandidateJobFeedback.Query().
			Where(candidatejobfeedback.ID(id))
		query, err := query.CollectFields(ctx, "CandidateJobFeedback")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case hiringjob.Table:
		query := c.HiringJob.Query().
			Where(hiringjob.ID(id))
		query, err := query.CollectFields(ctx, "HiringJob")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case team.Table:
		query := c.Team.Query().
			Where(team.ID(id))
		query, err := query.CollectFields(ctx, "Team")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case teammanager.Table:
		query := c.TeamManager.Query().
			Where(teammanager.ID(id))
		query, err := query.CollectFields(ctx, "TeamManager")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]uuid.UUID)
	id2idx := make(map[uuid.UUID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []uuid.UUID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[uuid.UUID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case attachment.Table:
		query := c.Attachment.Query().
			Where(attachment.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Attachment")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case audittrail.Table:
		query := c.AuditTrail.Query().
			Where(audittrail.IDIn(ids...))
		query, err := query.CollectFields(ctx, "AuditTrail")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case candidate.Table:
		query := c.Candidate.Query().
			Where(candidate.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Candidate")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case candidateinterview.Table:
		query := c.CandidateInterview.Query().
			Where(candidateinterview.IDIn(ids...))
		query, err := query.CollectFields(ctx, "CandidateInterview")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case candidateinterviewer.Table:
		query := c.CandidateInterviewer.Query().
			Where(candidateinterviewer.IDIn(ids...))
		query, err := query.CollectFields(ctx, "CandidateInterviewer")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case candidatejob.Table:
		query := c.CandidateJob.Query().
			Where(candidatejob.IDIn(ids...))
		query, err := query.CollectFields(ctx, "CandidateJob")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case candidatejobfeedback.Table:
		query := c.CandidateJobFeedback.Query().
			Where(candidatejobfeedback.IDIn(ids...))
		query, err := query.CollectFields(ctx, "CandidateJobFeedback")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case hiringjob.Table:
		query := c.HiringJob.Query().
			Where(hiringjob.IDIn(ids...))
		query, err := query.CollectFields(ctx, "HiringJob")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case team.Table:
		query := c.Team.Query().
			Where(team.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Team")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case teammanager.Table:
		query := c.TeamManager.Query().
			Where(teammanager.IDIn(ids...))
		query, err := query.CollectFields(ctx, "TeamManager")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
