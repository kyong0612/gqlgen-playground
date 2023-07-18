// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Pullrequest is an object representing the database table.
type Pullrequest struct {
	ID          string `boil:"id" json:"id" toml:"id" yaml:"id"`
	BaseRefName string `boil:"base_ref_name" json:"base_ref_name" toml:"base_ref_name" yaml:"base_ref_name"`
	Closed      int    `boil:"closed" json:"closed" toml:"closed" yaml:"closed"`
	HeadRefName string `boil:"head_ref_name" json:"head_ref_name" toml:"head_ref_name" yaml:"head_ref_name"`
	URL         string `boil:"url" json:"url" toml:"url" yaml:"url"`
	Number      int    `boil:"number" json:"number" toml:"number" yaml:"number"`
	Repository  string `boil:"repository" json:"repository" toml:"repository" yaml:"repository"`

	R *pullrequestR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pullrequestL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PullrequestColumns = struct {
	ID          string
	BaseRefName string
	Closed      string
	HeadRefName string
	URL         string
	Number      string
	Repository  string
}{
	ID:          "id",
	BaseRefName: "base_ref_name",
	Closed:      "closed",
	HeadRefName: "head_ref_name",
	URL:         "url",
	Number:      "number",
	Repository:  "repository",
}

var PullrequestTableColumns = struct {
	ID          string
	BaseRefName string
	Closed      string
	HeadRefName string
	URL         string
	Number      string
	Repository  string
}{
	ID:          "pullrequests.id",
	BaseRefName: "pullrequests.base_ref_name",
	Closed:      "pullrequests.closed",
	HeadRefName: "pullrequests.head_ref_name",
	URL:         "pullrequests.url",
	Number:      "pullrequests.number",
	Repository:  "pullrequests.repository",
}

// Generated where

var PullrequestWhere = struct {
	ID          whereHelperstring
	BaseRefName whereHelperstring
	Closed      whereHelperint
	HeadRefName whereHelperstring
	URL         whereHelperstring
	Number      whereHelperint
	Repository  whereHelperstring
}{
	ID:          whereHelperstring{field: "\"pullrequests\".\"id\""},
	BaseRefName: whereHelperstring{field: "\"pullrequests\".\"base_ref_name\""},
	Closed:      whereHelperint{field: "\"pullrequests\".\"closed\""},
	HeadRefName: whereHelperstring{field: "\"pullrequests\".\"head_ref_name\""},
	URL:         whereHelperstring{field: "\"pullrequests\".\"url\""},
	Number:      whereHelperint{field: "\"pullrequests\".\"number\""},
	Repository:  whereHelperstring{field: "\"pullrequests\".\"repository\""},
}

// PullrequestRels is where relationship names are stored.
var PullrequestRels = struct {
	PullrequestRepository string
	Projectcards          string
}{
	PullrequestRepository: "PullrequestRepository",
	Projectcards:          "Projectcards",
}

// pullrequestR is where relationships are stored.
type pullrequestR struct {
	PullrequestRepository *Repository      `boil:"PullrequestRepository" json:"PullrequestRepository" toml:"PullrequestRepository" yaml:"PullrequestRepository"`
	Projectcards          ProjectcardSlice `boil:"Projectcards" json:"Projectcards" toml:"Projectcards" yaml:"Projectcards"`
}

// NewStruct creates a new relationship struct
func (*pullrequestR) NewStruct() *pullrequestR {
	return &pullrequestR{}
}

func (r *pullrequestR) GetPullrequestRepository() *Repository {
	if r == nil {
		return nil
	}
	return r.PullrequestRepository
}

func (r *pullrequestR) GetProjectcards() ProjectcardSlice {
	if r == nil {
		return nil
	}
	return r.Projectcards
}

// pullrequestL is where Load methods for each relationship are stored.
type pullrequestL struct{}

var (
	pullrequestAllColumns            = []string{"id", "base_ref_name", "closed", "head_ref_name", "url", "number", "repository"}
	pullrequestColumnsWithoutDefault = []string{"id", "base_ref_name", "head_ref_name", "url", "number", "repository"}
	pullrequestColumnsWithDefault    = []string{"closed"}
	pullrequestPrimaryKeyColumns     = []string{"id"}
	pullrequestGeneratedColumns      = []string{}
)

type (
	// PullrequestSlice is an alias for a slice of pointers to Pullrequest.
	// This should almost always be used instead of []Pullrequest.
	PullrequestSlice []*Pullrequest
	// PullrequestHook is the signature for custom Pullrequest hook methods
	PullrequestHook func(context.Context, boil.ContextExecutor, *Pullrequest) error

	pullrequestQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pullrequestType                 = reflect.TypeOf(&Pullrequest{})
	pullrequestMapping              = queries.MakeStructMapping(pullrequestType)
	pullrequestPrimaryKeyMapping, _ = queries.BindMapping(pullrequestType, pullrequestMapping, pullrequestPrimaryKeyColumns)
	pullrequestInsertCacheMut       sync.RWMutex
	pullrequestInsertCache          = make(map[string]insertCache)
	pullrequestUpdateCacheMut       sync.RWMutex
	pullrequestUpdateCache          = make(map[string]updateCache)
	pullrequestUpsertCacheMut       sync.RWMutex
	pullrequestUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var pullrequestAfterSelectHooks []PullrequestHook

var pullrequestBeforeInsertHooks []PullrequestHook
var pullrequestAfterInsertHooks []PullrequestHook

var pullrequestBeforeUpdateHooks []PullrequestHook
var pullrequestAfterUpdateHooks []PullrequestHook

var pullrequestBeforeDeleteHooks []PullrequestHook
var pullrequestAfterDeleteHooks []PullrequestHook

var pullrequestBeforeUpsertHooks []PullrequestHook
var pullrequestAfterUpsertHooks []PullrequestHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Pullrequest) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pullrequestAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Pullrequest) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pullrequestBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Pullrequest) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pullrequestAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Pullrequest) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pullrequestBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Pullrequest) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pullrequestAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Pullrequest) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pullrequestBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Pullrequest) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pullrequestAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Pullrequest) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pullrequestBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Pullrequest) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pullrequestAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPullrequestHook registers your hook function for all future operations.
func AddPullrequestHook(hookPoint boil.HookPoint, pullrequestHook PullrequestHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		pullrequestAfterSelectHooks = append(pullrequestAfterSelectHooks, pullrequestHook)
	case boil.BeforeInsertHook:
		pullrequestBeforeInsertHooks = append(pullrequestBeforeInsertHooks, pullrequestHook)
	case boil.AfterInsertHook:
		pullrequestAfterInsertHooks = append(pullrequestAfterInsertHooks, pullrequestHook)
	case boil.BeforeUpdateHook:
		pullrequestBeforeUpdateHooks = append(pullrequestBeforeUpdateHooks, pullrequestHook)
	case boil.AfterUpdateHook:
		pullrequestAfterUpdateHooks = append(pullrequestAfterUpdateHooks, pullrequestHook)
	case boil.BeforeDeleteHook:
		pullrequestBeforeDeleteHooks = append(pullrequestBeforeDeleteHooks, pullrequestHook)
	case boil.AfterDeleteHook:
		pullrequestAfterDeleteHooks = append(pullrequestAfterDeleteHooks, pullrequestHook)
	case boil.BeforeUpsertHook:
		pullrequestBeforeUpsertHooks = append(pullrequestBeforeUpsertHooks, pullrequestHook)
	case boil.AfterUpsertHook:
		pullrequestAfterUpsertHooks = append(pullrequestAfterUpsertHooks, pullrequestHook)
	}
}

