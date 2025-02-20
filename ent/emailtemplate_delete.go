// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"trec/ent/emailtemplate"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmailTemplateDelete is the builder for deleting a EmailTemplate entity.
type EmailTemplateDelete struct {
	config
	hooks    []Hook
	mutation *EmailTemplateMutation
}

// Where appends a list predicates to the EmailTemplateDelete builder.
func (etd *EmailTemplateDelete) Where(ps ...predicate.EmailTemplate) *EmailTemplateDelete {
	etd.mutation.Where(ps...)
	return etd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (etd *EmailTemplateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(etd.hooks) == 0 {
		affected, err = etd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmailTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			etd.mutation = mutation
			affected, err = etd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(etd.hooks) - 1; i >= 0; i-- {
			if etd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = etd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, etd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (etd *EmailTemplateDelete) ExecX(ctx context.Context) int {
	n, err := etd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (etd *EmailTemplateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: emailtemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailtemplate.FieldID,
			},
		},
	}
	if ps := etd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, etd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// EmailTemplateDeleteOne is the builder for deleting a single EmailTemplate entity.
type EmailTemplateDeleteOne struct {
	etd *EmailTemplateDelete
}

// Exec executes the deletion query.
func (etdo *EmailTemplateDeleteOne) Exec(ctx context.Context) error {
	n, err := etdo.etd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{emailtemplate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (etdo *EmailTemplateDeleteOne) ExecX(ctx context.Context) {
	etdo.etd.ExecX(ctx)
}
