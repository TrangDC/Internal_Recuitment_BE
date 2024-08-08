// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidateaward"
	"trec/ent/candidatecertificate"
	"trec/ent/candidateeducate"
	"trec/ent/candidateexp"
	"trec/ent/candidatehistorycall"
	"trec/ent/candidatejob"
	"trec/ent/entityskill"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateCreate is the builder for creating a Candidate entity.
type CandidateCreate struct {
	config
	mutation *CandidateMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CandidateCreate) SetCreatedAt(t time.Time) *CandidateCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableCreatedAt(t *time.Time) *CandidateCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CandidateCreate) SetUpdatedAt(t time.Time) *CandidateCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableUpdatedAt(t *time.Time) *CandidateCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CandidateCreate) SetDeletedAt(t time.Time) *CandidateCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableDeletedAt(t *time.Time) *CandidateCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CandidateCreate) SetName(s string) *CandidateCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetEmail sets the "email" field.
func (cc *CandidateCreate) SetEmail(s string) *CandidateCreate {
	cc.mutation.SetEmail(s)
	return cc
}

// SetPhone sets the "phone" field.
func (cc *CandidateCreate) SetPhone(s string) *CandidateCreate {
	cc.mutation.SetPhone(s)
	return cc
}

// SetDob sets the "dob" field.
func (cc *CandidateCreate) SetDob(t time.Time) *CandidateCreate {
	cc.mutation.SetDob(t)
	return cc
}

// SetNillableDob sets the "dob" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableDob(t *time.Time) *CandidateCreate {
	if t != nil {
		cc.SetDob(*t)
	}
	return cc
}

// SetIsBlacklist sets the "is_blacklist" field.
func (cc *CandidateCreate) SetIsBlacklist(b bool) *CandidateCreate {
	cc.mutation.SetIsBlacklist(b)
	return cc
}

// SetNillableIsBlacklist sets the "is_blacklist" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableIsBlacklist(b *bool) *CandidateCreate {
	if b != nil {
		cc.SetIsBlacklist(*b)
	}
	return cc
}

// SetLastApplyDate sets the "last_apply_date" field.
func (cc *CandidateCreate) SetLastApplyDate(t time.Time) *CandidateCreate {
	cc.mutation.SetLastApplyDate(t)
	return cc
}

// SetNillableLastApplyDate sets the "last_apply_date" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableLastApplyDate(t *time.Time) *CandidateCreate {
	if t != nil {
		cc.SetLastApplyDate(*t)
	}
	return cc
}

// SetReferenceType sets the "reference_type" field.
func (cc *CandidateCreate) SetReferenceType(ct candidate.ReferenceType) *CandidateCreate {
	cc.mutation.SetReferenceType(ct)
	return cc
}

// SetNillableReferenceType sets the "reference_type" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableReferenceType(ct *candidate.ReferenceType) *CandidateCreate {
	if ct != nil {
		cc.SetReferenceType(*ct)
	}
	return cc
}

// SetReferenceValue sets the "reference_value" field.
func (cc *CandidateCreate) SetReferenceValue(s string) *CandidateCreate {
	cc.mutation.SetReferenceValue(s)
	return cc
}

// SetNillableReferenceValue sets the "reference_value" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableReferenceValue(s *string) *CandidateCreate {
	if s != nil {
		cc.SetReferenceValue(*s)
	}
	return cc
}

// SetReferenceUID sets the "reference_uid" field.
func (cc *CandidateCreate) SetReferenceUID(u uuid.UUID) *CandidateCreate {
	cc.mutation.SetReferenceUID(u)
	return cc
}

// SetNillableReferenceUID sets the "reference_uid" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableReferenceUID(u *uuid.UUID) *CandidateCreate {
	if u != nil {
		cc.SetReferenceUID(*u)
	}
	return cc
}

// SetRecruitTime sets the "recruit_time" field.
func (cc *CandidateCreate) SetRecruitTime(t time.Time) *CandidateCreate {
	cc.mutation.SetRecruitTime(t)
	return cc
}

// SetNillableRecruitTime sets the "recruit_time" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableRecruitTime(t *time.Time) *CandidateCreate {
	if t != nil {
		cc.SetRecruitTime(*t)
	}
	return cc
}