// One returns a single pullrequest record from the query.
func (q pullrequestQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Pullrequest, error) {
	o := &Pullrequest{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pullrequests")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Pullrequest records from the query.
func (q pullrequestQuery) All(ctx context.Context, exec boil.ContextExecutor) (PullrequestSlice, error) {
	var o []*Pullrequest

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Pullrequest slice")
	}

	if len(pullrequestAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Pullrequest records in the query.
func (q pullrequestQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pullrequests rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pullrequestQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pullrequests exists")
	}

	return count > 0, nil
}

// PullrequestRepository pointed to by the foreign key.
func (o *Pullrequest) PullrequestRepository(mods ...qm.QueryMod) repositoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.Repository),
	}

	queryMods = append(queryMods, mods...)

	return Repositories(queryMods...)
}

// Projectcards retrieves all the projectcard's Projectcards with an executor.
func (o *Pullrequest) Projectcards(mods ...qm.QueryMod) projectcardQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"projectcards\".\"pullrequest\"=?", o.ID),
	)

	return Projectcards(queryMods...)
}

// LoadPullrequestRepository allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pullrequestL) LoadPullrequestRepository(ctx context.Context, e boil.ContextExecutor, singular bool, maybePullrequest interface{}, mods queries.Applicator) error {
	var slice []*Pullrequest
	var object *Pullrequest

	if singular {
		var ok bool
		object, ok = maybePullrequest.(*Pullrequest)
		if !ok {
			object = new(Pullrequest)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePullrequest)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePullrequest))
			}
		}
	} else {
		s, ok := maybePullrequest.(*[]*Pullrequest)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePullrequest)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePullrequest))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pullrequestR{}
		}
		args = append(args, object.Repository)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pullrequestR{}
			}

			for _, a := range args {
				if a == obj.Repository {
					continue Outer
				}
			}

			args = append(args, obj.Repository)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`repositories`),
		qm.WhereIn(`repositories.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Repository")
	}

	var resultSlice []*Repository
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Repository")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for repositories")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for repositories")
	}

	if len(repositoryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.PullrequestRepository = foreign
		if foreign.R == nil {
			foreign.R = &repositoryR{}
		}
		foreign.R.Pullrequests = append(foreign.R.Pullrequests, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Repository == foreign.ID {
				local.R.PullrequestRepository = foreign
				if foreign.R == nil {
					foreign.R = &repositoryR{}
				}
				foreign.R.Pullrequests = append(foreign.R.Pullrequests, local)
				break
			}
		}
	}

	return nil
}

// LoadProjectcards allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pullrequestL) LoadProjectcards(ctx context.Context, e boil.ContextExecutor, singular bool, maybePullrequest interface{}, mods queries.Applicator) error {
	var slice []*Pullrequest
	var object *Pullrequest

	if singular {
		var ok bool
		object, ok = maybePullrequest.(*Pullrequest)
		if !ok {
			object = new(Pullrequest)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePullrequest)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePullrequest))
			}
		}
	} else {
		s, ok := maybePullrequest.(*[]*Pullrequest)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePullrequest)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePullrequest))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pullrequestR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pullrequestR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`projectcards`),
		qm.WhereIn(`projectcards.pullrequest in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load projectcards")
	}

	var resultSlice []*Projectcard
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice projectcards")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on projectcards")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for projectcards")
	}

	if len(projectcardAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Projectcards = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &projectcardR{}
			}
			foreign.R.ProjectcardPullrequest = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.Pullrequest) {
				local.R.Projectcards = append(local.R.Projectcards, foreign)
				if foreign.R == nil {
					foreign.R = &projectcardR{}
				}
				foreign.R.ProjectcardPullrequest = local
				break
			}
		}
	}

	return nil
}

