// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"trec/ent/permissiongroup"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionGroupDelete is the builder for deleting a PermissionGroup entity.
type PermissionGroupDelete struct {
	config
	hooks    []Hook
	mutation *PermissionGroupMutation
}

// Where appends a list predicates to the PermissionGroupDelete builder.
func (pgd *PermissionGroupDelete) Where(ps ...predicate.PermissionGroup) *PermissionGroupDelete {
	pgd.mutation.Where(ps...)
	return pgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pgd *PermissionGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pgd.hooks) == 0 {
		affected, err = pgd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pgd.mutation = mutation
			affected, err = pgd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pgd.hooks) - 1; i >= 0; i-- {
			if pgd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pgd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pgd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pgd *PermissionGroupDelete) ExecX(ctx context.Context) int {
	n, err := pgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pgd *PermissionGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: permissiongroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: permissiongroup.FieldID,
			},
		},
	}
	if ps := pgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pgd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// PermissionGroupDeleteOne is the builder for deleting a single PermissionGroup entity.
type PermissionGroupDeleteOne struct {
	pgd *PermissionGroupDelete
}

// Exec executes the deletion query.
func (pgdo *PermissionGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := pgdo.pgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{permissiongroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pgdo *PermissionGroupDeleteOne) ExecX(ctx context.Context) {
	pgdo.pgd.ExecX(ctx)
}
