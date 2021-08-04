package conn

import (
	"context"
	// "log"
	"math"
	// "net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	"github.com/th3outcast/licht/errors"
	"github.com/th3outcast/licht/protobuf"
)

type GRPCClient struct {
	grpc_addr string
	conn      *grpc.ClientConn
	ctx       context.Context
	cancel    context.CancelFunc

	client  protobuf.LichtClient
	nodeNum int32
}

func NewGRPCClient(grpc_addr string, nodeNum int32) (*GRPCClient, error) {
	return NewGRPCClientWithContext(grpc_addr, context.Background(), nodeNum)
}

func NewGRPCClientWithContext(grpc_addr string, baseCtx context.Context, nodeNum int32) (*GRPCClient, error) {
	return NewGRPCClientWithContextTLS(grpc_addr, baseCtx, "", "", nodeNum)
}

func NewGRPCClientWithContextTLS(grpc_addr string, baseCtx context.Context, certificateFile string, commonName string, nodeNum int32) (*GRPCClient, error) {
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

func (c *GRPCClient) RequestKey(req *protobuf.SearchKey, opts ...grpc.CallOption) (*protobuf.ReturnValue, error) {
	// if not correct nodeID return empty
	if req.Node != c.nodeNum {
		return &protobuf.ReturnValue{}, nil
	}
	resp, err := c.client.RequestKey(c.ctx, req, opts...)
	if err != nil {
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
}
