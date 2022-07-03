// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package order

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Order interface {
		Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
		CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
		CreateOrderCheck(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
		RollbackOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.Orders(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrderCheck(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.CreateOrderCheck(ctx, in, opts...)
}

func (m *defaultOrder) RollbackOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := NewOrderClient(m.cli.Conn())
	return client.RollbackOrder(ctx, in, opts...)
}
