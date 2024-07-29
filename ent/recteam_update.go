// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/predicate"
	"trec/ent/recteam"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RecTeamUpdate is the builder for updating RecTeam entities.
type RecTeamUpdate struct {
	config
	hooks    []Hook
	mutation *RecTeamMutation
}

// Where appends a list predicates to the RecTeamUpdate builder.
func (rtu *RecTeamUpdate) Where(ps ...predicate.RecTeam) *RecTeamUpdate {
	rtu.mutation.Where(ps...)
	return rtu
}

// SetUpdatedAt sets the "updated_at" field.
func (rtu *RecTeamUpdate) SetUpdatedAt(t time.Time) *RecTeamUpdate {
	rtu.mutation.SetUpdatedAt(t)
	return rtu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rtu *RecTeamUpdate) SetNillableUpdatedAt(t *time.Time) *RecTeamUpdate {
	if t != nil {
		rtu.SetUpdatedAt(*t)
	}
	return rtu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (rtu *RecTeamUpdate) ClearUpdatedAt() *RecTeamUpdate {
	rtu.mutation.ClearUpdatedAt()
	return rtu
}

// SetDeletedAt sets the "deleted_at" field.
func (rtu *RecTeamUpdate) SetDeletedAt(t time.Time) *RecTeamUpdate {
	rtu.mutation.SetDeletedAt(t)
	return rtu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rtu *RecTeamUpdate) SetNillableDeletedAt(t *time.Time) *RecTeamUpdate {
	if t != nil {
		rtu.SetDeletedAt(*t)
	}
	return rtu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (rtu *RecTeamUpdate) ClearDeletedAt() *RecTeamUpdate {
	rtu.mutation.ClearDeletedAt()
	return rtu
}

// SetName sets the "name" field.
func (rtu *RecTeamUpdate) SetName(s string) *RecTeamUpdate {
	rtu.mutation.SetName(s)
	return rtu
}

// SetDescription sets the "description" field.
func (rtu *RecTeamUpdate) SetDescription(s string) *RecTeamUpdate {
	rtu.mutation.SetDescription(s)
	return rtu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rtu *RecTeamUpdate) SetNillableDescription(s *string) *RecTeamUpdate {
	if s != nil {
		rtu.SetDescription(*s)
	}
	return rtu
}

// ClearDescription clears the value of the "description" field.
func (rtu *RecTeamUpdate) ClearDescription() *RecTeamUpdate {
	rtu.mutation.ClearDescription()
	return rtu
}

// SetLeaderID sets the "leader_id" field.
func (rtu *RecTeamUpdate) SetLeaderID(u uuid.UUID) *RecTeamUpdate {
	rtu.mutation.SetLeaderID(u)
	return rtu
}

// SetRecLeaderEdgeID sets the "rec_leader_edge" edge to the User entity by ID.
func (rtu *RecTeamUpdate) SetRecLeaderEdgeID(id uuid.UUID) *RecTeamUpdate {
	rtu.mutation.SetRecLeaderEdgeID(id)
	return rtu
}

// SetRecLeaderEdge sets the "rec_leader_edge" edge to the User entity.
func (rtu *RecTeamUpdate) SetRecLeaderEdge(u *User) *RecTeamUpdate {
	return rtu.SetRecLeaderEdgeID(u.ID)
}

// AddRecMemberEdgeIDs adds the "rec_member_edges" edge to the User entity by IDs.
func (rtu *RecTeamUpdate) AddRecMemberEdgeIDs(ids ...uuid.UUID) *RecTeamUpdate {
	rtu.mutation.AddRecMemberEdgeIDs(ids...)
	return rtu
}

// AddRecMemberEdges adds the "rec_member_edges" edges to the User entity.
func (rtu *RecTeamUpdate) AddRecMemberEdges(u ...*User) *RecTeamUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rtu.AddRecMemberEdgeIDs(ids...)
}

// Mutation returns the RecTeamMutation object of the builder.
func (rtu *RecTeamUpdate) Mutation() *RecTeamMutation {
	return rtu.mutation
}

// ClearRecLeaderEdge clears the "rec_leader_edge" edge to the User entity.
func (rtu *RecTeamUpdate) ClearRecLeaderEdge() *RecTeamUpdate {
	rtu.mutation.ClearRecLeaderEdge()
	return rtu
}

// ClearRecMemberEdges clears all "rec_member_edges" edges to the User entity.
func (rtu *RecTeamUpdate) ClearRecMemberEdges() *RecTeamUpdate {
	rtu.mutation.ClearRecMemberEdges()
	return rtu
}

// RemoveRecMemberEdgeIDs removes the "rec_member_edges" edge to User entities by IDs.
func (rtu *RecTeamUpdate) RemoveRecMemberEdgeIDs(ids ...uuid.UUID) *RecTeamUpdate {
	rtu.mutation.RemoveRecMemberEdgeIDs(ids...)
	return rtu
}

// RemoveRecMemberEdges removes "rec_member_edges" edges to User entities.
func (rtu *RecTeamUpdate) RemoveRecMemberEdges(u ...*User) *RecTeamUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rtu.RemoveRecMemberEdgeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtu *RecTeamUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rtu.hooks) == 0 {
		if err = rtu.check(); err != nil {
			return 0, err
		}
		affected, err = rtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecTeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtu.check(); err != nil {
				return 0, err
			}
			rtu.mutation = mutation
			affected, err = rtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rtu.hooks) - 1; i >= 0; i-- {
			if rtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtu *RecTeamUpdate) SaveX(ctx context.Context) int {
	affected, err := rtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtu *RecTeamUpdate) Exec(ctx context.Context) error {
	_, err := rtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtu *RecTeamUpdate) ExecX(ctx context.Context) {
	if err := rtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtu *RecTeamUpdate) check() error {
	if v, ok := rtu.mutation.Name(); ok {
		if err := recteam.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "RecTeam.name": %w`, err)}
		}
	}
	if v, ok := rtu.mutation.Description(); ok {
		if err := recteam.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "RecTeam.description": %w`, err)}
		}
	}
	if _, ok := rtu.mutation.RecLeaderEdgeID(); rtu.mutation.RecLeaderEdgeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RecTeam.rec_leader_edge"`)
	}
	return nil
}

