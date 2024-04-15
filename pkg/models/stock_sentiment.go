// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// StockSentiment is an object representing the database table.
type StockSentiment struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Ticker    string    `boil:"ticker" json:"ticker" toml:"ticker" yaml:"ticker"`
	DailyIci  float64   `boil:"daily_ici" json:"daily_ici" toml:"daily_ici" yaml:"daily_ici"`
	Chatter   int       `boil:"chatter" json:"chatter" toml:"chatter" yaml:"chatter"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *stockSentimentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockSentimentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StockSentimentColumns = struct {
	ID        string
	Ticker    string
	DailyIci  string
	Chatter   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Ticker:    "ticker",
	DailyIci:  "daily_ici",
	Chatter:   "chatter",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var StockSentimentTableColumns = struct {
	ID        string
	Ticker    string
	DailyIci  string
	Chatter   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "stock_sentiment.id",
	Ticker:    "stock_sentiment.ticker",
	DailyIci:  "stock_sentiment.daily_ici",
	Chatter:   "stock_sentiment.chatter",
	CreatedAt: "stock_sentiment.created_at",
	UpdatedAt: "stock_sentiment.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var StockSentimentWhere = struct {
	ID        whereHelperint
	Ticker    whereHelperstring
	DailyIci  whereHelperfloat64
	Chatter   whereHelperint
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "\"stock_sentiment\".\"id\""},
	Ticker:    whereHelperstring{field: "\"stock_sentiment\".\"ticker\""},
	DailyIci:  whereHelperfloat64{field: "\"stock_sentiment\".\"daily_ici\""},
	Chatter:   whereHelperint{field: "\"stock_sentiment\".\"chatter\""},
	CreatedAt: whereHelpernull_Time{field: "\"stock_sentiment\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"stock_sentiment\".\"updated_at\""},
}

// StockSentimentRels is where relationship names are stored.
var StockSentimentRels = struct {
}{}

// stockSentimentR is where relationships are stored.
type stockSentimentR struct {
}

// NewStruct creates a new relationship struct
func (*stockSentimentR) NewStruct() *stockSentimentR {
	return &stockSentimentR{}
}

// stockSentimentL is where Load methods for each relationship are stored.
type stockSentimentL struct{}

var (
	stockSentimentAllColumns            = []string{"id", "ticker", "daily_ici", "chatter", "created_at", "updated_at"}
	stockSentimentColumnsWithoutDefault = []string{"ticker", "daily_ici", "chatter"}
	stockSentimentColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	stockSentimentPrimaryKeyColumns     = []string{"id"}
	stockSentimentGeneratedColumns      = []string{}
)

type (
	// StockSentimentSlice is an alias for a slice of pointers to StockSentiment.
	// This should almost always be used instead of []StockSentiment.
	StockSentimentSlice []*StockSentiment
	// StockSentimentHook is the signature for custom StockSentiment hook methods
	StockSentimentHook func(context.Context, boil.ContextExecutor, *StockSentiment) error

	stockSentimentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockSentimentType                 = reflect.TypeOf(&StockSentiment{})
	stockSentimentMapping              = queries.MakeStructMapping(stockSentimentType)
	stockSentimentPrimaryKeyMapping, _ = queries.BindMapping(stockSentimentType, stockSentimentMapping, stockSentimentPrimaryKeyColumns)
	stockSentimentInsertCacheMut       sync.RWMutex
	stockSentimentInsertCache          = make(map[string]insertCache)
	stockSentimentUpdateCacheMut       sync.RWMutex
	stockSentimentUpdateCache          = make(map[string]updateCache)
	stockSentimentUpsertCacheMut       sync.RWMutex
	stockSentimentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var stockSentimentAfterSelectMu sync.Mutex
var stockSentimentAfterSelectHooks []StockSentimentHook

var stockSentimentBeforeInsertMu sync.Mutex
var stockSentimentBeforeInsertHooks []StockSentimentHook
var stockSentimentAfterInsertMu sync.Mutex
var stockSentimentAfterInsertHooks []StockSentimentHook

var stockSentimentBeforeUpdateMu sync.Mutex
var stockSentimentBeforeUpdateHooks []StockSentimentHook
var stockSentimentAfterUpdateMu sync.Mutex
var stockSentimentAfterUpdateHooks []StockSentimentHook

var stockSentimentBeforeDeleteMu sync.Mutex
var stockSentimentBeforeDeleteHooks []StockSentimentHook
var stockSentimentAfterDeleteMu sync.Mutex
var stockSentimentAfterDeleteHooks []StockSentimentHook

var stockSentimentBeforeUpsertMu sync.Mutex
var stockSentimentBeforeUpsertHooks []StockSentimentHook
var stockSentimentAfterUpsertMu sync.Mutex
var stockSentimentAfterUpsertHooks []StockSentimentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockSentiment) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockSentimentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockSentiment) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockSentimentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockSentiment) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockSentimentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockSentiment) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockSentimentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockSentiment) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockSentimentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockSentiment) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockSentimentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockSentiment) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockSentimentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockSentiment) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockSentimentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockSentiment) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockSentimentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockSentimentHook registers your hook function for all future operations.
func AddStockSentimentHook(hookPoint boil.HookPoint, stockSentimentHook StockSentimentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		stockSentimentAfterSelectMu.Lock()
		stockSentimentAfterSelectHooks = append(stockSentimentAfterSelectHooks, stockSentimentHook)
		stockSentimentAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		stockSentimentBeforeInsertMu.Lock()
		stockSentimentBeforeInsertHooks = append(stockSentimentBeforeInsertHooks, stockSentimentHook)
		stockSentimentBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		stockSentimentAfterInsertMu.Lock()
		stockSentimentAfterInsertHooks = append(stockSentimentAfterInsertHooks, stockSentimentHook)
		stockSentimentAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		stockSentimentBeforeUpdateMu.Lock()
		stockSentimentBeforeUpdateHooks = append(stockSentimentBeforeUpdateHooks, stockSentimentHook)
		stockSentimentBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		stockSentimentAfterUpdateMu.Lock()
		stockSentimentAfterUpdateHooks = append(stockSentimentAfterUpdateHooks, stockSentimentHook)
		stockSentimentAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		stockSentimentBeforeDeleteMu.Lock()
		stockSentimentBeforeDeleteHooks = append(stockSentimentBeforeDeleteHooks, stockSentimentHook)
		stockSentimentBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		stockSentimentAfterDeleteMu.Lock()
		stockSentimentAfterDeleteHooks = append(stockSentimentAfterDeleteHooks, stockSentimentHook)
		stockSentimentAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		stockSentimentBeforeUpsertMu.Lock()
		stockSentimentBeforeUpsertHooks = append(stockSentimentBeforeUpsertHooks, stockSentimentHook)
		stockSentimentBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		stockSentimentAfterUpsertMu.Lock()
		stockSentimentAfterUpsertHooks = append(stockSentimentAfterUpsertHooks, stockSentimentHook)
		stockSentimentAfterUpsertMu.Unlock()
	}
}

// OneG returns a single stockSentiment record from the query using the global executor.
func (q stockSentimentQuery) OneG(ctx context.Context) (*StockSentiment, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single stockSentiment record from the query.
func (q stockSentimentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*StockSentiment, error) {
	o := &StockSentiment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock_sentiment")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all StockSentiment records from the query using the global executor.
func (q stockSentimentQuery) AllG(ctx context.Context) (StockSentimentSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all StockSentiment records from the query.
func (q stockSentimentQuery) All(ctx context.Context, exec boil.ContextExecutor) (StockSentimentSlice, error) {
	var o []*StockSentiment

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockSentiment slice")
	}

	if len(stockSentimentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all StockSentiment records in the query using the global executor
func (q stockSentimentQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all StockSentiment records in the query.
func (q stockSentimentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock_sentiment rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q stockSentimentQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q stockSentimentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock_sentiment exists")
	}

	return count > 0, nil
}

// StockSentiments retrieves all the records using an executor.
func StockSentiments(mods ...qm.QueryMod) stockSentimentQuery {
	mods = append(mods, qm.From("\"stock_sentiment\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"stock_sentiment\".*"})
	}

	return stockSentimentQuery{q}
}

// FindStockSentimentG retrieves a single record by ID.
func FindStockSentimentG(ctx context.Context, iD int, selectCols ...string) (*StockSentiment, error) {
	return FindStockSentiment(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindStockSentiment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockSentiment(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*StockSentiment, error) {
	stockSentimentObj := &StockSentiment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_sentiment\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, stockSentimentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock_sentiment")
	}

	if err = stockSentimentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return stockSentimentObj, err
	}

	return stockSentimentObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockSentiment) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *StockSentiment) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stock_sentiment provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockSentimentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stockSentimentInsertCacheMut.RLock()
	cache, cached := stockSentimentInsertCache[key]
	stockSentimentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stockSentimentAllColumns,
			stockSentimentColumnsWithDefault,
			stockSentimentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stockSentimentType, stockSentimentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockSentimentType, stockSentimentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"stock_sentiment\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"stock_sentiment\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into stock_sentiment")
	}

	if !cached {
		stockSentimentInsertCacheMut.Lock()
		stockSentimentInsertCache[key] = cache
		stockSentimentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single StockSentiment record using the global executor.
// See Update for more documentation.
func (o *StockSentiment) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the StockSentiment.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *StockSentiment) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	stockSentimentUpdateCacheMut.RLock()
	cache, cached := stockSentimentUpdateCache[key]
	stockSentimentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stockSentimentAllColumns,
			stockSentimentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stock_sentiment, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_sentiment\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockSentimentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockSentimentType, stockSentimentMapping, append(wl, stockSentimentPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update stock_sentiment row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stock_sentiment")
	}

	if !cached {
		stockSentimentUpdateCacheMut.Lock()
		stockSentimentUpdateCache[key] = cache
		stockSentimentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q stockSentimentQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q stockSentimentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stock_sentiment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stock_sentiment")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockSentimentSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockSentimentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockSentimentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"stock_sentiment\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, stockSentimentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stockSentiment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stockSentiment")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockSentiment) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *StockSentiment) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no stock_sentiment provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockSentimentColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	stockSentimentUpsertCacheMut.RLock()
	cache, cached := stockSentimentUpsertCache[key]
	stockSentimentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			stockSentimentAllColumns,
			stockSentimentColumnsWithDefault,
			stockSentimentColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			stockSentimentAllColumns,
			stockSentimentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert stock_sentiment, could not build update column list")
		}

		ret := strmangle.SetComplement(stockSentimentAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(stockSentimentPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert stock_sentiment, could not build conflict column list")
			}

			conflict = make([]string, len(stockSentimentPrimaryKeyColumns))
			copy(conflict, stockSentimentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"stock_sentiment\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(stockSentimentType, stockSentimentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockSentimentType, stockSentimentMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert stock_sentiment")
	}

	if !cached {
		stockSentimentUpsertCacheMut.Lock()
		stockSentimentUpsertCache[key] = cache
		stockSentimentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single StockSentiment record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockSentiment) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single StockSentiment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockSentiment) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no StockSentiment provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockSentimentPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_sentiment\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stock_sentiment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stock_sentiment")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q stockSentimentQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q stockSentimentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stockSentimentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stock_sentiment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stock_sentiment")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o StockSentimentSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockSentimentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(stockSentimentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockSentimentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"stock_sentiment\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stockSentimentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stockSentiment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stock_sentiment")
	}

	if len(stockSentimentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockSentiment) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no StockSentiment provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockSentiment) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStockSentiment(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockSentimentSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty StockSentimentSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockSentimentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StockSentimentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockSentimentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"stock_sentiment\".* FROM \"stock_sentiment\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stockSentimentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockSentimentSlice")
	}

	*o = slice

	return nil
}

// StockSentimentExistsG checks if the StockSentiment row exists.
func StockSentimentExistsG(ctx context.Context, iD int) (bool, error) {
	return StockSentimentExists(ctx, boil.GetContextDB(), iD)
}

// StockSentimentExists checks if the StockSentiment row exists.
func StockSentimentExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"stock_sentiment\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock_sentiment exists")
	}

	return exists, nil
}

// Exists checks if the StockSentiment row exists.
func (o *StockSentiment) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return StockSentimentExists(ctx, exec, o.ID)
}