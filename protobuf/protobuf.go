package protobuf

import (
	// "log"
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/th3outcast/licht/lib"
	"github.com/th3outcast/licht/errors"
)

type Server struct {
	lib.App
}

func (s *Server) RequestKey(ctx context.Context, sk *SearchKey) (*ReturnValue, error) {
	var key []byte
	key = sk.GetKey()
	val := s.GetValue(key)

	// Turn to record
	var rec lib.Record
	rec = lib.ToRecord(val)

	// Return record fields
	rv := &ReturnValue{
		Hash: []byte(rec.Hash),
		Data: rec.Data,
	}
	return rv, nil
}

func (s *Server) SetKV(ctx context.Context, sk *SetKey) (*empty.Empty, error) {
  resp := &empty.Empty{}

  var data []byte
  hash := sk.GetHash()
  data = sk.GetData()

  rec := lib.Record{
    Hash: string(hash),
    Data: data,
  }

  data = lib.FromRecord(rec)
  var success bool
  success = s.SetValue(key, data)
  if success != true {
    return resp, errors.ErrSetKV
  }
  return resp, nil
}
