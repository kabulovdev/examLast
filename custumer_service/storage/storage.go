package storage

import (
	"examLast/custumer_service/storage/postgres"
	"examLast/custumer_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Custum() repo.CustumerInfoI
}

type storagePg struct {
	db         *sqlx.DB
	customRepo repo.CustumerInfoI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:         db,
		customRepo: postgres.NewCustumRepo(db),
	}
}
func (s storagePg) Custum() repo.CustumerInfoI {
	return s.customRepo
}
