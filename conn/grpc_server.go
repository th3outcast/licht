package conn

import (
  "fmt"
  "net"
  "log"
  "math"
  "time"
  //"context"

  "google.golang.org/grpc"
  "google.golang.org/grpc/keepalive"

  "github.com/th3outcast/licht/protobuf"
)

type GRPCServer struct {
  port string
  server *grpc.Server
  listener net.Listener
}

func NewGRPCServer(port string) (*GRPCServer, error) {
  opts := []grpc.ServerOption {
    grpc.MaxRecvMsgSize(math.MaxInt64),
    grpc.MaxRecvMsgSize(math.MaxInt64),
    grpc.KeepaliveParams(
      keepalive.ServerParameters {
        Time: 5 * time.Second,
        Timeout: 5 * time.Second,
      },
    ),
  }

  server := grpc.NewServer(
    opts...,
  )

  s := protobuf.Server{}
  protobuf.RegisterLichtServer(server, &s)

  listener, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to create listener over port %s", port)
  }

  return &GRPCServer{
    port: port,
    server: server,
    listener: listener,
  }, nil
}

func (s *GRPCServer) Start() error {
  /* 
  // Not entirely necessary?
  if err := s.server.Serve(s.listener); err != nil {
    log.Fatalf("Failed to start gRPC server: %s", err)
  }
  */

  go func() {
    _ = s.server.Serve(s.listener)
  }()
  fmt.Printf("gRPC server started over port: %s\n", s.port)
  return nil
}

func (s *GRPCServer) Stop() error {
  s.server.Stop()

  fmt.Printf("gRPC server stopped.")
  return nil
}

/*
func (s *GRPCServer) ServerRequest(ctx context.Context, req *protobuf.SearchKey) (*protobuf.ReturnValue, error) {
  rv := &protobuf.ReturnValue{
    Hash: []byte("test"),
    Data: []byte("run"),
  }
  return rv, nil
}
*/
