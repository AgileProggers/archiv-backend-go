// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/clip"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/game"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/vod"
)

// VodCreate is the builder for creating a Vod entity.
type VodCreate struct {
	config
	mutation *VodMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (vc *VodCreate) SetTitle(s string) *VodCreate {
	vc.mutation.SetTitle(s)
	return vc
}

// SetDuration sets the "duration" field.
func (vc *VodCreate) SetDuration(i int) *VodCreate {
	vc.mutation.SetDuration(i)
	return vc
}

// SetDate sets the "date" field.
func (vc *VodCreate) SetDate(t time.Time) *VodCreate {
	vc.mutation.SetDate(t)
	return vc
}

// SetFilename sets the "filename" field.
func (vc *VodCreate) SetFilename(s string) *VodCreate {
	vc.mutation.SetFilename(s)
	return vc
}

// SetResolution sets the "resolution" field.
func (vc *VodCreate) SetResolution(s string) *VodCreate {
	vc.mutation.SetResolution(s)
	return vc
}

// SetFps sets the "fps" field.
func (vc *VodCreate) SetFps(f float64) *VodCreate {
	vc.mutation.SetFps(f)
	return vc
}

// SetSize sets the "size" field.
func (vc *VodCreate) SetSize(i int) *VodCreate {
	vc.mutation.SetSize(i)
	return vc
}

// SetPublish sets the "publish" field.
func (vc *VodCreate) SetPublish(b bool) *VodCreate {
	vc.mutation.SetPublish(b)
	return vc
}

// AddClipIDs adds the "clips" edge to the Clip entity by IDs.
func (vc *VodCreate) AddClipIDs(ids ...int) *VodCreate {
	vc.mutation.AddClipIDs(ids...)
	return vc
}

// AddClips adds the "clips" edges to the Clip entity.
func (vc *VodCreate) AddClips(c ...*Clip) *VodCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return vc.AddClipIDs(ids...)
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (vc *VodCreate) AddGameIDs(ids ...int) *VodCreate {
	vc.mutation.AddGameIDs(ids...)
	return vc
}

// AddGame adds the "game" edges to the Game entity.
func (vc *VodCreate) AddGame(g ...*Game) *VodCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return vc.AddGameIDs(ids...)
}

// Mutation returns the VodMutation object of the builder.
func (vc *VodCreate) Mutation() *VodMutation {
	return vc.mutation
}

// Save creates the Vod in the database.
func (vc *VodCreate) Save(ctx context.Context) (*Vod, error) {
	var (
		err  error
		node *Vod
	)
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VodCreate) SaveX(ctx context.Context) *Vod {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VodCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VodCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VodCreate) check() error {
	if _, ok := vc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Vod.title"`)}
	}
	if _, ok := vc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "Vod.duration"`)}
	}
	if _, ok := vc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Vod.date"`)}
	}
	if _, ok := vc.mutation.Filename(); !ok {
		return &ValidationError{Name: "filename", err: errors.New(`ent: missing required field "Vod.filename"`)}
	}
	if _, ok := vc.mutation.Resolution(); !ok {
		return &ValidationError{Name: "resolution", err: errors.New(`ent: missing required field "Vod.resolution"`)}
	}
	if _, ok := vc.mutation.Fps(); !ok {
		return &ValidationError{Name: "fps", err: errors.New(`ent: missing required field "Vod.fps"`)}
	}
	if _, ok := vc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Vod.size"`)}
	}
	if _, ok := vc.mutation.Publish(); !ok {
		return &ValidationError{Name: "publish", err: errors.New(`ent: missing required field "Vod.publish"`)}
	}
	return nil
}

func (vc *VodCreate) sqlSave(ctx context.Context) (*Vod, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vc *VodCreate) createSpec() (*Vod, *sqlgraph.CreateSpec) {
	var (
		_node = &Vod{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vod.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vod.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vod.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := vc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vod.FieldDuration,
		})
		_node.Duration = value
	}
	if value, ok := vc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vod.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := vc.mutation.Filename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vod.FieldFilename,
		})
		_node.Filename = value
	}
	if value, ok := vc.mutation.Resolution(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vod.FieldResolution,
		})
		_node.Resolution = value
	}
	if value, ok := vc.mutation.Fps(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: vod.FieldFps,
		})
		_node.Fps = value
	}
	if value, ok := vc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: vod.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := vc.mutation.Publish(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: vod.FieldPublish,
		})
		_node.Publish = value
	}
	if nodes := vc.mutation.ClipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   vod.ClipsTable,
			Columns: vod.ClipsPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vod.GameTable,
			Columns: []string{vod.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VodCreateBulk is the builder for creating many Vod entities in bulk.
type VodCreateBulk struct {
	config
	builders []*VodCreate
}

// Save creates the Vod entities in the database.
func (vcb *VodCreateBulk) Save(ctx context.Context) ([]*Vod, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vod, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VodMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VodCreateBulk) SaveX(ctx context.Context) []*Vod {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VodCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VodCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
