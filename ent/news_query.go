// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"intern_traning/ent/news"
	"intern_traning/ent/predicate"
	"intern_traning/ent/user"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// NewsQuery is the builder for querying News entities.
type NewsQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	predicates     []predicate.News
	withAuthorEdge *UserQuery
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*News) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NewsQuery builder.
func (nq *NewsQuery) Where(ps ...predicate.News) *NewsQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NewsQuery) Limit(limit int) *NewsQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NewsQuery) Offset(offset int) *NewsQuery {
	nq.offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NewsQuery) Unique(unique bool) *NewsQuery {
	nq.unique = &unique
	return nq
}

// Order adds an order step to the query.
func (nq *NewsQuery) Order(o ...OrderFunc) *NewsQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryAuthorEdge chains the current query on the "author_edge" edge.
func (nq *NewsQuery) QueryAuthorEdge() *UserQuery {
	query := &UserQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(news.Table, news.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, news.AuthorEdgeTable, news.AuthorEdgeColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first News entity from the query.
// Returns a *NotFoundError when no News was found.
func (nq *NewsQuery) First(ctx context.Context) (*News, error) {
	nodes, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{news.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NewsQuery) FirstX(ctx context.Context) *News {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first News ID from the query.
// Returns a *NotFoundError when no News ID was found.
func (nq *NewsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{news.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NewsQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single News entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one News entity is found.
// Returns a *NotFoundError when no News entities are found.
func (nq *NewsQuery) Only(ctx context.Context) (*News, error) {
	nodes, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{news.Label}
	default:
		return nil, &NotSingularError{news.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NewsQuery) OnlyX(ctx context.Context) *News {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only News ID in the query.
// Returns a *NotSingularError when more than one News ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NewsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{news.Label}
	default:
		err = &NotSingularError{news.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NewsQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NewsSlice.
func (nq *NewsQuery) All(ctx context.Context) ([]*News, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NewsQuery) AllX(ctx context.Context) []*News {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of News IDs.
func (nq *NewsQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := nq.Select(news.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NewsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NewsQuery) Count(ctx context.Context) (int, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NewsQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NewsQuery) Exist(ctx context.Context) (bool, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NewsQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NewsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NewsQuery) Clone() *NewsQuery {
	if nq == nil {
		return nil
	}
	return &NewsQuery{
		config:         nq.config,
		limit:          nq.limit,
		offset:         nq.offset,
		order:          append([]OrderFunc{}, nq.order...),
		predicates:     append([]predicate.News{}, nq.predicates...),
		withAuthorEdge: nq.withAuthorEdge.Clone(),
		// clone intermediate query.
		sql:    nq.sql.Clone(),
		path:   nq.path,
		unique: nq.unique,
	}
}

// WithAuthorEdge tells the query-builder to eager-load the nodes that are connected to
// the "author_edge" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NewsQuery) WithAuthorEdge(opts ...func(*UserQuery)) *NewsQuery {
	query := &UserQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withAuthorEdge = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.News.Query().
//		GroupBy(news.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NewsQuery) GroupBy(field string, fields ...string) *NewsGroupBy {
	grbuild := &NewsGroupBy{config: nq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nq.sqlQuery(ctx), nil
	}
	grbuild.label = news.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.News.Query().
//		Select(news.FieldCreatedAt).
//		Scan(ctx, &v)
func (nq *NewsQuery) Select(fields ...string) *NewsSelect {
	nq.fields = append(nq.fields, fields...)
	selbuild := &NewsSelect{NewsQuery: nq}
	selbuild.label = news.Label
	selbuild.flds, selbuild.scan = &nq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a NewsSelect configured with the given aggregations.
func (nq *NewsQuery) Aggregate(fns ...AggregateFunc) *NewsSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NewsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nq.fields {
		if !news.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NewsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*News, error) {
	var (
		nodes       = []*News{}
		_spec       = nq.querySpec()
		loadedTypes = [1]bool{
			nq.withAuthorEdge != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*News).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &News{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withAuthorEdge; query != nil {
		if err := nq.loadAuthorEdge(ctx, query, nodes, nil,
			func(n *News, e *User) { n.Edges.AuthorEdge = e }); err != nil {
			return nil, err
		}
	}
	for i := range nq.loadTotal {
		if err := nq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NewsQuery) loadAuthorEdge(ctx context.Context, query *UserQuery, nodes []*News, init func(*News), assign func(*News, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*News)
	for i := range nodes {
		fk := nodes[i].AuthorID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "author_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (nq *NewsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	_spec.Node.Columns = nq.fields
	if len(nq.fields) > 0 {
		_spec.Unique = nq.unique != nil && *nq.unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NewsQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (nq *NewsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   news.Table,
			Columns: news.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: news.FieldID,
			},
		},
		From:   nq.sql,
		Unique: true,
	}
	if unique := nq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, news.FieldID)
		for i := range fields {
			if fields[i] != news.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NewsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(news.Table)
	columns := nq.fields
	if len(columns) == 0 {
		columns = news.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.unique != nil && *nq.unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NewsGroupBy is the group-by builder for News entities.
type NewsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NewsGroupBy) Aggregate(fns ...AggregateFunc) *NewsGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the group-by query and scans the result into the given value.
func (ngb *NewsGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ngb.path(ctx)
	if err != nil {
		return err
	}
	ngb.sql = query
	return ngb.sqlScan(ctx, v)
}

func (ngb *NewsGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ngb.fields {
		if !news.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ngb *NewsGroupBy) sqlQuery() *sql.Selector {
	selector := ngb.sql.Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ngb.fields)+len(ngb.fns))
		for _, f := range ngb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ngb.fields...)...)
}

// NewsSelect is the builder for selecting fields of News entities.
type NewsSelect struct {
	*NewsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NewsSelect) Aggregate(fns ...AggregateFunc) *NewsSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NewsSelect) Scan(ctx context.Context, v any) error {
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	ns.sql = ns.NewsQuery.sqlQuery(ctx)
	return ns.sqlScan(ctx, v)
}

func (ns *NewsSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(ns.sql))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ns.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ns.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ns.sql.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