// SetDescription sets the "description" field.
func (cc *CandidateCreate) SetDescription(s string) *CandidateCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableDescription(s *string) *CandidateCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetAvatar sets the "avatar" field.
func (cc *CandidateCreate) SetAvatar(u uuid.UUID) *CandidateCreate {
	cc.mutation.SetAvatar(u)
	return cc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableAvatar(u *uuid.UUID) *CandidateCreate {
	if u != nil {
		cc.SetAvatar(*u)
	}
	return cc
}

// SetCountry sets the "country" field.
func (cc *CandidateCreate) SetCountry(s string) *CandidateCreate {
	cc.mutation.SetCountry(s)
	return cc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableCountry(s *string) *CandidateCreate {
	if s != nil {
		cc.SetCountry(*s)
	}
	return cc
}

// SetAddress sets the "address" field.
func (cc *CandidateCreate) SetAddress(s string) *CandidateCreate {
	cc.mutation.SetAddress(s)
	return cc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (cc *CandidateCreate) SetNillableAddress(s *string) *CandidateCreate {
	if s != nil {
		cc.SetAddress(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CandidateCreate) SetID(u uuid.UUID) *CandidateCreate {
	cc.mutation.SetID(u)
	return cc
}

// AddCandidateJobEdgeIDs adds the "candidate_job_edges" edge to the CandidateJob entity by IDs.
func (cc *CandidateCreate) AddCandidateJobEdgeIDs(ids ...uuid.UUID) *CandidateCreate {
	cc.mutation.AddCandidateJobEdgeIDs(ids...)
	return cc
}

// AddCandidateJobEdges adds the "candidate_job_edges" edges to the CandidateJob entity.
func (cc *CandidateCreate) AddCandidateJobEdges(c ...*CandidateJob) *CandidateCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCandidateJobEdgeIDs(ids...)
}

// SetReferenceUserEdgeID sets the "reference_user_edge" edge to the User entity by ID.
func (cc *CandidateCreate) SetReferenceUserEdgeID(id uuid.UUID) *CandidateCreate {
	cc.mutation.SetReferenceUserEdgeID(id)
	return cc
}

// SetNillableReferenceUserEdgeID sets the "reference_user_edge" edge to the User entity by ID if the given value is not nil.
func (cc *CandidateCreate) SetNillableReferenceUserEdgeID(id *uuid.UUID) *CandidateCreate {
	if id != nil {
		cc = cc.SetReferenceUserEdgeID(*id)
	}
	return cc
}

// SetReferenceUserEdge sets the "reference_user_edge" edge to the User entity.
func (cc *CandidateCreate) SetReferenceUserEdge(u *User) *CandidateCreate {
	return cc.SetReferenceUserEdgeID(u.ID)
}

// AddAttachmentEdgeIDs adds the "attachment_edges" edge to the Attachment entity by IDs.
func (cc *CandidateCreate) AddAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateCreate {
	cc.mutation.AddAttachmentEdgeIDs(ids...)
	return cc
}

// AddAttachmentEdges adds the "attachment_edges" edges to the Attachment entity.
func (cc *CandidateCreate) AddAttachmentEdges(a ...*Attachment) *CandidateCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddAttachmentEdgeIDs(ids...)
}

// AddCandidateSkillEdgeIDs adds the "candidate_skill_edges" edge to the EntitySkill entity by IDs.
func (cc *CandidateCreate) AddCandidateSkillEdgeIDs(ids ...uuid.UUID) *CandidateCreate {
	cc.mutation.AddCandidateSkillEdgeIDs(ids...)
	return cc
}

// AddCandidateSkillEdges adds the "candidate_skill_edges" edges to the EntitySkill entity.
func (cc *CandidateCreate) AddCandidateSkillEdges(e ...*EntitySkill) *CandidateCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cc.AddCandidateSkillEdgeIDs(ids...)
}

// AddCandidateExpEdgeIDs adds the "candidate_exp_edges" edge to the CandidateExp entity by IDs.
func (cc *CandidateCreate) AddCandidateExpEdgeIDs(ids ...uuid.UUID) *CandidateCreate {
	cc.mutation.AddCandidateExpEdgeIDs(ids...)
	return cc
}

// AddCandidateExpEdges adds the "candidate_exp_edges" edges to the CandidateExp entity.
func (cc *CandidateCreate) AddCandidateExpEdges(c ...*CandidateExp) *CandidateCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCandidateExpEdgeIDs(ids...)
}

// AddCandidateEducateEdgeIDs adds the "candidate_educate_edges" edge to the CandidateEducate entity by IDs.
func (cc *CandidateCreate) AddCandidateEducateEdgeIDs(ids ...uuid.UUID) *CandidateCreate {
	cc.mutation.AddCandidateEducateEdgeIDs(ids...)
	return cc
}

// AddCandidateEducateEdges adds the "candidate_educate_edges" edges to the CandidateEducate entity.
func (cc *CandidateCreate) AddCandidateEducateEdges(c ...*CandidateEducate) *CandidateCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCandidateEducateEdgeIDs(ids...)
}

