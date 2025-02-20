// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"trec/ent/candidatehistorycall"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CandidateHistoryCallDelete is the builder for deleting a CandidateHistoryCall entity.
type CandidateHistoryCallDelete struct {
	config
	hooks    []Hook
	mutation *CandidateHistoryCallMutation
}

// Where appends a list predicates to the CandidateHistoryCallDelete builder.
func (chcd *CandidateHistoryCallDelete) Where(ps ...predicate.CandidateHistoryCall) *CandidateHistoryCallDelete {
	chcd.mutation.Where(ps...)
	return chcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (chcd *CandidateHistoryCallDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(chcd.hooks) == 0 {
		affected, err = chcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateHistoryCallMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			chcd.mutation = mutation
			affected, err = chcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(chcd.hooks) - 1; i >= 0; i-- {
			if chcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = chcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, chcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (chcd *CandidateHistoryCallDelete) ExecX(ctx context.Context) int {
	n, err := chcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (chcd *CandidateHistoryCallDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: candidatehistorycall.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidatehistorycall.FieldID,
			},
		},
	}
	if ps := chcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, chcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CandidateHistoryCallDeleteOne is the builder for deleting a single CandidateHistoryCall entity.
type CandidateHistoryCallDeleteOne struct {
	chcd *CandidateHistoryCallDelete
}

// Exec executes the deletion query.
func (chcdo *CandidateHistoryCallDeleteOne) Exec(ctx context.Context) error {
	n, err := chcdo.chcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{candidatehistorycall.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (chcdo *CandidateHistoryCallDeleteOne) ExecX(ctx context.Context) {
	chcdo.chcd.ExecX(ctx)
}
