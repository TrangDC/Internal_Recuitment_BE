// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/hiringjob"
	"trec/ent/hiringjobstep"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HiringJobStepCreate is the builder for creating a HiringJobStep entity.
type HiringJobStepCreate struct {
	config
	mutation *HiringJobStepMutation
	hooks    []Hook
}

// SetHiringJobID sets the "hiring_job_id" field.
func (hjsc *HiringJobStepCreate) SetHiringJobID(u uuid.UUID) *HiringJobStepCreate {
	hjsc.mutation.SetHiringJobID(u)
	return hjsc
}

// SetUserID sets the "user_id" field.
func (hjsc *HiringJobStepCreate) SetUserID(u uuid.UUID) *HiringJobStepCreate {
	hjsc.mutation.SetUserID(u)
	return hjsc
}

// SetStatus sets the "status" field.
func (hjsc *HiringJobStepCreate) SetStatus(h hiringjobstep.Status) *HiringJobStepCreate {
	hjsc.mutation.SetStatus(h)
	return hjsc
}

// SetOrderID sets the "order_id" field.
func (hjsc *HiringJobStepCreate) SetOrderID(i int) *HiringJobStepCreate {
	hjsc.mutation.SetOrderID(i)
	return hjsc
}

// SetCreatedAt sets the "created_at" field.
func (hjsc *HiringJobStepCreate) SetCreatedAt(t time.Time) *HiringJobStepCreate {
	hjsc.mutation.SetCreatedAt(t)
	return hjsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hjsc *HiringJobStepCreate) SetNillableCreatedAt(t *time.Time) *HiringJobStepCreate {
	if t != nil {
		hjsc.SetCreatedAt(*t)
	}
	return hjsc
}

// SetUpdatedAt sets the "updated_at" field.
func (hjsc *HiringJobStepCreate) SetUpdatedAt(t time.Time) *HiringJobStepCreate {
	hjsc.mutation.SetUpdatedAt(t)
	return hjsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hjsc *HiringJobStepCreate) SetNillableUpdatedAt(t *time.Time) *HiringJobStepCreate {
	if t != nil {
		hjsc.SetUpdatedAt(*t)
	}
	return hjsc
}

// SetID sets the "id" field.
func (hjsc *HiringJobStepCreate) SetID(u uuid.UUID) *HiringJobStepCreate {
	hjsc.mutation.SetID(u)
	return hjsc
}

// SetApprovalJobID sets the "approval_job" edge to the HiringJob entity by ID.
func (hjsc *HiringJobStepCreate) SetApprovalJobID(id uuid.UUID) *HiringJobStepCreate {
	hjsc.mutation.SetApprovalJobID(id)
	return hjsc
}

// SetApprovalJob sets the "approval_job" edge to the HiringJob entity.
func (hjsc *HiringJobStepCreate) SetApprovalJob(h *HiringJob) *HiringJobStepCreate {
	return hjsc.SetApprovalJobID(h.ID)
}

// SetApprovalUserID sets the "approval_user" edge to the User entity by ID.
func (hjsc *HiringJobStepCreate) SetApprovalUserID(id uuid.UUID) *HiringJobStepCreate {
	hjsc.mutation.SetApprovalUserID(id)
	return hjsc
}

// SetApprovalUser sets the "approval_user" edge to the User entity.
func (hjsc *HiringJobStepCreate) SetApprovalUser(u *User) *HiringJobStepCreate {
	return hjsc.SetApprovalUserID(u.ID)
}

// Mutation returns the HiringJobStepMutation object of the builder.
func (hjsc *HiringJobStepCreate) Mutation() *HiringJobStepMutation {
	return hjsc.mutation
}

