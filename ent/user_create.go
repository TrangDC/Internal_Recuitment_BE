// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/audittrail"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidateinterviewer"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/candidatenote"
	"trec/ent/entitypermission"
	"trec/ent/hiringjob"
	"trec/ent/hiringteam"
	"trec/ent/hiringteamapprover"
	"trec/ent/hiringteammanager"
	"trec/ent/recteam"
	"trec/ent/role"
	"trec/ent/user"
	"trec/ent/userrole"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetDeletedAt sets the "deleted_at" field.
func (uc *UserCreate) SetDeletedAt(t time.Time) *UserCreate {
	uc.mutation.SetDeletedAt(t)
	return uc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeletedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetDeletedAt(*t)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetWorkEmail sets the "work_email" field.
func (uc *UserCreate) SetWorkEmail(s string) *UserCreate {
	uc.mutation.SetWorkEmail(s)
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(u user.Status) *UserCreate {
	uc.mutation.SetStatus(u)
	return uc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(u *user.Status) *UserCreate {
	if u != nil {
		uc.SetStatus(*u)
	}
	return uc
}

// SetOid sets the "oid" field.
func (uc *UserCreate) SetOid(s string) *UserCreate {
	uc.mutation.SetOid(s)
	return uc
}

// SetRecTeamID sets the "rec_team_id" field.
func (uc *UserCreate) SetRecTeamID(u uuid.UUID) *UserCreate {
	uc.mutation.SetRecTeamID(u)
	return uc
}

// SetNillableRecTeamID sets the "rec_team_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableRecTeamID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetRecTeamID(*u)
	}
	return uc
}

// SetLocation sets the "location" field.
func (uc *UserCreate) SetLocation(s string) *UserCreate {
	uc.mutation.SetLocation(s)
	return uc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (uc *UserCreate) SetNillableLocation(s *string) *UserCreate {
	if s != nil {
		uc.SetLocation(*s)
	}
	return uc
}

// SetHiringTeamID sets the "hiring_team_id" field.
func (uc *UserCreate) SetHiringTeamID(u uuid.UUID) *UserCreate {
	uc.mutation.SetHiringTeamID(u)
	return uc
}

// SetNillableHiringTeamID sets the "hiring_team_id" field if the given value is not nil.
func (uc *UserCreate) SetNillableHiringTeamID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetHiringTeamID(*u)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// AddAuditEdgeIDs adds the "audit_edge" edge to the AuditTrail entity by IDs.
func (uc *UserCreate) AddAuditEdgeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddAuditEdgeIDs(ids...)
	return uc
}

// AddAuditEdge adds the "audit_edge" edges to the AuditTrail entity.
func (uc *UserCreate) AddAuditEdge(a ...*AuditTrail) *UserCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAuditEdgeIDs(ids...)
}

// AddHiringOwnerIDs adds the "hiring_owner" edge to the HiringJob entity by IDs.
func (uc *UserCreate) AddHiringOwnerIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddHiringOwnerIDs(ids...)
	return uc
}

// AddHiringOwner adds the "hiring_owner" edges to the HiringJob entity.
func (uc *UserCreate) AddHiringOwner(h ...*HiringJob) *UserCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uc.AddHiringOwnerIDs(ids...)
}

// AddCandidateJobFeedbackIDs adds the "candidate_job_feedback" edge to the CandidateJobFeedback entity by IDs.
func (uc *UserCreate) AddCandidateJobFeedbackIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddCandidateJobFeedbackIDs(ids...)
	return uc
}

// AddCandidateJobFeedback adds the "candidate_job_feedback" edges to the CandidateJobFeedback entity.
func (uc *UserCreate) AddCandidateJobFeedback(c ...*CandidateJobFeedback) *UserCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCandidateJobFeedbackIDs(ids...)
}

// AddInterviewEdgeIDs adds the "interview_edges" edge to the CandidateInterview entity by IDs.
func (uc *UserCreate) AddInterviewEdgeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddInterviewEdgeIDs(ids...)
	return uc
}

// AddInterviewEdges adds the "interview_edges" edges to the CandidateInterview entity.
func (uc *UserCreate) AddInterviewEdges(c ...*CandidateInterview) *UserCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddInterviewEdgeIDs(ids...)
}

