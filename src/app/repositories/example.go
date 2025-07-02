package repositories

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/entity"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/database"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/util"
)

type example struct {
	db database.Adapter
	// mongo mongodb.Adapter
}

// func NewExample(db database.Adapter, mongo mongodb.Adapter) Example {
func NewExample(db database.Adapter) Example {
	return &example{
		db: db,
		// mongo: mongo,
	}
}

func (r *example) Find(ctx context.Context, limit, page uint64) ([]entity.Example, interface{}, error) {

	p := []entity.Example{}

	prepareQuery := r.db.Builder().From(TABLE_NAME_EXAMPLE).Where(goqu.C("deleted_at").IsNull()).Order(
		goqu.C("id").Desc(),
	)

	meta := r.db.Meta(ctx, prepareQuery, limit, page)
	limit, page = util.Pagination(limit, page)
	q, args, e := prepareQuery.Limit(uint(limit)).Offset(uint(page)).Prepared(true).ToSQL()
	if e != nil {
		return p, nil, e
	}
	e = r.db.Fetch(ctx, &p, q, args...)
	return p, meta, e
}

func (r *example) Fetch(ctx context.Context, id uint64) (*entity.Example, error) {
	p := &entity.Example{}
	q, args, e := r.db.Builder().From(TABLE_NAME_EXAMPLE).
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("deleted_at").IsNull()).Prepared(true).ToSQL()
	if e != nil {
		return p, e
	}
	e = r.db.FetchRow(ctx, p, q, args...)
	return p, e
}

// Upsert data no tx
func (r *example) Upsert(ctx context.Context, param entity.Example) (uint64, error) {
	return r.db.Upsert(ctx, TABLE_NAME_EXAMPLE, map[string]interface{}{
		"id":         param.ID,
		"name":       param.Name,
		"email":      param.Email,
		"phone":      param.Phone,
		"created_at": param.CreatedAt,
		"updated_at": param.UpdatedAt,
		"deleted_at": param.DeletedAt,
	})
}

// Upsert  data with tx
func (r *example) UpsertWithTx(ctx context.Context, tx *sql.Tx, param entity.Example) (uint64, error) {
	return r.db.UpsertWithTx(ctx, tx, TABLE_NAME_EXAMPLE, map[string]interface{}{
		"id":         param.ID,
		"name":       param.Name,
		"email":      param.Email,
		"phone":      param.Phone,
		"created_at": param.CreatedAt,
		"updated_at": param.UpdatedAt,
		"deleted_at": param.DeletedAt,
	})
}

func (r *example) Delete(ctx context.Context, id uint64) error {
	// Hard delete
	// return r.db.Delete(ctx, TABLE_NAME_EXAMPLE, "id", id)
	// If you use soft delete please mark deleted_at as
	return r.db.SofDelete(ctx, TABLE_NAME_EXAMPLE, id)

}
