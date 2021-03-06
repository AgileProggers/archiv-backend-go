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
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/game"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/predicate"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/vod"
)

// GameUpdate is the builder for updating Game entities.
type GameUpdate struct {
	config
	hooks    []Hook
	mutation *GameMutation
}

// Where appends a list predicates to the GameUpdate builder.
func (gu *GameUpdate) Where(ps ...predicate.Game) *GameUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetGameID sets the "game_id" field.
func (gu *GameUpdate) SetGameID(i int) *GameUpdate {
	gu.mutation.ResetGameID()
	gu.mutation.SetGameID(i)
	return gu
}

// AddGameID adds i to the "game_id" field.
func (gu *GameUpdate) AddGameID(i int) *GameUpdate {
	gu.mutation.AddGameID(i)
	return gu
}

// SetName sets the "name" field.
func (gu *GameUpdate) SetName(s string) *GameUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetBoxArt sets the "box_art" field.
func (gu *GameUpdate) SetBoxArt(s string) *GameUpdate {
	gu.mutation.SetBoxArt(s)
	return gu
}

// SetClipID sets the "clip" edge to the Clip entity by ID.
func (gu *GameUpdate) SetClipID(id int) *GameUpdate {
	gu.mutation.SetClipID(id)
	return gu
}

// SetNillableClipID sets the "clip" edge to the Clip entity by ID if the given value is not nil.
func (gu *GameUpdate) SetNillableClipID(id *int) *GameUpdate {
	if id != nil {
		gu = gu.SetClipID(*id)
	}
	return gu
}

// SetClip sets the "clip" edge to the Clip entity.
func (gu *GameUpdate) SetClip(c *Clip) *GameUpdate {
	return gu.SetClipID(c.ID)
}

// SetVodID sets the "vod" edge to the Vod entity by ID.
func (gu *GameUpdate) SetVodID(id int) *GameUpdate {
	gu.mutation.SetVodID(id)
	return gu
}

// SetNillableVodID sets the "vod" edge to the Vod entity by ID if the given value is not nil.
func (gu *GameUpdate) SetNillableVodID(id *int) *GameUpdate {
	if id != nil {
		gu = gu.SetVodID(*id)
	}
	return gu
}

// SetVod sets the "vod" edge to the Vod entity.
func (gu *GameUpdate) SetVod(v *Vod) *GameUpdate {
	return gu.SetVodID(v.ID)
}

// Mutation returns the GameMutation object of the builder.
func (gu *GameUpdate) Mutation() *GameMutation {
	return gu.mutation
}

// ClearClip clears the "clip" edge to the Clip entity.
func (gu *GameUpdate) ClearClip() *GameUpdate {
	gu.mutation.ClearClip()
	return gu
}

// ClearVod clears the "vod" edge to the Vod entity.
func (gu *GameUpdate) ClearVod() *GameUpdate {
	gu.mutation.ClearVod()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GameUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GameUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GameUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GameUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GameUpdate) check() error {
	if v, ok := gu.mutation.GameID(); ok {
		if err := game.GameIDValidator(v); err != nil {
			return &ValidationError{Name: "game_id", err: fmt.Errorf(`ent: validator failed for field "Game.game_id": %w`, err)}
		}
	}
	return nil
}

