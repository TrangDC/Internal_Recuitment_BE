// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/entityskill"
	"trec/ent/skill"
	"trec/ent/skilltype"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SkillCreate is the builder for creating a Skill entity.
type SkillCreate struct {
	config
	mutation *SkillMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SkillCreate) SetCreatedAt(t time.Time) *SkillCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SkillCreate) SetNillableCreatedAt(t *time.Time) *SkillCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SkillCreate) SetUpdatedAt(t time.Time) *SkillCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SkillCreate) SetNillableUpdatedAt(t *time.Time) *SkillCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *SkillCreate) SetDeletedAt(t time.Time) *SkillCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *SkillCreate) SetNillableDeletedAt(t *time.Time) *SkillCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SkillCreate) SetName(s string) *SkillCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SkillCreate) SetDescription(s string) *SkillCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *SkillCreate) SetNillableDescription(s *string) *SkillCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetSkillTypeID sets the "skill_type_id" field.
func (sc *SkillCreate) SetSkillTypeID(u uuid.UUID) *SkillCreate {
	sc.mutation.SetSkillTypeID(u)
	return sc
}

// SetNillableSkillTypeID sets the "skill_type_id" field if the given value is not nil.
func (sc *SkillCreate) SetNillableSkillTypeID(u *uuid.UUID) *SkillCreate {
	if u != nil {
		sc.SetSkillTypeID(*u)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SkillCreate) SetID(u uuid.UUID) *SkillCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetSkillTypeEdgeID sets the "skill_type_edge" edge to the SkillType entity by ID.
func (sc *SkillCreate) SetSkillTypeEdgeID(id uuid.UUID) *SkillCreate {
	sc.mutation.SetSkillTypeEdgeID(id)
	return sc
}

// SetNillableSkillTypeEdgeID sets the "skill_type_edge" edge to the SkillType entity by ID if the given value is not nil.
func (sc *SkillCreate) SetNillableSkillTypeEdgeID(id *uuid.UUID) *SkillCreate {
	if id != nil {
		sc = sc.SetSkillTypeEdgeID(*id)
	}
	return sc
}

// SetSkillTypeEdge sets the "skill_type_edge" edge to the SkillType entity.
func (sc *SkillCreate) SetSkillTypeEdge(s *SkillType) *SkillCreate {
	return sc.SetSkillTypeEdgeID(s.ID)
}

// AddEntitySkillEdgeIDs adds the "entity_skill_edges" edge to the EntitySkill entity by IDs.
func (sc *SkillCreate) AddEntitySkillEdgeIDs(ids ...uuid.UUID) *SkillCreate {
	sc.mutation.AddEntitySkillEdgeIDs(ids...)
	return sc
}

// AddEntitySkillEdges adds the "entity_skill_edges" edges to the EntitySkill entity.
func (sc *SkillCreate) AddEntitySkillEdges(e ...*EntitySkill) *SkillCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return sc.AddEntitySkillEdgeIDs(ids...)
}

// Mutation returns the SkillMutation object of the builder.
func (sc *SkillCreate) Mutation() *SkillMutation {
	return sc.mutation
}

// Save creates the Skill in the database.
func (sc *SkillCreate) Save(ctx context.Context) (*Skill, error) {
	var (
		err  error
		node *Skill
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Skill)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SkillMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SkillCreate) SaveX(ctx context.Context) *Skill {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SkillCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SkillCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SkillCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := skill.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SkillCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Skill.created_at"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Skill.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := skill.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Skill.name": %w`, err)}
		}
	}
	if v, ok := sc.mutation.Description(); ok {
		if err := skill.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Skill.description": %w`, err)}
		}
	}
	return nil
}

func (sc *SkillCreate) sqlSave(ctx context.Context) (*Skill, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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

func (sc *SkillCreate) createSpec() (*Skill, *sqlgraph.CreateSpec) {
	var (
		_node = &Skill{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: skill.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: skill.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(skill.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(skill.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.SetField(skill.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(skill.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(skill.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := sc.mutation.SkillTypeEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   skill.SkillTypeEdgeTable,
			Columns: []string{skill.SkillTypeEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skilltype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SkillTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.EntitySkillEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skill.EntitySkillEdgesTable,
			Columns: []string{skill.EntitySkillEdgesColumn},
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
	return _node, _spec
}

// SkillCreateBulk is the builder for creating many Skill entities in bulk.
type SkillCreateBulk struct {
	config
	builders []*SkillCreate
}

// Save creates the Skill entities in the database.
func (scb *SkillCreateBulk) Save(ctx context.Context) ([]*Skill, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Skill, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkillMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SkillCreateBulk) SaveX(ctx context.Context) []*Skill {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SkillCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SkillCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
