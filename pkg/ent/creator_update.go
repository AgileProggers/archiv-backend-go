// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/clip"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/creator"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/predicate"
)

// CreatorUpdate is the builder for updating Creator entities.
type CreatorUpdate struct {
	config
	hooks    []Hook
	mutation *CreatorMutation
}

// Where appends a list predicates to the CreatorUpdate builder.
func (cu *CreatorUpdate) Where(ps ...predicate.Creator) *CreatorUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CreatorUpdate) SetName(s string) *CreatorUpdate {
	cu.mutation.SetName(s)
	return cu
}

// AddClipIDs adds the "clips" edge to the Clip entity by IDs.
func (cu *CreatorUpdate) AddClipIDs(ids ...int) *CreatorUpdate {
	cu.mutation.AddClipIDs(ids...)
	return cu
}

// AddClips adds the "clips" edges to the Clip entity.
func (cu *CreatorUpdate) AddClips(c ...*Clip) *CreatorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddClipIDs(ids...)
}

// Mutation returns the CreatorMutation object of the builder.
func (cu *CreatorUpdate) Mutation() *CreatorMutation {
	return cu.mutation
}

// ClearClips clears all "clips" edges to the Clip entity.
func (cu *CreatorUpdate) ClearClips() *CreatorUpdate {
	cu.mutation.ClearClips()
	return cu
}

// RemoveClipIDs removes the "clips" edge to Clip entities by IDs.
func (cu *CreatorUpdate) RemoveClipIDs(ids ...int) *CreatorUpdate {
	cu.mutation.RemoveClipIDs(ids...)
	return cu
}

// RemoveClips removes "clips" edges to Clip entities.
func (cu *CreatorUpdate) RemoveClips(c ...*Clip) *CreatorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveClipIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CreatorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CreatorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CreatorUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CreatorUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CreatorUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CreatorUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := creator.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Creator.name": %w`, err)}
		}
	}
	return nil
}

func (cu *CreatorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   creator.Table,
			Columns: creator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: creator.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: creator.FieldName,
		})
	}
	if cu.mutation.ClipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   creator.ClipsTable,
			Columns: []string{creator.ClipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clip.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedClipsIDs(); len(nodes) > 0 && !cu.mutation.ClipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   creator.ClipsTable,
			Columns: []string{creator.ClipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clip.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   creator.ClipsTable,
			Columns: []string{creator.ClipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clip.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{creator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CreatorUpdateOne is the builder for updating a single Creator entity.
type CreatorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CreatorMutation
}

// SetName sets the "name" field.
func (cuo *CreatorUpdateOne) SetName(s string) *CreatorUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// AddClipIDs adds the "clips" edge to the Clip entity by IDs.
func (cuo *CreatorUpdateOne) AddClipIDs(ids ...int) *CreatorUpdateOne {
	cuo.mutation.AddClipIDs(ids...)
	return cuo
}

// AddClips adds the "clips" edges to the Clip entity.
func (cuo *CreatorUpdateOne) AddClips(c ...*Clip) *CreatorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddClipIDs(ids...)
}

// Mutation returns the CreatorMutation object of the builder.
func (cuo *CreatorUpdateOne) Mutation() *CreatorMutation {
	return cuo.mutation
}

// ClearClips clears all "clips" edges to the Clip entity.
func (cuo *CreatorUpdateOne) ClearClips() *CreatorUpdateOne {
	cuo.mutation.ClearClips()
	return cuo
}

// RemoveClipIDs removes the "clips" edge to Clip entities by IDs.
func (cuo *CreatorUpdateOne) RemoveClipIDs(ids ...int) *CreatorUpdateOne {
	cuo.mutation.RemoveClipIDs(ids...)
	return cuo
}

// RemoveClips removes "clips" edges to Clip entities.
func (cuo *CreatorUpdateOne) RemoveClips(c ...*Clip) *CreatorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveClipIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CreatorUpdateOne) Select(field string, fields ...string) *CreatorUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Creator entity.
func (cuo *CreatorUpdateOne) Save(ctx context.Context) (*Creator, error) {
	var (
		err  error
		node *Creator
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CreatorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CreatorUpdateOne) SaveX(ctx context.Context) *Creator {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CreatorUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CreatorUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CreatorUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := creator.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Creator.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *CreatorUpdateOne) sqlSave(ctx context.Context) (_node *Creator, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   creator.Table,
			Columns: creator.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: creator.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Creator.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, creator.FieldID)
		for _, f := range fields {
			if !creator.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != creator.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: creator.FieldName,
		})
	}
	if cuo.mutation.ClipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   creator.ClipsTable,
			Columns: []string{creator.ClipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clip.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedClipsIDs(); len(nodes) > 0 && !cuo.mutation.ClipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   creator.ClipsTable,
			Columns: []string{creator.ClipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clip.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   creator.ClipsTable,
			Columns: []string{creator.ClipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clip.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Creator{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{creator.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}