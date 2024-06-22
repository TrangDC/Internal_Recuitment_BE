// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/skilltype"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SkillTypeCreate is the builder for creating a SkillType entity.
type SkillTypeCreate struct {
	config
	mutation *SkillTypeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (stc *SkillTypeCreate) SetCreatedAt(t time.Time) *SkillTypeCreate {
	stc.mutation.SetCreatedAt(t)
	return stc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (stc *SkillTypeCreate) SetNillableCreatedAt(t *time.Time) *SkillTypeCreate {
	if t != nil {
		stc.SetCreatedAt(*t)
	}
	return stc
}

// SetUpdatedAt sets the "updated_at" field.
func (stc *SkillTypeCreate) SetUpdatedAt(t time.Time) *SkillTypeCreate {
	stc.mutation.SetUpdatedAt(t)
	return stc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (stc *SkillTypeCreate) SetNillableUpdatedAt(t *time.Time) *SkillTypeCreate {
	if t != nil {
		stc.SetUpdatedAt(*t)
	}
	return stc
}

// SetDeletedAt sets the "deleted_at" field.
func (stc *SkillTypeCreate) SetDeletedAt(t time.Time) *SkillTypeCreate {
	stc.mutation.SetDeletedAt(t)
	return stc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (stc *SkillTypeCreate) SetNillableDeletedAt(t *time.Time) *SkillTypeCreate {
	if t != nil {
		stc.SetDeletedAt(*t)
	}
	return stc
}

// SetName sets the "name" field.
func (stc *SkillTypeCreate) SetName(s string) *SkillTypeCreate {
	stc.mutation.SetName(s)
	return stc
}

// SetDescription sets the "description" field.
func (stc *SkillTypeCreate) SetDescription(s string) *SkillTypeCreate {
	stc.mutation.SetDescription(s)
	return stc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (stc *SkillTypeCreate) SetNillableDescription(s *string) *SkillTypeCreate {
	if s != nil {
		stc.SetDescription(*s)
	}
	return stc
}

// SetID sets the "id" field.
func (stc *SkillTypeCreate) SetID(u uuid.UUID) *SkillTypeCreate {
	stc.mutation.SetID(u)
	return stc
}

// Mutation returns the SkillTypeMutation object of the builder.
func (stc *SkillTypeCreate) Mutation() *SkillTypeMutation {
	return stc.mutation
}

// Save creates the SkillType in the database.
func (stc *SkillTypeCreate) Save(ctx context.Context) (*SkillType, error) {
	var (
		err  error
		node *SkillType
	)
	stc.defaults()
	if len(stc.hooks) == 0 {
		if err = stc.check(); err != nil {
			return nil, err
		}
		node, err = stc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkillTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stc.check(); err != nil {
				return nil, err
			}
			stc.mutation = mutation
			if node, err = stc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(stc.hooks) - 1; i >= 0; i-- {
			if stc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, stc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SkillType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SkillTypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (stc *SkillTypeCreate) SaveX(ctx context.Context) *SkillType {
	v, err := stc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stc *SkillTypeCreate) Exec(ctx context.Context) error {
	_, err := stc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stc *SkillTypeCreate) ExecX(ctx context.Context) {
	if err := stc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stc *SkillTypeCreate) defaults() {
	if _, ok := stc.mutation.CreatedAt(); !ok {
		v := skilltype.DefaultCreatedAt()
		stc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stc *SkillTypeCreate) check() error {
	if _, ok := stc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SkillType.created_at"`)}
	}
	if _, ok := stc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "SkillType.name"`)}
	}
	if v, ok := stc.mutation.Name(); ok {
		if err := skilltype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SkillType.name": %w`, err)}
		}
	}
	if v, ok := stc.mutation.Description(); ok {
		if err := skilltype.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "SkillType.description": %w`, err)}
		}
	}
	return nil
}

func (stc *SkillTypeCreate) sqlSave(ctx context.Context) (*SkillType, error) {
	_node, _spec := stc.createSpec()
	if err := sqlgraph.CreateNode(ctx, stc.driver, _spec); err != nil {
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

func (stc *SkillTypeCreate) createSpec() (*SkillType, *sqlgraph.CreateSpec) {
	var (
		_node = &SkillType{config: stc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: skilltype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: skilltype.FieldID,
			},
		}
	)
	if id, ok := stc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := stc.mutation.CreatedAt(); ok {
		_spec.SetField(skilltype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := stc.mutation.UpdatedAt(); ok {
		_spec.SetField(skilltype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := stc.mutation.DeletedAt(); ok {
		_spec.SetField(skilltype.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := stc.mutation.Name(); ok {
		_spec.SetField(skilltype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := stc.mutation.Description(); ok {
		_spec.SetField(skilltype.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	return _node, _spec
}

// SkillTypeCreateBulk is the builder for creating many SkillType entities in bulk.
type SkillTypeCreateBulk struct {
	config
	builders []*SkillTypeCreate
}

// Save creates the SkillType entities in the database.
func (stcb *SkillTypeCreateBulk) Save(ctx context.Context) ([]*SkillType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(stcb.builders))
	nodes := make([]*SkillType, len(stcb.builders))
	mutators := make([]Mutator, len(stcb.builders))
	for i := range stcb.builders {
		func(i int, root context.Context) {
			builder := stcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkillTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, stcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, stcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, stcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (stcb *SkillTypeCreateBulk) SaveX(ctx context.Context) []*SkillType {
	v, err := stcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stcb *SkillTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := stcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stcb *SkillTypeCreateBulk) ExecX(ctx context.Context) {
	if err := stcb.Exec(ctx); err != nil {
		panic(err)
	}
}