func (rtu *RecTeamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recteam.Table,
			Columns: recteam.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: recteam.FieldID,
			},
		},
	}
	if ps := rtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtu.mutation.UpdatedAt(); ok {
		_spec.SetField(recteam.FieldUpdatedAt, field.TypeTime, value)
	}
	if rtu.mutation.UpdatedAtCleared() {
		_spec.ClearField(recteam.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := rtu.mutation.DeletedAt(); ok {
		_spec.SetField(recteam.FieldDeletedAt, field.TypeTime, value)
	}
	if rtu.mutation.DeletedAtCleared() {
		_spec.ClearField(recteam.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := rtu.mutation.Name(); ok {
		_spec.SetField(recteam.FieldName, field.TypeString, value)
	}
	if value, ok := rtu.mutation.Description(); ok {
		_spec.SetField(recteam.FieldDescription, field.TypeString, value)
	}
	if rtu.mutation.DescriptionCleared() {
		_spec.ClearField(recteam.FieldDescription, field.TypeString)
	}
	if rtu.mutation.RecLeaderEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recteam.RecLeaderEdgeTable,
			Columns: []string{recteam.RecLeaderEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.RecLeaderEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recteam.RecLeaderEdgeTable,
			Columns: []string{recteam.RecLeaderEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rtu.mutation.RecMemberEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recteam.RecMemberEdgesTable,
			Columns: []string{recteam.RecMemberEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.RemovedRecMemberEdgesIDs(); len(nodes) > 0 && !rtu.mutation.RecMemberEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recteam.RecMemberEdgesTable,
			Columns: []string{recteam.RecMemberEdgesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.RecMemberEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recteam.RecMemberEdgesTable,
			Columns: []string{recteam.RecMemberEdgesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recteam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RecTeamUpdateOne is the builder for updating a single RecTeam entity.
type RecTeamUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecTeamMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (rtuo *RecTeamUpdateOne) SetUpdatedAt(t time.Time) *RecTeamUpdateOne {
	rtuo.mutation.SetUpdatedAt(t)
	return rtuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rtuo *RecTeamUpdateOne) SetNillableUpdatedAt(t *time.Time) *RecTeamUpdateOne {
	if t != nil {
		rtuo.SetUpdatedAt(*t)
	}
	return rtuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (rtuo *RecTeamUpdateOne) ClearUpdatedAt() *RecTeamUpdateOne {
	rtuo.mutation.ClearUpdatedAt()
	return rtuo
}

// SetDeletedAt sets the "deleted_at" field.
func (rtuo *RecTeamUpdateOne) SetDeletedAt(t time.Time) *RecTeamUpdateOne {
	rtuo.mutation.SetDeletedAt(t)
	return rtuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rtuo *RecTeamUpdateOne) SetNillableDeletedAt(t *time.Time) *RecTeamUpdateOne {
	if t != nil {
		rtuo.SetDeletedAt(*t)
	}
	return rtuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (rtuo *RecTeamUpdateOne) ClearDeletedAt() *RecTeamUpdateOne {
	rtuo.mutation.ClearDeletedAt()
	return rtuo
}

// SetName sets the "name" field.
func (rtuo *RecTeamUpdateOne) SetName(s string) *RecTeamUpdateOne {
	rtuo.mutation.SetName(s)
	return rtuo
}

// SetDescription sets the "description" field.
func (rtuo *RecTeamUpdateOne) SetDescription(s string) *RecTeamUpdateOne {
	rtuo.mutation.SetDescription(s)
	return rtuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rtuo *RecTeamUpdateOne) SetNillableDescription(s *string) *RecTeamUpdateOne {
	if s != nil {
		rtuo.SetDescription(*s)
	}
	return rtuo
}

// ClearDescription clears the value of the "description" field.
func (rtuo *RecTeamUpdateOne) ClearDescription() *RecTeamUpdateOne {
	rtuo.mutation.ClearDescription()
	return rtuo
}

// SetLeaderID sets the "leader_id" field.
func (rtuo *RecTeamUpdateOne) SetLeaderID(u uuid.UUID) *RecTeamUpdateOne {
	rtuo.mutation.SetLeaderID(u)
	return rtuo
}

// SetRecLeaderEdgeID sets the "rec_leader_edge" edge to the User entity by ID.
func (rtuo *RecTeamUpdateOne) SetRecLeaderEdgeID(id uuid.UUID) *RecTeamUpdateOne {
	rtuo.mutation.SetRecLeaderEdgeID(id)
	return rtuo
}

// SetRecLeaderEdge sets the "rec_leader_edge" edge to the User entity.
func (rtuo *RecTeamUpdateOne) SetRecLeaderEdge(u *User) *RecTeamUpdateOne {
	return rtuo.SetRecLeaderEdgeID(u.ID)
}

// AddRecMemberEdgeIDs adds the "rec_member_edges" edge to the User entity by IDs.
func (rtuo *RecTeamUpdateOne) AddRecMemberEdgeIDs(ids ...uuid.UUID) *RecTeamUpdateOne {
	rtuo.mutation.AddRecMemberEdgeIDs(ids...)
	return rtuo
}

// AddRecMemberEdges adds the "rec_member_edges" edges to the User entity.
func (rtuo *RecTeamUpdateOne) AddRecMemberEdges(u ...*User) *RecTeamUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rtuo.AddRecMemberEdgeIDs(ids...)
}

// Mutation returns the RecTeamMutation object of the builder.
func (rtuo *RecTeamUpdateOne) Mutation() *RecTeamMutation {
	return rtuo.mutation
}

// ClearRecLeaderEdge clears the "rec_leader_edge" edge to the User entity.
func (rtuo *RecTeamUpdateOne) ClearRecLeaderEdge() *RecTeamUpdateOne {
	rtuo.mutation.ClearRecLeaderEdge()
	return rtuo
}

// ClearRecMemberEdges clears all "rec_member_edges" edges to the User entity.
func (rtuo *RecTeamUpdateOne) ClearRecMemberEdges() *RecTeamUpdateOne {
	rtuo.mutation.ClearRecMemberEdges()
	return rtuo
}

// RemoveRecMemberEdgeIDs removes the "rec_member_edges" edge to User entities by IDs.
func (rtuo *RecTeamUpdateOne) RemoveRecMemberEdgeIDs(ids ...uuid.UUID) *RecTeamUpdateOne {
	rtuo.mutation.RemoveRecMemberEdgeIDs(ids...)
	return rtuo
}

// RemoveRecMemberEdges removes "rec_member_edges" edges to User entities.
func (rtuo *RecTeamUpdateOne) RemoveRecMemberEdges(u ...*User) *RecTeamUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rtuo.RemoveRecMemberEdgeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rtuo *RecTeamUpdateOne) Select(field string, fields ...string) *RecTeamUpdateOne {
	rtuo.fields = append([]string{field}, fields...)
	return rtuo
}

// Save executes the query and returns the updated RecTeam entity.
func (rtuo *RecTeamUpdateOne) Save(ctx context.Context) (*RecTeam, error) {
	var (
		err  error
		node *RecTeam
	)
	if len(rtuo.hooks) == 0 {
		if err = rtuo.check(); err != nil {
			return nil, err
		}
		node, err = rtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecTeamMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rtuo.check(); err != nil {
				return nil, err
			}
			rtuo.mutation = mutation
			node, err = rtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rtuo.hooks) - 1; i >= 0; i-- {
			if rtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rtuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rtuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RecTeam)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RecTeamMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rtuo *RecTeamUpdateOne) SaveX(ctx context.Context) *RecTeam {
	node, err := rtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtuo *RecTeamUpdateOne) Exec(ctx context.Context) error {
	_, err := rtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtuo *RecTeamUpdateOne) ExecX(ctx context.Context) {
	if err := rtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtuo *RecTeamUpdateOne) check() error {
	if v, ok := rtuo.mutation.Name(); ok {
		if err := recteam.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "RecTeam.name": %w`, err)}
		}
	}
	if v, ok := rtuo.mutation.Description(); ok {
		if err := recteam.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "RecTeam.description": %w`, err)}
		}
	}
	if _, ok := rtuo.mutation.RecLeaderEdgeID(); rtuo.mutation.RecLeaderEdgeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RecTeam.rec_leader_edge"`)
	}
	return nil
}

func (rtuo *RecTeamUpdateOne) sqlSave(ctx context.Context) (_node *RecTeam, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recteam.Table,
			Columns: recteam.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: recteam.FieldID,
			},
		},
	}
	id, ok := rtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RecTeam.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recteam.FieldID)
		for _, f := range fields {
			if !recteam.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != recteam.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtuo.mutation.UpdatedAt(); ok {
		_spec.SetField(recteam.FieldUpdatedAt, field.TypeTime, value)
	}
	if rtuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(recteam.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := rtuo.mutation.DeletedAt(); ok {
		_spec.SetField(recteam.FieldDeletedAt, field.TypeTime, value)
	}
	if rtuo.mutation.DeletedAtCleared() {
		_spec.ClearField(recteam.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := rtuo.mutation.Name(); ok {
		_spec.SetField(recteam.FieldName, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.Description(); ok {
		_spec.SetField(recteam.FieldDescription, field.TypeString, value)
	}
	if rtuo.mutation.DescriptionCleared() {
		_spec.ClearField(recteam.FieldDescription, field.TypeString)
	}
	if rtuo.mutation.RecLeaderEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recteam.RecLeaderEdgeTable,
			Columns: []string{recteam.RecLeaderEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.RecLeaderEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recteam.RecLeaderEdgeTable,
			Columns: []string{recteam.RecLeaderEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rtuo.mutation.RecMemberEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recteam.RecMemberEdgesTable,
			Columns: []string{recteam.RecMemberEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.RemovedRecMemberEdgesIDs(); len(nodes) > 0 && !rtuo.mutation.RecMemberEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recteam.RecMemberEdgesTable,
			Columns: []string{recteam.RecMemberEdgesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.RecMemberEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recteam.RecMemberEdgesTable,
			Columns: []string{recteam.RecMemberEdgesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RecTeam{config: rtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recteam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
