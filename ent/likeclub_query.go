// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-club/ent/club"
	"github.com/khu-dev/khumu-club/ent/likeclub"
	"github.com/khu-dev/khumu-club/ent/predicate"
)

// LikeClubQuery is the builder for querying LikeClub entities.
type LikeClubQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LikeClub
	// eager-loading edges.
	withClub *ClubQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LikeClubQuery builder.
func (lcq *LikeClubQuery) Where(ps ...predicate.LikeClub) *LikeClubQuery {
	lcq.predicates = append(lcq.predicates, ps...)
	return lcq
}

// Limit adds a limit step to the query.
func (lcq *LikeClubQuery) Limit(limit int) *LikeClubQuery {
	lcq.limit = &limit
	return lcq
}

// Offset adds an offset step to the query.
func (lcq *LikeClubQuery) Offset(offset int) *LikeClubQuery {
	lcq.offset = &offset
	return lcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lcq *LikeClubQuery) Unique(unique bool) *LikeClubQuery {
	lcq.unique = &unique
	return lcq
}

// Order adds an order step to the query.
func (lcq *LikeClubQuery) Order(o ...OrderFunc) *LikeClubQuery {
	lcq.order = append(lcq.order, o...)
	return lcq
}

// QueryClub chains the current query on the "club" edge.
func (lcq *LikeClubQuery) QueryClub() *ClubQuery {
	query := &ClubQuery{config: lcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(likeclub.Table, likeclub.FieldID, selector),
			sqlgraph.To(club.Table, club.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, likeclub.ClubTable, likeclub.ClubColumn),
		)
		fromU = sqlgraph.SetNeighbors(lcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LikeClub entity from the query.
// Returns a *NotFoundError when no LikeClub was found.
func (lcq *LikeClubQuery) First(ctx context.Context) (*LikeClub, error) {
	nodes, err := lcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{likeclub.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lcq *LikeClubQuery) FirstX(ctx context.Context) *LikeClub {
	node, err := lcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LikeClub ID from the query.
// Returns a *NotFoundError when no LikeClub ID was found.
func (lcq *LikeClubQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{likeclub.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lcq *LikeClubQuery) FirstIDX(ctx context.Context) int {
	id, err := lcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LikeClub entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one LikeClub entity is not found.
// Returns a *NotFoundError when no LikeClub entities are found.
func (lcq *LikeClubQuery) Only(ctx context.Context) (*LikeClub, error) {
	nodes, err := lcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{likeclub.Label}
	default:
		return nil, &NotSingularError{likeclub.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lcq *LikeClubQuery) OnlyX(ctx context.Context) *LikeClub {
	node, err := lcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LikeClub ID in the query.
// Returns a *NotSingularError when exactly one LikeClub ID is not found.
// Returns a *NotFoundError when no entities are found.
func (lcq *LikeClubQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{likeclub.Label}
	default:
		err = &NotSingularError{likeclub.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lcq *LikeClubQuery) OnlyIDX(ctx context.Context) int {
	id, err := lcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LikeClubs.
func (lcq *LikeClubQuery) All(ctx context.Context) ([]*LikeClub, error) {
	if err := lcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lcq *LikeClubQuery) AllX(ctx context.Context) []*LikeClub {
	nodes, err := lcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LikeClub IDs.
func (lcq *LikeClubQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := lcq.Select(likeclub.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lcq *LikeClubQuery) IDsX(ctx context.Context) []int {
	ids, err := lcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lcq *LikeClubQuery) Count(ctx context.Context) (int, error) {
	if err := lcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lcq *LikeClubQuery) CountX(ctx context.Context) int {
	count, err := lcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lcq *LikeClubQuery) Exist(ctx context.Context) (bool, error) {
	if err := lcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lcq *LikeClubQuery) ExistX(ctx context.Context) bool {
	exist, err := lcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LikeClubQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lcq *LikeClubQuery) Clone() *LikeClubQuery {
	if lcq == nil {
		return nil
	}
	return &LikeClubQuery{
		config:     lcq.config,
		limit:      lcq.limit,
		offset:     lcq.offset,
		order:      append([]OrderFunc{}, lcq.order...),
		predicates: append([]predicate.LikeClub{}, lcq.predicates...),
		withClub:   lcq.withClub.Clone(),
		// clone intermediate query.
		sql:  lcq.sql.Clone(),
		path: lcq.path,
	}
}

// WithClub tells the query-builder to eager-load the nodes that are connected to
// the "club" edge. The optional arguments are used to configure the query builder of the edge.
func (lcq *LikeClubQuery) WithClub(opts ...func(*ClubQuery)) *LikeClubQuery {
	query := &ClubQuery{config: lcq.config}
	for _, opt := range opts {
		opt(query)
	}
	lcq.withClub = query
	return lcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LikeClub.Query().
//		GroupBy(likeclub.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (lcq *LikeClubQuery) GroupBy(field string, fields ...string) *LikeClubGroupBy {
	group := &LikeClubGroupBy{config: lcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.LikeClub.Query().
//		Select(likeclub.FieldUsername).
//		Scan(ctx, &v)
//
func (lcq *LikeClubQuery) Select(field string, fields ...string) *LikeClubSelect {
	lcq.fields = append([]string{field}, fields...)
	return &LikeClubSelect{LikeClubQuery: lcq}
}

func (lcq *LikeClubQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lcq.fields {
		if !likeclub.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lcq.path != nil {
		prev, err := lcq.path(ctx)
		if err != nil {
			return err
		}
		lcq.sql = prev
	}
	return nil
}

func (lcq *LikeClubQuery) sqlAll(ctx context.Context) ([]*LikeClub, error) {
	var (
		nodes       = []*LikeClub{}
		withFKs     = lcq.withFKs
		_spec       = lcq.querySpec()
		loadedTypes = [1]bool{
			lcq.withClub != nil,
		}
	)
	if lcq.withClub != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, likeclub.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &LikeClub{config: lcq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, lcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := lcq.withClub; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*LikeClub)
		for i := range nodes {
			if nodes[i].club_likes == nil {
				continue
			}
			fk := *nodes[i].club_likes
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(club.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "club_likes" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Club = n
			}
		}
	}

	return nodes, nil
}

func (lcq *LikeClubQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lcq.querySpec()
	return sqlgraph.CountNodes(ctx, lcq.driver, _spec)
}

func (lcq *LikeClubQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (lcq *LikeClubQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   likeclub.Table,
			Columns: likeclub.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: likeclub.FieldID,
			},
		},
		From:   lcq.sql,
		Unique: true,
	}
	if unique := lcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, likeclub.FieldID)
		for i := range fields {
			if fields[i] != likeclub.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lcq *LikeClubQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lcq.driver.Dialect())
	t1 := builder.Table(likeclub.Table)
	selector := builder.Select(t1.Columns(likeclub.Columns...)...).From(t1)
	if lcq.sql != nil {
		selector = lcq.sql
		selector.Select(selector.Columns(likeclub.Columns...)...)
	}
	for _, p := range lcq.predicates {
		p(selector)
	}
	for _, p := range lcq.order {
		p(selector)
	}
	if offset := lcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LikeClubGroupBy is the group-by builder for LikeClub entities.
type LikeClubGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lcgb *LikeClubGroupBy) Aggregate(fns ...AggregateFunc) *LikeClubGroupBy {
	lcgb.fns = append(lcgb.fns, fns...)
	return lcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (lcgb *LikeClubGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lcgb.path(ctx)
	if err != nil {
		return err
	}
	lcgb.sql = query
	return lcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lcgb *LikeClubGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := lcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (lcgb *LikeClubGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(lcgb.fields) > 1 {
		return nil, errors.New("ent: LikeClubGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := lcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lcgb *LikeClubGroupBy) StringsX(ctx context.Context) []string {
	v, err := lcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lcgb *LikeClubGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{likeclub.Label}
	default:
		err = fmt.Errorf("ent: LikeClubGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lcgb *LikeClubGroupBy) StringX(ctx context.Context) string {
	v, err := lcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (lcgb *LikeClubGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(lcgb.fields) > 1 {
		return nil, errors.New("ent: LikeClubGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := lcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lcgb *LikeClubGroupBy) IntsX(ctx context.Context) []int {
	v, err := lcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lcgb *LikeClubGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{likeclub.Label}
	default:
		err = fmt.Errorf("ent: LikeClubGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lcgb *LikeClubGroupBy) IntX(ctx context.Context) int {
	v, err := lcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (lcgb *LikeClubGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(lcgb.fields) > 1 {
		return nil, errors.New("ent: LikeClubGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := lcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lcgb *LikeClubGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := lcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lcgb *LikeClubGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{likeclub.Label}
	default:
		err = fmt.Errorf("ent: LikeClubGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lcgb *LikeClubGroupBy) Float64X(ctx context.Context) float64 {
	v, err := lcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (lcgb *LikeClubGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(lcgb.fields) > 1 {
		return nil, errors.New("ent: LikeClubGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := lcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lcgb *LikeClubGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := lcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lcgb *LikeClubGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{likeclub.Label}
	default:
		err = fmt.Errorf("ent: LikeClubGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lcgb *LikeClubGroupBy) BoolX(ctx context.Context) bool {
	v, err := lcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lcgb *LikeClubGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lcgb.fields {
		if !likeclub.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lcgb *LikeClubGroupBy) sqlQuery() *sql.Selector {
	selector := lcgb.sql
	columns := make([]string, 0, len(lcgb.fields)+len(lcgb.fns))
	columns = append(columns, lcgb.fields...)
	for _, fn := range lcgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(lcgb.fields...)
}

// LikeClubSelect is the builder for selecting fields of LikeClub entities.
type LikeClubSelect struct {
	*LikeClubQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (lcs *LikeClubSelect) Scan(ctx context.Context, v interface{}) error {
	if err := lcs.prepareQuery(ctx); err != nil {
		return err
	}
	lcs.sql = lcs.LikeClubQuery.sqlQuery(ctx)
	return lcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lcs *LikeClubSelect) ScanX(ctx context.Context, v interface{}) {
	if err := lcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (lcs *LikeClubSelect) Strings(ctx context.Context) ([]string, error) {
	if len(lcs.fields) > 1 {
		return nil, errors.New("ent: LikeClubSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := lcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lcs *LikeClubSelect) StringsX(ctx context.Context) []string {
	v, err := lcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (lcs *LikeClubSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{likeclub.Label}
	default:
		err = fmt.Errorf("ent: LikeClubSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lcs *LikeClubSelect) StringX(ctx context.Context) string {
	v, err := lcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (lcs *LikeClubSelect) Ints(ctx context.Context) ([]int, error) {
	if len(lcs.fields) > 1 {
		return nil, errors.New("ent: LikeClubSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := lcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lcs *LikeClubSelect) IntsX(ctx context.Context) []int {
	v, err := lcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (lcs *LikeClubSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{likeclub.Label}
	default:
		err = fmt.Errorf("ent: LikeClubSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lcs *LikeClubSelect) IntX(ctx context.Context) int {
	v, err := lcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (lcs *LikeClubSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(lcs.fields) > 1 {
		return nil, errors.New("ent: LikeClubSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := lcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lcs *LikeClubSelect) Float64sX(ctx context.Context) []float64 {
	v, err := lcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (lcs *LikeClubSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{likeclub.Label}
	default:
		err = fmt.Errorf("ent: LikeClubSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lcs *LikeClubSelect) Float64X(ctx context.Context) float64 {
	v, err := lcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (lcs *LikeClubSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(lcs.fields) > 1 {
		return nil, errors.New("ent: LikeClubSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := lcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lcs *LikeClubSelect) BoolsX(ctx context.Context) []bool {
	v, err := lcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (lcs *LikeClubSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{likeclub.Label}
	default:
		err = fmt.Errorf("ent: LikeClubSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lcs *LikeClubSelect) BoolX(ctx context.Context) bool {
	v, err := lcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lcs *LikeClubSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := lcs.sqlQuery().Query()
	if err := lcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lcs *LikeClubSelect) sqlQuery() sql.Querier {
	selector := lcs.sql
	selector.Select(selector.Columns(lcs.fields...)...)
	return selector
}
