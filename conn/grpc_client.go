package conn

import (
  "context"
  // "fmt"
  "log"
  "math"
  // "net"
  "time"

  "google.golang.org/grpc"
  "google.golang.org/grpc/keepalive"
  "google.golang.org/grpc/credentials"

  "github.com/th3outcast/licht/protobuf"
)

type GRPCClient struct {
	grpc_addr string
	conn      *grpc.ClientConn
	ctx       context.Context
	cancel    context.CancelFunc

	client    protobuf.LichtClient
    nodeNum   int
}

func NewGRPCClient(grpc_addr string, nodeNum int) (*GRPCClient, error) {
	return NewGRPCClientWithContext(grpc_addr, context.Background(), nodeNum)
}

func NewGRPCClientWithContext(grpc_addr string, baseCtx context.Context, nodeNum int) (*GRPCClient, error) {
	return NewGRPCClientWithContextTLS(grpc_addr, baseCtx, "", "", nodeNum)
}

func NewGRPCClientWithContextTLS(grpc_addr string, baseCtx context.Context, certificateFile string, commonName string, nodeNum int) (*GRPCClient, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(math.MaxInt64),
			grpc.MaxCallRecvMsgSize(math.MaxInt64),
		),
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                1 * time.Second,
				Timeout:             5 * time.Second,
				PermitWithoutStream: true,
			},
		),
	}

	ctx, cancel := context.WithCancel(baseCtx)

	if certificateFile == "" {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile(certificateFile, commonName)
		if err != nil {
			return nil, err
		}
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.DialContext(ctx, grpc_addr, dialOpts...)
	if err != nil {
		cancel()
		return nil, err
	}

	return &GRPCClient{
		grpc_addr: grpc_addr,
		conn:      conn,
		ctx:       ctx,
		cancel:    cancel,
		client:    protobuf.NewLichtClient(conn),
        nodeNum:   nodeNum,
	}, nil
}

func (c *GRPCClient) ServerRequest(req *protobuf.SearchKey, opts ...grpc.CallOption) (*protobuf.ReturnValue, error) {
  
  if req.Node != int32(c.nodeNum) {
    return &protobuf.ReturnValue{}, nil
  }/*
  if resp, err := c.client.ServerRequest(c.ctx, req, opts...); err != nil {
    stat, _ := status.FromError(err)
    switch stat.Code() {
      case codes.NotFound:
        return nil, errors.ErrNotFound
      default:
        return nil, err
    }
  } else {
      return resp, nil
  }
  */

  resp, err := c.client.ServerRequest(c.ctx, req, opts...)
  if err != nil {
    log.Fatalf("Error when calling ServerRequest: %s\n", err)
  }
  log.Printf("Response from server: %s, %s\n", string(resp.Hash), string(resp.Data))
  return &protobuf.ReturnValue{
    Hash: []byte("client"),
    Data: []byte("side"),
  }, nil
}