// AddCandidateJobEdgeIDs adds the "candidate_job_edges" edge to the CandidateJob entity by IDs.
func (uc *UserCreate) AddCandidateJobEdgeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddCandidateJobEdgeIDs(ids...)
	return uc
}

// AddCandidateJobEdges adds the "candidate_job_edges" edges to the CandidateJob entity.
func (uc *UserCreate) AddCandidateJobEdges(c ...*CandidateJob) *UserCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCandidateJobEdgeIDs(ids...)
}

// AddCandidateInterviewEdgeIDs adds the "candidate_interview_edges" edge to the CandidateInterview entity by IDs.
func (uc *UserCreate) AddCandidateInterviewEdgeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddCandidateInterviewEdgeIDs(ids...)
	return uc
}

// AddCandidateInterviewEdges adds the "candidate_interview_edges" edges to the CandidateInterview entity.
func (uc *UserCreate) AddCandidateInterviewEdges(c ...*CandidateInterview) *UserCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCandidateInterviewEdgeIDs(ids...)
}

// AddCandidateReferenceEdgeIDs adds the "candidate_reference_edges" edge to the Candidate entity by IDs.
func (uc *UserCreate) AddCandidateReferenceEdgeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddCandidateReferenceEdgeIDs(ids...)
	return uc
}

// AddCandidateReferenceEdges adds the "candidate_reference_edges" edges to the Candidate entity.
func (uc *UserCreate) AddCandidateReferenceEdges(c ...*Candidate) *UserCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCandidateReferenceEdgeIDs(ids...)
}

// AddUserPermissionEdgeIDs adds the "user_permission_edges" edge to the EntityPermission entity by IDs.
func (uc *UserCreate) AddUserPermissionEdgeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddUserPermissionEdgeIDs(ids...)
	return uc
}

// AddUserPermissionEdges adds the "user_permission_edges" edges to the EntityPermission entity.
func (uc *UserCreate) AddUserPermissionEdges(e ...*EntityPermission) *UserCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uc.AddUserPermissionEdgeIDs(ids...)
}

// AddRoleEdgeIDs adds the "role_edges" edge to the Role entity by IDs.
func (uc *UserCreate) AddRoleEdgeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddRoleEdgeIDs(ids...)
	return uc
}

// AddRoleEdges adds the "role_edges" edges to the Role entity.
func (uc *UserCreate) AddRoleEdges(r ...*Role) *UserCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRoleEdgeIDs(ids...)
}

// AddHiringTeamEdgeIDs adds the "hiring_team_edges" edge to the HiringTeam entity by IDs.
func (uc *UserCreate) AddHiringTeamEdgeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddHiringTeamEdgeIDs(ids...)
	return uc
}

// AddHiringTeamEdges adds the "hiring_team_edges" edges to the HiringTeam entity.
func (uc *UserCreate) AddHiringTeamEdges(h ...*HiringTeam) *UserCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uc.AddHiringTeamEdgeIDs(ids...)
}

// SetMemberOfHiringTeamEdgesID sets the "member_of_hiring_team_edges" edge to the HiringTeam entity by ID.
func (uc *UserCreate) SetMemberOfHiringTeamEdgesID(id uuid.UUID) *UserCreate {
	uc.mutation.SetMemberOfHiringTeamEdgesID(id)
	return uc
}

// SetNillableMemberOfHiringTeamEdgesID sets the "member_of_hiring_team_edges" edge to the HiringTeam entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableMemberOfHiringTeamEdgesID(id *uuid.UUID) *UserCreate {
	if id != nil {
		uc = uc.SetMemberOfHiringTeamEdgesID(*id)
	}
	return uc
}

// SetMemberOfHiringTeamEdges sets the "member_of_hiring_team_edges" edge to the HiringTeam entity.
func (uc *UserCreate) SetMemberOfHiringTeamEdges(h *HiringTeam) *UserCreate {
	return uc.SetMemberOfHiringTeamEdgesID(h.ID)
}

