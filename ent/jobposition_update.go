// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/hiringjob"
	"trec/ent/jobposition"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// JobPositionUpdate is the builder for updating JobPosition entities.
type JobPositionUpdate struct {
	config
	hooks    []Hook
	mutation *JobPositionMutation
}

// Where appends a list predicates to the JobPositionUpdate builder.
func (jpu *JobPositionUpdate) Where(ps ...predicate.JobPosition) *JobPositionUpdate {
	jpu.mutation.Where(ps...)
	return jpu
}

// SetUpdatedAt sets the "updated_at" field.
func (jpu *JobPositionUpdate) SetUpdatedAt(t time.Time) *JobPositionUpdate {
	jpu.mutation.SetUpdatedAt(t)
	return jpu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (jpu *JobPositionUpdate) SetNillableUpdatedAt(t *time.Time) *JobPositionUpdate {
	if t != nil {
		jpu.SetUpdatedAt(*t)
	}
	return jpu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (jpu *JobPositionUpdate) ClearUpdatedAt() *JobPositionUpdate {
	jpu.mutation.ClearUpdatedAt()
	return jpu
}

// SetDeletedAt sets the "deleted_at" field.
func (jpu *JobPositionUpdate) SetDeletedAt(t time.Time) *JobPositionUpdate {
	jpu.mutation.SetDeletedAt(t)
	return jpu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (jpu *JobPositionUpdate) SetNillableDeletedAt(t *time.Time) *JobPositionUpdate {
	if t != nil {
		jpu.SetDeletedAt(*t)
	}
	return jpu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (jpu *JobPositionUpdate) ClearDeletedAt() *JobPositionUpdate {
	jpu.mutation.ClearDeletedAt()
	return jpu
}

// SetName sets the "name" field.
func (jpu *JobPositionUpdate) SetName(s string) *JobPositionUpdate {
	jpu.mutation.SetName(s)
	return jpu
}

// SetDescription sets the "description" field.
func (jpu *JobPositionUpdate) SetDescription(s string) *JobPositionUpdate {
	jpu.mutation.SetDescription(s)
	return jpu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (jpu *JobPositionUpdate) SetNillableDescription(s *string) *JobPositionUpdate {
	if s != nil {
		jpu.SetDescription(*s)
	}
	return jpu
}

// ClearDescription clears the value of the "description" field.
func (jpu *JobPositionUpdate) ClearDescription() *JobPositionUpdate {
	jpu.mutation.ClearDescription()
	return jpu
}

// AddHiringJobPositionEdgeIDs adds the "hiring_job_position_edges" edge to the HiringJob entity by IDs.
func (jpu *JobPositionUpdate) AddHiringJobPositionEdgeIDs(ids ...uuid.UUID) *JobPositionUpdate {
	jpu.mutation.AddHiringJobPositionEdgeIDs(ids...)
	return jpu
}

// AddHiringJobPositionEdges adds the "hiring_job_position_edges" edges to the HiringJob entity.
func (jpu *JobPositionUpdate) AddHiringJobPositionEdges(h ...*HiringJob) *JobPositionUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return jpu.AddHiringJobPositionEdgeIDs(ids...)
}

// Mutation returns the JobPositionMutation object of the builder.
func (jpu *JobPositionUpdate) Mutation() *JobPositionMutation {
	return jpu.mutation
}

// ClearHiringJobPositionEdges clears all "hiring_job_position_edges" edges to the HiringJob entity.
func (jpu *JobPositionUpdate) ClearHiringJobPositionEdges() *JobPositionUpdate {
	jpu.mutation.ClearHiringJobPositionEdges()
	return jpu
}

// RemoveHiringJobPositionEdgeIDs removes the "hiring_job_position_edges" edge to HiringJob entities by IDs.
func (jpu *JobPositionUpdate) RemoveHiringJobPositionEdgeIDs(ids ...uuid.UUID) *JobPositionUpdate {
	jpu.mutation.RemoveHiringJobPositionEdgeIDs(ids...)
	return jpu
}

// RemoveHiringJobPositionEdges removes "hiring_job_position_edges" edges to HiringJob entities.
func (jpu *JobPositionUpdate) RemoveHiringJobPositionEdges(h ...*HiringJob) *JobPositionUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return jpu.RemoveHiringJobPositionEdgeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jpu *JobPositionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(jpu.hooks) == 0 {
		if err = jpu.check(); err != nil {
			return 0, err
		}
		affected, err = jpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = jpu.check(); err != nil {
				return 0, err
			}
			jpu.mutation = mutation
			affected, err = jpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(jpu.hooks) - 1; i >= 0; i-- {
			if jpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (jpu *JobPositionUpdate) SaveX(ctx context.Context) int {
	affected, err := jpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jpu *JobPositionUpdate) Exec(ctx context.Context) error {
	_, err := jpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jpu *JobPositionUpdate) ExecX(ctx context.Context) {
	if err := jpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jpu *JobPositionUpdate) check() error {
	if v, ok := jpu.mutation.Name(); ok {
		if err := jobposition.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "JobPosition.name": %w`, err)}
		}
	}
	if v, ok := jpu.mutation.Description(); ok {
		if err := jobposition.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "JobPosition.description": %w`, err)}
		}
	}
	return nil
}

func (jpu *JobPositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   jobposition.Table,
			Columns: jobposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: jobposition.FieldID,
			},
		},
	}
	if ps := jpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jpu.mutation.UpdatedAt(); ok {
		_spec.SetField(jobposition.FieldUpdatedAt, field.TypeTime, value)
	}
	if jpu.mutation.UpdatedAtCleared() {
		_spec.ClearField(jobposition.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := jpu.mutation.DeletedAt(); ok {
		_spec.SetField(jobposition.FieldDeletedAt, field.TypeTime, value)
	}
	if jpu.mutation.DeletedAtCleared() {
		_spec.ClearField(jobposition.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := jpu.mutation.Name(); ok {
		_spec.SetField(jobposition.FieldName, field.TypeString, value)
	}
	if value, ok := jpu.mutation.Description(); ok {
		_spec.SetField(jobposition.FieldDescription, field.TypeString, value)
	}
	if jpu.mutation.DescriptionCleared() {
		_spec.ClearField(jobposition.FieldDescription, field.TypeString)
	}
	if jpu.mutation.HiringJobPositionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.HiringJobPositionEdgesTable,
			Columns: []string{jobposition.HiringJobPositionEdgesColumn},
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
	if nodes := jpu.mutation.RemovedHiringJobPositionEdgesIDs(); len(nodes) > 0 && !jpu.mutation.HiringJobPositionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.HiringJobPositionEdgesTable,
			Columns: []string{jobposition.HiringJobPositionEdgesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jpu.mutation.HiringJobPositionEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.HiringJobPositionEdgesTable,
			Columns: []string{jobposition.HiringJobPositionEdgesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, jpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobposition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// JobPositionUpdateOne is the builder for updating a single JobPosition entity.
type JobPositionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobPositionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (jpuo *JobPositionUpdateOne) SetUpdatedAt(t time.Time) *JobPositionUpdateOne {
	jpuo.mutation.SetUpdatedAt(t)
	return jpuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (jpuo *JobPositionUpdateOne) SetNillableUpdatedAt(t *time.Time) *JobPositionUpdateOne {
	if t != nil {
		jpuo.SetUpdatedAt(*t)
	}
	return jpuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (jpuo *JobPositionUpdateOne) ClearUpdatedAt() *JobPositionUpdateOne {
	jpuo.mutation.ClearUpdatedAt()
	return jpuo
}

// SetDeletedAt sets the "deleted_at" field.
func (jpuo *JobPositionUpdateOne) SetDeletedAt(t time.Time) *JobPositionUpdateOne {
	jpuo.mutation.SetDeletedAt(t)
	return jpuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (jpuo *JobPositionUpdateOne) SetNillableDeletedAt(t *time.Time) *JobPositionUpdateOne {
	if t != nil {
		jpuo.SetDeletedAt(*t)
	}
	return jpuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (jpuo *JobPositionUpdateOne) ClearDeletedAt() *JobPositionUpdateOne {
	jpuo.mutation.ClearDeletedAt()
	return jpuo
}

// SetName sets the "name" field.
func (jpuo *JobPositionUpdateOne) SetName(s string) *JobPositionUpdateOne {
	jpuo.mutation.SetName(s)
	return jpuo
}

// SetDescription sets the "description" field.
func (jpuo *JobPositionUpdateOne) SetDescription(s string) *JobPositionUpdateOne {
	jpuo.mutation.SetDescription(s)
	return jpuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (jpuo *JobPositionUpdateOne) SetNillableDescription(s *string) *JobPositionUpdateOne {
	if s != nil {
		jpuo.SetDescription(*s)
	}
	return jpuo
}

// ClearDescription clears the value of the "description" field.
func (jpuo *JobPositionUpdateOne) ClearDescription() *JobPositionUpdateOne {
	jpuo.mutation.ClearDescription()
	return jpuo
}

// AddHiringJobPositionEdgeIDs adds the "hiring_job_position_edges" edge to the HiringJob entity by IDs.
func (jpuo *JobPositionUpdateOne) AddHiringJobPositionEdgeIDs(ids ...uuid.UUID) *JobPositionUpdateOne {
	jpuo.mutation.AddHiringJobPositionEdgeIDs(ids...)
	return jpuo
}

// AddHiringJobPositionEdges adds the "hiring_job_position_edges" edges to the HiringJob entity.
func (jpuo *JobPositionUpdateOne) AddHiringJobPositionEdges(h ...*HiringJob) *JobPositionUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return jpuo.AddHiringJobPositionEdgeIDs(ids...)
}

// Mutation returns the JobPositionMutation object of the builder.
func (jpuo *JobPositionUpdateOne) Mutation() *JobPositionMutation {
	return jpuo.mutation
}

// ClearHiringJobPositionEdges clears all "hiring_job_position_edges" edges to the HiringJob entity.
func (jpuo *JobPositionUpdateOne) ClearHiringJobPositionEdges() *JobPositionUpdateOne {
	jpuo.mutation.ClearHiringJobPositionEdges()
	return jpuo
}

// RemoveHiringJobPositionEdgeIDs removes the "hiring_job_position_edges" edge to HiringJob entities by IDs.
func (jpuo *JobPositionUpdateOne) RemoveHiringJobPositionEdgeIDs(ids ...uuid.UUID) *JobPositionUpdateOne {
	jpuo.mutation.RemoveHiringJobPositionEdgeIDs(ids...)
	return jpuo
}

// RemoveHiringJobPositionEdges removes "hiring_job_position_edges" edges to HiringJob entities.
func (jpuo *JobPositionUpdateOne) RemoveHiringJobPositionEdges(h ...*HiringJob) *JobPositionUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return jpuo.RemoveHiringJobPositionEdgeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jpuo *JobPositionUpdateOne) Select(field string, fields ...string) *JobPositionUpdateOne {
	jpuo.fields = append([]string{field}, fields...)
	return jpuo
}

// Save executes the query and returns the updated JobPosition entity.
func (jpuo *JobPositionUpdateOne) Save(ctx context.Context) (*JobPosition, error) {
	var (
		err  error
		node *JobPosition
	)
	if len(jpuo.hooks) == 0 {
		if err = jpuo.check(); err != nil {
			return nil, err
		}
		node, err = jpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = jpuo.check(); err != nil {
				return nil, err
			}
			jpuo.mutation = mutation
			node, err = jpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(jpuo.hooks) - 1; i >= 0; i-- {
			if jpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jpuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, jpuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*JobPosition)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from JobPositionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (jpuo *JobPositionUpdateOne) SaveX(ctx context.Context) *JobPosition {
	node, err := jpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jpuo *JobPositionUpdateOne) Exec(ctx context.Context) error {
	_, err := jpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jpuo *JobPositionUpdateOne) ExecX(ctx context.Context) {
	if err := jpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jpuo *JobPositionUpdateOne) check() error {
	if v, ok := jpuo.mutation.Name(); ok {
		if err := jobposition.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "JobPosition.name": %w`, err)}
		}
	}
	if v, ok := jpuo.mutation.Description(); ok {
		if err := jobposition.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "JobPosition.description": %w`, err)}
		}
	}
	return nil
}

func (jpuo *JobPositionUpdateOne) sqlSave(ctx context.Context) (_node *JobPosition, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   jobposition.Table,
			Columns: jobposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: jobposition.FieldID,
			},
		},
	}
	id, ok := jpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobPosition.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobposition.FieldID)
		for _, f := range fields {
			if !jobposition.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jobposition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(jobposition.FieldUpdatedAt, field.TypeTime, value)
	}
	if jpuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(jobposition.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := jpuo.mutation.DeletedAt(); ok {
		_spec.SetField(jobposition.FieldDeletedAt, field.TypeTime, value)
	}
	if jpuo.mutation.DeletedAtCleared() {
		_spec.ClearField(jobposition.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := jpuo.mutation.Name(); ok {
		_spec.SetField(jobposition.FieldName, field.TypeString, value)
	}
	if value, ok := jpuo.mutation.Description(); ok {
		_spec.SetField(jobposition.FieldDescription, field.TypeString, value)
	}
	if jpuo.mutation.DescriptionCleared() {
		_spec.ClearField(jobposition.FieldDescription, field.TypeString)
	}
	if jpuo.mutation.HiringJobPositionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.HiringJobPositionEdgesTable,
			Columns: []string{jobposition.HiringJobPositionEdgesColumn},
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
	if nodes := jpuo.mutation.RemovedHiringJobPositionEdgesIDs(); len(nodes) > 0 && !jpuo.mutation.HiringJobPositionEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.HiringJobPositionEdgesTable,
			Columns: []string{jobposition.HiringJobPositionEdgesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jpuo.mutation.HiringJobPositionEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobposition.HiringJobPositionEdgesTable,
			Columns: []string{jobposition.HiringJobPositionEdgesColumn},
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
	_node = &JobPosition{config: jpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobposition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
