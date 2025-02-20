// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidateaward"
	"trec/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateAwardUpdate is the builder for updating CandidateAward entities.
type CandidateAwardUpdate struct {
	config
	hooks    []Hook
	mutation *CandidateAwardMutation
}

// Where appends a list predicates to the CandidateAwardUpdate builder.
func (cau *CandidateAwardUpdate) Where(ps ...predicate.CandidateAward) *CandidateAwardUpdate {
	cau.mutation.Where(ps...)
	return cau
}

// SetUpdatedAt sets the "updated_at" field.
func (cau *CandidateAwardUpdate) SetUpdatedAt(t time.Time) *CandidateAwardUpdate {
	cau.mutation.SetUpdatedAt(t)
	return cau
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cau *CandidateAwardUpdate) SetNillableUpdatedAt(t *time.Time) *CandidateAwardUpdate {
	if t != nil {
		cau.SetUpdatedAt(*t)
	}
	return cau
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cau *CandidateAwardUpdate) ClearUpdatedAt() *CandidateAwardUpdate {
	cau.mutation.ClearUpdatedAt()
	return cau
}

// SetDeletedAt sets the "deleted_at" field.
func (cau *CandidateAwardUpdate) SetDeletedAt(t time.Time) *CandidateAwardUpdate {
	cau.mutation.SetDeletedAt(t)
	return cau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cau *CandidateAwardUpdate) SetNillableDeletedAt(t *time.Time) *CandidateAwardUpdate {
	if t != nil {
		cau.SetDeletedAt(*t)
	}
	return cau
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cau *CandidateAwardUpdate) ClearDeletedAt() *CandidateAwardUpdate {
	cau.mutation.ClearDeletedAt()
	return cau
}

// SetCandidateID sets the "candidate_id" field.
func (cau *CandidateAwardUpdate) SetCandidateID(u uuid.UUID) *CandidateAwardUpdate {
	cau.mutation.SetCandidateID(u)
	return cau
}

// SetNillableCandidateID sets the "candidate_id" field if the given value is not nil.
func (cau *CandidateAwardUpdate) SetNillableCandidateID(u *uuid.UUID) *CandidateAwardUpdate {
	if u != nil {
		cau.SetCandidateID(*u)
	}
	return cau
}

// ClearCandidateID clears the value of the "candidate_id" field.
func (cau *CandidateAwardUpdate) ClearCandidateID() *CandidateAwardUpdate {
	cau.mutation.ClearCandidateID()
	return cau
}

// SetName sets the "name" field.
func (cau *CandidateAwardUpdate) SetName(s string) *CandidateAwardUpdate {
	cau.mutation.SetName(s)
	return cau
}

// SetAchievedDate sets the "achieved_date" field.
func (cau *CandidateAwardUpdate) SetAchievedDate(t time.Time) *CandidateAwardUpdate {
	cau.mutation.SetAchievedDate(t)
	return cau
}

// SetNillableAchievedDate sets the "achieved_date" field if the given value is not nil.
func (cau *CandidateAwardUpdate) SetNillableAchievedDate(t *time.Time) *CandidateAwardUpdate {
	if t != nil {
		cau.SetAchievedDate(*t)
	}
	return cau
}

// ClearAchievedDate clears the value of the "achieved_date" field.
func (cau *CandidateAwardUpdate) ClearAchievedDate() *CandidateAwardUpdate {
	cau.mutation.ClearAchievedDate()
	return cau
}

// SetOrderID sets the "order_id" field.
func (cau *CandidateAwardUpdate) SetOrderID(i int) *CandidateAwardUpdate {
	cau.mutation.ResetOrderID()
	cau.mutation.SetOrderID(i)
	return cau
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cau *CandidateAwardUpdate) SetNillableOrderID(i *int) *CandidateAwardUpdate {
	if i != nil {
		cau.SetOrderID(*i)
	}
	return cau
}

// AddOrderID adds i to the "order_id" field.
func (cau *CandidateAwardUpdate) AddOrderID(i int) *CandidateAwardUpdate {
	cau.mutation.AddOrderID(i)
	return cau
}

