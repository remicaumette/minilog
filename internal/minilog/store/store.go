package store

import (
	"github.com/remicaumette/minilog/pkg/record"
	uuid "github.com/satori/go.uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Store struct {
	db *leveldb.DB
}

func New(path string) (*Store, error) {
	db, err := leveldb.OpenFile(path, &opt.Options{})
	if err != nil {
		return nil, err
	}
	return &Store{db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Store(record *record.Record) error {
	key := uuid.NewV4().Bytes()
	value, err := record.ToBinary()
	if err != nil {
		return err
	}
	return s.db.Put(key, value, &opt.WriteOptions{})
}
