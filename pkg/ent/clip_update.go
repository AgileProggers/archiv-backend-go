// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/clip"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/creator"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/game"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/predicate"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/vod"
)

// ClipUpdate is the builder for updating Clip entities.
type ClipUpdate struct {
	config
	hooks    []Hook
	mutation *ClipMutation
}

// Where appends a list predicates to the ClipUpdate builder.
func (cu *ClipUpdate) Where(ps ...predicate.Clip) *ClipUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTitle sets the "title" field.
func (cu *ClipUpdate) SetTitle(s string) *ClipUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetDuration sets the "duration" field.
func (cu *ClipUpdate) SetDuration(i int) *ClipUpdate {
	cu.mutation.ResetDuration()
	cu.mutation.SetDuration(i)
	return cu
}

// AddDuration adds i to the "duration" field.
func (cu *ClipUpdate) AddDuration(i int) *ClipUpdate {
	cu.mutation.AddDuration(i)
	return cu
}

// SetDate sets the "date" field.
func (cu *ClipUpdate) SetDate(t time.Time) *ClipUpdate {
	cu.mutation.SetDate(t)
	return cu
}

// SetFilename sets the "filename" field.
func (cu *ClipUpdate) SetFilename(s string) *ClipUpdate {
	cu.mutation.SetFilename(s)
	return cu
}

// SetResolution sets the "resolution" field.
func (cu *ClipUpdate) SetResolution(s string) *ClipUpdate {
	cu.mutation.SetResolution(s)
	return cu
}

// SetSize sets the "size" field.
func (cu *ClipUpdate) SetSize(i int) *ClipUpdate {
	cu.mutation.ResetSize()
	cu.mutation.SetSize(i)
	return cu
}

// AddSize adds i to the "size" field.
func (cu *ClipUpdate) AddSize(i int) *ClipUpdate {
	cu.mutation.AddSize(i)
	return cu
}

// SetViewCount sets the "view_count" field.
func (cu *ClipUpdate) SetViewCount(i int) *ClipUpdate {
	cu.mutation.ResetViewCount()
	cu.mutation.SetViewCount(i)
	return cu
}

// AddViewCount adds i to the "view_count" field.
func (cu *ClipUpdate) AddViewCount(i int) *ClipUpdate {
	cu.mutation.AddViewCount(i)
	return cu
}

// SetCreatorID sets the "creator" edge to the Creator entity by ID.
func (cu *ClipUpdate) SetCreatorID(id int) *ClipUpdate {
	cu.mutation.SetCreatorID(id)
	return cu
}

// SetNillableCreatorID sets the "creator" edge to the Creator entity by ID if the given value is not nil.
func (cu *ClipUpdate) SetNillableCreatorID(id *int) *ClipUpdate {
	if id != nil {
		cu = cu.SetCreatorID(*id)
	}
	return cu
}

// SetCreator sets the "creator" edge to the Creator entity.
func (cu *ClipUpdate) SetCreator(c *Creator) *ClipUpdate {
	return cu.SetCreatorID(c.ID)
}

// AddVodIDs adds the "vod" edge to the Vod entity by IDs.
func (cu *ClipUpdate) AddVodIDs(ids ...int) *ClipUpdate {
	cu.mutation.AddVodIDs(ids...)
	return cu
}

// AddVod adds the "vod" edges to the Vod entity.
func (cu *ClipUpdate) AddVod(v ...*Vod) *ClipUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.AddVodIDs(ids...)
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (cu *ClipUpdate) AddGameIDs(ids ...int) *ClipUpdate {
	cu.mutation.AddGameIDs(ids...)
	return cu
}

// AddGame adds the "game" edges to the Game entity.
func (cu *ClipUpdate) AddGame(g ...*Game) *ClipUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cu.AddGameIDs(ids...)
}

// Mutation returns the ClipMutation object of the builder.
func (cu *ClipUpdate) Mutation() *ClipMutation {
	return cu.mutation
}

// ClearCreator clears the "creator" edge to the Creator entity.
func (cu *ClipUpdate) ClearCreator() *ClipUpdate {
	cu.mutation.ClearCreator()
	return cu
}

// ClearVod clears all "vod" edges to the Vod entity.
func (cu *ClipUpdate) ClearVod() *ClipUpdate {
	cu.mutation.ClearVod()
	return cu
}

// RemoveVodIDs removes the "vod" edge to Vod entities by IDs.
func (cu *ClipUpdate) RemoveVodIDs(ids ...int) *ClipUpdate {
	cu.mutation.RemoveVodIDs(ids...)
	return cu
}

// RemoveVod removes "vod" edges to Vod entities.
func (cu *ClipUpdate) RemoveVod(v ...*Vod) *ClipUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.RemoveVodIDs(ids...)
}