// SetPullrequestRepository of the pullrequest to the related item.
// Sets o.R.PullrequestRepository to related.
// Adds o to related.R.Pullrequests.
func (o *Pullrequest) SetPullrequestRepository(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Repository) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"pullrequests\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"repository"}),
		strmangle.WhereClause("\"", "\"", 2, pullrequestPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Repository = related.ID
	if o.R == nil {
		o.R = &pullrequestR{
			PullrequestRepository: related,
		}
	} else {
		o.R.PullrequestRepository = related
	}

	if related.R == nil {
		related.R = &repositoryR{
			Pullrequests: PullrequestSlice{o},
		}
	} else {
		related.R.Pullrequests = append(related.R.Pullrequests, o)
	}

	return nil
}

// AddProjectcards adds the given related objects to the existing relationships
// of the pullrequest, optionally inserting them as new records.
// Appends related to o.R.Projectcards.
// Sets related.R.ProjectcardPullrequest appropriately.
func (o *Pullrequest) AddProjectcards(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Projectcard) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.Pullrequest, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"projectcards\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"pullrequest"}),
				strmangle.WhereClause("\"", "\"", 2, projectcardPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.Pullrequest, o.ID)
		}
	}

	if o.R == nil {
		o.R = &pullrequestR{
			Projectcards: related,
		}
	} else {
		o.R.Projectcards = append(o.R.Projectcards, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &projectcardR{
				ProjectcardPullrequest: o,
			}
		} else {
			rel.R.ProjectcardPullrequest = o
		}
	}
	return nil
}

// SetProjectcards removes all previously related items of the
// pullrequest replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.ProjectcardPullrequest's Projectcards accordingly.
// Replaces o.R.Projectcards with related.
// Sets related.R.ProjectcardPullrequest's Projectcards accordingly.
func (o *Pullrequest) SetProjectcards(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Projectcard) error {
	query := "update \"projectcards\" set \"pullrequest\" = null where \"pullrequest\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Projectcards {
			queries.SetScanner(&rel.Pullrequest, nil)
			if rel.R == nil {
				continue
			}

			rel.R.ProjectcardPullrequest = nil
		}
		o.R.Projectcards = nil
	}

	return o.AddProjectcards(ctx, exec, insert, related...)
}

