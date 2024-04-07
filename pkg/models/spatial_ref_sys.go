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

// SpatialRefSy is an object representing the database table.
type SpatialRefSy struct {
	Srid      int         `boil:"srid" json:"srid" toml:"srid" yaml:"srid"`
	AuthName  null.String `boil:"auth_name" json:"auth_name,omitempty" toml:"auth_name" yaml:"auth_name,omitempty"`
	AuthSrid  null.Int    `boil:"auth_srid" json:"auth_srid,omitempty" toml:"auth_srid" yaml:"auth_srid,omitempty"`
	Srtext    null.String `boil:"srtext" json:"srtext,omitempty" toml:"srtext" yaml:"srtext,omitempty"`
	Proj4text null.String `boil:"proj4text" json:"proj4text,omitempty" toml:"proj4text" yaml:"proj4text,omitempty"`

	R *spatialRefSyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L spatialRefSyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SpatialRefSyColumns = struct {
	Srid      string
	AuthName  string
	AuthSrid  string
	Srtext    string
	Proj4text string
}{
	Srid:      "srid",
	AuthName:  "auth_name",
	AuthSrid:  "auth_srid",
	Srtext:    "srtext",
	Proj4text: "proj4text",
}

var SpatialRefSyTableColumns = struct {
	Srid      string
	AuthName  string
	AuthSrid  string
	Srtext    string
	Proj4text string
}{
	Srid:      "spatial_ref_sys.srid",
	AuthName:  "spatial_ref_sys.auth_name",
	AuthSrid:  "spatial_ref_sys.auth_srid",
	Srtext:    "spatial_ref_sys.srtext",
	Proj4text: "spatial_ref_sys.proj4text",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) ILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" ILIKE ?", x)
}
func (w whereHelpernull_String) NILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT ILIKE ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var SpatialRefSyWhere = struct {
	Srid      whereHelperint
	AuthName  whereHelpernull_String
	AuthSrid  whereHelpernull_Int
	Srtext    whereHelpernull_String
	Proj4text whereHelpernull_String
}{
	Srid:      whereHelperint{field: "\"spatial_ref_sys\".\"srid\""},
	AuthName:  whereHelpernull_String{field: "\"spatial_ref_sys\".\"auth_name\""},
	AuthSrid:  whereHelpernull_Int{field: "\"spatial_ref_sys\".\"auth_srid\""},
	Srtext:    whereHelpernull_String{field: "\"spatial_ref_sys\".\"srtext\""},
	Proj4text: whereHelpernull_String{field: "\"spatial_ref_sys\".\"proj4text\""},
}

// SpatialRefSyRels is where relationship names are stored.
var SpatialRefSyRels = struct {
}{}

// spatialRefSyR is where relationships are stored.
type spatialRefSyR struct {
}

// NewStruct creates a new relationship struct
func (*spatialRefSyR) NewStruct() *spatialRefSyR {
	return &spatialRefSyR{}
}

// spatialRefSyL is where Load methods for each relationship are stored.
type spatialRefSyL struct{}

var (
	spatialRefSyAllColumns            = []string{"srid", "auth_name", "auth_srid", "srtext", "proj4text"}
	spatialRefSyColumnsWithoutDefault = []string{"srid"}
	spatialRefSyColumnsWithDefault    = []string{"auth_name", "auth_srid", "srtext", "proj4text"}
	spatialRefSyPrimaryKeyColumns     = []string{"srid"}
	spatialRefSyGeneratedColumns      = []string{}
)