// ClearGame clears all "game" edges to the Game entity.
func (cu *ClipUpdate) ClearGame() *ClipUpdate {
	cu.mutation.ClearGame()
	return cu
}

// RemoveGameIDs removes the "game" edge to Game entities by IDs.
func (cu *ClipUpdate) RemoveGameIDs(ids ...int) *ClipUpdate {
	cu.mutation.RemoveGameIDs(ids...)
	return cu
}

// RemoveGame removes "game" edges to Game entities.
func (cu *ClipUpdate) RemoveGame(g ...*Game) *ClipUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cu.RemoveGameIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClipUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*ClipMutation)
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
func (cu *ClipUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClipUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClipUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ClipUpdate) check() error {
	if v, ok := cu.mutation.Title(); ok {
		if err := clip.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Clip.title": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Duration(); ok {
		if err := clip.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Clip.duration": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Filename(); ok {
		if err := clip.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf(`ent: validator failed for field "Clip.filename": %w`, err)}
		}
	}
	return nil
}

func (cu *ClipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clip.Table,
			Columns: clip.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clip.FieldID,
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
	if value, ok := cu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clip.FieldTitle,
		})
	}
	if value, ok := cu.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldDuration,
		})
	}
	if value, ok := cu.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldDuration,
		})
	}
	if value, ok := cu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clip.FieldDate,
		})
	}
	if value, ok := cu.mutation.Filename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clip.FieldFilename,
		})
	}
	if value, ok := cu.mutation.Resolution(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clip.FieldResolution,
		})
	}
	if value, ok := cu.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldSize,
		})
	}
	if value, ok := cu.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldSize,
		})
	}
	if value, ok := cu.mutation.ViewCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldViewCount,
		})
	}
	if value, ok := cu.mutation.AddedViewCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldViewCount,
		})
	}
	if cu.mutation.CreatorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CreatorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.VodCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedVodIDs(); len(nodes) > 0 && !cu.mutation.VodCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.VodIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.GameCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedGameIDs(); len(nodes) > 0 && !cu.mutation.GameCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.GameIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clip.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ClipUpdateOne is the builder for updating a single Clip entity.
type ClipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClipMutation
}

// SetTitle sets the "title" field.
func (cuo *ClipUpdateOne) SetTitle(s string) *ClipUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetDuration sets the "duration" field.
func (cuo *ClipUpdateOne) SetDuration(i int) *ClipUpdateOne {
	cuo.mutation.ResetDuration()
	cuo.mutation.SetDuration(i)
	return cuo
}

// AddDuration adds i to the "duration" field.
func (cuo *ClipUpdateOne) AddDuration(i int) *ClipUpdateOne {
	cuo.mutation.AddDuration(i)
	return cuo
}

// SetDate sets the "date" field.
func (cuo *ClipUpdateOne) SetDate(t time.Time) *ClipUpdateOne {
	cuo.mutation.SetDate(t)
	return cuo
}

// SetFilename sets the "filename" field.
func (cuo *ClipUpdateOne) SetFilename(s string) *ClipUpdateOne {
	cuo.mutation.SetFilename(s)
	return cuo
}

// SetResolution sets the "resolution" field.
func (cuo *ClipUpdateOne) SetResolution(s string) *ClipUpdateOne {
	cuo.mutation.SetResolution(s)
	return cuo
}

// SetSize sets the "size" field.
func (cuo *ClipUpdateOne) SetSize(i int) *ClipUpdateOne {
	cuo.mutation.ResetSize()
	cuo.mutation.SetSize(i)
	return cuo
}

// AddSize adds i to the "size" field.
func (cuo *ClipUpdateOne) AddSize(i int) *ClipUpdateOne {
	cuo.mutation.AddSize(i)
	return cuo
}

// SetViewCount sets the "view_count" field.
func (cuo *ClipUpdateOne) SetViewCount(i int) *ClipUpdateOne {
	cuo.mutation.ResetViewCount()
	cuo.mutation.SetViewCount(i)
	return cuo
}

// AddViewCount adds i to the "view_count" field.
func (cuo *ClipUpdateOne) AddViewCount(i int) *ClipUpdateOne {
	cuo.mutation.AddViewCount(i)
	return cuo
}

// SetCreatorID sets the "creator" edge to the Creator entity by ID.
func (cuo *ClipUpdateOne) SetCreatorID(id int) *ClipUpdateOne {
	cuo.mutation.SetCreatorID(id)
	return cuo
}

// SetNillableCreatorID sets the "creator" edge to the Creator entity by ID if the given value is not nil.
func (cuo *ClipUpdateOne) SetNillableCreatorID(id *int) *ClipUpdateOne {
	if id != nil {
		cuo = cuo.SetCreatorID(*id)
	}
	return cuo
}

