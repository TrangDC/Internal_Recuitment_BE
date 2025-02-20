// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"trec/ent/hiringjob"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HiringJobDelete is the builder for deleting a HiringJob entity.
type HiringJobDelete struct {
	config
	hooks    []Hook
	mutation *HiringJobMutation
}

// Where appends a list predicates to the HiringJobDelete builder.
func (hjd *HiringJobDelete) Where(ps ...predicate.HiringJob) *HiringJobDelete {
	hjd.mutation.Where(ps...)
	return hjd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hjd *HiringJobDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hjd.hooks) == 0 {
		affected, err = hjd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HiringJobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hjd.mutation = mutation
			affected, err = hjd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hjd.hooks) - 1; i >= 0; i-- {
			if hjd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hjd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hjd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hjd *HiringJobDelete) ExecX(ctx context.Context) int {
	n, err := hjd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hjd *HiringJobDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: hiringjob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hiringjob.FieldID,
			},
		},
	}
	if ps := hjd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hjd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// HiringJobDeleteOne is the builder for deleting a single HiringJob entity.
type HiringJobDeleteOne struct {
	hjd *HiringJobDelete
}

// Exec executes the deletion query.
func (hjdo *HiringJobDeleteOne) Exec(ctx context.Context) error {
	n, err := hjdo.hjd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hiringjob.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hjdo *HiringJobDeleteOne) ExecX(ctx context.Context) {
	hjdo.hjd.ExecX(ctx)
}
