package protobuf

import (
 // "log"
  "context"
)

type Server struct{}

func (s *Server) ServerRequest(ctx context.Context, sk *SearchKey) (*ReturnValue, error) {
  return &ReturnValue{}, nil
}