// AddApproversHiringTeamIDs adds the "approvers_hiring_teams" edge to the HiringTeam entity by IDs.
func (uc *UserCreate) AddApproversHiringTeamIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddApproversHiringTeamIDs(ids...)
	return uc
}

// AddApproversHiringTeams adds the "approvers_hiring_teams" edges to the HiringTeam entity.
func (uc *UserCreate) AddApproversHiringTeams(h ...*HiringTeam) *UserCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uc.AddApproversHiringTeamIDs(ids...)
}

// SetLeaderRecEdgeID sets the "leader_rec_edge" edge to the RecTeam entity by ID.
func (uc *UserCreate) SetLeaderRecEdgeID(id uuid.UUID) *UserCreate {
	uc.mutation.SetLeaderRecEdgeID(id)
	return uc
}

// SetNillableLeaderRecEdgeID sets the "leader_rec_edge" edge to the RecTeam entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableLeaderRecEdgeID(id *uuid.UUID) *UserCreate {
	if id != nil {
		uc = uc.SetLeaderRecEdgeID(*id)
	}
	return uc
}

// SetLeaderRecEdge sets the "leader_rec_edge" edge to the RecTeam entity.
func (uc *UserCreate) SetLeaderRecEdge(r *RecTeam) *UserCreate {
	return uc.SetLeaderRecEdgeID(r.ID)
}

// SetRecTeamsID sets the "rec_teams" edge to the RecTeam entity by ID.
func (uc *UserCreate) SetRecTeamsID(id uuid.UUID) *UserCreate {
	uc.mutation.SetRecTeamsID(id)
	return uc
}

// SetNillableRecTeamsID sets the "rec_teams" edge to the RecTeam entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableRecTeamsID(id *uuid.UUID) *UserCreate {
	if id != nil {
		uc = uc.SetRecTeamsID(*id)
	}
	return uc
}

// SetRecTeams sets the "rec_teams" edge to the RecTeam entity.
func (uc *UserCreate) SetRecTeams(r *RecTeam) *UserCreate {
	return uc.SetRecTeamsID(r.ID)
}

// AddCandidateNoteEdgeIDs adds the "candidate_note_edges" edge to the CandidateNote entity by IDs.
func (uc *UserCreate) AddCandidateNoteEdgeIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddCandidateNoteEdgeIDs(ids...)
	return uc
}

// AddCandidateNoteEdges adds the "candidate_note_edges" edges to the CandidateNote entity.
func (uc *UserCreate) AddCandidateNoteEdges(c ...*CandidateNote) *UserCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCandidateNoteEdgeIDs(ids...)
}

// AddInterviewUserIDs adds the "interview_users" edge to the CandidateInterviewer entity by IDs.
func (uc *UserCreate) AddInterviewUserIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddInterviewUserIDs(ids...)
	return uc
}

// AddInterviewUsers adds the "interview_users" edges to the CandidateInterviewer entity.
func (uc *UserCreate) AddInterviewUsers(c ...*CandidateInterviewer) *UserCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddInterviewUserIDs(ids...)
}

// AddRoleUserIDs adds the "role_users" edge to the UserRole entity by IDs.
func (uc *UserCreate) AddRoleUserIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddRoleUserIDs(ids...)
	return uc
}

// AddRoleUsers adds the "role_users" edges to the UserRole entity.
func (uc *UserCreate) AddRoleUsers(u ...*UserRole) *UserCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddRoleUserIDs(ids...)
}

// AddHiringTeamUserIDs adds the "hiring_team_users" edge to the HiringTeamManager entity by IDs.
func (uc *UserCreate) AddHiringTeamUserIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddHiringTeamUserIDs(ids...)
	return uc
}

// AddHiringTeamUsers adds the "hiring_team_users" edges to the HiringTeamManager entity.
func (uc *UserCreate) AddHiringTeamUsers(h ...*HiringTeamManager) *UserCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uc.AddHiringTeamUserIDs(ids...)
}

// AddHiringTeamApproverIDs adds the "hiring_team_approvers" edge to the HiringTeamApprover entity by IDs.
func (uc *UserCreate) AddHiringTeamApproverIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddHiringTeamApproverIDs(ids...)
	return uc
}

