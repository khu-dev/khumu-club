// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/khu-dev/khumu-club/ent/club"
	"github.com/khu-dev/khumu-club/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ClubDelete is the builder for deleting a Club entity.
type ClubDelete struct {
	config
	hooks    []Hook
	mutation *ClubMutation
}

// Where adds a new predicate to the ClubDelete builder.
func (cd *ClubDelete) Where(ps ...predicate.Club) *ClubDelete {
	cd.mutation.predicates = append(cd.mutation.predicates, ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ClubDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cd.hooks) == 0 {
		affected, err = cd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cd.mutation = mutation
			affected, err = cd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cd.hooks) - 1; i >= 0; i-- {
			mut = cd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ClubDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *ClubDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: club.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: club.FieldID,
			},
		},
	}
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
}

// ClubDeleteOne is the builder for deleting a single Club entity.
type ClubDeleteOne struct {
	cd *ClubDelete
}

// Exec executes the deletion query.
func (cdo *ClubDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{club.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ClubDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