// AddCandidateAwardEdgeIDs adds the "candidate_award_edges" edge to the CandidateAward entity by IDs.
func (cc *CandidateCreate) AddCandidateAwardEdgeIDs(ids ...uuid.UUID) *CandidateCreate {
	cc.mutation.AddCandidateAwardEdgeIDs(ids...)
	return cc
}

// AddCandidateAwardEdges adds the "candidate_award_edges" edges to the CandidateAward entity.
func (cc *CandidateCreate) AddCandidateAwardEdges(c ...*CandidateAward) *CandidateCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCandidateAwardEdgeIDs(ids...)
}

// AddCandidateCertificateEdgeIDs adds the "candidate_certificate_edges" edge to the CandidateCertificate entity by IDs.
func (cc *CandidateCreate) AddCandidateCertificateEdgeIDs(ids ...uuid.UUID) *CandidateCreate {
	cc.mutation.AddCandidateCertificateEdgeIDs(ids...)
	return cc
}

// AddCandidateCertificateEdges adds the "candidate_certificate_edges" edges to the CandidateCertificate entity.
func (cc *CandidateCreate) AddCandidateCertificateEdges(c ...*CandidateCertificate) *CandidateCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCandidateCertificateEdgeIDs(ids...)
}

// AddCandidateHistoryCallEdgeIDs adds the "candidate_history_call_edges" edge to the CandidateHistoryCall entity by IDs.
func (cc *CandidateCreate) AddCandidateHistoryCallEdgeIDs(ids ...uuid.UUID) *CandidateCreate {
	cc.mutation.AddCandidateHistoryCallEdgeIDs(ids...)
	return cc
}

// AddCandidateHistoryCallEdges adds the "candidate_history_call_edges" edges to the CandidateHistoryCall entity.
func (cc *CandidateCreate) AddCandidateHistoryCallEdges(c ...*CandidateHistoryCall) *CandidateCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCandidateHistoryCallEdgeIDs(ids...)
}

// Mutation returns the CandidateMutation object of the builder.
func (cc *CandidateCreate) Mutation() *CandidateMutation {
	return cc.mutation
}

// Save creates the Candidate in the database.
func (cc *CandidateCreate) Save(ctx context.Context) (*Candidate, error) {
	var (
		err  error
		node *Candidate
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Candidate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CandidateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CandidateCreate) SaveX(ctx context.Context) *Candidate {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CandidateCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CandidateCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CandidateCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := candidate.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.IsBlacklist(); !ok {
		v := candidate.DefaultIsBlacklist
		cc.mutation.SetIsBlacklist(v)
	}
	if _, ok := cc.mutation.ReferenceType(); !ok {
		v := candidate.DefaultReferenceType
		cc.mutation.SetReferenceType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CandidateCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Candidate.created_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Candidate.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := candidate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Candidate.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Candidate.email"`)}
	}
	if v, ok := cc.mutation.Email(); ok {
		if err := candidate.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Candidate.email": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Candidate.phone"`)}
	}
	if v, ok := cc.mutation.Phone(); ok {
		if err := candidate.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Candidate.phone": %w`, err)}
		}
	}
	if _, ok := cc.mutation.IsBlacklist(); !ok {
		return &ValidationError{Name: "is_blacklist", err: errors.New(`ent: missing required field "Candidate.is_blacklist"`)}
	}
	if _, ok := cc.mutation.ReferenceType(); !ok {
		return &ValidationError{Name: "reference_type", err: errors.New(`ent: missing required field "Candidate.reference_type"`)}
	}
	if v, ok := cc.mutation.ReferenceType(); ok {
		if err := candidate.ReferenceTypeValidator(v); err != nil {
			return &ValidationError{Name: "reference_type", err: fmt.Errorf(`ent: validator failed for field "Candidate.reference_type": %w`, err)}
		}
	}
	if v, ok := cc.mutation.ReferenceValue(); ok {
		if err := candidate.ReferenceValueValidator(v); err != nil {
			return &ValidationError{Name: "reference_value", err: fmt.Errorf(`ent: validator failed for field "Candidate.reference_value": %w`, err)}
		}
	}
	if v, ok := cc.mutation.Description(); ok {
		if err := candidate.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Candidate.description": %w`, err)}
		}
	}
	if v, ok := cc.mutation.Country(); ok {
		if err := candidate.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "Candidate.country": %w`, err)}
		}
	}
	if v, ok := cc.mutation.Address(); ok {
		if err := candidate.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Candidate.address": %w`, err)}
		}
	}
	return nil
}