// AddHiringTeamApprovers adds the "hiring_team_approvers" edges to the HiringTeamApprover entity.
func (uc *UserCreate) AddHiringTeamApprovers(h ...*HiringTeamApprover) *UserCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uc.AddHiringTeamApproverIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.Status(); !ok {
		v := user.DefaultStatus
		uc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.WorkEmail(); !ok {
		return &ValidationError{Name: "work_email", err: errors.New(`ent: missing required field "User.work_email"`)}
	}
	if v, ok := uc.mutation.WorkEmail(); ok {
		if err := user.WorkEmailValidator(v); err != nil {
			return &ValidationError{Name: "work_email", err: fmt.Errorf(`ent: validator failed for field "User.work_email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "User.status"`)}
	}
	if v, ok := uc.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Oid(); !ok {
		return &ValidationError{Name: "oid", err: errors.New(`ent: missing required field "User.oid"`)}
	}
	if v, ok := uc.mutation.Oid(); ok {
		if err := user.OidValidator(v); err != nil {
			return &ValidationError{Name: "oid", err: fmt.Errorf(`ent: validator failed for field "User.oid": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Location(); ok {
		if err := user.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "User.location": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.WorkEmail(); ok {
		_spec.SetField(user.FieldWorkEmail, field.TypeString, value)
		_node.WorkEmail = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.Oid(); ok {
		_spec.SetField(user.FieldOid, field.TypeString, value)
		_node.Oid = value
	}
	if value, ok := uc.mutation.Location(); ok {
		_spec.SetField(user.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if nodes := uc.mutation.AuditEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuditEdgeTable,
			Columns: []string{user.AuditEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: audittrail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.HiringOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.HiringOwnerTable,
			Columns: []string{user.HiringOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringjob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CandidateJobFeedbackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CandidateJobFeedbackTable,
			Columns: []string{user.CandidateJobFeedbackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatejobfeedback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.InterviewEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.InterviewEdgesTable,
			Columns: user.InterviewEdgesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateinterview.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &CandidateInterviewerCreate{config: uc.config, mutation: newCandidateInterviewerMutation(uc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CandidateJobEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CandidateJobEdgesTable,
			Columns: []string{user.CandidateJobEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatejob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CandidateInterviewEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CandidateInterviewEdgesTable,
			Columns: []string{user.CandidateInterviewEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateinterview.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CandidateReferenceEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CandidateReferenceEdgesTable,
			Columns: []string{user.CandidateReferenceEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserPermissionEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserPermissionEdgesTable,
			Columns: []string{user.UserPermissionEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entitypermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RoleEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RoleEdgesTable,
			Columns: user.RoleEdgesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserRoleCreate{config: uc.config, mutation: newUserRoleMutation(uc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.HiringTeamEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.HiringTeamEdgesTable,
			Columns: user.HiringTeamEdgesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteam.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &HiringTeamManagerCreate{config: uc.config, mutation: newHiringTeamManagerMutation(uc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.MemberOfHiringTeamEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.MemberOfHiringTeamEdgesTable,
			Columns: []string{user.MemberOfHiringTeamEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteam.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.HiringTeamID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ApproversHiringTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.ApproversHiringTeamsTable,
			Columns: user.ApproversHiringTeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteam.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &HiringTeamApproverCreate{config: uc.config, mutation: newHiringTeamApproverMutation(uc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.LeaderRecEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.LeaderRecEdgeTable,
			Columns: []string{user.LeaderRecEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: recteam.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RecTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RecTeamsTable,
			Columns: []string{user.RecTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: recteam.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RecTeamID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.CandidateNoteEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CandidateNoteEdgesTable,
			Columns: []string{user.CandidateNoteEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatenote.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.InterviewUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.InterviewUsersTable,
			Columns: []string{user.InterviewUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateinterviewer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RoleUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.RoleUsersTable,
			Columns: []string{user.RoleUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: userrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.HiringTeamUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.HiringTeamUsersTable,
			Columns: []string{user.HiringTeamUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteammanager.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.HiringTeamApproversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.HiringTeamApproversTable,
			Columns: []string{user.HiringTeamApproversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringteamapprover.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