// RemoveProjectcards relationships from objects passed in.
// Removes related items from R.Projectcards (uses pointer comparison, removal does not keep order)
// Sets related.R.ProjectcardPullrequest.
func (o *Pullrequest) RemoveProjectcards(ctx context.Context, exec boil.ContextExecutor, related ...*Projectcard) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.Pullrequest, nil)
		if rel.R != nil {
			rel.R.ProjectcardPullrequest = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("pullrequest")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Projectcards {
			if rel != ri {
				continue
			}

			ln := len(o.R.Projectcards)
			if ln > 1 && i < ln-1 {
				o.R.Projectcards[i] = o.R.Projectcards[ln-1]
			}
			o.R.Projectcards = o.R.Projectcards[:ln-1]
			break
		}
	}

	return nil
}

// Pullrequests retrieves all the records using an executor.
func Pullrequests(mods ...qm.QueryMod) pullrequestQuery {
	mods = append(mods, qm.From("\"pullrequests\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"pullrequests\".*"})
	}

	return pullrequestQuery{q}
}

// FindPullrequest retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPullrequest(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Pullrequest, error) {
	pullrequestObj := &Pullrequest{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pullrequests\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pullrequestObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pullrequests")
	}

	if err = pullrequestObj.doAfterSelectHooks(ctx, exec); err != nil {
		return pullrequestObj, err
	}

	return pullrequestObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Pullrequest) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pullrequests provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pullrequestColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pullrequestInsertCacheMut.RLock()
	cache, cached := pullrequestInsertCache[key]
	pullrequestInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pullrequestAllColumns,
			pullrequestColumnsWithDefault,
			pullrequestColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pullrequestType, pullrequestMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pullrequestType, pullrequestMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pullrequests\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pullrequests\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into pullrequests")
	}

	if !cached {
		pullrequestInsertCacheMut.Lock()
		pullrequestInsertCache[key] = cache
		pullrequestInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Pullrequest.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Pullrequest) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	pullrequestUpdateCacheMut.RLock()
	cache, cached := pullrequestUpdateCache[key]
	pullrequestUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pullrequestAllColumns,
			pullrequestPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pullrequests, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pullrequests\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pullrequestPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pullrequestType, pullrequestMapping, append(wl, pullrequestPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update pullrequests row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pullrequests")
	}

	if !cached {
		pullrequestUpdateCacheMut.Lock()
		pullrequestUpdateCache[key] = cache
		pullrequestUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q pullrequestQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pullrequests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pullrequests")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PullrequestSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pullrequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pullrequests\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pullrequestPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pullrequest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pullrequest")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Pullrequest) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pullrequests provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pullrequestColumnsWithDefault, o)

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

	pullrequestUpsertCacheMut.RLock()
	cache, cached := pullrequestUpsertCache[key]
	pullrequestUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pullrequestAllColumns,
			pullrequestColumnsWithDefault,
			pullrequestColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			pullrequestAllColumns,
			pullrequestPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert pullrequests, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pullrequestPrimaryKeyColumns))
			copy(conflict, pullrequestPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"pullrequests\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pullrequestType, pullrequestMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pullrequestType, pullrequestMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert pullrequests")
	}

	if !cached {
		pullrequestUpsertCacheMut.Lock()
		pullrequestUpsertCache[key] = cache
		pullrequestUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Pullrequest record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pullrequest) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Pullrequest provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pullrequestPrimaryKeyMapping)
	sql := "DELETE FROM \"pullrequests\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pullrequests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pullrequests")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pullrequestQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pullrequestQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pullrequests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pullrequests")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PullrequestSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(pullrequestBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pullrequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pullrequests\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pullrequestPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pullrequest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pullrequests")
	}

	if len(pullrequestAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Pullrequest) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPullrequest(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PullrequestSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PullrequestSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pullrequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pullrequests\".* FROM \"pullrequests\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pullrequestPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PullrequestSlice")
	}

	*o = slice

	return nil
}

// PullrequestExists checks if the Pullrequest row exists.
func PullrequestExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pullrequests\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pullrequests exists")
	}

	return exists, nil
}

// Exists checks if the Pullrequest row exists.
func (o *Pullrequest) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PullrequestExists(ctx, exec, o.ID)
}
