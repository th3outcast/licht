package lib

import (
	_ "net/http"
	"strings"

	"github.com/dgraph-io/badger"
)

type App struct {
	Db *badger.DB
}

func (a *App) GetValue(key []byte) []byte {
	var valCopy []byte
	_ = a.Db.View(func(txn *badger.Txn) error {
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

func (a *App) SetValue(key, value []byte) bool {
	success := true
	_ = a.Db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		if err != nil {
			success = false
		}
		return err
	})
	return success
}

type Record struct {
	Hash string
	Data []byte
}

func ToRecord(data []byte) Record {
	var rec Record
	ss := string(data)
	if strings.HasPrefix(ss, "HASH") {
		rec.Hash = ss[4:36]
		ss = ss[36:]
	}
	rec.Data = []byte(ss)
	return rec
}

func FromRecord(rec Record) []byte {
	cc := ""
	if len(rec.Hash) == 32 {
		cc += "HASH" + rec.Hash
	}
	return []byte(cc + string(rec.Data))
}
