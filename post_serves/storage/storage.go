package storage

import (
	"examLast/post_serves/storage/postgres"
	"examLast/post_serves/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Post() repo.PostInfoI
}

type storagePg struct {
	db         *sqlx.DB
	postRepo repo.PostInfoI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:         db,
		postRepo: postgres.NewPostRepo(db),
	}
}
func (s storagePg) Post() repo.PostInfoI {
	return s.postRepo
}
