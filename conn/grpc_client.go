package conn

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/th3outcast/licht/protobuf"
)

type GRPCClient struct {
	grpc_addr string
	conn      *grpc.ClientConn
	ctx       context.Context
	cancel    context.CancelFunc

	client protobuf.LichtClient
}

func NewGRPCClient(grpc_addr string) (*GRPCClient, error) {
	return NewGRPCClientWithContext(grpc_addr, context.Background())
}

func NewGRPCClientWithContext(grpc_addr string, baseCtx context.Context) (*GRPCClient, error) {
	return NewGRPCClientWithContextTLS(grpc_addr, baseCtx, "", "")
}

func NewGRPCClientWithContextTLS(grpc_addr string, baseCtx context.Context, certificateFile string, commonName string) (*GRPCClient, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithDefaultOptions(
			grpc.MaxCallSendMsgSize(math.MaxInt64),
			grpc.MaxCallRecvMsgSize(math.MaxInt64),
		),
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                1 * time.Second,
				TImeout:             5 * time.Second,
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
	}, nil
}

func (c *GRPCClient) Start()
