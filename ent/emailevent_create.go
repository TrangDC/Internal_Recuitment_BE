// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/emailevent"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmailEventCreate is the builder for creating a EmailEvent entity.
type EmailEventCreate struct {
	config
	mutation *EmailEventMutation
	hooks    []Hook
}

// SetModule sets the "module" field.
func (eec *EmailEventCreate) SetModule(e emailevent.Module) *EmailEventCreate {
	eec.mutation.SetModule(e)
	return eec
}

// SetAction sets the "action" field.
func (eec *EmailEventCreate) SetAction(e emailevent.Action) *EmailEventCreate {
	eec.mutation.SetAction(e)
	return eec
}

// SetName sets the "name" field.
func (eec *EmailEventCreate) SetName(s string) *EmailEventCreate {
	eec.mutation.SetName(s)
	return eec
}

// SetCreatedAt sets the "created_at" field.
func (eec *EmailEventCreate) SetCreatedAt(t time.Time) *EmailEventCreate {
	eec.mutation.SetCreatedAt(t)
	return eec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eec *EmailEventCreate) SetNillableCreatedAt(t *time.Time) *EmailEventCreate {
	if t != nil {
		eec.SetCreatedAt(*t)
	}
	return eec
}

// SetUpdatedAt sets the "updated_at" field.
func (eec *EmailEventCreate) SetUpdatedAt(t time.Time) *EmailEventCreate {
	eec.mutation.SetUpdatedAt(t)
	return eec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eec *EmailEventCreate) SetNillableUpdatedAt(t *time.Time) *EmailEventCreate {
	if t != nil {
		eec.SetUpdatedAt(*t)
	}
	return eec
}

// SetID sets the "id" field.
func (eec *EmailEventCreate) SetID(u uuid.UUID) *EmailEventCreate {
	eec.mutation.SetID(u)
	return eec
}

// Mutation returns the EmailEventMutation object of the builder.
func (eec *EmailEventCreate) Mutation() *EmailEventMutation {
	return eec.mutation
}

// Save creates the EmailEvent in the database.
func (eec *EmailEventCreate) Save(ctx context.Context) (*EmailEvent, error) {
	var (
		err  error
		node *EmailEvent
	)
	eec.defaults()
	if len(eec.hooks) == 0 {
		if err = eec.check(); err != nil {
			return nil, err
		}
		node, err = eec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmailEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eec.check(); err != nil {
				return nil, err
			}
			eec.mutation = mutation
			if node, err = eec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(eec.hooks) - 1; i >= 0; i-- {
			if eec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, eec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EmailEvent)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EmailEventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (eec *EmailEventCreate) SaveX(ctx context.Context) *EmailEvent {
	v, err := eec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eec *EmailEventCreate) Exec(ctx context.Context) error {
	_, err := eec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eec *EmailEventCreate) ExecX(ctx context.Context) {
	if err := eec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eec *EmailEventCreate) defaults() {
	if _, ok := eec.mutation.CreatedAt(); !ok {
		v := emailevent.DefaultCreatedAt()
		eec.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eec *EmailEventCreate) check() error {
	if _, ok := eec.mutation.Module(); !ok {
		return &ValidationError{Name: "module", err: errors.New(`ent: missing required field "EmailEvent.module"`)}
	}
	if v, ok := eec.mutation.Module(); ok {
		if err := emailevent.ModuleValidator(v); err != nil {
			return &ValidationError{Name: "module", err: fmt.Errorf(`ent: validator failed for field "EmailEvent.module": %w`, err)}
		}
	}
	if _, ok := eec.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required field "EmailEvent.action"`)}
	}
	if v, ok := eec.mutation.Action(); ok {
		if err := emailevent.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "EmailEvent.action": %w`, err)}
		}
	}
	if _, ok := eec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "EmailEvent.name"`)}
	}
	if _, ok := eec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EmailEvent.created_at"`)}
	}
	return nil
}

func (eec *EmailEventCreate) sqlSave(ctx context.Context) (*EmailEvent, error) {
	_node, _spec := eec.createSpec()
	if err := sqlgraph.CreateNode(ctx, eec.driver, _spec); err != nil {
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

func (eec *EmailEventCreate) createSpec() (*EmailEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailEvent{config: eec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: emailevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailevent.FieldID,
			},
		}
	)
	if id, ok := eec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := eec.mutation.Module(); ok {
		_spec.SetField(emailevent.FieldModule, field.TypeEnum, value)
		_node.Module = value
	}
	if value, ok := eec.mutation.Action(); ok {
		_spec.SetField(emailevent.FieldAction, field.TypeEnum, value)
		_node.Action = value
	}
	if value, ok := eec.mutation.Name(); ok {
		_spec.SetField(emailevent.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := eec.mutation.CreatedAt(); ok {
		_spec.SetField(emailevent.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := eec.mutation.UpdatedAt(); ok {
		_spec.SetField(emailevent.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// EmailEventCreateBulk is the builder for creating many EmailEvent entities in bulk.
type EmailEventCreateBulk struct {
	config
	builders []*EmailEventCreate
}

// Save creates the EmailEvent entities in the database.
func (eecb *EmailEventCreateBulk) Save(ctx context.Context) ([]*EmailEvent, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eecb.builders))
	nodes := make([]*EmailEvent, len(eecb.builders))
	mutators := make([]Mutator, len(eecb.builders))
	for i := range eecb.builders {
		func(i int, root context.Context) {
			builder := eecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailEventMutation)
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
					_, err = mutators[i+1].Mutate(root, eecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eecb *EmailEventCreateBulk) SaveX(ctx context.Context) []*EmailEvent {
	v, err := eecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eecb *EmailEventCreateBulk) Exec(ctx context.Context) error {
	_, err := eecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eecb *EmailEventCreateBulk) ExecX(ctx context.Context) {
	if err := eecb.Exec(ctx); err != nil {
		panic(err)
	}
}
