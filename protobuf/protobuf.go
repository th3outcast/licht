package protobuf

import (
 // "log"
  "context"
)

type Server struct{}

func (s *Server) ServerRequest(ctx context.Context, sk *SearchKey) (*ReturnValue, error) {
  rv := &ReturnValue{
    Hash: []byte("test"),
    Data: []byte("run"),
  }
  return rv, nil
}