// Save creates the HiringJobStep in the database.
func (hjsc *HiringJobStepCreate) Save(ctx context.Context) (*HiringJobStep, error) {
	var (
		err  error
		node *HiringJobStep
	)
	hjsc.defaults()
	if len(hjsc.hooks) == 0 {
		if err = hjsc.check(); err != nil {
			return nil, err
		}
		node, err = hjsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HiringJobStepMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hjsc.check(); err != nil {
				return nil, err
			}
			hjsc.mutation = mutation
			if node, err = hjsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hjsc.hooks) - 1; i >= 0; i-- {
			if hjsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hjsc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, hjsc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*HiringJobStep)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HiringJobStepMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hjsc *HiringJobStepCreate) SaveX(ctx context.Context) *HiringJobStep {
	v, err := hjsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hjsc *HiringJobStepCreate) Exec(ctx context.Context) error {
	_, err := hjsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hjsc *HiringJobStepCreate) ExecX(ctx context.Context) {
	if err := hjsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hjsc *HiringJobStepCreate) defaults() {
	if _, ok := hjsc.mutation.CreatedAt(); !ok {
		v := hiringjobstep.DefaultCreatedAt()
		hjsc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hjsc *HiringJobStepCreate) check() error {
	if _, ok := hjsc.mutation.HiringJobID(); !ok {
		return &ValidationError{Name: "hiring_job_id", err: errors.New(`ent: missing required field "HiringJobStep.hiring_job_id"`)}
	}
	if _, ok := hjsc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "HiringJobStep.user_id"`)}
	}
	if _, ok := hjsc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "HiringJobStep.status"`)}
	}
	if v, ok := hjsc.mutation.Status(); ok {
		if err := hiringjobstep.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HiringJobStep.status": %w`, err)}
		}
	}
	if _, ok := hjsc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "HiringJobStep.order_id"`)}
	}
	if v, ok := hjsc.mutation.OrderID(); ok {
		if err := hiringjobstep.OrderIDValidator(v); err != nil {
			return &ValidationError{Name: "order_id", err: fmt.Errorf(`ent: validator failed for field "HiringJobStep.order_id": %w`, err)}
		}
	}
	if _, ok := hjsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "HiringJobStep.created_at"`)}
	}
	if _, ok := hjsc.mutation.ApprovalJobID(); !ok {
		return &ValidationError{Name: "approval_job", err: errors.New(`ent: missing required edge "HiringJobStep.approval_job"`)}
	}
	if _, ok := hjsc.mutation.ApprovalUserID(); !ok {
		return &ValidationError{Name: "approval_user", err: errors.New(`ent: missing required edge "HiringJobStep.approval_user"`)}
	}
	return nil
}

func (hjsc *HiringJobStepCreate) sqlSave(ctx context.Context) (*HiringJobStep, error) {
	_node, _spec := hjsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hjsc.driver, _spec); err != nil {
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

func (hjsc *HiringJobStepCreate) createSpec() (*HiringJobStep, *sqlgraph.CreateSpec) {
	var (
		_node = &HiringJobStep{config: hjsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hiringjobstep.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hiringjobstep.FieldID,
			},
		}
	)
	if id, ok := hjsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := hjsc.mutation.Status(); ok {
		_spec.SetField(hiringjobstep.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := hjsc.mutation.OrderID(); ok {
		_spec.SetField(hiringjobstep.FieldOrderID, field.TypeInt, value)
		_node.OrderID = value
	}
	if value, ok := hjsc.mutation.CreatedAt(); ok {
		_spec.SetField(hiringjobstep.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hjsc.mutation.UpdatedAt(); ok {
		_spec.SetField(hiringjobstep.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := hjsc.mutation.ApprovalJobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hiringjobstep.ApprovalJobTable,
			Columns: []string{hiringjobstep.ApprovalJobColumn},
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
		_node.HiringJobID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hjsc.mutation.ApprovalUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hiringjobstep.ApprovalUserTable,
			Columns: []string{hiringjobstep.ApprovalUserColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HiringJobStepCreateBulk is the builder for creating many HiringJobStep entities in bulk.
type HiringJobStepCreateBulk struct {
	config
	builders []*HiringJobStepCreate
}

// Save creates the HiringJobStep entities in the database.
func (hjscb *HiringJobStepCreateBulk) Save(ctx context.Context) ([]*HiringJobStep, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hjscb.builders))
	nodes := make([]*HiringJobStep, len(hjscb.builders))
	mutators := make([]Mutator, len(hjscb.builders))
	for i := range hjscb.builders {
		func(i int, root context.Context) {
			builder := hjscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HiringJobStepMutation)
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
					_, err = mutators[i+1].Mutate(root, hjscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hjscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hjscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hjscb *HiringJobStepCreateBulk) SaveX(ctx context.Context) []*HiringJobStep {
	v, err := hjscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hjscb *HiringJobStepCreateBulk) Exec(ctx context.Context) error {
	_, err := hjscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hjscb *HiringJobStepCreateBulk) ExecX(ctx context.Context) {
	if err := hjscb.Exec(ctx); err != nil {
		panic(err)
	}
}
