// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/hiringjob"
	"trec/ent/hiringjobstep"
	"trec/ent/predicate"
	"trec/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HiringJobStepUpdate is the builder for updating HiringJobStep entities.
type HiringJobStepUpdate struct {
	config
	hooks    []Hook
	mutation *HiringJobStepMutation
}

// Where appends a list predicates to the HiringJobStepUpdate builder.
func (hjsu *HiringJobStepUpdate) Where(ps ...predicate.HiringJobStep) *HiringJobStepUpdate {
	hjsu.mutation.Where(ps...)
	return hjsu
}

// SetHiringJobID sets the "hiring_job_id" field.
func (hjsu *HiringJobStepUpdate) SetHiringJobID(u uuid.UUID) *HiringJobStepUpdate {
	hjsu.mutation.SetHiringJobID(u)
	return hjsu
}

// SetUserID sets the "user_id" field.
func (hjsu *HiringJobStepUpdate) SetUserID(u uuid.UUID) *HiringJobStepUpdate {
	hjsu.mutation.SetUserID(u)
	return hjsu
}

// SetStatus sets the "status" field.
func (hjsu *HiringJobStepUpdate) SetStatus(h hiringjobstep.Status) *HiringJobStepUpdate {
	hjsu.mutation.SetStatus(h)
	return hjsu
}

// SetOrderID sets the "order_id" field.
func (hjsu *HiringJobStepUpdate) SetOrderID(i int) *HiringJobStepUpdate {
	hjsu.mutation.ResetOrderID()
	hjsu.mutation.SetOrderID(i)
	return hjsu
}

// AddOrderID adds i to the "order_id" field.
func (hjsu *HiringJobStepUpdate) AddOrderID(i int) *HiringJobStepUpdate {
	hjsu.mutation.AddOrderID(i)
	return hjsu
}

// SetUpdatedAt sets the "updated_at" field.
func (hjsu *HiringJobStepUpdate) SetUpdatedAt(t time.Time) *HiringJobStepUpdate {
	hjsu.mutation.SetUpdatedAt(t)
	return hjsu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hjsu *HiringJobStepUpdate) SetNillableUpdatedAt(t *time.Time) *HiringJobStepUpdate {
	if t != nil {
		hjsu.SetUpdatedAt(*t)
	}
	return hjsu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (hjsu *HiringJobStepUpdate) ClearUpdatedAt() *HiringJobStepUpdate {
	hjsu.mutation.ClearUpdatedAt()
	return hjsu
}

// SetApprovalJobID sets the "approval_job" edge to the HiringJob entity by ID.
func (hjsu *HiringJobStepUpdate) SetApprovalJobID(id uuid.UUID) *HiringJobStepUpdate {
	hjsu.mutation.SetApprovalJobID(id)
	return hjsu
}

// SetApprovalJob sets the "approval_job" edge to the HiringJob entity.
func (hjsu *HiringJobStepUpdate) SetApprovalJob(h *HiringJob) *HiringJobStepUpdate {
	return hjsu.SetApprovalJobID(h.ID)
}

// SetApprovalUserID sets the "approval_user" edge to the User entity by ID.
func (hjsu *HiringJobStepUpdate) SetApprovalUserID(id uuid.UUID) *HiringJobStepUpdate {
	hjsu.mutation.SetApprovalUserID(id)
	return hjsu
}

// SetApprovalUser sets the "approval_user" edge to the User entity.
func (hjsu *HiringJobStepUpdate) SetApprovalUser(u *User) *HiringJobStepUpdate {
	return hjsu.SetApprovalUserID(u.ID)
}

// Mutation returns the HiringJobStepMutation object of the builder.
func (hjsu *HiringJobStepUpdate) Mutation() *HiringJobStepMutation {
	return hjsu.mutation
}

// ClearApprovalJob clears the "approval_job" edge to the HiringJob entity.
func (hjsu *HiringJobStepUpdate) ClearApprovalJob() *HiringJobStepUpdate {
	hjsu.mutation.ClearApprovalJob()
	return hjsu
}

// ClearApprovalUser clears the "approval_user" edge to the User entity.
func (hjsu *HiringJobStepUpdate) ClearApprovalUser() *HiringJobStepUpdate {
	hjsu.mutation.ClearApprovalUser()
	return hjsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hjsu *HiringJobStepUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hjsu.hooks) == 0 {
		if err = hjsu.check(); err != nil {
			return 0, err
		}
		affected, err = hjsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HiringJobStepMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hjsu.check(); err != nil {
				return 0, err
			}
			hjsu.mutation = mutation
			affected, err = hjsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hjsu.hooks) - 1; i >= 0; i-- {
			if hjsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hjsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hjsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hjsu *HiringJobStepUpdate) SaveX(ctx context.Context) int {
	affected, err := hjsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hjsu *HiringJobStepUpdate) Exec(ctx context.Context) error {
	_, err := hjsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hjsu *HiringJobStepUpdate) ExecX(ctx context.Context) {
	if err := hjsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hjsu *HiringJobStepUpdate) check() error {
	if v, ok := hjsu.mutation.Status(); ok {
		if err := hiringjobstep.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HiringJobStep.status": %w`, err)}
		}
	}
	if v, ok := hjsu.mutation.OrderID(); ok {
		if err := hiringjobstep.OrderIDValidator(v); err != nil {
			return &ValidationError{Name: "order_id", err: fmt.Errorf(`ent: validator failed for field "HiringJobStep.order_id": %w`, err)}
		}
	}
	if _, ok := hjsu.mutation.ApprovalJobID(); hjsu.mutation.ApprovalJobCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HiringJobStep.approval_job"`)
	}
	if _, ok := hjsu.mutation.ApprovalUserID(); hjsu.mutation.ApprovalUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HiringJobStep.approval_user"`)
	}
	return nil
}

