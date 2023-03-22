// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Friendship is an object representing the database table.
type Friendship struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	FriendID  string    `boil:"friend_id" json:"friend_id" toml:"friend_id" yaml:"friend_id"`
	Status    int       `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *friendshipR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L friendshipL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FriendshipColumns = struct {
	ID        string
	UserID    string
	FriendID  string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	FriendID:  "friend_id",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var FriendshipTableColumns = struct {
	ID        string
	UserID    string
	FriendID  string
	Status    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "friendships.id",
	UserID:    "friendships.user_id",
	FriendID:  "friendships.friend_id",
	Status:    "friendships.status",
	CreatedAt: "friendships.created_at",
	UpdatedAt: "friendships.updated_at",
}

// Generated where

var FriendshipWhere = struct {
	ID        whereHelperstring
	UserID    whereHelperstring
	FriendID  whereHelperstring
	Status    whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperstring{field: "\"friendships\".\"id\""},
	UserID:    whereHelperstring{field: "\"friendships\".\"user_id\""},
	FriendID:  whereHelperstring{field: "\"friendships\".\"friend_id\""},
	Status:    whereHelperint{field: "\"friendships\".\"status\""},
	CreatedAt: whereHelpertime_Time{field: "\"friendships\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"friendships\".\"updated_at\""},
}

// FriendshipRels is where relationship names are stored.
var FriendshipRels = struct {
}{}

// friendshipR is where relationships are stored.
type friendshipR struct {
}

// NewStruct creates a new relationship struct
func (*friendshipR) NewStruct() *friendshipR {
	return &friendshipR{}
}

// friendshipL is where Load methods for each relationship are stored.
type friendshipL struct{}

var (
	friendshipAllColumns            = []string{"id", "user_id", "friend_id", "status", "created_at", "updated_at"}
	friendshipColumnsWithoutDefault = []string{"id", "user_id", "friend_id", "created_at", "updated_at"}
	friendshipColumnsWithDefault    = []string{"status"}
	friendshipPrimaryKeyColumns     = []string{"id"}
	friendshipGeneratedColumns      = []string{}
)

type (
	// FriendshipSlice is an alias for a slice of pointers to Friendship.
	// This should almost always be used instead of []Friendship.
	FriendshipSlice []*Friendship
	// FriendshipHook is the signature for custom Friendship hook methods
	FriendshipHook func(context.Context, boil.ContextExecutor, *Friendship) error

	friendshipQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	friendshipType                 = reflect.TypeOf(&Friendship{})
	friendshipMapping              = queries.MakeStructMapping(friendshipType)
	friendshipPrimaryKeyMapping, _ = queries.BindMapping(friendshipType, friendshipMapping, friendshipPrimaryKeyColumns)
	friendshipInsertCacheMut       sync.RWMutex
	friendshipInsertCache          = make(map[string]insertCache)
	friendshipUpdateCacheMut       sync.RWMutex
	friendshipUpdateCache          = make(map[string]updateCache)
	friendshipUpsertCacheMut       sync.RWMutex
	friendshipUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var friendshipAfterSelectHooks []FriendshipHook

var friendshipBeforeInsertHooks []FriendshipHook
var friendshipAfterInsertHooks []FriendshipHook

var friendshipBeforeUpdateHooks []FriendshipHook
var friendshipAfterUpdateHooks []FriendshipHook

var friendshipBeforeDeleteHooks []FriendshipHook
var friendshipAfterDeleteHooks []FriendshipHook

var friendshipBeforeUpsertHooks []FriendshipHook
var friendshipAfterUpsertHooks []FriendshipHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Friendship) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Friendship) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Friendship) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Friendship) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Friendship) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Friendship) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Friendship) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Friendship) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Friendship) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFriendshipHook registers your hook function for all future operations.
func AddFriendshipHook(hookPoint boil.HookPoint, friendshipHook FriendshipHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		friendshipAfterSelectHooks = append(friendshipAfterSelectHooks, friendshipHook)
	case boil.BeforeInsertHook:
		friendshipBeforeInsertHooks = append(friendshipBeforeInsertHooks, friendshipHook)
	case boil.AfterInsertHook:
		friendshipAfterInsertHooks = append(friendshipAfterInsertHooks, friendshipHook)
	case boil.BeforeUpdateHook:
		friendshipBeforeUpdateHooks = append(friendshipBeforeUpdateHooks, friendshipHook)
	case boil.AfterUpdateHook:
		friendshipAfterUpdateHooks = append(friendshipAfterUpdateHooks, friendshipHook)
	case boil.BeforeDeleteHook:
		friendshipBeforeDeleteHooks = append(friendshipBeforeDeleteHooks, friendshipHook)
	case boil.AfterDeleteHook:
		friendshipAfterDeleteHooks = append(friendshipAfterDeleteHooks, friendshipHook)
	case boil.BeforeUpsertHook:
		friendshipBeforeUpsertHooks = append(friendshipBeforeUpsertHooks, friendshipHook)
	case boil.AfterUpsertHook:
		friendshipAfterUpsertHooks = append(friendshipAfterUpsertHooks, friendshipHook)
	}
}

// OneG returns a single friendship record from the query using the global executor.
func (q friendshipQuery) OneG(ctx context.Context) (*Friendship, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single friendship record from the query.
func (q friendshipQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Friendship, error) {
	o := &Friendship{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for friendships")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Friendship records from the query using the global executor.
func (q friendshipQuery) AllG(ctx context.Context) (FriendshipSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Friendship records from the query.
func (q friendshipQuery) All(ctx context.Context, exec boil.ContextExecutor) (FriendshipSlice, error) {
	var o []*Friendship

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Friendship slice")
	}

	if len(friendshipAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Friendship records in the query using the global executor
func (q friendshipQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Friendship records in the query.
func (q friendshipQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count friendships rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q friendshipQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q friendshipQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if friendships exists")
	}

	return count > 0, nil
}

// Friendships retrieves all the records using an executor.
func Friendships(mods ...qm.QueryMod) friendshipQuery {
	mods = append(mods, qm.From("\"friendships\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"friendships\".*"})
	}

	return friendshipQuery{q}
}

// FindFriendshipG retrieves a single record by ID.
func FindFriendshipG(ctx context.Context, iD string, selectCols ...string) (*Friendship, error) {
	return FindFriendship(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindFriendship retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFriendship(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Friendship, error) {
	friendshipObj := &Friendship{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"friendships\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, friendshipObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from friendships")
	}

	if err = friendshipObj.doAfterSelectHooks(ctx, exec); err != nil {
		return friendshipObj, err
	}

	return friendshipObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Friendship) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Friendship) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no friendships provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(friendshipColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	friendshipInsertCacheMut.RLock()
	cache, cached := friendshipInsertCache[key]
	friendshipInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			friendshipAllColumns,
			friendshipColumnsWithDefault,
			friendshipColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(friendshipType, friendshipMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(friendshipType, friendshipMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"friendships\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"friendships\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "model: unable to insert into friendships")
	}

	if !cached {
		friendshipInsertCacheMut.Lock()
		friendshipInsertCache[key] = cache
		friendshipInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Friendship record using the global executor.
// See Update for more documentation.
func (o *Friendship) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Friendship.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Friendship) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	friendshipUpdateCacheMut.RLock()
	cache, cached := friendshipUpdateCache[key]
	friendshipUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			friendshipAllColumns,
			friendshipPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update friendships, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"friendships\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, friendshipPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(friendshipType, friendshipMapping, append(wl, friendshipPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update friendships row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for friendships")
	}

	if !cached {
		friendshipUpdateCacheMut.Lock()
		friendshipUpdateCache[key] = cache
		friendshipUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q friendshipQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q friendshipQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for friendships")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for friendships")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FriendshipSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FriendshipSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), friendshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"friendships\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, friendshipPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in friendship slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all friendship")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Friendship) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Friendship) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no friendships provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(friendshipColumnsWithDefault, o)

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

	friendshipUpsertCacheMut.RLock()
	cache, cached := friendshipUpsertCache[key]
	friendshipUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			friendshipAllColumns,
			friendshipColumnsWithDefault,
			friendshipColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			friendshipAllColumns,
			friendshipPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("model: unable to upsert friendships, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(friendshipPrimaryKeyColumns))
			copy(conflict, friendshipPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"friendships\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(friendshipType, friendshipMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(friendshipType, friendshipMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert friendships")
	}

	if !cached {
		friendshipUpsertCacheMut.Lock()
		friendshipUpsertCache[key] = cache
		friendshipUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Friendship record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Friendship) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Friendship record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Friendship) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no Friendship provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), friendshipPrimaryKeyMapping)
	sql := "DELETE FROM \"friendships\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from friendships")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for friendships")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q friendshipQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q friendshipQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no friendshipQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from friendships")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for friendships")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o FriendshipSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FriendshipSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(friendshipBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), friendshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"friendships\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, friendshipPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from friendship slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for friendships")
	}

	if len(friendshipAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Friendship) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("model: no Friendship provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Friendship) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFriendship(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FriendshipSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("model: empty FriendshipSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FriendshipSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FriendshipSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), friendshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"friendships\".* FROM \"friendships\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, friendshipPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in FriendshipSlice")
	}

	*o = slice

	return nil
}

// FriendshipExistsG checks if the Friendship row exists.
func FriendshipExistsG(ctx context.Context, iD string) (bool, error) {
	return FriendshipExists(ctx, boil.GetContextDB(), iD)
}

// FriendshipExists checks if the Friendship row exists.
func FriendshipExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"friendships\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if friendships exists")
	}

	return exists, nil
}

// Exists checks if the Friendship row exists.
func (o *Friendship) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FriendshipExists(ctx, exec, o.ID)
}
