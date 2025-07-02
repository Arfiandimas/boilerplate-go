// Package repositories
package repositories

import (
	"context"
	"database/sql"

	"github.com/kiriminaja/kaj-rest-engine-go/src/app/entity"
)

// Contoh penulisan table
const (
	TABLE_NAME_EXAMPLE = `example`
)

/*
Repositories Contract
Adalah Cerminan fungsi yang diakses langsung oleh usecase
Seluruh Fungsi yang terdapat pada repositories wajib dicerminkan
Hal ini untuk memudahkan anda dalam melakukan proses mockup atau
Mencegah terjadinya proses missing anda delete parameter pada funsi
Repositories

Untuk membuat repositories anda dapat membuat contract pada file
repositories/contract.go
*/
type Example interface {
	Find(ctx context.Context, limit, page uint64) ([]entity.Example, interface{}, error)
	Fetch(ctx context.Context, id uint64) (*entity.Example, error)
	Upsert(ctx context.Context, p entity.Example) (uint64, error)
	UpsertWithTx(ctx context.Context, tx *sql.Tx, param entity.Example) (uint64, error)
	Delete(ctx context.Context, id uint64) error
}