func (hjsu *HiringJobStepUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hiringjobstep.Table,
			Columns: hiringjobstep.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hiringjobstep.FieldID,
			},
		},
	}
	if ps := hjsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hjsu.mutation.Status(); ok {
		_spec.SetField(hiringjobstep.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := hjsu.mutation.OrderID(); ok {
		_spec.SetField(hiringjobstep.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := hjsu.mutation.AddedOrderID(); ok {
		_spec.AddField(hiringjobstep.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := hjsu.mutation.UpdatedAt(); ok {
		_spec.SetField(hiringjobstep.FieldUpdatedAt, field.TypeTime, value)
	}
	if hjsu.mutation.UpdatedAtCleared() {
		_spec.ClearField(hiringjobstep.FieldUpdatedAt, field.TypeTime)
	}
	if hjsu.mutation.ApprovalJobCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hjsu.mutation.ApprovalJobIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hjsu.mutation.ApprovalUserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hjsu.mutation.ApprovalUserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hjsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hiringjobstep.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// HiringJobStepUpdateOne is the builder for updating a single HiringJobStep entity.
type HiringJobStepUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HiringJobStepMutation
}

// SetHiringJobID sets the "hiring_job_id" field.
func (hjsuo *HiringJobStepUpdateOne) SetHiringJobID(u uuid.UUID) *HiringJobStepUpdateOne {
	hjsuo.mutation.SetHiringJobID(u)
	return hjsuo
}

// SetUserID sets the "user_id" field.
func (hjsuo *HiringJobStepUpdateOne) SetUserID(u uuid.UUID) *HiringJobStepUpdateOne {
	hjsuo.mutation.SetUserID(u)
	return hjsuo
}

// SetStatus sets the "status" field.
func (hjsuo *HiringJobStepUpdateOne) SetStatus(h hiringjobstep.Status) *HiringJobStepUpdateOne {
	hjsuo.mutation.SetStatus(h)
	return hjsuo
}

// SetOrderID sets the "order_id" field.
func (hjsuo *HiringJobStepUpdateOne) SetOrderID(i int) *HiringJobStepUpdateOne {
	hjsuo.mutation.ResetOrderID()
	hjsuo.mutation.SetOrderID(i)
	return hjsuo
}

// AddOrderID adds i to the "order_id" field.
func (hjsuo *HiringJobStepUpdateOne) AddOrderID(i int) *HiringJobStepUpdateOne {
	hjsuo.mutation.AddOrderID(i)
	return hjsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (hjsuo *HiringJobStepUpdateOne) SetUpdatedAt(t time.Time) *HiringJobStepUpdateOne {
	hjsuo.mutation.SetUpdatedAt(t)
	return hjsuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hjsuo *HiringJobStepUpdateOne) SetNillableUpdatedAt(t *time.Time) *HiringJobStepUpdateOne {
	if t != nil {
		hjsuo.SetUpdatedAt(*t)
	}
	return hjsuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (hjsuo *HiringJobStepUpdateOne) ClearUpdatedAt() *HiringJobStepUpdateOne {
	hjsuo.mutation.ClearUpdatedAt()
	return hjsuo
}

// SetApprovalJobID sets the "approval_job" edge to the HiringJob entity by ID.
func (hjsuo *HiringJobStepUpdateOne) SetApprovalJobID(id uuid.UUID) *HiringJobStepUpdateOne {
	hjsuo.mutation.SetApprovalJobID(id)
	return hjsuo
}

// SetApprovalJob sets the "approval_job" edge to the HiringJob entity.
func (hjsuo *HiringJobStepUpdateOne) SetApprovalJob(h *HiringJob) *HiringJobStepUpdateOne {
	return hjsuo.SetApprovalJobID(h.ID)
}

// SetApprovalUserID sets the "approval_user" edge to the User entity by ID.
func (hjsuo *HiringJobStepUpdateOne) SetApprovalUserID(id uuid.UUID) *HiringJobStepUpdateOne {
	hjsuo.mutation.SetApprovalUserID(id)
	return hjsuo
}

// SetApprovalUser sets the "approval_user" edge to the User entity.
func (hjsuo *HiringJobStepUpdateOne) SetApprovalUser(u *User) *HiringJobStepUpdateOne {
	return hjsuo.SetApprovalUserID(u.ID)
}

// Mutation returns the HiringJobStepMutation object of the builder.
func (hjsuo *HiringJobStepUpdateOne) Mutation() *HiringJobStepMutation {
	return hjsuo.mutation
}

// ClearApprovalJob clears the "approval_job" edge to the HiringJob entity.
func (hjsuo *HiringJobStepUpdateOne) ClearApprovalJob() *HiringJobStepUpdateOne {
	hjsuo.mutation.ClearApprovalJob()
	return hjsuo
}

// ClearApprovalUser clears the "approval_user" edge to the User entity.
func (hjsuo *HiringJobStepUpdateOne) ClearApprovalUser() *HiringJobStepUpdateOne {
	hjsuo.mutation.ClearApprovalUser()
	return hjsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hjsuo *HiringJobStepUpdateOne) Select(field string, fields ...string) *HiringJobStepUpdateOne {
	hjsuo.fields = append([]string{field}, fields...)
	return hjsuo
}

// Save executes the query and returns the updated HiringJobStep entity.
func (hjsuo *HiringJobStepUpdateOne) Save(ctx context.Context) (*HiringJobStep, error) {
	var (
		err  error
		node *HiringJobStep
	)
	if len(hjsuo.hooks) == 0 {
		if err = hjsuo.check(); err != nil {
			return nil, err
		}
		node, err = hjsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HiringJobStepMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hjsuo.check(); err != nil {
				return nil, err
			}
			hjsuo.mutation = mutation
			node, err = hjsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hjsuo.hooks) - 1; i >= 0; i-- {
			if hjsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hjsuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, hjsuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (hjsuo *HiringJobStepUpdateOne) SaveX(ctx context.Context) *HiringJobStep {
	node, err := hjsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hjsuo *HiringJobStepUpdateOne) Exec(ctx context.Context) error {
	_, err := hjsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hjsuo *HiringJobStepUpdateOne) ExecX(ctx context.Context) {
	if err := hjsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hjsuo *HiringJobStepUpdateOne) check() error {
	if v, ok := hjsuo.mutation.Status(); ok {
		if err := hiringjobstep.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HiringJobStep.status": %w`, err)}
		}
	}
	if v, ok := hjsuo.mutation.OrderID(); ok {
		if err := hiringjobstep.OrderIDValidator(v); err != nil {
			return &ValidationError{Name: "order_id", err: fmt.Errorf(`ent: validator failed for field "HiringJobStep.order_id": %w`, err)}
		}
	}
	if _, ok := hjsuo.mutation.ApprovalJobID(); hjsuo.mutation.ApprovalJobCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HiringJobStep.approval_job"`)
	}
	if _, ok := hjsuo.mutation.ApprovalUserID(); hjsuo.mutation.ApprovalUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HiringJobStep.approval_user"`)
	}
	return nil
}

func (hjsuo *HiringJobStepUpdateOne) sqlSave(ctx context.Context) (_node *HiringJobStep, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hiringjobstep.Table,
			Columns: hiringjobstep.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hiringjobstep.FieldID,
			},
		},
	}
	id, ok := hjsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HiringJobStep.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hjsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hiringjobstep.FieldID)
		for _, f := range fields {
			if !hiringjobstep.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hiringjobstep.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hjsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hjsuo.mutation.Status(); ok {
		_spec.SetField(hiringjobstep.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := hjsuo.mutation.OrderID(); ok {
		_spec.SetField(hiringjobstep.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := hjsuo.mutation.AddedOrderID(); ok {
		_spec.AddField(hiringjobstep.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := hjsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(hiringjobstep.FieldUpdatedAt, field.TypeTime, value)
	}
	if hjsuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(hiringjobstep.FieldUpdatedAt, field.TypeTime)
	}
	if hjsuo.mutation.ApprovalJobCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hjsuo.mutation.ApprovalJobIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hjsuo.mutation.ApprovalUserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hjsuo.mutation.ApprovalUserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HiringJobStep{config: hjsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hjsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hiringjobstep.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
