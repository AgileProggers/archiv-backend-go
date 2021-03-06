// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/emote"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/provider"
)

// EmoteCreate is the builder for creating a Emote entity.
type EmoteCreate struct {
	config
	mutation *EmoteMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ec *EmoteCreate) SetName(s string) *EmoteCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetURL sets the "url" field.
func (ec *EmoteCreate) SetURL(s string) *EmoteCreate {
	ec.mutation.SetURL(s)
	return ec
}

// SetProviderID sets the "provider" edge to the Provider entity by ID.
func (ec *EmoteCreate) SetProviderID(id int) *EmoteCreate {
	ec.mutation.SetProviderID(id)
	return ec
}

// SetNillableProviderID sets the "provider" edge to the Provider entity by ID if the given value is not nil.
func (ec *EmoteCreate) SetNillableProviderID(id *int) *EmoteCreate {
	if id != nil {
		ec = ec.SetProviderID(*id)
	}
	return ec
}

// SetProvider sets the "provider" edge to the Provider entity.
func (ec *EmoteCreate) SetProvider(p *Provider) *EmoteCreate {
	return ec.SetProviderID(p.ID)
}

// Mutation returns the EmoteMutation object of the builder.
func (ec *EmoteCreate) Mutation() *EmoteMutation {
	return ec.mutation
}

// Save creates the Emote in the database.
func (ec *EmoteCreate) Save(ctx context.Context) (*Emote, error) {
	var (
		err  error
		node *Emote
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmoteCreate) SaveX(ctx context.Context) *Emote {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EmoteCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EmoteCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EmoteCreate) check() error {
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Emote.name"`)}
	}
	if _, ok := ec.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Emote.url"`)}
	}
	return nil
}

func (ec *EmoteCreate) sqlSave(ctx context.Context) (*Emote, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *EmoteCreate) createSpec() (*Emote, *sqlgraph.CreateSpec) {
	var (
		_node = &Emote{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: emote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: emote.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emote.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ec.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emote.FieldURL,
		})
		_node.URL = value
	}
	if nodes := ec.mutation.ProviderIDs(); len(nodes) > 0 {
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
		_node.provider_emotes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EmoteCreateBulk is the builder for creating many Emote entities in bulk.
type EmoteCreateBulk struct {
	config
	builders []*EmoteCreate
}

// Save creates the Emote entities in the database.
func (ecb *EmoteCreateBulk) Save(ctx context.Context) ([]*Emote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Emote, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmoteMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EmoteCreateBulk) SaveX(ctx context.Context) []*Emote {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EmoteCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EmoteCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
