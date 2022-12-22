// Code generated by protoc-gen-psrpc v0.0.1, DO NOT EDIT.
// source: service_method_same_name.proto

package service_method_same_name

import context "context"

import psrpc "github.com/livekit/psrpc"

// =====================
// Echo Client Interface
// =====================

type EchoClient interface {
	Echo(context.Context, *Msg, ...psrpc.RequestOpt) (*Msg, error)
}

// =========================
// Echo ServerImpl Interface
// =========================

type EchoServerImpl interface {
	Echo(context.Context, *Msg) (*Msg, error)
}

// =====================
// Echo Server Interface
// =====================

type EchoServer interface {
}

// ===========
// Echo Client
// ===========

type echoClient struct {
	client psrpc.RPCClient
}

// NewEchoClient creates a psrpc client that implements the EchoClient interface.
func NewEchoClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOpt) (EchoClient, error) {
	rpcClient, err := psrpc.NewRPCClient("Echo", clientID, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &echoClient{
		client: rpcClient,
	}, nil
}

func (c *echoClient) Echo(ctx context.Context, req *Msg, opts ...psrpc.RequestOpt) (*Msg, error) {
	return psrpc.RequestTopicSingle[*Msg](ctx, c.client, "Echo", "", req, opts...)
}

// ===========
// Echo Server
// ===========

type echoServer struct {
	svc EchoServerImpl
	rpc psrpc.RPCServer
}

// NewEchoServer builds a RPCServer that can be used to handle
// requests that are routed to the right method in the provided svc implementation.
func NewEchoServer(serverID string, svc EchoServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOpt) (EchoServer, error) {
	rpcServer := psrpc.NewRPCServer("Echo", serverID, bus, opts...)

	var err error
	err = rpcServer.RegisterHandler(psrpc.NewHandler("Echo", svc.Echo))
	if err != nil {
		rpcServer.Close()
		return nil, err
	}

	return &echoServer{
		svc: svc,
		rpc: rpcServer,
	}, nil
}

var psrpcFileDescriptor0 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0xcf, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0x89, 0x2f, 0x4e, 0xcc, 0x4d, 0x8d,
	0xcf, 0x4b, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x62, 0xe5, 0x62, 0xf6, 0x2d,
	0x4e, 0x37, 0x92, 0xe1, 0x62, 0x71, 0x4d, 0xce, 0xc8, 0x17, 0x12, 0x81, 0xd2, 0x2c, 0x7a, 0xbe,
	0xc5, 0xe9, 0x52, 0x60, 0x52, 0x89, 0xc1, 0x49, 0x3a, 0x4a, 0x52, 0x1f, 0x97, 0x39, 0x49, 0x6c,
	0x60, 0x83, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97, 0x74, 0x2e, 0x74, 0x6a, 0x00, 0x00,
	0x00,
}
