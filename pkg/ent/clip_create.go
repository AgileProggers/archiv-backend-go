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
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/creator"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/game"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/vod"
)

// ClipCreate is the builder for creating a Clip entity.
type ClipCreate struct {
	config
	mutation *ClipMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (cc *ClipCreate) SetTitle(s string) *ClipCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetDuration sets the "duration" field.
func (cc *ClipCreate) SetDuration(i int) *ClipCreate {
	cc.mutation.SetDuration(i)
	return cc
}

// SetDate sets the "date" field.
func (cc *ClipCreate) SetDate(t time.Time) *ClipCreate {
	cc.mutation.SetDate(t)
	return cc
}

// SetFilename sets the "filename" field.
func (cc *ClipCreate) SetFilename(s string) *ClipCreate {
	cc.mutation.SetFilename(s)
	return cc
}

// SetResolution sets the "resolution" field.
func (cc *ClipCreate) SetResolution(s string) *ClipCreate {
	cc.mutation.SetResolution(s)
	return cc
}

// SetSize sets the "size" field.
func (cc *ClipCreate) SetSize(i int) *ClipCreate {
	cc.mutation.SetSize(i)
	return cc
}

// SetViewCount sets the "view_count" field.
func (cc *ClipCreate) SetViewCount(i int) *ClipCreate {
	cc.mutation.SetViewCount(i)
	return cc
}

// SetCreatorID sets the "creator" edge to the Creator entity by ID.
func (cc *ClipCreate) SetCreatorID(id int) *ClipCreate {
	cc.mutation.SetCreatorID(id)
	return cc
}

// SetNillableCreatorID sets the "creator" edge to the Creator entity by ID if the given value is not nil.
func (cc *ClipCreate) SetNillableCreatorID(id *int) *ClipCreate {
	if id != nil {
		cc = cc.SetCreatorID(*id)
	}
	return cc
}

// SetCreator sets the "creator" edge to the Creator entity.
func (cc *ClipCreate) SetCreator(c *Creator) *ClipCreate {
	return cc.SetCreatorID(c.ID)
}

// AddVodIDs adds the "vod" edge to the Vod entity by IDs.
func (cc *ClipCreate) AddVodIDs(ids ...int) *ClipCreate {
	cc.mutation.AddVodIDs(ids...)
	return cc
}

// AddVod adds the "vod" edges to the Vod entity.
func (cc *ClipCreate) AddVod(v ...*Vod) *ClipCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cc.AddVodIDs(ids...)
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (cc *ClipCreate) AddGameIDs(ids ...int) *ClipCreate {
	cc.mutation.AddGameIDs(ids...)
	return cc
}

// AddGame adds the "game" edges to the Game entity.
func (cc *ClipCreate) AddGame(g ...*Game) *ClipCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cc.AddGameIDs(ids...)
}

// Mutation returns the ClipMutation object of the builder.
func (cc *ClipCreate) Mutation() *ClipMutation {
	return cc.mutation
}

// Save creates the Clip in the database.
func (cc *ClipCreate) Save(ctx context.Context) (*Clip, error) {
	var (
		err  error
		node *Clip
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClipCreate) SaveX(ctx context.Context) *Clip {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClipCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClipCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClipCreate) check() error {
	if _, ok := cc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Clip.title"`)}
	}
	if v, ok := cc.mutation.Title(); ok {
		if err := clip.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Clip.title": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "Clip.duration"`)}
	}
	if v, ok := cc.mutation.Duration(); ok {
		if err := clip.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Clip.duration": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Clip.date"`)}
	}
	if _, ok := cc.mutation.Filename(); !ok {
		return &ValidationError{Name: "filename", err: errors.New(`ent: missing required field "Clip.filename"`)}
	}
	if v, ok := cc.mutation.Filename(); ok {
		if err := clip.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf(`ent: validator failed for field "Clip.filename": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Resolution(); !ok {
		return &ValidationError{Name: "resolution", err: errors.New(`ent: missing required field "Clip.resolution"`)}
	}
	if _, ok := cc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Clip.size"`)}
	}
	if _, ok := cc.mutation.ViewCount(); !ok {
		return &ValidationError{Name: "view_count", err: errors.New(`ent: missing required field "Clip.view_count"`)}
	}
	return nil
}

func (cc *ClipCreate) sqlSave(ctx context.Context) (*Clip, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *ClipCreate) createSpec() (*Clip, *sqlgraph.CreateSpec) {
	var (
		_node = &Clip{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: clip.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clip.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clip.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := cc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldDuration,
		})
		_node.Duration = value
	}
	if value, ok := cc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clip.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := cc.mutation.Filename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clip.FieldFilename,
		})
		_node.Filename = value
	}
	if value, ok := cc.mutation.Resolution(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clip.FieldResolution,
		})
		_node.Resolution = value
	}
	if value, ok := cc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := cc.mutation.ViewCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldViewCount,
		})
		_node.ViewCount = value
	}
	if nodes := cc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   clip.CreatorTable,
			Columns: []string{clip.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: creator.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.creator_clips = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.VodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   clip.VodTable,
			Columns: clip.VodPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   clip.GameTable,
			Columns: []string{clip.GameColumn},
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

// ClipCreateBulk is the builder for creating many Clip entities in bulk.
type ClipCreateBulk struct {
	config
	builders []*ClipCreate
}

// Save creates the Clip entities in the database.
func (ccb *ClipCreateBulk) Save(ctx context.Context) ([]*Clip, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Clip, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClipMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClipCreateBulk) SaveX(ctx context.Context) []*Clip {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClipCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClipCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