func (gu *GameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: game.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.GameID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: game.FieldGameID,
		})
	}
	if value, ok := gu.mutation.AddedGameID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: game.FieldGameID,
		})
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldName,
		})
	}
	if value, ok := gu.mutation.BoxArt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldBoxArt,
		})
	}
	if gu.mutation.ClipCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.ClipTable,
			Columns: []string{game.ClipColumn},
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
	if nodes := gu.mutation.ClipIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.ClipTable,
			Columns: []string{game.ClipColumn},
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
	if gu.mutation.VodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.VodTable,
			Columns: []string{game.VodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.VodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.VodTable,
			Columns: []string{game.VodColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GameUpdateOne is the builder for updating a single Game entity.
type GameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameMutation
}

// SetGameID sets the "game_id" field.
func (guo *GameUpdateOne) SetGameID(i int) *GameUpdateOne {
	guo.mutation.ResetGameID()
	guo.mutation.SetGameID(i)
	return guo
}

// AddGameID adds i to the "game_id" field.
func (guo *GameUpdateOne) AddGameID(i int) *GameUpdateOne {
	guo.mutation.AddGameID(i)
	return guo
}

// SetName sets the "name" field.
func (guo *GameUpdateOne) SetName(s string) *GameUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetBoxArt sets the "box_art" field.
func (guo *GameUpdateOne) SetBoxArt(s string) *GameUpdateOne {
	guo.mutation.SetBoxArt(s)
	return guo
}

// SetClipID sets the "clip" edge to the Clip entity by ID.
func (guo *GameUpdateOne) SetClipID(id int) *GameUpdateOne {
	guo.mutation.SetClipID(id)
	return guo
}

// SetNillableClipID sets the "clip" edge to the Clip entity by ID if the given value is not nil.
func (guo *GameUpdateOne) SetNillableClipID(id *int) *GameUpdateOne {
	if id != nil {
		guo = guo.SetClipID(*id)
	}
	return guo
}

// SetClip sets the "clip" edge to the Clip entity.
func (guo *GameUpdateOne) SetClip(c *Clip) *GameUpdateOne {
	return guo.SetClipID(c.ID)
}

// SetVodID sets the "vod" edge to the Vod entity by ID.
func (guo *GameUpdateOne) SetVodID(id int) *GameUpdateOne {
	guo.mutation.SetVodID(id)
	return guo
}

// SetNillableVodID sets the "vod" edge to the Vod entity by ID if the given value is not nil.
func (guo *GameUpdateOne) SetNillableVodID(id *int) *GameUpdateOne {
	if id != nil {
		guo = guo.SetVodID(*id)
	}
	return guo
}

// SetVod sets the "vod" edge to the Vod entity.
func (guo *GameUpdateOne) SetVod(v *Vod) *GameUpdateOne {
	return guo.SetVodID(v.ID)
}

// Mutation returns the GameMutation object of the builder.
func (guo *GameUpdateOne) Mutation() *GameMutation {
	return guo.mutation
}

// ClearClip clears the "clip" edge to the Clip entity.
func (guo *GameUpdateOne) ClearClip() *GameUpdateOne {
	guo.mutation.ClearClip()
	return guo
}

// ClearVod clears the "vod" edge to the Vod entity.
func (guo *GameUpdateOne) ClearVod() *GameUpdateOne {
	guo.mutation.ClearVod()
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GameUpdateOne) Select(field string, fields ...string) *GameUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Game entity.
func (guo *GameUpdateOne) Save(ctx context.Context) (*Game, error) {
	var (
		err  error
		node *Game
	)
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GameUpdateOne) SaveX(ctx context.Context) *Game {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GameUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GameUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GameUpdateOne) check() error {
	if v, ok := guo.mutation.GameID(); ok {
		if err := game.GameIDValidator(v); err != nil {
			return &ValidationError{Name: "game_id", err: fmt.Errorf(`ent: validator failed for field "Game.game_id": %w`, err)}
		}
	}
	return nil
}

func (guo *GameUpdateOne) sqlSave(ctx context.Context) (_node *Game, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: game.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Game.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, game.FieldID)
		for _, f := range fields {
			if !game.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != game.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.GameID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: game.FieldGameID,
		})
	}
	if value, ok := guo.mutation.AddedGameID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: game.FieldGameID,
		})
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldName,
		})
	}
	if value, ok := guo.mutation.BoxArt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: game.FieldBoxArt,
		})
	}
	if guo.mutation.ClipCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.ClipTable,
			Columns: []string{game.ClipColumn},
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
	if nodes := guo.mutation.ClipIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.ClipTable,
			Columns: []string{game.ClipColumn},
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
	if guo.mutation.VodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.VodTable,
			Columns: []string{game.VodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: vod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.VodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.VodTable,
			Columns: []string{game.VodColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Game{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