func (cc *CandidateCreate) sqlSave(ctx context.Context) (*Candidate, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *CandidateCreate) createSpec() (*Candidate, *sqlgraph.CreateSpec) {
	var (
		_node = &Candidate{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: candidate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidate.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(candidate.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(candidate.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(candidate.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(candidate.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Email(); ok {
		_spec.SetField(candidate.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := cc.mutation.Phone(); ok {
		_spec.SetField(candidate.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := cc.mutation.Dob(); ok {
		_spec.SetField(candidate.FieldDob, field.TypeTime, value)
		_node.Dob = value
	}
	if value, ok := cc.mutation.IsBlacklist(); ok {
		_spec.SetField(candidate.FieldIsBlacklist, field.TypeBool, value)
		_node.IsBlacklist = value
	}
	if value, ok := cc.mutation.LastApplyDate(); ok {
		_spec.SetField(candidate.FieldLastApplyDate, field.TypeTime, value)
		_node.LastApplyDate = value
	}
	if value, ok := cc.mutation.ReferenceType(); ok {
		_spec.SetField(candidate.FieldReferenceType, field.TypeEnum, value)
		_node.ReferenceType = value
	}
	if value, ok := cc.mutation.ReferenceValue(); ok {
		_spec.SetField(candidate.FieldReferenceValue, field.TypeString, value)
		_node.ReferenceValue = value
	}
	if value, ok := cc.mutation.RecruitTime(); ok {
		_spec.SetField(candidate.FieldRecruitTime, field.TypeTime, value)
		_node.RecruitTime = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(candidate.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.Avatar(); ok {
		_spec.SetField(candidate.FieldAvatar, field.TypeUUID, value)
		_node.Avatar = value
	}
	if value, ok := cc.mutation.Country(); ok {
		_spec.SetField(candidate.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := cc.mutation.Address(); ok {
		_spec.SetField(candidate.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if nodes := cc.mutation.CandidateJobEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidate.CandidateJobEdgesTable,
			Columns: []string{candidate.CandidateJobEdgesColumn},
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
	if nodes := cc.mutation.ReferenceUserEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidate.ReferenceUserEdgeTable,
			Columns: []string{candidate.ReferenceUserEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReferenceUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AttachmentEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidate.AttachmentEdgesTable,
			Columns: []string{candidate.AttachmentEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CandidateSkillEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidate.CandidateSkillEdgesTable,
			Columns: []string{candidate.CandidateSkillEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entityskill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CandidateExpEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidate.CandidateExpEdgesTable,
			Columns: []string{candidate.CandidateExpEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateexp.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CandidateEducateEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidate.CandidateEducateEdgesTable,
			Columns: []string{candidate.CandidateEducateEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateeducate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CandidateAwardEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidate.CandidateAwardEdgesTable,
			Columns: []string{candidate.CandidateAwardEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidateaward.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CandidateCertificateEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidate.CandidateCertificateEdgesTable,
			Columns: []string{candidate.CandidateCertificateEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatecertificate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CandidateHistoryCallEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidate.CandidateHistoryCallEdgesTable,
			Columns: []string{candidate.CandidateHistoryCallEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidatehistorycall.FieldID,
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

// CandidateCreateBulk is the builder for creating many Candidate entities in bulk.
type CandidateCreateBulk struct {
	config
	builders []*CandidateCreate
}

// Save creates the Candidate entities in the database.
func (ccb *CandidateCreateBulk) Save(ctx context.Context) ([]*Candidate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Candidate, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CandidateMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CandidateCreateBulk) SaveX(ctx context.Context) []*Candidate {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CandidateCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CandidateCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