// SetCreator sets the "creator" edge to the Creator entity.
func (cuo *ClipUpdateOne) SetCreator(c *Creator) *ClipUpdateOne {
	return cuo.SetCreatorID(c.ID)
}

// AddVodIDs adds the "vod" edge to the Vod entity by IDs.
func (cuo *ClipUpdateOne) AddVodIDs(ids ...int) *ClipUpdateOne {
	cuo.mutation.AddVodIDs(ids...)
	return cuo
}

// AddVod adds the "vod" edges to the Vod entity.
func (cuo *ClipUpdateOne) AddVod(v ...*Vod) *ClipUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.AddVodIDs(ids...)
}

// AddGameIDs adds the "game" edge to the Game entity by IDs.
func (cuo *ClipUpdateOne) AddGameIDs(ids ...int) *ClipUpdateOne {
	cuo.mutation.AddGameIDs(ids...)
	return cuo
}

// AddGame adds the "game" edges to the Game entity.
func (cuo *ClipUpdateOne) AddGame(g ...*Game) *ClipUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cuo.AddGameIDs(ids...)
}

// Mutation returns the ClipMutation object of the builder.
func (cuo *ClipUpdateOne) Mutation() *ClipMutation {
	return cuo.mutation
}

// ClearCreator clears the "creator" edge to the Creator entity.
func (cuo *ClipUpdateOne) ClearCreator() *ClipUpdateOne {
	cuo.mutation.ClearCreator()
	return cuo
}

// ClearVod clears all "vod" edges to the Vod entity.
func (cuo *ClipUpdateOne) ClearVod() *ClipUpdateOne {
	cuo.mutation.ClearVod()
	return cuo
}

// RemoveVodIDs removes the "vod" edge to Vod entities by IDs.
func (cuo *ClipUpdateOne) RemoveVodIDs(ids ...int) *ClipUpdateOne {
	cuo.mutation.RemoveVodIDs(ids...)
	return cuo
}

// RemoveVod removes "vod" edges to Vod entities.
func (cuo *ClipUpdateOne) RemoveVod(v ...*Vod) *ClipUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.RemoveVodIDs(ids...)
}

// ClearGame clears all "game" edges to the Game entity.
func (cuo *ClipUpdateOne) ClearGame() *ClipUpdateOne {
	cuo.mutation.ClearGame()
	return cuo
}

// RemoveGameIDs removes the "game" edge to Game entities by IDs.
func (cuo *ClipUpdateOne) RemoveGameIDs(ids ...int) *ClipUpdateOne {
	cuo.mutation.RemoveGameIDs(ids...)
	return cuo
}

// RemoveGame removes "game" edges to Game entities.
func (cuo *ClipUpdateOne) RemoveGame(g ...*Game) *ClipUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cuo.RemoveGameIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ClipUpdateOne) Select(field string, fields ...string) *ClipUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Clip entity.
func (cuo *ClipUpdateOne) Save(ctx context.Context) (*Clip, error) {
	var (
		err  error
		node *Clip
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClipMutation)
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
func (cuo *ClipUpdateOne) SaveX(ctx context.Context) *Clip {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClipUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClipUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ClipUpdateOne) check() error {
	if v, ok := cuo.mutation.Title(); ok {
		if err := clip.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Clip.title": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Duration(); ok {
		if err := clip.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Clip.duration": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Filename(); ok {
		if err := clip.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf(`ent: validator failed for field "Clip.filename": %w`, err)}
		}
	}
	return nil
}

func (cuo *ClipUpdateOne) sqlSave(ctx context.Context) (_node *Clip, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clip.Table,
			Columns: clip.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: clip.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Clip.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clip.FieldID)
		for _, f := range fields {
			if !clip.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != clip.FieldID {
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
	if value, ok := cuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clip.FieldTitle,
		})
	}
	if value, ok := cuo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldDuration,
		})
	}
	if value, ok := cuo.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldDuration,
		})
	}
	if value, ok := cuo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clip.FieldDate,
		})
	}
	if value, ok := cuo.mutation.Filename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clip.FieldFilename,
		})
	}
	if value, ok := cuo.mutation.Resolution(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clip.FieldResolution,
		})
	}
	if value, ok := cuo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldSize,
		})
	}
	if value, ok := cuo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldSize,
		})
	}
	if value, ok := cuo.mutation.ViewCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldViewCount,
		})
	}
	if value, ok := cuo.mutation.AddedViewCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: clip.FieldViewCount,
		})
	}
	if cuo.mutation.CreatorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CreatorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.VodCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedVodIDs(); len(nodes) > 0 && !cuo.mutation.VodCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.VodIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.GameCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedGameIDs(); len(nodes) > 0 && !cuo.mutation.GameCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.GameIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Clip{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clip.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
