// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/predicate"
	"trec/ent/skill"
	"trec/ent/skilltype"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SkillTypeUpdate is the builder for updating SkillType entities.
type SkillTypeUpdate struct {
	config
	hooks    []Hook
	mutation *SkillTypeMutation
}

// Where appends a list predicates to the SkillTypeUpdate builder.
func (stu *SkillTypeUpdate) Where(ps ...predicate.SkillType) *SkillTypeUpdate {
	stu.mutation.Where(ps...)
	return stu
}

// SetUpdatedAt sets the "updated_at" field.
func (stu *SkillTypeUpdate) SetUpdatedAt(t time.Time) *SkillTypeUpdate {
	stu.mutation.SetUpdatedAt(t)
	return stu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (stu *SkillTypeUpdate) SetNillableUpdatedAt(t *time.Time) *SkillTypeUpdate {
	if t != nil {
		stu.SetUpdatedAt(*t)
	}
	return stu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (stu *SkillTypeUpdate) ClearUpdatedAt() *SkillTypeUpdate {
	stu.mutation.ClearUpdatedAt()
	return stu
}

// SetDeletedAt sets the "deleted_at" field.
func (stu *SkillTypeUpdate) SetDeletedAt(t time.Time) *SkillTypeUpdate {
	stu.mutation.SetDeletedAt(t)
	return stu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (stu *SkillTypeUpdate) SetNillableDeletedAt(t *time.Time) *SkillTypeUpdate {
	if t != nil {
		stu.SetDeletedAt(*t)
	}
	return stu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (stu *SkillTypeUpdate) ClearDeletedAt() *SkillTypeUpdate {
	stu.mutation.ClearDeletedAt()
	return stu
}

// SetName sets the "name" field.
func (stu *SkillTypeUpdate) SetName(s string) *SkillTypeUpdate {
	stu.mutation.SetName(s)
	return stu
}

// SetDescription sets the "description" field.
func (stu *SkillTypeUpdate) SetDescription(s string) *SkillTypeUpdate {
	stu.mutation.SetDescription(s)
	return stu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (stu *SkillTypeUpdate) SetNillableDescription(s *string) *SkillTypeUpdate {
	if s != nil {
		stu.SetDescription(*s)
	}
	return stu
}

// ClearDescription clears the value of the "description" field.
func (stu *SkillTypeUpdate) ClearDescription() *SkillTypeUpdate {
	stu.mutation.ClearDescription()
	return stu
}

// AddSkillEdgeIDs adds the "skill_edges" edge to the Skill entity by IDs.
func (stu *SkillTypeUpdate) AddSkillEdgeIDs(ids ...uuid.UUID) *SkillTypeUpdate {
	stu.mutation.AddSkillEdgeIDs(ids...)
	return stu
}

// AddSkillEdges adds the "skill_edges" edges to the Skill entity.
func (stu *SkillTypeUpdate) AddSkillEdges(s ...*Skill) *SkillTypeUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.AddSkillEdgeIDs(ids...)
}

// Mutation returns the SkillTypeMutation object of the builder.
func (stu *SkillTypeUpdate) Mutation() *SkillTypeMutation {
	return stu.mutation
}

// ClearSkillEdges clears all "skill_edges" edges to the Skill entity.
func (stu *SkillTypeUpdate) ClearSkillEdges() *SkillTypeUpdate {
	stu.mutation.ClearSkillEdges()
	return stu
}

// RemoveSkillEdgeIDs removes the "skill_edges" edge to Skill entities by IDs.
func (stu *SkillTypeUpdate) RemoveSkillEdgeIDs(ids ...uuid.UUID) *SkillTypeUpdate {
	stu.mutation.RemoveSkillEdgeIDs(ids...)
	return stu
}

// RemoveSkillEdges removes "skill_edges" edges to Skill entities.
func (stu *SkillTypeUpdate) RemoveSkillEdges(s ...*Skill) *SkillTypeUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.RemoveSkillEdgeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (stu *SkillTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(stu.hooks) == 0 {
		if err = stu.check(); err != nil {
			return 0, err
		}
		affected, err = stu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkillTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stu.check(); err != nil {
				return 0, err
			}
			stu.mutation = mutation
			affected, err = stu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(stu.hooks) - 1; i >= 0; i-- {
			if stu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (stu *SkillTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *SkillTypeUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *SkillTypeUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stu *SkillTypeUpdate) check() error {
	if v, ok := stu.mutation.Name(); ok {
		if err := skilltype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SkillType.name": %w`, err)}
		}
	}
	if v, ok := stu.mutation.Description(); ok {
		if err := skilltype.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "SkillType.description": %w`, err)}
		}
	}
	return nil
}

func (stu *SkillTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skilltype.Table,
			Columns: skilltype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: skilltype.FieldID,
			},
		},
	}
	if ps := stu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.UpdatedAt(); ok {
		_spec.SetField(skilltype.FieldUpdatedAt, field.TypeTime, value)
	}
	if stu.mutation.UpdatedAtCleared() {
		_spec.ClearField(skilltype.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := stu.mutation.DeletedAt(); ok {
		_spec.SetField(skilltype.FieldDeletedAt, field.TypeTime, value)
	}
	if stu.mutation.DeletedAtCleared() {
		_spec.ClearField(skilltype.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := stu.mutation.Name(); ok {
		_spec.SetField(skilltype.FieldName, field.TypeString, value)
	}
	if value, ok := stu.mutation.Description(); ok {
		_spec.SetField(skilltype.FieldDescription, field.TypeString, value)
	}
	if stu.mutation.DescriptionCleared() {
		_spec.ClearField(skilltype.FieldDescription, field.TypeString)
	}
	if stu.mutation.SkillEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltype.SkillEdgesTable,
			Columns: []string{skilltype.SkillEdgesColumn},
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
	if nodes := stu.mutation.RemovedSkillEdgesIDs(); len(nodes) > 0 && !stu.mutation.SkillEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltype.SkillEdgesTable,
			Columns: []string{skilltype.SkillEdgesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.SkillEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltype.SkillEdgesTable,
			Columns: []string{skilltype.SkillEdgesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skilltype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SkillTypeUpdateOne is the builder for updating a single SkillType entity.
type SkillTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SkillTypeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (stuo *SkillTypeUpdateOne) SetUpdatedAt(t time.Time) *SkillTypeUpdateOne {
	stuo.mutation.SetUpdatedAt(t)
	return stuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (stuo *SkillTypeUpdateOne) SetNillableUpdatedAt(t *time.Time) *SkillTypeUpdateOne {
	if t != nil {
		stuo.SetUpdatedAt(*t)
	}
	return stuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (stuo *SkillTypeUpdateOne) ClearUpdatedAt() *SkillTypeUpdateOne {
	stuo.mutation.ClearUpdatedAt()
	return stuo
}

// SetDeletedAt sets the "deleted_at" field.
func (stuo *SkillTypeUpdateOne) SetDeletedAt(t time.Time) *SkillTypeUpdateOne {
	stuo.mutation.SetDeletedAt(t)
	return stuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (stuo *SkillTypeUpdateOne) SetNillableDeletedAt(t *time.Time) *SkillTypeUpdateOne {
	if t != nil {
		stuo.SetDeletedAt(*t)
	}
	return stuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (stuo *SkillTypeUpdateOne) ClearDeletedAt() *SkillTypeUpdateOne {
	stuo.mutation.ClearDeletedAt()
	return stuo
}

// SetName sets the "name" field.
func (stuo *SkillTypeUpdateOne) SetName(s string) *SkillTypeUpdateOne {
	stuo.mutation.SetName(s)
	return stuo
}

// SetDescription sets the "description" field.
func (stuo *SkillTypeUpdateOne) SetDescription(s string) *SkillTypeUpdateOne {
	stuo.mutation.SetDescription(s)
	return stuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (stuo *SkillTypeUpdateOne) SetNillableDescription(s *string) *SkillTypeUpdateOne {
	if s != nil {
		stuo.SetDescription(*s)
	}
	return stuo
}

// ClearDescription clears the value of the "description" field.
func (stuo *SkillTypeUpdateOne) ClearDescription() *SkillTypeUpdateOne {
	stuo.mutation.ClearDescription()
	return stuo
}

// AddSkillEdgeIDs adds the "skill_edges" edge to the Skill entity by IDs.
func (stuo *SkillTypeUpdateOne) AddSkillEdgeIDs(ids ...uuid.UUID) *SkillTypeUpdateOne {
	stuo.mutation.AddSkillEdgeIDs(ids...)
	return stuo
}

// AddSkillEdges adds the "skill_edges" edges to the Skill entity.
func (stuo *SkillTypeUpdateOne) AddSkillEdges(s ...*Skill) *SkillTypeUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.AddSkillEdgeIDs(ids...)
}

// Mutation returns the SkillTypeMutation object of the builder.
func (stuo *SkillTypeUpdateOne) Mutation() *SkillTypeMutation {
	return stuo.mutation
}

// ClearSkillEdges clears all "skill_edges" edges to the Skill entity.
func (stuo *SkillTypeUpdateOne) ClearSkillEdges() *SkillTypeUpdateOne {
	stuo.mutation.ClearSkillEdges()
	return stuo
}

// RemoveSkillEdgeIDs removes the "skill_edges" edge to Skill entities by IDs.
func (stuo *SkillTypeUpdateOne) RemoveSkillEdgeIDs(ids ...uuid.UUID) *SkillTypeUpdateOne {
	stuo.mutation.RemoveSkillEdgeIDs(ids...)
	return stuo
}

// RemoveSkillEdges removes "skill_edges" edges to Skill entities.
func (stuo *SkillTypeUpdateOne) RemoveSkillEdges(s ...*Skill) *SkillTypeUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.RemoveSkillEdgeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (stuo *SkillTypeUpdateOne) Select(field string, fields ...string) *SkillTypeUpdateOne {
	stuo.fields = append([]string{field}, fields...)
	return stuo
}

// Save executes the query and returns the updated SkillType entity.
func (stuo *SkillTypeUpdateOne) Save(ctx context.Context) (*SkillType, error) {
	var (
		err  error
		node *SkillType
	)
	if len(stuo.hooks) == 0 {
		if err = stuo.check(); err != nil {
			return nil, err
		}
		node, err = stuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkillTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stuo.check(); err != nil {
				return nil, err
			}
			stuo.mutation = mutation
			node, err = stuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(stuo.hooks) - 1; i >= 0; i-- {
			if stuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, stuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (stuo *SkillTypeUpdateOne) SaveX(ctx context.Context) *SkillType {
	node, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (stuo *SkillTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *SkillTypeUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stuo *SkillTypeUpdateOne) check() error {
	if v, ok := stuo.mutation.Name(); ok {
		if err := skilltype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SkillType.name": %w`, err)}
		}
	}
	if v, ok := stuo.mutation.Description(); ok {
		if err := skilltype.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "SkillType.description": %w`, err)}
		}
	}
	return nil
}

func (stuo *SkillTypeUpdateOne) sqlSave(ctx context.Context) (_node *SkillType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skilltype.Table,
			Columns: skilltype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: skilltype.FieldID,
			},
		},
	}
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SkillType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := stuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skilltype.FieldID)
		for _, f := range fields {
			if !skilltype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != skilltype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := stuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stuo.mutation.UpdatedAt(); ok {
		_spec.SetField(skilltype.FieldUpdatedAt, field.TypeTime, value)
	}
	if stuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(skilltype.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := stuo.mutation.DeletedAt(); ok {
		_spec.SetField(skilltype.FieldDeletedAt, field.TypeTime, value)
	}
	if stuo.mutation.DeletedAtCleared() {
		_spec.ClearField(skilltype.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := stuo.mutation.Name(); ok {
		_spec.SetField(skilltype.FieldName, field.TypeString, value)
	}
	if value, ok := stuo.mutation.Description(); ok {
		_spec.SetField(skilltype.FieldDescription, field.TypeString, value)
	}
	if stuo.mutation.DescriptionCleared() {
		_spec.ClearField(skilltype.FieldDescription, field.TypeString)
	}
	if stuo.mutation.SkillEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltype.SkillEdgesTable,
			Columns: []string{skilltype.SkillEdgesColumn},
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
	if nodes := stuo.mutation.RemovedSkillEdgesIDs(); len(nodes) > 0 && !stuo.mutation.SkillEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltype.SkillEdgesTable,
			Columns: []string{skilltype.SkillEdgesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.SkillEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skilltype.SkillEdgesTable,
			Columns: []string{skilltype.SkillEdgesColumn},
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
	_node = &SkillType{config: stuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skilltype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
