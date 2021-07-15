package lib

import (
  "fmt"
)

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
