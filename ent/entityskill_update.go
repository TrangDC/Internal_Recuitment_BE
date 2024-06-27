// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/candidate"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/ent/skill"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntitySkillUpdate is the builder for updating EntitySkill entities.
type EntitySkillUpdate struct {
	config
	hooks    []Hook
	mutation *EntitySkillMutation
}

// Where appends a list predicates to the EntitySkillUpdate builder.
func (esu *EntitySkillUpdate) Where(ps ...predicate.EntitySkill) *EntitySkillUpdate {
	esu.mutation.Where(ps...)
	return esu
}

// SetUpdatedAt sets the "updated_at" field.
func (esu *EntitySkillUpdate) SetUpdatedAt(t time.Time) *EntitySkillUpdate {
	esu.mutation.SetUpdatedAt(t)
	return esu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (esu *EntitySkillUpdate) SetNillableUpdatedAt(t *time.Time) *EntitySkillUpdate {
	if t != nil {
		esu.SetUpdatedAt(*t)
	}
	return esu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (esu *EntitySkillUpdate) ClearUpdatedAt() *EntitySkillUpdate {
	esu.mutation.ClearUpdatedAt()
	return esu
}

// SetDeletedAt sets the "deleted_at" field.
func (esu *EntitySkillUpdate) SetDeletedAt(t time.Time) *EntitySkillUpdate {
	esu.mutation.SetDeletedAt(t)
	return esu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (esu *EntitySkillUpdate) SetNillableDeletedAt(t *time.Time) *EntitySkillUpdate {
	if t != nil {
		esu.SetDeletedAt(*t)
	}
	return esu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (esu *EntitySkillUpdate) ClearDeletedAt() *EntitySkillUpdate {
	esu.mutation.ClearDeletedAt()
	return esu
}

// SetEntityType sets the "entity_type" field.
func (esu *EntitySkillUpdate) SetEntityType(et entityskill.EntityType) *EntitySkillUpdate {
	esu.mutation.SetEntityType(et)
	return esu
}

// SetEntityID sets the "entity_id" field.
func (esu *EntitySkillUpdate) SetEntityID(u uuid.UUID) *EntitySkillUpdate {
	esu.mutation.SetEntityID(u)
	return esu
}

// SetNillableEntityID sets the "entity_id" field if the given value is not nil.
func (esu *EntitySkillUpdate) SetNillableEntityID(u *uuid.UUID) *EntitySkillUpdate {
	if u != nil {
		esu.SetEntityID(*u)
	}
	return esu
}

// ClearEntityID clears the value of the "entity_id" field.
func (esu *EntitySkillUpdate) ClearEntityID() *EntitySkillUpdate {
	esu.mutation.ClearEntityID()
	return esu
}

// SetSkillID sets the "skill_id" field.
func (esu *EntitySkillUpdate) SetSkillID(u uuid.UUID) *EntitySkillUpdate {
	esu.mutation.SetSkillID(u)
	return esu
}

// SetNillableSkillID sets the "skill_id" field if the given value is not nil.
func (esu *EntitySkillUpdate) SetNillableSkillID(u *uuid.UUID) *EntitySkillUpdate {
	if u != nil {
		esu.SetSkillID(*u)
	}
	return esu
}

// ClearSkillID clears the value of the "skill_id" field.
func (esu *EntitySkillUpdate) ClearSkillID() *EntitySkillUpdate {
	esu.mutation.ClearSkillID()
	return esu
}

// SetOrderID sets the "order_id" field.
func (esu *EntitySkillUpdate) SetOrderID(i int) *EntitySkillUpdate {
	esu.mutation.ResetOrderID()
	esu.mutation.SetOrderID(i)
	return esu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (esu *EntitySkillUpdate) SetNillableOrderID(i *int) *EntitySkillUpdate {
	if i != nil {
		esu.SetOrderID(*i)
	}
	return esu
}

// AddOrderID adds i to the "order_id" field.
func (esu *EntitySkillUpdate) AddOrderID(i int) *EntitySkillUpdate {
	esu.mutation.AddOrderID(i)
	return esu
}

// ClearOrderID clears the value of the "order_id" field.
func (esu *EntitySkillUpdate) ClearOrderID() *EntitySkillUpdate {
	esu.mutation.ClearOrderID()
	return esu
}

// SetSkillEdgeID sets the "skill_edge" edge to the Skill entity by ID.
func (esu *EntitySkillUpdate) SetSkillEdgeID(id uuid.UUID) *EntitySkillUpdate {
	esu.mutation.SetSkillEdgeID(id)
	return esu
}

// SetNillableSkillEdgeID sets the "skill_edge" edge to the Skill entity by ID if the given value is not nil.
func (esu *EntitySkillUpdate) SetNillableSkillEdgeID(id *uuid.UUID) *EntitySkillUpdate {
	if id != nil {
		esu = esu.SetSkillEdgeID(*id)
	}
	return esu
}

// SetSkillEdge sets the "skill_edge" edge to the Skill entity.
func (esu *EntitySkillUpdate) SetSkillEdge(s *Skill) *EntitySkillUpdate {
	return esu.SetSkillEdgeID(s.ID)
}

// SetHiringJobEdgeID sets the "hiring_job_edge" edge to the HiringJob entity by ID.
func (esu *EntitySkillUpdate) SetHiringJobEdgeID(id uuid.UUID) *EntitySkillUpdate {
	esu.mutation.SetHiringJobEdgeID(id)
	return esu
}

// SetNillableHiringJobEdgeID sets the "hiring_job_edge" edge to the HiringJob entity by ID if the given value is not nil.
func (esu *EntitySkillUpdate) SetNillableHiringJobEdgeID(id *uuid.UUID) *EntitySkillUpdate {
	if id != nil {
		esu = esu.SetHiringJobEdgeID(*id)
	}
	return esu
}

// SetHiringJobEdge sets the "hiring_job_edge" edge to the HiringJob entity.
func (esu *EntitySkillUpdate) SetHiringJobEdge(h *HiringJob) *EntitySkillUpdate {
	return esu.SetHiringJobEdgeID(h.ID)
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (esu *EntitySkillUpdate) SetCandidateEdgeID(id uuid.UUID) *EntitySkillUpdate {
	esu.mutation.SetCandidateEdgeID(id)
	return esu
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (esu *EntitySkillUpdate) SetNillableCandidateEdgeID(id *uuid.UUID) *EntitySkillUpdate {
	if id != nil {
		esu = esu.SetCandidateEdgeID(*id)
	}
	return esu
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (esu *EntitySkillUpdate) SetCandidateEdge(c *Candidate) *EntitySkillUpdate {
	return esu.SetCandidateEdgeID(c.ID)
}

// Mutation returns the EntitySkillMutation object of the builder.
func (esu *EntitySkillUpdate) Mutation() *EntitySkillMutation {
	return esu.mutation
}

// ClearSkillEdge clears the "skill_edge" edge to the Skill entity.
func (esu *EntitySkillUpdate) ClearSkillEdge() *EntitySkillUpdate {
	esu.mutation.ClearSkillEdge()
	return esu
}

// ClearHiringJobEdge clears the "hiring_job_edge" edge to the HiringJob entity.
func (esu *EntitySkillUpdate) ClearHiringJobEdge() *EntitySkillUpdate {
	esu.mutation.ClearHiringJobEdge()
	return esu
}

// ClearCandidateEdge clears the "candidate_edge" edge to the Candidate entity.
func (esu *EntitySkillUpdate) ClearCandidateEdge() *EntitySkillUpdate {
	esu.mutation.ClearCandidateEdge()
	return esu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (esu *EntitySkillUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(esu.hooks) == 0 {
		if err = esu.check(); err != nil {
			return 0, err
		}
		affected, err = esu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntitySkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = esu.check(); err != nil {
				return 0, err
			}
			esu.mutation = mutation
			affected, err = esu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(esu.hooks) - 1; i >= 0; i-- {
			if esu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = esu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, esu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (esu *EntitySkillUpdate) SaveX(ctx context.Context) int {
	affected, err := esu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (esu *EntitySkillUpdate) Exec(ctx context.Context) error {
	_, err := esu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esu *EntitySkillUpdate) ExecX(ctx context.Context) {
	if err := esu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (esu *EntitySkillUpdate) check() error {
	if v, ok := esu.mutation.EntityType(); ok {
		if err := entityskill.EntityTypeValidator(v); err != nil {
			return &ValidationError{Name: "entity_type", err: fmt.Errorf(`ent: validator failed for field "EntitySkill.entity_type": %w`, err)}
		}
	}
	return nil
}

func (esu *EntitySkillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entityskill.Table,
			Columns: entityskill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entityskill.FieldID,
			},
		},
	}
	if ps := esu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := esu.mutation.UpdatedAt(); ok {
		_spec.SetField(entityskill.FieldUpdatedAt, field.TypeTime, value)
	}
	if esu.mutation.UpdatedAtCleared() {
		_spec.ClearField(entityskill.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := esu.mutation.DeletedAt(); ok {
		_spec.SetField(entityskill.FieldDeletedAt, field.TypeTime, value)
	}
	if esu.mutation.DeletedAtCleared() {
		_spec.ClearField(entityskill.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := esu.mutation.EntityType(); ok {
		_spec.SetField(entityskill.FieldEntityType, field.TypeEnum, value)
	}
	if value, ok := esu.mutation.OrderID(); ok {
		_spec.SetField(entityskill.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := esu.mutation.AddedOrderID(); ok {
		_spec.AddField(entityskill.FieldOrderID, field.TypeInt, value)
	}
	if esu.mutation.OrderIDCleared() {
		_spec.ClearField(entityskill.FieldOrderID, field.TypeInt)
	}
	if esu.mutation.SkillEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.SkillEdgeTable,
			Columns: []string{entityskill.SkillEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esu.mutation.SkillEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.SkillEdgeTable,
			Columns: []string{entityskill.SkillEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esu.mutation.HiringJobEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.HiringJobEdgeTable,
			Columns: []string{entityskill.HiringJobEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringjob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esu.mutation.HiringJobEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.HiringJobEdgeTable,
			Columns: []string{entityskill.HiringJobEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esu.mutation.CandidateEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.CandidateEdgeTable,
			Columns: []string{entityskill.CandidateEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esu.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.CandidateEdgeTable,
			Columns: []string{entityskill.CandidateEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, esu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entityskill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EntitySkillUpdateOne is the builder for updating a single EntitySkill entity.
type EntitySkillUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntitySkillMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (esuo *EntitySkillUpdateOne) SetUpdatedAt(t time.Time) *EntitySkillUpdateOne {
	esuo.mutation.SetUpdatedAt(t)
	return esuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (esuo *EntitySkillUpdateOne) SetNillableUpdatedAt(t *time.Time) *EntitySkillUpdateOne {
	if t != nil {
		esuo.SetUpdatedAt(*t)
	}
	return esuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (esuo *EntitySkillUpdateOne) ClearUpdatedAt() *EntitySkillUpdateOne {
	esuo.mutation.ClearUpdatedAt()
	return esuo
}

// SetDeletedAt sets the "deleted_at" field.
func (esuo *EntitySkillUpdateOne) SetDeletedAt(t time.Time) *EntitySkillUpdateOne {
	esuo.mutation.SetDeletedAt(t)
	return esuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (esuo *EntitySkillUpdateOne) SetNillableDeletedAt(t *time.Time) *EntitySkillUpdateOne {
	if t != nil {
		esuo.SetDeletedAt(*t)
	}
	return esuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (esuo *EntitySkillUpdateOne) ClearDeletedAt() *EntitySkillUpdateOne {
	esuo.mutation.ClearDeletedAt()
	return esuo
}

// SetEntityType sets the "entity_type" field.
func (esuo *EntitySkillUpdateOne) SetEntityType(et entityskill.EntityType) *EntitySkillUpdateOne {
	esuo.mutation.SetEntityType(et)
	return esuo
}

// SetEntityID sets the "entity_id" field.
func (esuo *EntitySkillUpdateOne) SetEntityID(u uuid.UUID) *EntitySkillUpdateOne {
	esuo.mutation.SetEntityID(u)
	return esuo
}

// SetNillableEntityID sets the "entity_id" field if the given value is not nil.
func (esuo *EntitySkillUpdateOne) SetNillableEntityID(u *uuid.UUID) *EntitySkillUpdateOne {
	if u != nil {
		esuo.SetEntityID(*u)
	}
	return esuo
}

// ClearEntityID clears the value of the "entity_id" field.
func (esuo *EntitySkillUpdateOne) ClearEntityID() *EntitySkillUpdateOne {
	esuo.mutation.ClearEntityID()
	return esuo
}

// SetSkillID sets the "skill_id" field.
func (esuo *EntitySkillUpdateOne) SetSkillID(u uuid.UUID) *EntitySkillUpdateOne {
	esuo.mutation.SetSkillID(u)
	return esuo
}

// SetNillableSkillID sets the "skill_id" field if the given value is not nil.
func (esuo *EntitySkillUpdateOne) SetNillableSkillID(u *uuid.UUID) *EntitySkillUpdateOne {
	if u != nil {
		esuo.SetSkillID(*u)
	}
	return esuo
}

// ClearSkillID clears the value of the "skill_id" field.
func (esuo *EntitySkillUpdateOne) ClearSkillID() *EntitySkillUpdateOne {
	esuo.mutation.ClearSkillID()
	return esuo
}

// SetOrderID sets the "order_id" field.
func (esuo *EntitySkillUpdateOne) SetOrderID(i int) *EntitySkillUpdateOne {
	esuo.mutation.ResetOrderID()
	esuo.mutation.SetOrderID(i)
	return esuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (esuo *EntitySkillUpdateOne) SetNillableOrderID(i *int) *EntitySkillUpdateOne {
	if i != nil {
		esuo.SetOrderID(*i)
	}
	return esuo
}

// AddOrderID adds i to the "order_id" field.
func (esuo *EntitySkillUpdateOne) AddOrderID(i int) *EntitySkillUpdateOne {
	esuo.mutation.AddOrderID(i)
	return esuo
}

// ClearOrderID clears the value of the "order_id" field.
func (esuo *EntitySkillUpdateOne) ClearOrderID() *EntitySkillUpdateOne {
	esuo.mutation.ClearOrderID()
	return esuo
}

// SetSkillEdgeID sets the "skill_edge" edge to the Skill entity by ID.
func (esuo *EntitySkillUpdateOne) SetSkillEdgeID(id uuid.UUID) *EntitySkillUpdateOne {
	esuo.mutation.SetSkillEdgeID(id)
	return esuo
}

// SetNillableSkillEdgeID sets the "skill_edge" edge to the Skill entity by ID if the given value is not nil.
func (esuo *EntitySkillUpdateOne) SetNillableSkillEdgeID(id *uuid.UUID) *EntitySkillUpdateOne {
	if id != nil {
		esuo = esuo.SetSkillEdgeID(*id)
	}
	return esuo
}

// SetSkillEdge sets the "skill_edge" edge to the Skill entity.
func (esuo *EntitySkillUpdateOne) SetSkillEdge(s *Skill) *EntitySkillUpdateOne {
	return esuo.SetSkillEdgeID(s.ID)
}

// SetHiringJobEdgeID sets the "hiring_job_edge" edge to the HiringJob entity by ID.
func (esuo *EntitySkillUpdateOne) SetHiringJobEdgeID(id uuid.UUID) *EntitySkillUpdateOne {
	esuo.mutation.SetHiringJobEdgeID(id)
	return esuo
}

// SetNillableHiringJobEdgeID sets the "hiring_job_edge" edge to the HiringJob entity by ID if the given value is not nil.
func (esuo *EntitySkillUpdateOne) SetNillableHiringJobEdgeID(id *uuid.UUID) *EntitySkillUpdateOne {
	if id != nil {
		esuo = esuo.SetHiringJobEdgeID(*id)
	}
	return esuo
}

// SetHiringJobEdge sets the "hiring_job_edge" edge to the HiringJob entity.
func (esuo *EntitySkillUpdateOne) SetHiringJobEdge(h *HiringJob) *EntitySkillUpdateOne {
	return esuo.SetHiringJobEdgeID(h.ID)
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (esuo *EntitySkillUpdateOne) SetCandidateEdgeID(id uuid.UUID) *EntitySkillUpdateOne {
	esuo.mutation.SetCandidateEdgeID(id)
	return esuo
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (esuo *EntitySkillUpdateOne) SetNillableCandidateEdgeID(id *uuid.UUID) *EntitySkillUpdateOne {
	if id != nil {
		esuo = esuo.SetCandidateEdgeID(*id)
	}
	return esuo
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (esuo *EntitySkillUpdateOne) SetCandidateEdge(c *Candidate) *EntitySkillUpdateOne {
	return esuo.SetCandidateEdgeID(c.ID)
}

// Mutation returns the EntitySkillMutation object of the builder.
func (esuo *EntitySkillUpdateOne) Mutation() *EntitySkillMutation {
	return esuo.mutation
}

// ClearSkillEdge clears the "skill_edge" edge to the Skill entity.
func (esuo *EntitySkillUpdateOne) ClearSkillEdge() *EntitySkillUpdateOne {
	esuo.mutation.ClearSkillEdge()
	return esuo
}

// ClearHiringJobEdge clears the "hiring_job_edge" edge to the HiringJob entity.
func (esuo *EntitySkillUpdateOne) ClearHiringJobEdge() *EntitySkillUpdateOne {
	esuo.mutation.ClearHiringJobEdge()
	return esuo
}

// ClearCandidateEdge clears the "candidate_edge" edge to the Candidate entity.
func (esuo *EntitySkillUpdateOne) ClearCandidateEdge() *EntitySkillUpdateOne {
	esuo.mutation.ClearCandidateEdge()
	return esuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (esuo *EntitySkillUpdateOne) Select(field string, fields ...string) *EntitySkillUpdateOne {
	esuo.fields = append([]string{field}, fields...)
	return esuo
}

// Save executes the query and returns the updated EntitySkill entity.
func (esuo *EntitySkillUpdateOne) Save(ctx context.Context) (*EntitySkill, error) {
	var (
		err  error
		node *EntitySkill
	)
	if len(esuo.hooks) == 0 {
		if err = esuo.check(); err != nil {
			return nil, err
		}
		node, err = esuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntitySkillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = esuo.check(); err != nil {
				return nil, err
			}
			esuo.mutation = mutation
			node, err = esuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(esuo.hooks) - 1; i >= 0; i-- {
			if esuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = esuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, esuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EntitySkill)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EntitySkillMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (esuo *EntitySkillUpdateOne) SaveX(ctx context.Context) *EntitySkill {
	node, err := esuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (esuo *EntitySkillUpdateOne) Exec(ctx context.Context) error {
	_, err := esuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esuo *EntitySkillUpdateOne) ExecX(ctx context.Context) {
	if err := esuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (esuo *EntitySkillUpdateOne) check() error {
	if v, ok := esuo.mutation.EntityType(); ok {
		if err := entityskill.EntityTypeValidator(v); err != nil {
			return &ValidationError{Name: "entity_type", err: fmt.Errorf(`ent: validator failed for field "EntitySkill.entity_type": %w`, err)}
		}
	}
	return nil
}

func (esuo *EntitySkillUpdateOne) sqlSave(ctx context.Context) (_node *EntitySkill, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entityskill.Table,
			Columns: entityskill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entityskill.FieldID,
			},
		},
	}
	id, ok := esuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EntitySkill.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := esuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entityskill.FieldID)
		for _, f := range fields {
			if !entityskill.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entityskill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := esuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := esuo.mutation.UpdatedAt(); ok {
		_spec.SetField(entityskill.FieldUpdatedAt, field.TypeTime, value)
	}
	if esuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(entityskill.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := esuo.mutation.DeletedAt(); ok {
		_spec.SetField(entityskill.FieldDeletedAt, field.TypeTime, value)
	}
	if esuo.mutation.DeletedAtCleared() {
		_spec.ClearField(entityskill.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := esuo.mutation.EntityType(); ok {
		_spec.SetField(entityskill.FieldEntityType, field.TypeEnum, value)
	}
	if value, ok := esuo.mutation.OrderID(); ok {
		_spec.SetField(entityskill.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := esuo.mutation.AddedOrderID(); ok {
		_spec.AddField(entityskill.FieldOrderID, field.TypeInt, value)
	}
	if esuo.mutation.OrderIDCleared() {
		_spec.ClearField(entityskill.FieldOrderID, field.TypeInt)
	}
	if esuo.mutation.SkillEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.SkillEdgeTable,
			Columns: []string{entityskill.SkillEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esuo.mutation.SkillEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.SkillEdgeTable,
			Columns: []string{entityskill.SkillEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: skill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esuo.mutation.HiringJobEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.HiringJobEdgeTable,
			Columns: []string{entityskill.HiringJobEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hiringjob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esuo.mutation.HiringJobEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.HiringJobEdgeTable,
			Columns: []string{entityskill.HiringJobEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if esuo.mutation.CandidateEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.CandidateEdgeTable,
			Columns: []string{entityskill.CandidateEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := esuo.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entityskill.CandidateEdgeTable,
			Columns: []string{entityskill.CandidateEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EntitySkill{config: esuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, esuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entityskill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}