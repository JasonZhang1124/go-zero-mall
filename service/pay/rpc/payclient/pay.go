// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package payclient

import (
	"context"

	"go-zero-mall/service/pay/rpc/types/pay"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CallbackRequest  = pay.CallbackRequest
	CallbackResponse = pay.CallbackResponse
	CreateRequest    = pay.CreateRequest
	CreateResponse   = pay.CreateResponse
	DetailRequest    = pay.DetailRequest
	DetailResponse   = pay.DetailResponse

	Pay interface {
		Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
		Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
		Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error)
	}

	defaultPay struct {
		cli zrpc.Client
	}
)

func NewPay(cli zrpc.Client) Pay {
	return &defaultPay{
		cli: cli,
	}
}

func (m *defaultPay) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultPay) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}

func (m *defaultPay) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.Callback(ctx, in, opts...)
}
