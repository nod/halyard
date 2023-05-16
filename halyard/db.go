package halyard

import (
  "log"
  badger "github.com/dgraph-io/badger/v4"
)

func OpenDB(path string) (*badger.DB,error) {
    db, err := badger.Open(badger.DefaultOptions(path))
    if err != nil {
      log.Fatal(err)
    }
    return db, err
}


