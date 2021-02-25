// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"smpp/ent/predicate"
	"smpp/ent/price"
	"smpp/ent/rate"
	"smpp/ent/rateprice"
	"smpp/ent/user"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// RatePriceQuery is the builder for querying RatePrice entities.
type RatePriceQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.RatePrice
	// eager-loading edges.
	withIDRate  *RateQuery
	withIDPrice *PriceQuery
	withUser    *UserQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RatePriceQuery builder.
func (rpq *RatePriceQuery) Where(ps ...predicate.RatePrice) *RatePriceQuery {
	rpq.predicates = append(rpq.predicates, ps...)
	return rpq
}

// Limit adds a limit step to the query.
func (rpq *RatePriceQuery) Limit(limit int) *RatePriceQuery {
	rpq.limit = &limit
	return rpq
}

// Offset adds an offset step to the query.
func (rpq *RatePriceQuery) Offset(offset int) *RatePriceQuery {
	rpq.offset = &offset
	return rpq
}

// Order adds an order step to the query.
func (rpq *RatePriceQuery) Order(o ...OrderFunc) *RatePriceQuery {
	rpq.order = append(rpq.order, o...)
	return rpq
}

// QueryIDRate chains the current query on the "id_rate" edge.
func (rpq *RatePriceQuery) QueryIDRate() *RateQuery {
	query := &RateQuery{config: rpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rateprice.Table, rateprice.FieldID, selector),
			sqlgraph.To(rate.Table, rate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rateprice.IDRateTable, rateprice.IDRateColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIDPrice chains the current query on the "id_price" edge.
func (rpq *RatePriceQuery) QueryIDPrice() *PriceQuery {
	query := &PriceQuery{config: rpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rateprice.Table, rateprice.FieldID, selector),
			sqlgraph.To(price.Table, price.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rateprice.IDPriceTable, rateprice.IDPriceColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (rpq *RatePriceQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: rpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rateprice.Table, rateprice.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, rateprice.UserTable, rateprice.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RatePrice entity from the query.
// Returns a *NotFoundError when no RatePrice was found.
func (rpq *RatePriceQuery) First(ctx context.Context) (*RatePrice, error) {
	nodes, err := rpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rateprice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rpq *RatePriceQuery) FirstX(ctx context.Context) *RatePrice {
	node, err := rpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RatePrice ID from the query.
// Returns a *NotFoundError when no RatePrice ID was found.
func (rpq *RatePriceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rateprice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rpq *RatePriceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RatePrice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one RatePrice entity is not found.
// Returns a *NotFoundError when no RatePrice entities are found.
func (rpq *RatePriceQuery) Only(ctx context.Context) (*RatePrice, error) {
	nodes, err := rpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rateprice.Label}
	default:
		return nil, &NotSingularError{rateprice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rpq *RatePriceQuery) OnlyX(ctx context.Context) *RatePrice {
	node, err := rpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RatePrice ID in the query.
// Returns a *NotSingularError when exactly one RatePrice ID is not found.
// Returns a *NotFoundError when no entities are found.
func (rpq *RatePriceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rateprice.Label}
	default:
		err = &NotSingularError{rateprice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rpq *RatePriceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RatePrices.
func (rpq *RatePriceQuery) All(ctx context.Context) ([]*RatePrice, error) {
	if err := rpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rpq *RatePriceQuery) AllX(ctx context.Context) []*RatePrice {
	nodes, err := rpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RatePrice IDs.
func (rpq *RatePriceQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := rpq.Select(rateprice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rpq *RatePriceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rpq *RatePriceQuery) Count(ctx context.Context) (int, error) {
	if err := rpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rpq *RatePriceQuery) CountX(ctx context.Context) int {
	count, err := rpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rpq *RatePriceQuery) Exist(ctx context.Context) (bool, error) {
	if err := rpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rpq *RatePriceQuery) ExistX(ctx context.Context) bool {
	exist, err := rpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RatePriceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rpq *RatePriceQuery) Clone() *RatePriceQuery {
	if rpq == nil {
		return nil
	}
	return &RatePriceQuery{
		config:      rpq.config,
		limit:       rpq.limit,
		offset:      rpq.offset,
		order:       append([]OrderFunc{}, rpq.order...),
		predicates:  append([]predicate.RatePrice{}, rpq.predicates...),
		withIDRate:  rpq.withIDRate.Clone(),
		withIDPrice: rpq.withIDPrice.Clone(),
		withUser:    rpq.withUser.Clone(),
		// clone intermediate query.
		sql:  rpq.sql.Clone(),
		path: rpq.path,
	}
}

// WithIDRate tells the query-builder to eager-load the nodes that are connected to
// the "id_rate" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *RatePriceQuery) WithIDRate(opts ...func(*RateQuery)) *RatePriceQuery {
	query := &RateQuery{config: rpq.config}
	for _, opt := range opts {
		opt(query)
	}
	rpq.withIDRate = query
	return rpq
}

// WithIDPrice tells the query-builder to eager-load the nodes that are connected to
// the "id_price" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *RatePriceQuery) WithIDPrice(opts ...func(*PriceQuery)) *RatePriceQuery {
	query := &PriceQuery{config: rpq.config}
	for _, opt := range opts {
		opt(query)
	}
	rpq.withIDPrice = query
	return rpq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *RatePriceQuery) WithUser(opts ...func(*UserQuery)) *RatePriceQuery {
	query := &UserQuery{config: rpq.config}
	for _, opt := range opts {
		opt(query)
	}
	rpq.withUser = query
	return rpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateAt time.Time `json:"create_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RatePrice.Query().
//		GroupBy(rateprice.FieldCreateAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rpq *RatePriceQuery) GroupBy(field string, fields ...string) *RatePriceGroupBy {
	group := &RatePriceGroupBy{config: rpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rpq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateAt time.Time `json:"create_at,omitempty"`
//	}
//
//	client.RatePrice.Query().
//		Select(rateprice.FieldCreateAt).
//		Scan(ctx, &v)
//
func (rpq *RatePriceQuery) Select(field string, fields ...string) *RatePriceSelect {
	rpq.fields = append([]string{field}, fields...)
	return &RatePriceSelect{RatePriceQuery: rpq}
}

func (rpq *RatePriceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rpq.fields {
		if !rateprice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rpq.path != nil {
		prev, err := rpq.path(ctx)
		if err != nil {
			return err
		}
		rpq.sql = prev
	}
	return nil
}

func (rpq *RatePriceQuery) sqlAll(ctx context.Context) ([]*RatePrice, error) {
	var (
		nodes       = []*RatePrice{}
		withFKs     = rpq.withFKs
		_spec       = rpq.querySpec()
		loadedTypes = [3]bool{
			rpq.withIDRate != nil,
			rpq.withIDPrice != nil,
			rpq.withUser != nil,
		}
	)
	if rpq.withIDRate != nil || rpq.withIDPrice != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, rateprice.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &RatePrice{config: rpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, rpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rpq.withIDRate; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*RatePrice)
		for i := range nodes {
			if fk := nodes[i].rate_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(rate.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "rate_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.IDRate = n
			}
		}
	}

	if query := rpq.withIDPrice; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*RatePrice)
		for i := range nodes {
			if fk := nodes[i].price_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(price.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "price_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.IDPrice = n
			}
		}
	}

	if query := rpq.withUser; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*RatePrice)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.User = []*User{}
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(rateprice.UserColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.rate_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "rate_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "rate_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.User = append(node.Edges.User, n)
		}
	}

	return nodes, nil
}

func (rpq *RatePriceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rpq.querySpec()
	return sqlgraph.CountNodes(ctx, rpq.driver, _spec)
}

func (rpq *RatePriceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rpq *RatePriceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rateprice.Table,
			Columns: rateprice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: rateprice.FieldID,
			},
		},
		From:   rpq.sql,
		Unique: true,
	}
	if fields := rpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rateprice.FieldID)
		for i := range fields {
			if fields[i] != rateprice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, rateprice.ValidColumn)
			}
		}
	}
	return _spec
}

func (rpq *RatePriceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rpq.driver.Dialect())
	t1 := builder.Table(rateprice.Table)
	selector := builder.Select(t1.Columns(rateprice.Columns...)...).From(t1)
	if rpq.sql != nil {
		selector = rpq.sql
		selector.Select(selector.Columns(rateprice.Columns...)...)
	}
	for _, p := range rpq.predicates {
		p(selector)
	}
	for _, p := range rpq.order {
		p(selector, rateprice.ValidColumn)
	}
	if offset := rpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RatePriceGroupBy is the group-by builder for RatePrice entities.
type RatePriceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rpgb *RatePriceGroupBy) Aggregate(fns ...AggregateFunc) *RatePriceGroupBy {
	rpgb.fns = append(rpgb.fns, fns...)
	return rpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rpgb *RatePriceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rpgb.path(ctx)
	if err != nil {
		return err
	}
	rpgb.sql = query
	return rpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rpgb *RatePriceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RatePriceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rpgb.fields) > 1 {
		return nil, errors.New("ent: RatePriceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rpgb *RatePriceGroupBy) StringsX(ctx context.Context) []string {
	v, err := rpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RatePriceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rateprice.Label}
	default:
		err = fmt.Errorf("ent: RatePriceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rpgb *RatePriceGroupBy) StringX(ctx context.Context) string {
	v, err := rpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RatePriceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rpgb.fields) > 1 {
		return nil, errors.New("ent: RatePriceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rpgb *RatePriceGroupBy) IntsX(ctx context.Context) []int {
	v, err := rpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RatePriceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rateprice.Label}
	default:
		err = fmt.Errorf("ent: RatePriceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rpgb *RatePriceGroupBy) IntX(ctx context.Context) int {
	v, err := rpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RatePriceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rpgb.fields) > 1 {
		return nil, errors.New("ent: RatePriceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rpgb *RatePriceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RatePriceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rateprice.Label}
	default:
		err = fmt.Errorf("ent: RatePriceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rpgb *RatePriceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RatePriceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rpgb.fields) > 1 {
		return nil, errors.New("ent: RatePriceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rpgb *RatePriceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rpgb *RatePriceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rateprice.Label}
	default:
		err = fmt.Errorf("ent: RatePriceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rpgb *RatePriceGroupBy) BoolX(ctx context.Context) bool {
	v, err := rpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rpgb *RatePriceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rpgb.fields {
		if !rateprice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rpgb *RatePriceGroupBy) sqlQuery() *sql.Selector {
	selector := rpgb.sql
	columns := make([]string, 0, len(rpgb.fields)+len(rpgb.fns))
	columns = append(columns, rpgb.fields...)
	for _, fn := range rpgb.fns {
		columns = append(columns, fn(selector, rateprice.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(rpgb.fields...)
}

// RatePriceSelect is the builder for selecting fields of RatePrice entities.
type RatePriceSelect struct {
	*RatePriceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rps *RatePriceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rps.prepareQuery(ctx); err != nil {
		return err
	}
	rps.sql = rps.RatePriceQuery.sqlQuery()
	return rps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rps *RatePriceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rps *RatePriceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rps.fields) > 1 {
		return nil, errors.New("ent: RatePriceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rps *RatePriceSelect) StringsX(ctx context.Context) []string {
	v, err := rps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rps *RatePriceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rateprice.Label}
	default:
		err = fmt.Errorf("ent: RatePriceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rps *RatePriceSelect) StringX(ctx context.Context) string {
	v, err := rps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rps *RatePriceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rps.fields) > 1 {
		return nil, errors.New("ent: RatePriceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rps *RatePriceSelect) IntsX(ctx context.Context) []int {
	v, err := rps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rps *RatePriceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rateprice.Label}
	default:
		err = fmt.Errorf("ent: RatePriceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rps *RatePriceSelect) IntX(ctx context.Context) int {
	v, err := rps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rps *RatePriceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rps.fields) > 1 {
		return nil, errors.New("ent: RatePriceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rps *RatePriceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rps *RatePriceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rateprice.Label}
	default:
		err = fmt.Errorf("ent: RatePriceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rps *RatePriceSelect) Float64X(ctx context.Context) float64 {
	v, err := rps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rps *RatePriceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rps.fields) > 1 {
		return nil, errors.New("ent: RatePriceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rps *RatePriceSelect) BoolsX(ctx context.Context) []bool {
	v, err := rps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rps *RatePriceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rateprice.Label}
	default:
		err = fmt.Errorf("ent: RatePriceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rps *RatePriceSelect) BoolX(ctx context.Context) bool {
	v, err := rps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rps *RatePriceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rps.sqlQuery().Query()
	if err := rps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rps *RatePriceSelect) sqlQuery() sql.Querier {
	selector := rps.sql
	selector.Select(selector.Columns(rps.fields...)...)
	return selector
}
