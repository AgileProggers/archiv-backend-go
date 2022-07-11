// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/emote"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/predicate"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/provider"
)

// EmoteUpdate is the builder for updating Emote entities.
type EmoteUpdate struct {
	config
	hooks    []Hook
	mutation *EmoteMutation
}

// Where appends a list predicates to the EmoteUpdate builder.
func (eu *EmoteUpdate) Where(ps ...predicate.Emote) *EmoteUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EmoteUpdate) SetName(s string) *EmoteUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetURL sets the "url" field.
func (eu *EmoteUpdate) SetURL(s string) *EmoteUpdate {
	eu.mutation.SetURL(s)
	return eu
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (eu *EmoteUpdate) SetProviderID(id int) *EmoteUpdate {
	eu.mutation.SetProviderID(id)
	return eu
}

// SetNillableProviderID sets the "provider" edge to the Provider entity by ID if the given value is not nil.
func (eu *EmoteUpdate) SetNillableProviderID(id *int) *EmoteUpdate {
	if id != nil {
		eu = eu.SetProviderID(*id)
	}
	return eu
}

// SetProvider sets the "provider" edge to the Provider entity.
func (eu *EmoteUpdate) SetProvider(p *Provider) *EmoteUpdate {
	return eu.SetProviderID(p.ID)
}

// Mutation returns the EmoteMutation object of the builder.
func (eu *EmoteUpdate) Mutation() *EmoteMutation {
	return eu.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (eu *EmoteUpdate) ClearProvider() *EmoteUpdate {
	eu.mutation.ClearProvider()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EmoteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EmoteUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EmoteUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EmoteUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EmoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emote.Table,
			Columns: emote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: emote.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emote.FieldName,
		})
	}
	if value, ok := eu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emote.FieldURL,
		})
	}
	if eu.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emote.ProviderTable,
			Columns: []string{emote.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: provider.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emote.ProviderTable,
			Columns: []string{emote.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: provider.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EmoteUpdateOne is the builder for updating a single Emote entity.
type EmoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmoteMutation
}

// SetName sets the "name" field.
func (euo *EmoteUpdateOne) SetName(s string) *EmoteUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetURL sets the "url" field.
func (euo *EmoteUpdateOne) SetURL(s string) *EmoteUpdateOne {
	euo.mutation.SetURL(s)
	return euo
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (euo *EmoteUpdateOne) SetProviderID(id int) *EmoteUpdateOne {
	euo.mutation.SetProviderID(id)
	return euo
}

// SetNillableProviderID sets the "provider" edge to the Provider entity by ID if the given value is not nil.
func (euo *EmoteUpdateOne) SetNillableProviderID(id *int) *EmoteUpdateOne {
	if id != nil {
		euo = euo.SetProviderID(*id)
	}
	return euo
}

// SetProvider sets the "provider" edge to the Provider entity.
func (euo *EmoteUpdateOne) SetProvider(p *Provider) *EmoteUpdateOne {
	return euo.SetProviderID(p.ID)
}

// Mutation returns the EmoteMutation object of the builder.
func (euo *EmoteUpdateOne) Mutation() *EmoteMutation {
	return euo.mutation
}

// ClearProvider clears the "provider" edge to the Provider entity.
func (euo *EmoteUpdateOne) ClearProvider() *EmoteUpdateOne {
	euo.mutation.ClearProvider()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EmoteUpdateOne) Select(field string, fields ...string) *EmoteUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Emote entity.
func (euo *EmoteUpdateOne) Save(ctx context.Context) (*Emote, error) {
	var (
		err  error
		node *Emote
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EmoteUpdateOne) SaveX(ctx context.Context) *Emote {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EmoteUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EmoteUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EmoteUpdateOne) sqlSave(ctx context.Context) (_node *Emote, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emote.Table,
			Columns: emote.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: emote.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Emote.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emote.FieldID)
		for _, f := range fields {
			if !emote.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emote.FieldName,
		})
	}
	if value, ok := euo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emote.FieldURL,
		})
	}
	if euo.mutation.ProviderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emote.ProviderTable,
			Columns: []string{emote.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: provider.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emote.ProviderTable,
			Columns: []string{emote.ProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: provider.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Emote{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}