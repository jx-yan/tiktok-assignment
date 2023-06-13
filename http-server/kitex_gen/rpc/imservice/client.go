package imservice

import (
	"context"
	rpc "github.com/TikTokTechImmersion/assignment_demo_2023/http-server/kitex_gen/rpc"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Send(ctx context.Context, req *rpc.SendRequest, callOptions ...callopt.Option) (r *rpc.SendResponse, err error)
	Pull(ctx context.Context, req *rpc.PullRequest, callOptions ...callopt.Option) (r *rpc.PullResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kIMServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kIMServiceClient struct {
	*kClient
}

func (p *kIMServiceClient) Send(ctx context.Context, req *rpc.SendRequest, callOptions ...callopt.Option) (r *rpc.SendResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Send(ctx, req)
}

func (p *kIMServiceClient) Pull(ctx context.Context, req *rpc.PullRequest, callOptions ...callopt.Option) (r *rpc.PullResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Pull(ctx, req)
}