type (
	// SpatialRefSySlice is an alias for a slice of pointers to SpatialRefSy.
	// This should almost always be used instead of []SpatialRefSy.
	SpatialRefSySlice []*SpatialRefSy
	// SpatialRefSyHook is the signature for custom SpatialRefSy hook methods
	SpatialRefSyHook func(context.Context, boil.ContextExecutor, *SpatialRefSy) error

	spatialRefSyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	spatialRefSyType                 = reflect.TypeOf(&SpatialRefSy{})
	spatialRefSyMapping              = queries.MakeStructMapping(spatialRefSyType)
	spatialRefSyPrimaryKeyMapping, _ = queries.BindMapping(spatialRefSyType, spatialRefSyMapping, spatialRefSyPrimaryKeyColumns)
	spatialRefSyInsertCacheMut       sync.RWMutex
	spatialRefSyInsertCache          = make(map[string]insertCache)
	spatialRefSyUpdateCacheMut       sync.RWMutex
	spatialRefSyUpdateCache          = make(map[string]updateCache)
	spatialRefSyUpsertCacheMut       sync.RWMutex
	spatialRefSyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var spatialRefSyAfterSelectMu sync.Mutex
var spatialRefSyAfterSelectHooks []SpatialRefSyHook

var spatialRefSyBeforeInsertMu sync.Mutex
var spatialRefSyBeforeInsertHooks []SpatialRefSyHook
var spatialRefSyAfterInsertMu sync.Mutex
var spatialRefSyAfterInsertHooks []SpatialRefSyHook

var spatialRefSyBeforeUpdateMu sync.Mutex
var spatialRefSyBeforeUpdateHooks []SpatialRefSyHook
var spatialRefSyAfterUpdateMu sync.Mutex
var spatialRefSyAfterUpdateHooks []SpatialRefSyHook

var spatialRefSyBeforeDeleteMu sync.Mutex
var spatialRefSyBeforeDeleteHooks []SpatialRefSyHook
var spatialRefSyAfterDeleteMu sync.Mutex
var spatialRefSyAfterDeleteHooks []SpatialRefSyHook

var spatialRefSyBeforeUpsertMu sync.Mutex
var spatialRefSyBeforeUpsertHooks []SpatialRefSyHook
var spatialRefSyAfterUpsertMu sync.Mutex
var spatialRefSyAfterUpsertHooks []SpatialRefSyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SpatialRefSy) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spatialRefSyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SpatialRefSy) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spatialRefSyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SpatialRefSy) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spatialRefSyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SpatialRefSy) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spatialRefSyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SpatialRefSy) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spatialRefSyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SpatialRefSy) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spatialRefSyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SpatialRefSy) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spatialRefSyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SpatialRefSy) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spatialRefSyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SpatialRefSy) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spatialRefSyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSpatialRefSyHook registers your hook function for all future operations.
func AddSpatialRefSyHook(hookPoint boil.HookPoint, spatialRefSyHook SpatialRefSyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		spatialRefSyAfterSelectMu.Lock()
		spatialRefSyAfterSelectHooks = append(spatialRefSyAfterSelectHooks, spatialRefSyHook)
		spatialRefSyAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		spatialRefSyBeforeInsertMu.Lock()
		spatialRefSyBeforeInsertHooks = append(spatialRefSyBeforeInsertHooks, spatialRefSyHook)
		spatialRefSyBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		spatialRefSyAfterInsertMu.Lock()
		spatialRefSyAfterInsertHooks = append(spatialRefSyAfterInsertHooks, spatialRefSyHook)
		spatialRefSyAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		spatialRefSyBeforeUpdateMu.Lock()
		spatialRefSyBeforeUpdateHooks = append(spatialRefSyBeforeUpdateHooks, spatialRefSyHook)
		spatialRefSyBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		spatialRefSyAfterUpdateMu.Lock()
		spatialRefSyAfterUpdateHooks = append(spatialRefSyAfterUpdateHooks, spatialRefSyHook)
		spatialRefSyAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		spatialRefSyBeforeDeleteMu.Lock()
		spatialRefSyBeforeDeleteHooks = append(spatialRefSyBeforeDeleteHooks, spatialRefSyHook)
		spatialRefSyBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		spatialRefSyAfterDeleteMu.Lock()
		spatialRefSyAfterDeleteHooks = append(spatialRefSyAfterDeleteHooks, spatialRefSyHook)
		spatialRefSyAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		spatialRefSyBeforeUpsertMu.Lock()
		spatialRefSyBeforeUpsertHooks = append(spatialRefSyBeforeUpsertHooks, spatialRefSyHook)
		spatialRefSyBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		spatialRefSyAfterUpsertMu.Lock()
		spatialRefSyAfterUpsertHooks = append(spatialRefSyAfterUpsertHooks, spatialRefSyHook)
		spatialRefSyAfterUpsertMu.Unlock()
	}
}