// ClearOrderID clears the value of the "order_id" field.
func (cau *CandidateAwardUpdate) ClearOrderID() *CandidateAwardUpdate {
	cau.mutation.ClearOrderID()
	return cau
}

// AddAttachmentEdgeIDs adds the "attachment_edges" edge to the Attachment entity by IDs.
func (cau *CandidateAwardUpdate) AddAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateAwardUpdate {
	cau.mutation.AddAttachmentEdgeIDs(ids...)
	return cau
}

// AddAttachmentEdges adds the "attachment_edges" edges to the Attachment entity.
func (cau *CandidateAwardUpdate) AddAttachmentEdges(a ...*Attachment) *CandidateAwardUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cau.AddAttachmentEdgeIDs(ids...)
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (cau *CandidateAwardUpdate) SetCandidateEdgeID(id uuid.UUID) *CandidateAwardUpdate {
	cau.mutation.SetCandidateEdgeID(id)
	return cau
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (cau *CandidateAwardUpdate) SetNillableCandidateEdgeID(id *uuid.UUID) *CandidateAwardUpdate {
	if id != nil {
		cau = cau.SetCandidateEdgeID(*id)
	}
	return cau
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (cau *CandidateAwardUpdate) SetCandidateEdge(c *Candidate) *CandidateAwardUpdate {
	return cau.SetCandidateEdgeID(c.ID)
}

// Mutation returns the CandidateAwardMutation object of the builder.
func (cau *CandidateAwardUpdate) Mutation() *CandidateAwardMutation {
	return cau.mutation
}

// ClearAttachmentEdges clears all "attachment_edges" edges to the Attachment entity.
func (cau *CandidateAwardUpdate) ClearAttachmentEdges() *CandidateAwardUpdate {
	cau.mutation.ClearAttachmentEdges()
	return cau
}

// RemoveAttachmentEdgeIDs removes the "attachment_edges" edge to Attachment entities by IDs.
func (cau *CandidateAwardUpdate) RemoveAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateAwardUpdate {
	cau.mutation.RemoveAttachmentEdgeIDs(ids...)
	return cau
}

// RemoveAttachmentEdges removes "attachment_edges" edges to Attachment entities.
func (cau *CandidateAwardUpdate) RemoveAttachmentEdges(a ...*Attachment) *CandidateAwardUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cau.RemoveAttachmentEdgeIDs(ids...)
}

// ClearCandidateEdge clears the "candidate_edge" edge to the Candidate entity.
func (cau *CandidateAwardUpdate) ClearCandidateEdge() *CandidateAwardUpdate {
	cau.mutation.ClearCandidateEdge()
	return cau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cau *CandidateAwardUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cau.hooks) == 0 {
		if err = cau.check(); err != nil {
			return 0, err
		}
		affected, err = cau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateAwardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cau.check(); err != nil {
				return 0, err
			}
			cau.mutation = mutation
			affected, err = cau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cau.hooks) - 1; i >= 0; i-- {
			if cau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cau *CandidateAwardUpdate) SaveX(ctx context.Context) int {
	affected, err := cau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cau *CandidateAwardUpdate) Exec(ctx context.Context) error {
	_, err := cau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cau *CandidateAwardUpdate) ExecX(ctx context.Context) {
	if err := cau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cau *CandidateAwardUpdate) check() error {
	if v, ok := cau.mutation.Name(); ok {
		if err := candidateaward.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CandidateAward.name": %w`, err)}
		}
	}
	return nil
}

