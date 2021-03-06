// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-club/ent/likeclub"
	"github.com/khu-dev/khumu-club/ent/predicate"
)

// LikeClubDelete is the builder for deleting a LikeClub entity.
type LikeClubDelete struct {
	config
	hooks    []Hook
	mutation *LikeClubMutation
}

// Where adds a new predicate to the LikeClubDelete builder.
func (lcd *LikeClubDelete) Where(ps ...predicate.LikeClub) *LikeClubDelete {
	lcd.mutation.predicates = append(lcd.mutation.predicates, ps...)
	return lcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lcd *LikeClubDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lcd.hooks) == 0 {
		affected, err = lcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeClubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lcd.mutation = mutation
			affected, err = lcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lcd.hooks) - 1; i >= 0; i-- {
			mut = lcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcd *LikeClubDelete) ExecX(ctx context.Context) int {
	n, err := lcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lcd *LikeClubDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: likeclub.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: likeclub.FieldID,
			},
		},
	}
	if ps := lcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, lcd.driver, _spec)
}

// LikeClubDeleteOne is the builder for deleting a single LikeClub entity.
type LikeClubDeleteOne struct {
	lcd *LikeClubDelete
}

// Exec executes the deletion query.
func (lcdo *LikeClubDeleteOne) Exec(ctx context.Context) error {
	n, err := lcdo.lcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{likeclub.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lcdo *LikeClubDeleteOne) ExecX(ctx context.Context) {
	lcdo.lcd.ExecX(ctx)
}
