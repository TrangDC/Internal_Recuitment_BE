// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobstep"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateJobStepCreate is the builder for creating a CandidateJobStep entity.
type CandidateJobStepCreate struct {
	config
	mutation *CandidateJobStepMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cjsc *CandidateJobStepCreate) SetCreatedAt(t time.Time) *CandidateJobStepCreate {
	cjsc.mutation.SetCreatedAt(t)
	return cjsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cjsc *CandidateJobStepCreate) SetNillableCreatedAt(t *time.Time) *CandidateJobStepCreate {
	if t != nil {
		cjsc.SetCreatedAt(*t)
	}
	return cjsc
}

// SetUpdatedAt sets the "updated_at" field.
func (cjsc *CandidateJobStepCreate) SetUpdatedAt(t time.Time) *CandidateJobStepCreate {
	cjsc.mutation.SetUpdatedAt(t)
	return cjsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cjsc *CandidateJobStepCreate) SetNillableUpdatedAt(t *time.Time) *CandidateJobStepCreate {
	if t != nil {
		cjsc.SetUpdatedAt(*t)
	}
	return cjsc
}

// SetDeletedAt sets the "deleted_at" field.
func (cjsc *CandidateJobStepCreate) SetDeletedAt(t time.Time) *CandidateJobStepCreate {
	cjsc.mutation.SetDeletedAt(t)
	return cjsc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cjsc *CandidateJobStepCreate) SetNillableDeletedAt(t *time.Time) *CandidateJobStepCreate {
	if t != nil {
		cjsc.SetDeletedAt(*t)
	}
	return cjsc
}

// SetCandidateJobStatus sets the "candidate_job_status" field.
func (cjsc *CandidateJobStepCreate) SetCandidateJobStatus(cjs candidatejobstep.CandidateJobStatus) *CandidateJobStepCreate {
	cjsc.mutation.SetCandidateJobStatus(cjs)
	return cjsc
}

// SetNillableCandidateJobStatus sets the "candidate_job_status" field if the given value is not nil.
func (cjsc *CandidateJobStepCreate) SetNillableCandidateJobStatus(cjs *candidatejobstep.CandidateJobStatus) *CandidateJobStepCreate {
	if cjs != nil {
		cjsc.SetCandidateJobStatus(*cjs)
	}
	return cjsc
}

// SetCandidateJobID sets the "candidate_job_id" field.
func (cjsc *CandidateJobStepCreate) SetCandidateJobID(u uuid.UUID) *CandidateJobStepCreate {
	cjsc.mutation.SetCandidateJobID(u)
	return cjsc
}

// SetNillableCandidateJobID sets the "candidate_job_id" field if the given value is not nil.
func (cjsc *CandidateJobStepCreate) SetNillableCandidateJobID(u *uuid.UUID) *CandidateJobStepCreate {
	if u != nil {
		cjsc.SetCandidateJobID(*u)
	}
	return cjsc
}

// SetID sets the "id" field.
func (cjsc *CandidateJobStepCreate) SetID(u uuid.UUID) *CandidateJobStepCreate {
	cjsc.mutation.SetID(u)
	return cjsc
}

// SetCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID.
func (cjsc *CandidateJobStepCreate) SetCandidateJobEdgeID(id uuid.UUID) *CandidateJobStepCreate {
	cjsc.mutation.SetCandidateJobEdgeID(id)
	return cjsc
}

// SetNillableCandidateJobEdgeID sets the "candidate_job_edge" edge to the CandidateJob entity by ID if the given value is not nil.
func (cjsc *CandidateJobStepCreate) SetNillableCandidateJobEdgeID(id *uuid.UUID) *CandidateJobStepCreate {
	if id != nil {
		cjsc = cjsc.SetCandidateJobEdgeID(*id)
	}
	return cjsc
}

// SetCandidateJobEdge sets the "candidate_job_edge" edge to the CandidateJob entity.
func (cjsc *CandidateJobStepCreate) SetCandidateJobEdge(c *CandidateJob) *CandidateJobStepCreate {
	return cjsc.SetCandidateJobEdgeID(c.ID)
}

// Mutation returns the CandidateJobStepMutation object of the builder.
func (cjsc *CandidateJobStepCreate) Mutation() *CandidateJobStepMutation {
	return cjsc.mutation
}

