// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"trec/ent/entitypermission"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntityPermissionDelete is the builder for deleting a EntityPermission entity.
type EntityPermissionDelete struct {
	config
	hooks    []Hook
	mutation *EntityPermissionMutation
}

// Where appends a list predicates to the EntityPermissionDelete builder.
func (epd *EntityPermissionDelete) Where(ps ...predicate.EntityPermission) *EntityPermissionDelete {
	epd.mutation.Where(ps...)
	return epd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (epd *EntityPermissionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(epd.hooks) == 0 {
		affected, err = epd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityPermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			epd.mutation = mutation
			affected, err = epd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(epd.hooks) - 1; i >= 0; i-- {
			if epd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = epd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, epd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (epd *EntityPermissionDelete) ExecX(ctx context.Context) int {
	n, err := epd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (epd *EntityPermissionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: entitypermission.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entitypermission.FieldID,
			},
		},
	}
	if ps := epd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, epd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// EntityPermissionDeleteOne is the builder for deleting a single EntityPermission entity.
type EntityPermissionDeleteOne struct {
	epd *EntityPermissionDelete
}

// Exec executes the deletion query.
func (epdo *EntityPermissionDeleteOne) Exec(ctx context.Context) error {
	n, err := epdo.epd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{entitypermission.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (epdo *EntityPermissionDeleteOne) ExecX(ctx context.Context) {
	epdo.epd.ExecX(ctx)
}
