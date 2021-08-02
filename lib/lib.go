package lib

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

type Record struct {
  hash string
  data []byte
}

func toRecord(data []byte) Record {
  var rec Record
  ss := string(data)
  if strings.HasPrefix(ss, "HASH") {
    rec.hash = ss[4:36]
    ss = ss[36:]
  }
  rec.data = []byte(ss)
  return rec
}

func fromRecord(rec Record)  []byte {
  cc := ""
  if len(rec.hash) == 32 {
    cc += "HASH" + rec.hash
  }
  return []byte(cc + string(rec.data))
}
