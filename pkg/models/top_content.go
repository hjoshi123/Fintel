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

// TopContent is an object representing the database table.
type TopContent struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Ticker    string    `boil:"ticker" json:"ticker" toml:"ticker" yaml:"ticker"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	URL       string    `boil:"url" json:"url" toml:"url" yaml:"url"`
	SRC       string    `boil:"src" json:"src" toml:"src" yaml:"src"`
	Info      null.JSON `boil:"info" json:"info,omitempty" toml:"info" yaml:"info,omitempty"`

	R *topContentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L topContentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TopContentColumns = struct {
	ID        string
	Ticker    string
	CreatedAt string
	UpdatedAt string
	URL       string
	SRC       string
	Info      string
}{
	ID:        "id",
	Ticker:    "ticker",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	URL:       "url",
	SRC:       "src",
	Info:      "info",
}

var TopContentTableColumns = struct {
	ID        string
	Ticker    string
	CreatedAt string
	UpdatedAt string
	URL       string
	SRC       string
	Info      string
}{
	ID:        "top_content.id",
	Ticker:    "top_content.ticker",
	CreatedAt: "top_content.created_at",
	UpdatedAt: "top_content.updated_at",
	URL:       "top_content.url",
	SRC:       "top_content.src",
	Info:      "top_content.info",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_JSON struct{ field string }

func (w whereHelpernull_JSON) EQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_JSON) NEQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_JSON) LT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_JSON) LTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_JSON) GT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_JSON) GTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_JSON) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_JSON) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var TopContentWhere = struct {
	ID        whereHelperint
	Ticker    whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	URL       whereHelperstring
	SRC       whereHelperstring
	Info      whereHelpernull_JSON
}{
	ID:        whereHelperint{field: "\"top_content\".\"id\""},
	Ticker:    whereHelperstring{field: "\"top_content\".\"ticker\""},
	CreatedAt: whereHelpertime_Time{field: "\"top_content\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"top_content\".\"updated_at\""},
	URL:       whereHelperstring{field: "\"top_content\".\"url\""},
	SRC:       whereHelperstring{field: "\"top_content\".\"src\""},
	Info:      whereHelpernull_JSON{field: "\"top_content\".\"info\""},
}

// TopContentRels is where relationship names are stored.
var TopContentRels = struct {
}{}

// topContentR is where relationships are stored.
type topContentR struct {
}

// NewStruct creates a new relationship struct
func (*topContentR) NewStruct() *topContentR {
	return &topContentR{}
}

// topContentL is where Load methods for each relationship are stored.
type topContentL struct{}

var (
	topContentAllColumns            = []string{"id", "ticker", "created_at", "updated_at", "url", "src", "info"}
	topContentColumnsWithoutDefault = []string{"ticker", "url", "src"}
	topContentColumnsWithDefault    = []string{"id", "created_at", "updated_at", "info"}
	topContentPrimaryKeyColumns     = []string{"id"}
	topContentGeneratedColumns      = []string{}
)

type (
	// TopContentSlice is an alias for a slice of pointers to TopContent.
	// This should almost always be used instead of []TopContent.
	TopContentSlice []*TopContent
	// TopContentHook is the signature for custom TopContent hook methods
	TopContentHook func(context.Context, boil.ContextExecutor, *TopContent) error

	topContentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	topContentType                 = reflect.TypeOf(&TopContent{})
	topContentMapping              = queries.MakeStructMapping(topContentType)
	topContentPrimaryKeyMapping, _ = queries.BindMapping(topContentType, topContentMapping, topContentPrimaryKeyColumns)
	topContentInsertCacheMut       sync.RWMutex
	topContentInsertCache          = make(map[string]insertCache)
	topContentUpdateCacheMut       sync.RWMutex
	topContentUpdateCache          = make(map[string]updateCache)
	topContentUpsertCacheMut       sync.RWMutex
	topContentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var topContentAfterSelectMu sync.Mutex
var topContentAfterSelectHooks []TopContentHook

var topContentBeforeInsertMu sync.Mutex
var topContentBeforeInsertHooks []TopContentHook
var topContentAfterInsertMu sync.Mutex
var topContentAfterInsertHooks []TopContentHook

var topContentBeforeUpdateMu sync.Mutex
var topContentBeforeUpdateHooks []TopContentHook
var topContentAfterUpdateMu sync.Mutex
var topContentAfterUpdateHooks []TopContentHook

var topContentBeforeDeleteMu sync.Mutex
var topContentBeforeDeleteHooks []TopContentHook
var topContentAfterDeleteMu sync.Mutex
var topContentAfterDeleteHooks []TopContentHook

var topContentBeforeUpsertMu sync.Mutex
var topContentBeforeUpsertHooks []TopContentHook
var topContentAfterUpsertMu sync.Mutex
var topContentAfterUpsertHooks []TopContentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TopContent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topContentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TopContent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topContentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TopContent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topContentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TopContent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topContentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TopContent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topContentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TopContent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topContentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TopContent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topContentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TopContent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topContentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TopContent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range topContentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTopContentHook registers your hook function for all future operations.
func AddTopContentHook(hookPoint boil.HookPoint, topContentHook TopContentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		topContentAfterSelectMu.Lock()
		topContentAfterSelectHooks = append(topContentAfterSelectHooks, topContentHook)
		topContentAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		topContentBeforeInsertMu.Lock()
		topContentBeforeInsertHooks = append(topContentBeforeInsertHooks, topContentHook)
		topContentBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		topContentAfterInsertMu.Lock()
		topContentAfterInsertHooks = append(topContentAfterInsertHooks, topContentHook)
		topContentAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		topContentBeforeUpdateMu.Lock()
		topContentBeforeUpdateHooks = append(topContentBeforeUpdateHooks, topContentHook)
		topContentBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		topContentAfterUpdateMu.Lock()
		topContentAfterUpdateHooks = append(topContentAfterUpdateHooks, topContentHook)
		topContentAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		topContentBeforeDeleteMu.Lock()
		topContentBeforeDeleteHooks = append(topContentBeforeDeleteHooks, topContentHook)
		topContentBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		topContentAfterDeleteMu.Lock()
		topContentAfterDeleteHooks = append(topContentAfterDeleteHooks, topContentHook)
		topContentAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		topContentBeforeUpsertMu.Lock()
		topContentBeforeUpsertHooks = append(topContentBeforeUpsertHooks, topContentHook)
		topContentBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		topContentAfterUpsertMu.Lock()
		topContentAfterUpsertHooks = append(topContentAfterUpsertHooks, topContentHook)
		topContentAfterUpsertMu.Unlock()
	}
}

// OneG returns a single topContent record from the query using the global executor.
func (q topContentQuery) OneG(ctx context.Context) (*TopContent, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single topContent record from the query.
func (q topContentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TopContent, error) {
	o := &TopContent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for top_content")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TopContent records from the query using the global executor.
func (q topContentQuery) AllG(ctx context.Context) (TopContentSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TopContent records from the query.
func (q topContentQuery) All(ctx context.Context, exec boil.ContextExecutor) (TopContentSlice, error) {
	var o []*TopContent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TopContent slice")
	}

	if len(topContentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TopContent records in the query using the global executor
func (q topContentQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TopContent records in the query.
func (q topContentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count top_content rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q topContentQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q topContentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if top_content exists")
	}

	return count > 0, nil
}

// TopContents retrieves all the records using an executor.
func TopContents(mods ...qm.QueryMod) topContentQuery {
	mods = append(mods, qm.From("\"top_content\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"top_content\".*"})
	}

	return topContentQuery{q}
}

// FindTopContentG retrieves a single record by ID.
func FindTopContentG(ctx context.Context, iD int, selectCols ...string) (*TopContent, error) {
	return FindTopContent(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTopContent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTopContent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TopContent, error) {
	topContentObj := &TopContent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"top_content\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, topContentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from top_content")
	}

	if err = topContentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return topContentObj, err
	}

	return topContentObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TopContent) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TopContent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no top_content provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(topContentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	topContentInsertCacheMut.RLock()
	cache, cached := topContentInsertCache[key]
	topContentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			topContentAllColumns,
			topContentColumnsWithDefault,
			topContentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(topContentType, topContentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(topContentType, topContentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"top_content\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"top_content\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into top_content")
	}

	if !cached {
		topContentInsertCacheMut.Lock()
		topContentInsertCache[key] = cache
		topContentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TopContent record using the global executor.
// See Update for more documentation.
func (o *TopContent) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TopContent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TopContent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	topContentUpdateCacheMut.RLock()
	cache, cached := topContentUpdateCache[key]
	topContentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			topContentAllColumns,
			topContentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update top_content, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"top_content\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, topContentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(topContentType, topContentMapping, append(wl, topContentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update top_content row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for top_content")
	}

	if !cached {
		topContentUpdateCacheMut.Lock()
		topContentUpdateCache[key] = cache
		topContentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q topContentQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q topContentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for top_content")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for top_content")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TopContentSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TopContentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), topContentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"top_content\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, topContentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in topContent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all topContent")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TopContent) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TopContent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no top_content provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(topContentColumnsWithDefault, o)

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

	topContentUpsertCacheMut.RLock()
	cache, cached := topContentUpsertCache[key]
	topContentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			topContentAllColumns,
			topContentColumnsWithDefault,
			topContentColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			topContentAllColumns,
			topContentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert top_content, could not build update column list")
		}

		ret := strmangle.SetComplement(topContentAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(topContentPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert top_content, could not build conflict column list")
			}

			conflict = make([]string, len(topContentPrimaryKeyColumns))
			copy(conflict, topContentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"top_content\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(topContentType, topContentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(topContentType, topContentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert top_content")
	}

	if !cached {
		topContentUpsertCacheMut.Lock()
		topContentUpsertCache[key] = cache
		topContentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TopContent record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TopContent) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TopContent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TopContent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TopContent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), topContentPrimaryKeyMapping)
	sql := "DELETE FROM \"top_content\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from top_content")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for top_content")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q topContentQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q topContentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no topContentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from top_content")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for top_content")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TopContentSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TopContentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(topContentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), topContentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"top_content\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, topContentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from topContent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for top_content")
	}

	if len(topContentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TopContent) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no TopContent provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TopContent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTopContent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TopContentSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TopContentSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TopContentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TopContentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), topContentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"top_content\".* FROM \"top_content\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, topContentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TopContentSlice")
	}

	*o = slice

	return nil
}

// TopContentExistsG checks if the TopContent row exists.
func TopContentExistsG(ctx context.Context, iD int) (bool, error) {
	return TopContentExists(ctx, boil.GetContextDB(), iD)
}

// TopContentExists checks if the TopContent row exists.
func TopContentExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"top_content\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if top_content exists")
	}

	return exists, nil
}

// Exists checks if the TopContent row exists.
func (o *TopContent) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TopContentExists(ctx, exec, o.ID)
}