// Save creates the CandidateJobStep in the database.
func (cjsc *CandidateJobStepCreate) Save(ctx context.Context) (*CandidateJobStep, error) {
	var (
		err  error
		node *CandidateJobStep
	)
	cjsc.defaults()
	if len(cjsc.hooks) == 0 {
		if err = cjsc.check(); err != nil {
			return nil, err
		}
		node, err = cjsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateJobStepMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cjsc.check(); err != nil {
				return nil, err
			}
			cjsc.mutation = mutation
			if node, err = cjsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cjsc.hooks) - 1; i >= 0; i-- {
			if cjsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cjsc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cjsc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CandidateJobStep)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CandidateJobStepMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cjsc *CandidateJobStepCreate) SaveX(ctx context.Context) *CandidateJobStep {
	v, err := cjsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cjsc *CandidateJobStepCreate) Exec(ctx context.Context) error {
	_, err := cjsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cjsc *CandidateJobStepCreate) ExecX(ctx context.Context) {
	if err := cjsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cjsc *CandidateJobStepCreate) defaults() {
	if _, ok := cjsc.mutation.CreatedAt(); !ok {
		v := candidatejobstep.DefaultCreatedAt()
		cjsc.mutation.SetCreatedAt(v)
	}
	if _, ok := cjsc.mutation.CandidateJobStatus(); !ok {
		v := candidatejobstep.DefaultCandidateJobStatus
		cjsc.mutation.SetCandidateJobStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cjsc *CandidateJobStepCreate) check() error {
	if _, ok := cjsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CandidateJobStep.created_at"`)}
	}
	if _, ok := cjsc.mutation.CandidateJobStatus(); !ok {
		return &ValidationError{Name: "candidate_job_status", err: errors.New(`ent: missing required field "CandidateJobStep.candidate_job_status"`)}
	}
	if v, ok := cjsc.mutation.CandidateJobStatus(); ok {
		if err := candidatejobstep.CandidateJobStatusValidator(v); err != nil {
			return &ValidationError{Name: "candidate_job_status", err: fmt.Errorf(`ent: validator failed for field "CandidateJobStep.candidate_job_status": %w`, err)}
		}
	}
	return nil
}

func (cjsc *CandidateJobStepCreate) sqlSave(ctx context.Context) (*CandidateJobStep, error) {
	_node, _spec := cjsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cjsc.driver, _spec); err != nil {
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

func (cjsc *CandidateJobStepCreate) createSpec() (*CandidateJobStep, *sqlgraph.CreateSpec) {
	var (
		_node = &CandidateJobStep{config: cjsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: candidatejobstep.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatejobstep.FieldID,
			},
		}
	)
	if id, ok := cjsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cjsc.mutation.CreatedAt(); ok {
		_spec.SetField(candidatejobstep.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cjsc.mutation.UpdatedAt(); ok {
		_spec.SetField(candidatejobstep.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cjsc.mutation.DeletedAt(); ok {
		_spec.SetField(candidatejobstep.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cjsc.mutation.CandidateJobStatus(); ok {
		_spec.SetField(candidatejobstep.FieldCandidateJobStatus, field.TypeEnum, value)
		_node.CandidateJobStatus = value
	}
	if nodes := cjsc.mutation.CandidateJobEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidatejobstep.CandidateJobEdgeTable,
			Columns: []string{candidatejobstep.CandidateJobEdgeColumn},
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
		_node.CandidateJobID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CandidateJobStepCreateBulk is the builder for creating many CandidateJobStep entities in bulk.
type CandidateJobStepCreateBulk struct {
	config
	builders []*CandidateJobStepCreate
}

// Save creates the CandidateJobStep entities in the database.
func (cjscb *CandidateJobStepCreateBulk) Save(ctx context.Context) ([]*CandidateJobStep, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cjscb.builders))
	nodes := make([]*CandidateJobStep, len(cjscb.builders))
	mutators := make([]Mutator, len(cjscb.builders))
	for i := range cjscb.builders {
		func(i int, root context.Context) {
			builder := cjscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CandidateJobStepMutation)
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
					_, err = mutators[i+1].Mutate(root, cjscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cjscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cjscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cjscb *CandidateJobStepCreateBulk) SaveX(ctx context.Context) []*CandidateJobStep {
	v, err := cjscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cjscb *CandidateJobStepCreateBulk) Exec(ctx context.Context) error {
	_, err := cjscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cjscb *CandidateJobStepCreateBulk) ExecX(ctx context.Context) {
	if err := cjscb.Exec(ctx); err != nil {
		panic(err)
	}
}