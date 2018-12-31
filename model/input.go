// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Input is an object representing the database table.
type Input struct {
	ID                  uint64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	TransactionID       uint64       `boil:"transaction_id" json:"transaction_id" toml:"transaction_id" yaml:"transaction_id"`
	TransactionHash     string       `boil:"transaction_hash" json:"transaction_hash" toml:"transaction_hash" yaml:"transaction_hash"`
	InputAddressID      null.Uint64  `boil:"input_address_id" json:"input_address_id,omitempty" toml:"input_address_id" yaml:"input_address_id,omitempty"`
	IsCoinbase          bool         `boil:"is_coinbase" json:"is_coinbase" toml:"is_coinbase" yaml:"is_coinbase"`
	Coinbase            null.String  `boil:"coinbase" json:"coinbase,omitempty" toml:"coinbase" yaml:"coinbase,omitempty"`
	PrevoutHash         null.String  `boil:"prevout_hash" json:"prevout_hash,omitempty" toml:"prevout_hash" yaml:"prevout_hash,omitempty"`
	PrevoutN            null.Uint    `boil:"prevout_n" json:"prevout_n,omitempty" toml:"prevout_n" yaml:"prevout_n,omitempty"`
	PrevoutSpendUpdated bool         `boil:"prevout_spend_updated" json:"prevout_spend_updated" toml:"prevout_spend_updated" yaml:"prevout_spend_updated"`
	Sequence            uint         `boil:"sequence" json:"sequence" toml:"sequence" yaml:"sequence"`
	Value               null.Float64 `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	ScriptSigAsm        null.String  `boil:"script_sig_asm" json:"script_sig_asm,omitempty" toml:"script_sig_asm" yaml:"script_sig_asm,omitempty"`
	ScriptSigHex        null.String  `boil:"script_sig_hex" json:"script_sig_hex,omitempty" toml:"script_sig_hex" yaml:"script_sig_hex,omitempty"`
	Created             time.Time    `boil:"created" json:"created" toml:"created" yaml:"created"`
	Modified            time.Time    `boil:"modified" json:"modified" toml:"modified" yaml:"modified"`

	R *inputR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L inputL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InputColumns = struct {
	ID                  string
	TransactionID       string
	TransactionHash     string
	InputAddressID      string
	IsCoinbase          string
	Coinbase            string
	PrevoutHash         string
	PrevoutN            string
	PrevoutSpendUpdated string
	Sequence            string
	Value               string
	ScriptSigAsm        string
	ScriptSigHex        string
	Created             string
	Modified            string
}{
	ID:                  "id",
	TransactionID:       "transaction_id",
	TransactionHash:     "transaction_hash",
	InputAddressID:      "input_address_id",
	IsCoinbase:          "is_coinbase",
	Coinbase:            "coinbase",
	PrevoutHash:         "prevout_hash",
	PrevoutN:            "prevout_n",
	PrevoutSpendUpdated: "prevout_spend_updated",
	Sequence:            "sequence",
	Value:               "value",
	ScriptSigAsm:        "script_sig_asm",
	ScriptSigHex:        "script_sig_hex",
	Created:             "created",
	Modified:            "modified",
}

// InputRels is where relationship names are stored.
var InputRels = struct {
	Transaction string
}{
	Transaction: "Transaction",
}

// inputR is where relationships are stored.
type inputR struct {
	Transaction *Transaction
}

// NewStruct creates a new relationship struct
func (*inputR) NewStruct() *inputR {
	return &inputR{}
}

// inputL is where Load methods for each relationship are stored.
type inputL struct{}

var (
	inputColumns               = []string{"id", "transaction_id", "transaction_hash", "input_address_id", "is_coinbase", "coinbase", "prevout_hash", "prevout_n", "prevout_spend_updated", "sequence", "value", "script_sig_asm", "script_sig_hex", "created", "modified"}
	inputColumnsWithoutDefault = []string{"transaction_id", "transaction_hash", "input_address_id", "coinbase", "prevout_hash", "prevout_n", "sequence", "value", "script_sig_asm", "script_sig_hex"}
	inputColumnsWithDefault    = []string{"id", "is_coinbase", "prevout_spend_updated", "created", "modified"}
	inputPrimaryKeyColumns     = []string{"id"}
)

type (
	// InputSlice is an alias for a slice of pointers to Input.
	// This should generally be used opposed to []Input.
	InputSlice []*Input

	inputQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	inputType                 = reflect.TypeOf(&Input{})
	inputMapping              = queries.MakeStructMapping(inputType)
	inputPrimaryKeyMapping, _ = queries.BindMapping(inputType, inputMapping, inputPrimaryKeyColumns)
	inputInsertCacheMut       sync.RWMutex
	inputInsertCache          = make(map[string]insertCache)
	inputUpdateCacheMut       sync.RWMutex
	inputUpdateCache          = make(map[string]updateCache)
	inputUpsertCacheMut       sync.RWMutex
	inputUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

// OneG returns a single input record from the query using the global executor.
func (q inputQuery) OneG() (*Input, error) {
	return q.One(boil.GetDB())
}

// OneGP returns a single input record from the query using the global executor, and panics on error.
func (q inputQuery) OneGP() *Input {
	o, err := q.One(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// OneP returns a single input record from the query, and panics on error.
func (q inputQuery) OneP(exec boil.Executor) *Input {
	o, err := q.One(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single input record from the query.
func (q inputQuery) One(exec boil.Executor) (*Input, error) {
	o := &Input{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for input")
	}

	return o, nil
}

// AllG returns all Input records from the query using the global executor.
func (q inputQuery) AllG() (InputSlice, error) {
	return q.All(boil.GetDB())
}

// AllGP returns all Input records from the query using the global executor, and panics on error.
func (q inputQuery) AllGP() InputSlice {
	o, err := q.All(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// AllP returns all Input records from the query, and panics on error.
func (q inputQuery) AllP(exec boil.Executor) InputSlice {
	o, err := q.All(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Input records from the query.
func (q inputQuery) All(exec boil.Executor) (InputSlice, error) {
	var o []*Input

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Input slice")
	}

	return o, nil
}

// CountG returns the count of all Input records in the query, and panics on error.
func (q inputQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// CountGP returns the count of all Input records in the query using the global executor, and panics on error.
func (q inputQuery) CountGP() int64 {
	c, err := q.Count(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// CountP returns the count of all Input records in the query, and panics on error.
func (q inputQuery) CountP(exec boil.Executor) int64 {
	c, err := q.Count(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Input records in the query.
func (q inputQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count input rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q inputQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// ExistsGP checks if the row exists in the table using the global executor, and panics on error.
func (q inputQuery) ExistsGP() bool {
	e, err := q.Exists(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ExistsP checks if the row exists in the table, and panics on error.
func (q inputQuery) ExistsP(exec boil.Executor) bool {
	e, err := q.Exists(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q inputQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if input exists")
	}

	return count > 0, nil
}

// Transaction pointed to by the foreign key.
func (o *Input) Transaction(mods ...qm.QueryMod) transactionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TransactionID),
	}

	queryMods = append(queryMods, mods...)

	query := Transactions(queryMods...)
	queries.SetFrom(query.Query, "`transaction`")

	return query
}

// LoadTransaction allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (inputL) LoadTransaction(e boil.Executor, singular bool, maybeInput interface{}, mods queries.Applicator) error {
	var slice []*Input
	var object *Input

	if singular {
		object = maybeInput.(*Input)
	} else {
		slice = *maybeInput.(*[]*Input)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &inputR{}
		}
		args = append(args, object.TransactionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &inputR{}
			}

			for _, a := range args {
				if a == obj.TransactionID {
					continue Outer
				}
			}

			args = append(args, obj.TransactionID)

		}
	}

	query := NewQuery(qm.From(`transaction`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Transaction")
	}

	var resultSlice []*Transaction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Transaction")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for transaction")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for transaction")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Transaction = foreign
		if foreign.R == nil {
			foreign.R = &transactionR{}
		}
		foreign.R.Inputs = append(foreign.R.Inputs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TransactionID == foreign.ID {
				local.R.Transaction = foreign
				if foreign.R == nil {
					foreign.R = &transactionR{}
				}
				foreign.R.Inputs = append(foreign.R.Inputs, local)
				break
			}
		}
	}

	return nil
}

// SetTransactionG of the input to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.Inputs.
// Uses the global database handle.
func (o *Input) SetTransactionG(insert bool, related *Transaction) error {
	return o.SetTransaction(boil.GetDB(), insert, related)
}

// SetTransactionP of the input to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.Inputs.
// Panics on error.
func (o *Input) SetTransactionP(exec boil.Executor, insert bool, related *Transaction) {
	if err := o.SetTransaction(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetTransactionGP of the input to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.Inputs.
// Uses the global database handle and panics on error.
func (o *Input) SetTransactionGP(insert bool, related *Transaction) {
	if err := o.SetTransaction(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetTransaction of the input to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.Inputs.
func (o *Input) SetTransaction(exec boil.Executor, insert bool, related *Transaction) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `input` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"transaction_id"}),
		strmangle.WhereClause("`", "`", 0, inputPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TransactionID = related.ID
	if o.R == nil {
		o.R = &inputR{
			Transaction: related,
		}
	} else {
		o.R.Transaction = related
	}

	if related.R == nil {
		related.R = &transactionR{
			Inputs: InputSlice{o},
		}
	} else {
		related.R.Inputs = append(related.R.Inputs, o)
	}

	return nil
}

// Inputs retrieves all the records using an executor.
func Inputs(mods ...qm.QueryMod) inputQuery {
	mods = append(mods, qm.From("`input`"))
	return inputQuery{NewQuery(mods...)}
}

// FindInputG retrieves a single record by ID.
func FindInputG(iD uint64, selectCols ...string) (*Input, error) {
	return FindInput(boil.GetDB(), iD, selectCols...)
}

// FindInputP retrieves a single record by ID with an executor, and panics on error.
func FindInputP(exec boil.Executor, iD uint64, selectCols ...string) *Input {
	retobj, err := FindInput(exec, iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindInputGP retrieves a single record by ID, and panics on error.
func FindInputGP(iD uint64, selectCols ...string) *Input {
	retobj, err := FindInput(boil.GetDB(), iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindInput retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInput(exec boil.Executor, iD uint64, selectCols ...string) (*Input, error) {
	inputObj := &Input{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `input` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, inputObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from input")
	}

	return inputObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Input) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Input) InsertP(exec boil.Executor, columns boil.Columns) {
	if err := o.Insert(exec, columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Input) InsertGP(columns boil.Columns) {
	if err := o.Insert(boil.GetDB(), columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Input) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no input provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(inputColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	inputInsertCacheMut.RLock()
	cache, cached := inputInsertCache[key]
	inputInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			inputColumns,
			inputColumnsWithDefault,
			inputColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(inputType, inputMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(inputType, inputMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `input` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `input` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `input` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, inputPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into input")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == inputMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for input")
	}

CacheNoHooks:
	if !cached {
		inputInsertCacheMut.Lock()
		inputInsertCache[key] = cache
		inputInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Input record using the global executor.
// See Update for more documentation.
func (o *Input) UpdateG(columns boil.Columns) error {
	return o.Update(boil.GetDB(), columns)
}

// UpdateP uses an executor to update the Input, and panics on error.
// See Update for more documentation.
func (o *Input) UpdateP(exec boil.Executor, columns boil.Columns) {
	err := o.Update(exec, columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateGP a single Input record using the global executor. Panics on error.
// See Update for more documentation.
func (o *Input) UpdateGP(columns boil.Columns) {
	err := o.Update(boil.GetDB(), columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Input.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Input) Update(exec boil.Executor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	inputUpdateCacheMut.RLock()
	cache, cached := inputUpdateCache[key]
	inputUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			inputColumns,
			inputPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return errors.New("model: unable to update input, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `input` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, inputPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(inputType, inputMapping, append(wl, inputPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update input row")
	}

	if !cached {
		inputUpdateCacheMut.Lock()
		inputUpdateCache[key] = cache
		inputUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q inputQuery) UpdateAllP(exec boil.Executor, cols M) {
	err := q.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllG updates all rows with the specified column values.
func (q inputQuery) UpdateAllG(cols M) error {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q inputQuery) UpdateAll(exec boil.Executor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all for input")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o InputSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o InputSlice) UpdateAllGP(cols M) {
	err := o.UpdateAll(boil.GetDB(), cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o InputSlice) UpdateAllP(exec boil.Executor, cols M) {
	err := o.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InputSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), inputPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `input` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, inputPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all in input slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Input) UpsertG(updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateColumns, insertColumns)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Input) UpsertGP(updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(boil.GetDB(), updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Input) UpsertP(exec boil.Executor, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(exec, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

var mySQLInputUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Input) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no input provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(inputColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLInputUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	inputUpsertCacheMut.RLock()
	cache, cached := inputUpsertCache[key]
	inputUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			inputColumns,
			inputColumnsWithDefault,
			inputColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			inputColumns,
			inputPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert input, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "input", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `input` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(inputType, inputMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(inputType, inputMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for input")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == inputMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(inputType, inputMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for input")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for input")
	}

CacheNoHooks:
	if !cached {
		inputUpsertCacheMut.Lock()
		inputUpsertCache[key] = cache
		inputUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Input record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Input) DeleteG() error {
	return o.Delete(boil.GetDB())
}

// DeleteP deletes a single Input record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Input) DeleteP(exec boil.Executor) {
	err := o.Delete(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteGP deletes a single Input record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Input) DeleteGP() {
	err := o.Delete(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Input record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Input) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no Input provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), inputPrimaryKeyMapping)
	sql := "DELETE FROM `input` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete from input")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q inputQuery) DeleteAllP(exec boil.Executor) {
	err := q.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q inputQuery) DeleteAll(exec boil.Executor) error {
	if q.Query == nil {
		return errors.New("model: no inputQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from input")
	}

	return nil
}

// DeleteAllG deletes all rows in the slice.
func (o InputSlice) DeleteAllG() error {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o InputSlice) DeleteAllP(exec boil.Executor) {
	err := o.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o InputSlice) DeleteAllGP() {
	err := o.DeleteAll(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InputSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no Input slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), inputPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `input` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, inputPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from input slice")
	}

	return nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Input) ReloadG() error {
	if o == nil {
		return errors.New("model: no Input provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Input) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Input) ReloadGP() {
	if err := o.Reload(boil.GetDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Input) Reload(exec boil.Executor) error {
	ret, err := FindInput(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InputSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("model: empty InputSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *InputSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *InputSlice) ReloadAllGP() {
	if err := o.ReloadAll(boil.GetDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InputSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InputSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), inputPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `input`.* FROM `input` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, inputPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in InputSlice")
	}

	*o = slice

	return nil
}

// InputExistsG checks if the Input row exists.
func InputExistsG(iD uint64) (bool, error) {
	return InputExists(boil.GetDB(), iD)
}

// InputExistsP checks if the Input row exists. Panics on error.
func InputExistsP(exec boil.Executor, iD uint64) bool {
	e, err := InputExists(exec, iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// InputExistsGP checks if the Input row exists. Panics on error.
func InputExistsGP(iD uint64) bool {
	e, err := InputExists(boil.GetDB(), iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// InputExists checks if the Input row exists.
func InputExists(exec boil.Executor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `input` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if input exists")
	}

	return exists, nil
}
