// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/predicate"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/vod"
)

// VodDelete is the builder for deleting a Vod entity.
type VodDelete struct {
	config
	hooks    []Hook
	mutation *VodMutation
}

// Where appends a list predicates to the VodDelete builder.
func (vd *VodDelete) Where(ps ...predicate.Vod) *VodDelete {
	vd.mutation.Where(ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VodDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vd.hooks) == 0 {
		affected, err = vd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vd.mutation = mutation
			affected, err = vd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vd.hooks) - 1; i >= 0; i-- {
			if vd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VodDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VodDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vod.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vod.FieldID,
			},
		},
	}
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
}

// VodDeleteOne is the builder for deleting a single Vod entity.
type VodDeleteOne struct {
	vd *VodDelete
}

// Exec executes the deletion query.
func (vdo *VodDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vod.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VodDeleteOne) ExecX(ctx context.Context) {
	vdo.vd.ExecX(ctx)
}
