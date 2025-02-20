// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"trec/ent/jobposition"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobPositionDelete is the builder for deleting a JobPosition entity.
type JobPositionDelete struct {
	config
	hooks    []Hook
	mutation *JobPositionMutation
}

// Where appends a list predicates to the JobPositionDelete builder.
func (jpd *JobPositionDelete) Where(ps ...predicate.JobPosition) *JobPositionDelete {
	jpd.mutation.Where(ps...)
	return jpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jpd *JobPositionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(jpd.hooks) == 0 {
		affected, err = jpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jpd.mutation = mutation
			affected, err = jpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(jpd.hooks) - 1; i >= 0; i-- {
			if jpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (jpd *JobPositionDelete) ExecX(ctx context.Context) int {
	n, err := jpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jpd *JobPositionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: jobposition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: jobposition.FieldID,
			},
		},
	}
	if ps := jpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, jpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// JobPositionDeleteOne is the builder for deleting a single JobPosition entity.
type JobPositionDeleteOne struct {
	jpd *JobPositionDelete
}

// Exec executes the deletion query.
func (jpdo *JobPositionDeleteOne) Exec(ctx context.Context) error {
	n, err := jpdo.jpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{jobposition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jpdo *JobPositionDeleteOne) ExecX(ctx context.Context) {
	jpdo.jpd.ExecX(ctx)
}
