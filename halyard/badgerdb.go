package halyard

import (
	badger "github.com/dgraph-io/badger/v4"
	"log"
)

type DB interface {
	Get([]byte) ([]byte, error)
	Set([]byte, []byte) error
	AllForPrefix([]byte) ([]([]byte), error)
}

type BadgerDB struct {
	bdb *badger.DB
}

func NewBadgerDB(path string) (*BadgerDB, error) {
	dbc, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		log.Fatal(err)
	}
	db := BadgerDB{}
	db.bdb = dbc
	return &db, err
}

func (db *BadgerDB) Get(key []byte) ([]byte, error) {
	var valCopy []byte
	err := db.bdb.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		return nil
	})
	return valCopy, err
}

func (db *BadgerDB) Set(key []byte, val []byte) error {
	// do some checking on the key type, etc  TODO XXX
	err := db.bdb.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
	return err
}

func (db *BadgerDB) AllForPrefix(prefix []byte) ([]([]byte), error) {
	var ret []([]byte)
	err := db.bdb.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			itemv, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			ret = append(ret, itemv)
		}
		return nil
	})
	return ret, err
}
