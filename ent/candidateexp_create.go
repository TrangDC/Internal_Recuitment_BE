// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/candidate"
	"trec/ent/candidateexp"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateExpCreate is the builder for creating a CandidateExp entity.
type CandidateExpCreate struct {
	config
	mutation *CandidateExpMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cec *CandidateExpCreate) SetCreatedAt(t time.Time) *CandidateExpCreate {
	cec.mutation.SetCreatedAt(t)
	return cec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableCreatedAt(t *time.Time) *CandidateExpCreate {
	if t != nil {
		cec.SetCreatedAt(*t)
	}
	return cec
}

// SetUpdatedAt sets the "updated_at" field.
func (cec *CandidateExpCreate) SetUpdatedAt(t time.Time) *CandidateExpCreate {
	cec.mutation.SetUpdatedAt(t)
	return cec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableUpdatedAt(t *time.Time) *CandidateExpCreate {
	if t != nil {
		cec.SetUpdatedAt(*t)
	}
	return cec
}

// SetDeletedAt sets the "deleted_at" field.
func (cec *CandidateExpCreate) SetDeletedAt(t time.Time) *CandidateExpCreate {
	cec.mutation.SetDeletedAt(t)
	return cec
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableDeletedAt(t *time.Time) *CandidateExpCreate {
	if t != nil {
		cec.SetDeletedAt(*t)
	}
	return cec
}

// SetCandidateID sets the "candidate_id" field.
func (cec *CandidateExpCreate) SetCandidateID(u uuid.UUID) *CandidateExpCreate {
	cec.mutation.SetCandidateID(u)
	return cec
}

// SetNillableCandidateID sets the "candidate_id" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableCandidateID(u *uuid.UUID) *CandidateExpCreate {
	if u != nil {
		cec.SetCandidateID(*u)
	}
	return cec
}

// SetPosition sets the "position" field.
func (cec *CandidateExpCreate) SetPosition(s string) *CandidateExpCreate {
	cec.mutation.SetPosition(s)
	return cec
}

// SetCompany sets the "company" field.
func (cec *CandidateExpCreate) SetCompany(s string) *CandidateExpCreate {
	cec.mutation.SetCompany(s)
	return cec
}

// SetLocation sets the "location" field.
func (cec *CandidateExpCreate) SetLocation(s string) *CandidateExpCreate {
	cec.mutation.SetLocation(s)
	return cec
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableLocation(s *string) *CandidateExpCreate {
	if s != nil {
		cec.SetLocation(*s)
	}
	return cec
}

// SetDescription sets the "description" field.
func (cec *CandidateExpCreate) SetDescription(s string) *CandidateExpCreate {
	cec.mutation.SetDescription(s)
	return cec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableDescription(s *string) *CandidateExpCreate {
	if s != nil {
		cec.SetDescription(*s)
	}
	return cec
}

// SetStartDate sets the "start_date" field.
func (cec *CandidateExpCreate) SetStartDate(t time.Time) *CandidateExpCreate {
	cec.mutation.SetStartDate(t)
	return cec
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableStartDate(t *time.Time) *CandidateExpCreate {
	if t != nil {
		cec.SetStartDate(*t)
	}
	return cec
}

// SetEndDate sets the "end_date" field.
func (cec *CandidateExpCreate) SetEndDate(t time.Time) *CandidateExpCreate {
	cec.mutation.SetEndDate(t)
	return cec
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableEndDate(t *time.Time) *CandidateExpCreate {
	if t != nil {
		cec.SetEndDate(*t)
	}
	return cec
}

// SetOrderID sets the "order_id" field.
func (cec *CandidateExpCreate) SetOrderID(i int) *CandidateExpCreate {
	cec.mutation.SetOrderID(i)
	return cec
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableOrderID(i *int) *CandidateExpCreate {
	if i != nil {
		cec.SetOrderID(*i)
	}
	return cec
}

// SetIsCurrent sets the "is_current" field.
func (cec *CandidateExpCreate) SetIsCurrent(b bool) *CandidateExpCreate {
	cec.mutation.SetIsCurrent(b)
	return cec
}

// SetNillableIsCurrent sets the "is_current" field if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableIsCurrent(b *bool) *CandidateExpCreate {
	if b != nil {
		cec.SetIsCurrent(*b)
	}
	return cec
}

// SetID sets the "id" field.
func (cec *CandidateExpCreate) SetID(u uuid.UUID) *CandidateExpCreate {
	cec.mutation.SetID(u)
	return cec
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (cec *CandidateExpCreate) SetCandidateEdgeID(id uuid.UUID) *CandidateExpCreate {
	cec.mutation.SetCandidateEdgeID(id)
	return cec
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (cec *CandidateExpCreate) SetNillableCandidateEdgeID(id *uuid.UUID) *CandidateExpCreate {
	if id != nil {
		cec = cec.SetCandidateEdgeID(*id)
	}
	return cec
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (cec *CandidateExpCreate) SetCandidateEdge(c *Candidate) *CandidateExpCreate {
	return cec.SetCandidateEdgeID(c.ID)
}

// Mutation returns the CandidateExpMutation object of the builder.
func (cec *CandidateExpCreate) Mutation() *CandidateExpMutation {
	return cec.mutation
}

// Save creates the CandidateExp in the database.
func (cec *CandidateExpCreate) Save(ctx context.Context) (*CandidateExp, error) {
	var (
		err  error
		node *CandidateExp
	)
	cec.defaults()
	if len(cec.hooks) == 0 {
		if err = cec.check(); err != nil {
			return nil, err
		}
		node, err = cec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateExpMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cec.check(); err != nil {
				return nil, err
			}
			cec.mutation = mutation
			if node, err = cec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cec.hooks) - 1; i >= 0; i-- {
			if cec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CandidateExp)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CandidateExpMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cec *CandidateExpCreate) SaveX(ctx context.Context) *CandidateExp {
	v, err := cec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cec *CandidateExpCreate) Exec(ctx context.Context) error {
	_, err := cec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cec *CandidateExpCreate) ExecX(ctx context.Context) {
	if err := cec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cec *CandidateExpCreate) defaults() {
	if _, ok := cec.mutation.CreatedAt(); !ok {
		v := candidateexp.DefaultCreatedAt()
		cec.mutation.SetCreatedAt(v)
	}
	if _, ok := cec.mutation.IsCurrent(); !ok {
		v := candidateexp.DefaultIsCurrent
		cec.mutation.SetIsCurrent(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cec *CandidateExpCreate) check() error {
	if _, ok := cec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CandidateExp.created_at"`)}
	}
	if _, ok := cec.mutation.Position(); !ok {
		return &ValidationError{Name: "position", err: errors.New(`ent: missing required field "CandidateExp.position"`)}
	}
	if v, ok := cec.mutation.Position(); ok {
		if err := candidateexp.PositionValidator(v); err != nil {
			return &ValidationError{Name: "position", err: fmt.Errorf(`ent: validator failed for field "CandidateExp.position": %w`, err)}
		}
	}
	if _, ok := cec.mutation.Company(); !ok {
		return &ValidationError{Name: "company", err: errors.New(`ent: missing required field "CandidateExp.company"`)}
	}
	if v, ok := cec.mutation.Company(); ok {
		if err := candidateexp.CompanyValidator(v); err != nil {
			return &ValidationError{Name: "company", err: fmt.Errorf(`ent: validator failed for field "CandidateExp.company": %w`, err)}
		}
	}
	if _, ok := cec.mutation.IsCurrent(); !ok {
		return &ValidationError{Name: "is_current", err: errors.New(`ent: missing required field "CandidateExp.is_current"`)}
	}
	return nil
}

func (cec *CandidateExpCreate) sqlSave(ctx context.Context) (*CandidateExp, error) {
	_node, _spec := cec.createSpec()
	if err := sqlgraph.CreateNode(ctx, cec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cec *CandidateExpCreate) createSpec() (*CandidateExp, *sqlgraph.CreateSpec) {
	var (
		_node = &CandidateExp{config: cec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: candidateexp.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidateexp.FieldID,
			},
		}
	)
	if id, ok := cec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cec.mutation.CreatedAt(); ok {
		_spec.SetField(candidateexp.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cec.mutation.UpdatedAt(); ok {
		_spec.SetField(candidateexp.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cec.mutation.DeletedAt(); ok {
		_spec.SetField(candidateexp.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cec.mutation.Position(); ok {
		_spec.SetField(candidateexp.FieldPosition, field.TypeString, value)
		_node.Position = value
	}
	if value, ok := cec.mutation.Company(); ok {
		_spec.SetField(candidateexp.FieldCompany, field.TypeString, value)
		_node.Company = value
	}
	if value, ok := cec.mutation.Location(); ok {
		_spec.SetField(candidateexp.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := cec.mutation.Description(); ok {
		_spec.SetField(candidateexp.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cec.mutation.StartDate(); ok {
		_spec.SetField(candidateexp.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := cec.mutation.EndDate(); ok {
		_spec.SetField(candidateexp.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if value, ok := cec.mutation.OrderID(); ok {
		_spec.SetField(candidateexp.FieldOrderID, field.TypeInt, value)
		_node.OrderID = value
	}
	if value, ok := cec.mutation.IsCurrent(); ok {
		_spec.SetField(candidateexp.FieldIsCurrent, field.TypeBool, value)
		_node.IsCurrent = value
	}
	if nodes := cec.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidateexp.CandidateEdgeTable,
			Columns: []string{candidateexp.CandidateEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CandidateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CandidateExpCreateBulk is the builder for creating many CandidateExp entities in bulk.
type CandidateExpCreateBulk struct {
	config
	builders []*CandidateExpCreate
}

// Save creates the CandidateExp entities in the database.
func (cecb *CandidateExpCreateBulk) Save(ctx context.Context) ([]*CandidateExp, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cecb.builders))
	nodes := make([]*CandidateExp, len(cecb.builders))
	mutators := make([]Mutator, len(cecb.builders))
	for i := range cecb.builders {
		func(i int, root context.Context) {
			builder := cecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CandidateExpMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cecb *CandidateExpCreateBulk) SaveX(ctx context.Context) []*CandidateExp {
	v, err := cecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cecb *CandidateExpCreateBulk) Exec(ctx context.Context) error {
	_, err := cecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cecb *CandidateExpCreateBulk) ExecX(ctx context.Context) {
	if err := cecb.Exec(ctx); err != nil {
		panic(err)
	}
}
