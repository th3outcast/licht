package main

import (
  "fmt"
  _ "net/http"

  "github.com/dgraph-io/badger"
)

type App struct {
  db *badger.DB
}

func (a *App) GetValue(key []byte) []byte {
  var valCopy []byte
  _ = a.db.View(func(txn *badger.Txn) error {
    item, err := txn.Get(key)

    if err != badger.ErrKeyNotFound {
      // should not fail
      valCopy, err = item.ValueCopy(nil)
      return nil
    }
    return err
  })
  return valCopy
}

func (a *App) PutValue(key, value []byte) bool {
  success := true
  _ = a.db.Update(func(txn *badger.Txn) error {
    err := txn.Set([]byte(key), []byte(value))
    if err != nil {
      success = false
    }
    return err
  })
  return success
}

func main() {
  db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
  if err != nil {
    fmt.Printf("BadgerDB failed to open: %s", err)
  }
  defer db.Close()

  a := App{
    db: db,
  }
}
