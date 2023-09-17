// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: order.proto

package order

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	GetOrderByID(ctx context.Context, in *OrderID, opts ...client.CallOption) (*OrderInfo, error)
	GetAllOrder(ctx context.Context, in *AllOrderRequest, opts ...client.CallOption) (*AllOrder, error)
	CreateOrder(ctx context.Context, in *OrderInfo, opts ...client.CallOption) (*OrderID, error)
	DeleteOrderByID(ctx context.Context, in *OrderID, opts ...client.CallOption) (*Response, error)
	UpdateOrderPayStatus(ctx context.Context, in *PayStatus, opts ...client.CallOption) (*Response, error)
	UpdateOrderShipStatus(ctx context.Context, in *ShipStatus, opts ...client.CallOption) (*Response, error)
	UpdateOrder(ctx context.Context, in *OrderInfo, opts ...client.CallOption) (*Response, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) GetOrderByID(ctx context.Context, in *OrderID, opts ...client.CallOption) (*OrderInfo, error) {
	req := c.c.NewRequest(c.name, "Order.GetOrderByID", in)
	out := new(OrderInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetAllOrder(ctx context.Context, in *AllOrderRequest, opts ...client.CallOption) (*AllOrder, error) {
	req := c.c.NewRequest(c.name, "Order.GetAllOrder", in)
	out := new(AllOrder)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CreateOrder(ctx context.Context, in *OrderInfo, opts ...client.CallOption) (*OrderID, error) {
	req := c.c.NewRequest(c.name, "Order.CreateOrder", in)
	out := new(OrderID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) DeleteOrderByID(ctx context.Context, in *OrderID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.DeleteOrderByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrderPayStatus(ctx context.Context, in *PayStatus, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateOrderPayStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrderShipStatus(ctx context.Context, in *ShipStatus, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateOrderShipStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrder(ctx context.Context, in *OrderInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Order.UpdateOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	GetOrderByID(context.Context, *OrderID, *OrderInfo) error
	GetAllOrder(context.Context, *AllOrderRequest, *AllOrder) error
	CreateOrder(context.Context, *OrderInfo, *OrderID) error
	DeleteOrderByID(context.Context, *OrderID, *Response) error
	UpdateOrderPayStatus(context.Context, *PayStatus, *Response) error
	UpdateOrderShipStatus(context.Context, *ShipStatus, *Response) error
	UpdateOrder(context.Context, *OrderInfo, *Response) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		GetOrderByID(ctx context.Context, in *OrderID, out *OrderInfo) error
		GetAllOrder(ctx context.Context, in *AllOrderRequest, out *AllOrder) error
		CreateOrder(ctx context.Context, in *OrderInfo, out *OrderID) error
		DeleteOrderByID(ctx context.Context, in *OrderID, out *Response) error
		UpdateOrderPayStatus(ctx context.Context, in *PayStatus, out *Response) error
		UpdateOrderShipStatus(ctx context.Context, in *ShipStatus, out *Response) error
		UpdateOrder(ctx context.Context, in *OrderInfo, out *Response) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) GetOrderByID(ctx context.Context, in *OrderID, out *OrderInfo) error {
	return h.OrderHandler.GetOrderByID(ctx, in, out)
}

func (h *orderHandler) GetAllOrder(ctx context.Context, in *AllOrderRequest, out *AllOrder) error {
	return h.OrderHandler.GetAllOrder(ctx, in, out)
}

func (h *orderHandler) CreateOrder(ctx context.Context, in *OrderInfo, out *OrderID) error {
	return h.OrderHandler.CreateOrder(ctx, in, out)
}

func (h *orderHandler) DeleteOrderByID(ctx context.Context, in *OrderID, out *Response) error {
	return h.OrderHandler.DeleteOrderByID(ctx, in, out)
}

func (h *orderHandler) UpdateOrderPayStatus(ctx context.Context, in *PayStatus, out *Response) error {
	return h.OrderHandler.UpdateOrderPayStatus(ctx, in, out)
}

func (h *orderHandler) UpdateOrderShipStatus(ctx context.Context, in *ShipStatus, out *Response) error {
	return h.OrderHandler.UpdateOrderShipStatus(ctx, in, out)
}

func (h *orderHandler) UpdateOrder(ctx context.Context, in *OrderInfo, out *Response) error {
	return h.OrderHandler.UpdateOrder(ctx, in, out)
}
