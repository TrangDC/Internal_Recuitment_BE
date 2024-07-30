// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"trec/ent/hiringteamapprover"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HiringTeamApproverDelete is the builder for deleting a HiringTeamApprover entity.
type HiringTeamApproverDelete struct {
	config
	hooks    []Hook
	mutation *HiringTeamApproverMutation
}

// Where appends a list predicates to the HiringTeamApproverDelete builder.
func (htad *HiringTeamApproverDelete) Where(ps ...predicate.HiringTeamApprover) *HiringTeamApproverDelete {
	htad.mutation.Where(ps...)
	return htad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (htad *HiringTeamApproverDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(htad.hooks) == 0 {
		affected, err = htad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HiringTeamApproverMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			htad.mutation = mutation
			affected, err = htad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(htad.hooks) - 1; i >= 0; i-- {
			if htad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = htad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, htad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (htad *HiringTeamApproverDelete) ExecX(ctx context.Context) int {
	n, err := htad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (htad *HiringTeamApproverDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: hiringteamapprover.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hiringteamapprover.FieldID,
			},
		},
	}
	if ps := htad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, htad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// HiringTeamApproverDeleteOne is the builder for deleting a single HiringTeamApprover entity.
type HiringTeamApproverDeleteOne struct {
	htad *HiringTeamApproverDelete
}

// Exec executes the deletion query.
func (htado *HiringTeamApproverDeleteOne) Exec(ctx context.Context) error {
	n, err := htado.htad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hiringteamapprover.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (htado *HiringTeamApproverDeleteOne) ExecX(ctx context.Context) {
	htado.htad.ExecX(ctx)
}
