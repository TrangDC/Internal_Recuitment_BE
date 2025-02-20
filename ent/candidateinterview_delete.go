// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"trec/ent/candidateinterview"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CandidateInterviewDelete is the builder for deleting a CandidateInterview entity.
type CandidateInterviewDelete struct {
	config
	hooks    []Hook
	mutation *CandidateInterviewMutation
}

// Where appends a list predicates to the CandidateInterviewDelete builder.
func (cid *CandidateInterviewDelete) Where(ps ...predicate.CandidateInterview) *CandidateInterviewDelete {
	cid.mutation.Where(ps...)
	return cid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cid *CandidateInterviewDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cid.hooks) == 0 {
		affected, err = cid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateInterviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cid.mutation = mutation
			affected, err = cid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cid.hooks) - 1; i >= 0; i-- {
			if cid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cid *CandidateInterviewDelete) ExecX(ctx context.Context) int {
	n, err := cid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cid *CandidateInterviewDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: candidateinterview.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidateinterview.FieldID,
			},
		},
	}
	if ps := cid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CandidateInterviewDeleteOne is the builder for deleting a single CandidateInterview entity.
type CandidateInterviewDeleteOne struct {
	cid *CandidateInterviewDelete
}

// Exec executes the deletion query.
func (cido *CandidateInterviewDeleteOne) Exec(ctx context.Context) error {
	n, err := cido.cid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{candidateinterview.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cido *CandidateInterviewDeleteOne) ExecX(ctx context.Context) {
	cido.cid.ExecX(ctx)
}
