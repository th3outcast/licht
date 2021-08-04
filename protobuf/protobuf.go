package protobuf

import (
	// "log"
	"context"

	"github.com/th3outcast/licht/lib"
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