// OneG returns a single spatialRefSy record from the query using the global executor.
func (q spatialRefSyQuery) OneG(ctx context.Context) (*SpatialRefSy, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single spatialRefSy record from the query.
func (q spatialRefSyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SpatialRefSy, error) {
	o := &SpatialRefSy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for spatial_ref_sys")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all SpatialRefSy records from the query using the global executor.
func (q spatialRefSyQuery) AllG(ctx context.Context) (SpatialRefSySlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all SpatialRefSy records from the query.
func (q spatialRefSyQuery) All(ctx context.Context, exec boil.ContextExecutor) (SpatialRefSySlice, error) {
	var o []*SpatialRefSy

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SpatialRefSy slice")
	}

	if len(spatialRefSyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all SpatialRefSy records in the query using the global executor
func (q spatialRefSyQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all SpatialRefSy records in the query.
func (q spatialRefSyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count spatial_ref_sys rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q spatialRefSyQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q spatialRefSyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if spatial_ref_sys exists")
	}

	return count > 0, nil
}

// SpatialRefSys retrieves all the records using an executor.
func SpatialRefSys(mods ...qm.QueryMod) spatialRefSyQuery {
	mods = append(mods, qm.From("\"spatial_ref_sys\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"spatial_ref_sys\".*"})
	}

	return spatialRefSyQuery{q}
}

// FindSpatialRefSyG retrieves a single record by ID.
func FindSpatialRefSyG(ctx context.Context, srid int, selectCols ...string) (*SpatialRefSy, error) {
	return FindSpatialRefSy(ctx, boil.GetContextDB(), srid, selectCols...)
}

// FindSpatialRefSy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSpatialRefSy(ctx context.Context, exec boil.ContextExecutor, srid int, selectCols ...string) (*SpatialRefSy, error) {
	spatialRefSyObj := &SpatialRefSy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"spatial_ref_sys\" where \"srid\"=$1", sel,
	)

	q := queries.Raw(query, srid)

	err := q.Bind(ctx, exec, spatialRefSyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from spatial_ref_sys")
	}

	if err = spatialRefSyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return spatialRefSyObj, err
	}

	return spatialRefSyObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *SpatialRefSy) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SpatialRefSy) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no spatial_ref_sys provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(spatialRefSyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	spatialRefSyInsertCacheMut.RLock()
	cache, cached := spatialRefSyInsertCache[key]
	spatialRefSyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			spatialRefSyAllColumns,
			spatialRefSyColumnsWithDefault,
			spatialRefSyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(spatialRefSyType, spatialRefSyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(spatialRefSyType, spatialRefSyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"spatial_ref_sys\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"spatial_ref_sys\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into spatial_ref_sys")
	}

	if !cached {
		spatialRefSyInsertCacheMut.Lock()
		spatialRefSyInsertCache[key] = cache
		spatialRefSyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single SpatialRefSy record using the global executor.
// See Update for more documentation.
func (o *SpatialRefSy) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the SpatialRefSy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SpatialRefSy) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	spatialRefSyUpdateCacheMut.RLock()
	cache, cached := spatialRefSyUpdateCache[key]
	spatialRefSyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			spatialRefSyAllColumns,
			spatialRefSyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update spatial_ref_sys, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"spatial_ref_sys\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, spatialRefSyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(spatialRefSyType, spatialRefSyMapping, append(wl, spatialRefSyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update spatial_ref_sys row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for spatial_ref_sys")
	}

	if !cached {
		spatialRefSyUpdateCacheMut.Lock()
		spatialRefSyUpdateCache[key] = cache
		spatialRefSyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q spatialRefSyQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q spatialRefSyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for spatial_ref_sys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for spatial_ref_sys")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SpatialRefSySlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SpatialRefSySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spatialRefSyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"spatial_ref_sys\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, spatialRefSyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in spatialRefSy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all spatialRefSy")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *SpatialRefSy) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SpatialRefSy) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no spatial_ref_sys provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(spatialRefSyColumnsWithDefault, o)

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

	spatialRefSyUpsertCacheMut.RLock()
	cache, cached := spatialRefSyUpsertCache[key]
	spatialRefSyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			spatialRefSyAllColumns,
			spatialRefSyColumnsWithDefault,
			spatialRefSyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			spatialRefSyAllColumns,
			spatialRefSyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert spatial_ref_sys, could not build update column list")
		}

		ret := strmangle.SetComplement(spatialRefSyAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(spatialRefSyPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert spatial_ref_sys, could not build conflict column list")
			}

			conflict = make([]string, len(spatialRefSyPrimaryKeyColumns))
			copy(conflict, spatialRefSyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"spatial_ref_sys\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(spatialRefSyType, spatialRefSyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(spatialRefSyType, spatialRefSyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert spatial_ref_sys")
	}

	if !cached {
		spatialRefSyUpsertCacheMut.Lock()
		spatialRefSyUpsertCache[key] = cache
		spatialRefSyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single SpatialRefSy record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *SpatialRefSy) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single SpatialRefSy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SpatialRefSy) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SpatialRefSy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), spatialRefSyPrimaryKeyMapping)
	sql := "DELETE FROM \"spatial_ref_sys\" WHERE \"srid\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from spatial_ref_sys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for spatial_ref_sys")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q spatialRefSyQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q spatialRefSyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no spatialRefSyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spatial_ref_sys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spatial_ref_sys")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o SpatialRefSySlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SpatialRefSySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(spatialRefSyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spatialRefSyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"spatial_ref_sys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spatialRefSyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spatialRefSy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spatial_ref_sys")
	}

	if len(spatialRefSyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *SpatialRefSy) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no SpatialRefSy provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SpatialRefSy) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSpatialRefSy(ctx, exec, o.Srid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SpatialRefSySlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty SpatialRefSySlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SpatialRefSySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SpatialRefSySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spatialRefSyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"spatial_ref_sys\".* FROM \"spatial_ref_sys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spatialRefSyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SpatialRefSySlice")
	}

	*o = slice

	return nil
}

// SpatialRefSyExistsG checks if the SpatialRefSy row exists.
func SpatialRefSyExistsG(ctx context.Context, srid int) (bool, error) {
	return SpatialRefSyExists(ctx, boil.GetContextDB(), srid)
}

// SpatialRefSyExists checks if the SpatialRefSy row exists.
func SpatialRefSyExists(ctx context.Context, exec boil.ContextExecutor, srid int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"spatial_ref_sys\" where \"srid\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, srid)
	}
	row := exec.QueryRowContext(ctx, sql, srid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if spatial_ref_sys exists")
	}

	return exists, nil
}

// Exists checks if the SpatialRefSy row exists.
func (o *SpatialRefSy) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SpatialRefSyExists(ctx, exec, o.Srid)
}