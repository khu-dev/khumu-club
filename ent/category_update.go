// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-club/ent/category"
	"github.com/khu-dev/khumu-club/ent/club"
	"github.com/khu-dev/khumu-club/ent/predicate"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where adds a new predicate for the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// AddClubIDs adds the "clubs" edge to the Club entity by IDs.
func (cu *CategoryUpdate) AddClubIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddClubIDs(ids...)
	return cu
}

// AddClubs adds the "clubs" edges to the Club entity.
func (cu *CategoryUpdate) AddClubs(c ...*Club) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddClubIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearClubs clears all "clubs" edges to the Club entity.
func (cu *CategoryUpdate) ClearClubs() *CategoryUpdate {
	cu.mutation.ClearClubs()
	return cu
}

// RemoveClubIDs removes the "clubs" edge to Club entities by IDs.
func (cu *CategoryUpdate) RemoveClubIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveClubIDs(ids...)
	return cu
}

// RemoveClubs removes "clubs" edges to Club entities.
func (cu *CategoryUpdate) RemoveClubs(c ...*Club) *CategoryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveClubIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: category.FieldID,
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
	if cu.mutation.ClubsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ClubsTable,
			Columns: category.ClubsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedClubsIDs(); len(nodes) > 0 && !cu.mutation.ClubsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ClubsTable,
			Columns: category.ClubsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClubsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ClubsTable,
			Columns: category.ClubsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
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
			err = &NotFoundError{category.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// AddClubIDs adds the "clubs" edge to the Club entity by IDs.
func (cuo *CategoryUpdateOne) AddClubIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddClubIDs(ids...)
	return cuo
}

// AddClubs adds the "clubs" edges to the Club entity.
func (cuo *CategoryUpdateOne) AddClubs(c ...*Club) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddClubIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearClubs clears all "clubs" edges to the Club entity.
func (cuo *CategoryUpdateOne) ClearClubs() *CategoryUpdateOne {
	cuo.mutation.ClearClubs()
	return cuo
}

// RemoveClubIDs removes the "clubs" edge to Club entities by IDs.
func (cuo *CategoryUpdateOne) RemoveClubIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveClubIDs(ids...)
	return cuo
}

// RemoveClubs removes "clubs" edges to Club entities.
func (cuo *CategoryUpdateOne) RemoveClubs(c ...*Club) *CategoryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveClubIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	var (
		err  error
		node *Category
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: category.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Category.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
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
	if cuo.mutation.ClubsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ClubsTable,
			Columns: category.ClubsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedClubsIDs(); len(nodes) > 0 && !cuo.mutation.ClubsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ClubsTable,
			Columns: category.ClubsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClubsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ClubsTable,
			Columns: category.ClubsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
