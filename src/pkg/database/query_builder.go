package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func (d *mariaMasterSlave) Builder() *goqu.Database {
	return d.goqu
}

func (d *mariaMasterSlave) Table(name string, stmt bool) *goqu.SelectDataset {
	return d.goqu.From(name).Prepared(stmt)
}

func (d *mariaMasterSlave) InsertWithTx(ctx context.Context, tx *sql.Tx, table interface{}, params map[string]interface{}) (uint64, error) {
	query, args, e := d.goqu.From(table).Insert().Rows(params).ToSQL()
	if e != nil {
		return 0, e
	}
	d.debugInfo(query, args)
	result, e := tx.Exec(query, args...)
	if e != nil {
		return 0, e
	}
	id, _ := result.LastInsertId()
	return uint64(id), nil
}

func (d *mariaMasterSlave) Insert(ctx context.Context, table interface{}, params map[string]interface{}) (uint64, error) {
	prepareDs := d.goqu.From(table).Prepared(true)
	query, args, e := prepareDs.Insert().Rows(params).ToSQL()
	if e != nil {
		return 0, e
	}
	d.debugInfo(query, args)
	tx, e := d.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if e != nil {
		return 0, e
	}
	result, e := tx.Exec(query, args...)
	if e != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, e
	}
	e = tx.Commit()
	if e != nil {
		return 0, e
	}

	id, _ := result.LastInsertId()
	return uint64(id), nil
}

func (d *mariaMasterSlave) Upsert(ctx context.Context, table interface{}, params map[string]interface{}) (uint64, error) {
	prepareDs := d.goqu.From(table).Prepared(true)
	query, args, e := prepareDs.Insert().Rows(params).OnConflict(goqu.DoUpdate("key", params)).ToSQL()
	if e != nil {
		return 0, e
	}
	d.debugInfo(query, args)
	tx, e := d.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if e != nil {
		return 0, e
	}
	result, e := tx.Exec(query, args...)
	if e != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, e
	}
	e = tx.Commit()
	if e != nil {
		return 0, e
	}

	id, _ := result.LastInsertId()
	return uint64(id), nil
}

func (d *mariaMasterSlave) UpsertWithTx(ctx context.Context, tx *sql.Tx, table interface{}, params map[string]interface{}) (uint64, error) {
	query, args, e := d.goqu.From(table).Prepared(true).Insert().Rows(params).OnConflict(goqu.DoUpdate("key", params)).ToSQL()
	if e != nil {
		return 0, e
	}
	d.debugInfo(query, args)
	result, e := tx.Exec(query, args...)
	if e != nil {
		return 0, e
	}
	id, _ := result.LastInsertId()
	return uint64(id), nil
}

func (d *mariaMasterSlave) Delete(ctx context.Context, table, field string, value interface{}) error {
	query, args, e := d.Table(table, true).Where(
		goqu.C(field).Eq(value)).Delete().Prepared(true).ToSQL()
	if e != nil {
		return e
	}
	d.debugInfo(query, args)
	tx, e := d.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if e != nil {
		return e
	}
	_, e = tx.Exec(query, args...)
	if e != nil {
		tx.Rollback()
		return e
	}
	tx.Commit()
	return nil
}

func (d *mariaMasterSlave) SofDelete(ctx context.Context, table string, id interface{}) error {
	now := time.Now()
	query, args, e := d.goqu.Update(table).Set(goqu.Record{
		"deleted_at": &now,
	}).Where(goqu.C("id").Eq(id)).Prepared(true).ToSQL()
	if e != nil {
		return e
	}
	d.debugInfo(query, args)

	tx, e := d.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

	if e != nil {
		return e
	}

	_, e = tx.Exec(query, args...)

	if e != nil {
		tx.Rollback()
		return e
	}
	tx.Commit()
	return nil
}