func (cau *CandidateAwardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidateaward.Table,
			Columns: candidateaward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidateaward.FieldID,
			},
		},
	}
	if ps := cau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cau.mutation.UpdatedAt(); ok {
		_spec.SetField(candidateaward.FieldUpdatedAt, field.TypeTime, value)
	}
	if cau.mutation.UpdatedAtCleared() {
		_spec.ClearField(candidateaward.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cau.mutation.DeletedAt(); ok {
		_spec.SetField(candidateaward.FieldDeletedAt, field.TypeTime, value)
	}
	if cau.mutation.DeletedAtCleared() {
		_spec.ClearField(candidateaward.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cau.mutation.Name(); ok {
		_spec.SetField(candidateaward.FieldName, field.TypeString, value)
	}
	if value, ok := cau.mutation.AchievedDate(); ok {
		_spec.SetField(candidateaward.FieldAchievedDate, field.TypeTime, value)
	}
	if cau.mutation.AchievedDateCleared() {
		_spec.ClearField(candidateaward.FieldAchievedDate, field.TypeTime)
	}
	if value, ok := cau.mutation.OrderID(); ok {
		_spec.SetField(candidateaward.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := cau.mutation.AddedOrderID(); ok {
		_spec.AddField(candidateaward.FieldOrderID, field.TypeInt, value)
	}
	if cau.mutation.OrderIDCleared() {
		_spec.ClearField(candidateaward.FieldOrderID, field.TypeInt)
	}
	if cau.mutation.AttachmentEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidateaward.AttachmentEdgesTable,
			Columns: []string{candidateaward.AttachmentEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cau.mutation.RemovedAttachmentEdgesIDs(); len(nodes) > 0 && !cau.mutation.AttachmentEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidateaward.AttachmentEdgesTable,
			Columns: []string{candidateaward.AttachmentEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cau.mutation.AttachmentEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidateaward.AttachmentEdgesTable,
			Columns: []string{candidateaward.AttachmentEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cau.mutation.CandidateEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidateaward.CandidateEdgeTable,
			Columns: []string{candidateaward.CandidateEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cau.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidateaward.CandidateEdgeTable,
			Columns: []string{candidateaward.CandidateEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{candidateaward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CandidateAwardUpdateOne is the builder for updating a single CandidateAward entity.
type CandidateAwardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CandidateAwardMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cauo *CandidateAwardUpdateOne) SetUpdatedAt(t time.Time) *CandidateAwardUpdateOne {
	cauo.mutation.SetUpdatedAt(t)
	return cauo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cauo *CandidateAwardUpdateOne) SetNillableUpdatedAt(t *time.Time) *CandidateAwardUpdateOne {
	if t != nil {
		cauo.SetUpdatedAt(*t)
	}
	return cauo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cauo *CandidateAwardUpdateOne) ClearUpdatedAt() *CandidateAwardUpdateOne {
	cauo.mutation.ClearUpdatedAt()
	return cauo
}

// SetDeletedAt sets the "deleted_at" field.
func (cauo *CandidateAwardUpdateOne) SetDeletedAt(t time.Time) *CandidateAwardUpdateOne {
	cauo.mutation.SetDeletedAt(t)
	return cauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cauo *CandidateAwardUpdateOne) SetNillableDeletedAt(t *time.Time) *CandidateAwardUpdateOne {
	if t != nil {
		cauo.SetDeletedAt(*t)
	}
	return cauo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cauo *CandidateAwardUpdateOne) ClearDeletedAt() *CandidateAwardUpdateOne {
	cauo.mutation.ClearDeletedAt()
	return cauo
}

// SetCandidateID sets the "candidate_id" field.
func (cauo *CandidateAwardUpdateOne) SetCandidateID(u uuid.UUID) *CandidateAwardUpdateOne {
	cauo.mutation.SetCandidateID(u)
	return cauo
}

// SetNillableCandidateID sets the "candidate_id" field if the given value is not nil.
func (cauo *CandidateAwardUpdateOne) SetNillableCandidateID(u *uuid.UUID) *CandidateAwardUpdateOne {
	if u != nil {
		cauo.SetCandidateID(*u)
	}
	return cauo
}

// ClearCandidateID clears the value of the "candidate_id" field.
func (cauo *CandidateAwardUpdateOne) ClearCandidateID() *CandidateAwardUpdateOne {
	cauo.mutation.ClearCandidateID()
	return cauo
}

// SetName sets the "name" field.
func (cauo *CandidateAwardUpdateOne) SetName(s string) *CandidateAwardUpdateOne {
	cauo.mutation.SetName(s)
	return cauo
}

// SetAchievedDate sets the "achieved_date" field.
func (cauo *CandidateAwardUpdateOne) SetAchievedDate(t time.Time) *CandidateAwardUpdateOne {
	cauo.mutation.SetAchievedDate(t)
	return cauo
}

// SetNillableAchievedDate sets the "achieved_date" field if the given value is not nil.
func (cauo *CandidateAwardUpdateOne) SetNillableAchievedDate(t *time.Time) *CandidateAwardUpdateOne {
	if t != nil {
		cauo.SetAchievedDate(*t)
	}
	return cauo
}

// ClearAchievedDate clears the value of the "achieved_date" field.
func (cauo *CandidateAwardUpdateOne) ClearAchievedDate() *CandidateAwardUpdateOne {
	cauo.mutation.ClearAchievedDate()
	return cauo
}

// SetOrderID sets the "order_id" field.
func (cauo *CandidateAwardUpdateOne) SetOrderID(i int) *CandidateAwardUpdateOne {
	cauo.mutation.ResetOrderID()
	cauo.mutation.SetOrderID(i)
	return cauo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cauo *CandidateAwardUpdateOne) SetNillableOrderID(i *int) *CandidateAwardUpdateOne {
	if i != nil {
		cauo.SetOrderID(*i)
	}
	return cauo
}

// AddOrderID adds i to the "order_id" field.
func (cauo *CandidateAwardUpdateOne) AddOrderID(i int) *CandidateAwardUpdateOne {
	cauo.mutation.AddOrderID(i)
	return cauo
}

// ClearOrderID clears the value of the "order_id" field.
func (cauo *CandidateAwardUpdateOne) ClearOrderID() *CandidateAwardUpdateOne {
	cauo.mutation.ClearOrderID()
	return cauo
}

// AddAttachmentEdgeIDs adds the "attachment_edges" edge to the Attachment entity by IDs.
func (cauo *CandidateAwardUpdateOne) AddAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateAwardUpdateOne {
	cauo.mutation.AddAttachmentEdgeIDs(ids...)
	return cauo
}

// AddAttachmentEdges adds the "attachment_edges" edges to the Attachment entity.
func (cauo *CandidateAwardUpdateOne) AddAttachmentEdges(a ...*Attachment) *CandidateAwardUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cauo.AddAttachmentEdgeIDs(ids...)
}

// SetCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID.
func (cauo *CandidateAwardUpdateOne) SetCandidateEdgeID(id uuid.UUID) *CandidateAwardUpdateOne {
	cauo.mutation.SetCandidateEdgeID(id)
	return cauo
}

// SetNillableCandidateEdgeID sets the "candidate_edge" edge to the Candidate entity by ID if the given value is not nil.
func (cauo *CandidateAwardUpdateOne) SetNillableCandidateEdgeID(id *uuid.UUID) *CandidateAwardUpdateOne {
	if id != nil {
		cauo = cauo.SetCandidateEdgeID(*id)
	}
	return cauo
}

// SetCandidateEdge sets the "candidate_edge" edge to the Candidate entity.
func (cauo *CandidateAwardUpdateOne) SetCandidateEdge(c *Candidate) *CandidateAwardUpdateOne {
	return cauo.SetCandidateEdgeID(c.ID)
}

// Mutation returns the CandidateAwardMutation object of the builder.
func (cauo *CandidateAwardUpdateOne) Mutation() *CandidateAwardMutation {
	return cauo.mutation
}

// ClearAttachmentEdges clears all "attachment_edges" edges to the Attachment entity.
func (cauo *CandidateAwardUpdateOne) ClearAttachmentEdges() *CandidateAwardUpdateOne {
	cauo.mutation.ClearAttachmentEdges()
	return cauo
}

// RemoveAttachmentEdgeIDs removes the "attachment_edges" edge to Attachment entities by IDs.
func (cauo *CandidateAwardUpdateOne) RemoveAttachmentEdgeIDs(ids ...uuid.UUID) *CandidateAwardUpdateOne {
	cauo.mutation.RemoveAttachmentEdgeIDs(ids...)
	return cauo
}

// RemoveAttachmentEdges removes "attachment_edges" edges to Attachment entities.
func (cauo *CandidateAwardUpdateOne) RemoveAttachmentEdges(a ...*Attachment) *CandidateAwardUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cauo.RemoveAttachmentEdgeIDs(ids...)
}

// ClearCandidateEdge clears the "candidate_edge" edge to the Candidate entity.
func (cauo *CandidateAwardUpdateOne) ClearCandidateEdge() *CandidateAwardUpdateOne {
	cauo.mutation.ClearCandidateEdge()
	return cauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cauo *CandidateAwardUpdateOne) Select(field string, fields ...string) *CandidateAwardUpdateOne {
	cauo.fields = append([]string{field}, fields...)
	return cauo
}

// Save executes the query and returns the updated CandidateAward entity.
func (cauo *CandidateAwardUpdateOne) Save(ctx context.Context) (*CandidateAward, error) {
	var (
		err  error
		node *CandidateAward
	)
	if len(cauo.hooks) == 0 {
		if err = cauo.check(); err != nil {
			return nil, err
		}
		node, err = cauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CandidateAwardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cauo.check(); err != nil {
				return nil, err
			}
			cauo.mutation = mutation
			node, err = cauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cauo.hooks) - 1; i >= 0; i-- {
			if cauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CandidateAward)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CandidateAwardMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cauo *CandidateAwardUpdateOne) SaveX(ctx context.Context) *CandidateAward {
	node, err := cauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cauo *CandidateAwardUpdateOne) Exec(ctx context.Context) error {
	_, err := cauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cauo *CandidateAwardUpdateOne) ExecX(ctx context.Context) {
	if err := cauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cauo *CandidateAwardUpdateOne) check() error {
	if v, ok := cauo.mutation.Name(); ok {
		if err := candidateaward.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "CandidateAward.name": %w`, err)}
		}
	}
	return nil
}

func (cauo *CandidateAwardUpdateOne) sqlSave(ctx context.Context) (_node *CandidateAward, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   candidateaward.Table,
			Columns: candidateaward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: candidateaward.FieldID,
			},
		},
	}
	id, ok := cauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CandidateAward.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, candidateaward.FieldID)
		for _, f := range fields {
			if !candidateaward.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != candidateaward.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cauo.mutation.UpdatedAt(); ok {
		_spec.SetField(candidateaward.FieldUpdatedAt, field.TypeTime, value)
	}
	if cauo.mutation.UpdatedAtCleared() {
		_spec.ClearField(candidateaward.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cauo.mutation.DeletedAt(); ok {
		_spec.SetField(candidateaward.FieldDeletedAt, field.TypeTime, value)
	}
	if cauo.mutation.DeletedAtCleared() {
		_spec.ClearField(candidateaward.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cauo.mutation.Name(); ok {
		_spec.SetField(candidateaward.FieldName, field.TypeString, value)
	}
	if value, ok := cauo.mutation.AchievedDate(); ok {
		_spec.SetField(candidateaward.FieldAchievedDate, field.TypeTime, value)
	}
	if cauo.mutation.AchievedDateCleared() {
		_spec.ClearField(candidateaward.FieldAchievedDate, field.TypeTime)
	}
	if value, ok := cauo.mutation.OrderID(); ok {
		_spec.SetField(candidateaward.FieldOrderID, field.TypeInt, value)
	}
	if value, ok := cauo.mutation.AddedOrderID(); ok {
		_spec.AddField(candidateaward.FieldOrderID, field.TypeInt, value)
	}
	if cauo.mutation.OrderIDCleared() {
		_spec.ClearField(candidateaward.FieldOrderID, field.TypeInt)
	}
	if cauo.mutation.AttachmentEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidateaward.AttachmentEdgesTable,
			Columns: []string{candidateaward.AttachmentEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cauo.mutation.RemovedAttachmentEdgesIDs(); len(nodes) > 0 && !cauo.mutation.AttachmentEdgesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidateaward.AttachmentEdgesTable,
			Columns: []string{candidateaward.AttachmentEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cauo.mutation.AttachmentEdgesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   candidateaward.AttachmentEdgesTable,
			Columns: []string{candidateaward.AttachmentEdgesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cauo.mutation.CandidateEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidateaward.CandidateEdgeTable,
			Columns: []string{candidateaward.CandidateEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: candidate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cauo.mutation.CandidateEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   candidateaward.CandidateEdgeTable,
			Columns: []string{candidateaward.CandidateEdgeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CandidateAward{config: cauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{candidateaward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
