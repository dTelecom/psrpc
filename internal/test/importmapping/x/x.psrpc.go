// Code generated by protoc-gen-psrpc v0.2.0, DO NOT EDIT.
// source: x/x.proto

package x

import context "context"

import psrpc "github.com/livekit/psrpc"
import psrpc_internal_test_importmapping_y "github.com/livekit/psrpc/internal/test/importmapping/y"

// =====================
// Svc1 Client Interface
// =====================

type Svc1Client interface {
	Send(context.Context, *psrpc_internal_test_importmapping_y.MsgY, ...psrpc.RequestOption) (*psrpc_internal_test_importmapping_y.MsgY, error)
}

// =========================
// Svc1 ServerImpl Interface
// =========================

type Svc1ServerImpl interface {
	Send(context.Context, *psrpc_internal_test_importmapping_y.MsgY) (*psrpc_internal_test_importmapping_y.MsgY, error)
}

// =====================
// Svc1 Server Interface
// =====================

type Svc1Server interface {
	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ===========
// Svc1 Client
// ===========

type svc1Client struct {
	client *psrpc.RPCClient
}

// NewSvc1Client creates a psrpc client that implements the Svc1Client interface.
func NewSvc1Client(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (Svc1Client, error) {
	rpcClient, err := psrpc.NewRPCClient("Svc1", clientID, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &svc1Client{
		client: rpcClient,
	}, nil
}

func (c *svc1Client) Send(ctx context.Context, req *psrpc_internal_test_importmapping_y.MsgY, opts ...psrpc.RequestOption) (*psrpc_internal_test_importmapping_y.MsgY, error) {
	return psrpc.RequestSingle[*psrpc_internal_test_importmapping_y.MsgY](ctx, c.client, "Send", "", req, opts...)
}

// ===========
// Svc1 Server
// ===========

type svc1Server struct {
	svc Svc1ServerImpl
	rpc *psrpc.RPCServer
}

// NewSvc1Server builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewSvc1Server(serverID string, svc Svc1ServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (Svc1Server, error) {
	s := psrpc.NewRPCServer("Svc1", serverID, bus, opts...)

	var err error
	err = psrpc.RegisterHandler(s, "Send", "", svc.Send, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &svc1Server{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *svc1Server) Shutdown() {
	s.rpc.Close(false)
}

func (s *svc1Server) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor0 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xce, 0x3d, 0x0a, 0x02, 0x31,
	0x10, 0x40, 0xe1, 0x66, 0x11, 0x76, 0xcb, 0x94, 0x0b, 0x62, 0x61, 0x65, 0x33, 0x83, 0x7a, 0x03,
	0x7b, 0xab, 0xad, 0x14, 0x9b, 0x6c, 0x32, 0x84, 0x40, 0x7e, 0x86, 0x24, 0xe8, 0xe4, 0xf6, 0x82,
	0x1e, 0x40, 0xfb, 0x8f, 0xc7, 0x9b, 0x46, 0x41, 0x01, 0x2e, 0xb9, 0x65, 0xb5, 0xe7, 0x5a, 0xd8,
	0x80, 0x4f, 0x8d, 0x4a, 0xd2, 0x01, 0x1a, 0xd5, 0x06, 0x3e, 0x72, 0x2e, 0x2d, 0x6a, 0x66, 0x9f,
	0x1c, 0xc8, 0x3c, 0x76, 0xec, 0x5f, 0x7f, 0xb2, 0xd3, 0xb0, 0x3c, 0xcd, 0x51, 0x3d, 0xa6, 0x61,
	0xa1, 0x64, 0xd5, 0x01, 0x7e, 0x07, 0x3a, 0x5c, 0xab, 0xbb, 0xcd, 0xff, 0xd3, 0xcb, 0xee, 0xbe,
	0x25, 0xd1, 0x91, 0x03, 0x81, 0xc9, 0x11, 0x5f, 0x3e, 0x04, 0x5c, 0x09, 0xbd, 0x4b, 0xb9, 0x90,
	0x45, 0x59, 0x37, 0x9f, 0x9b, 0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x6b, 0x1a, 0xa1, 0xca,
	0x00, 0x00, 0x00,
}
