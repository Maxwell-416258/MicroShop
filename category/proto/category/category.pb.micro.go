// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: category.proto

package category

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

// Api Endpoints for Catagory service

func NewCatagoryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Catagory service

type CatagoryService interface {
	CreateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CreateCategoryResponse, error)
	UpdateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*UpdateCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...client.CallOption) (*DeleteCatoryResponse, error)
	FindCategoryByName(ctx context.Context, in *FindByNameRequest, opts ...client.CallOption) (*CategoryResponse, error)
	FindCategoryByID(ctx context.Context, in *FindByIdRequest, opts ...client.CallOption) (*CategoryResponse, error)
	FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, opts ...client.CallOption) (*FindAllResponse, error)
	FindCategoryByParent(ctx context.Context, in *FindByParentRequest, opts ...client.CallOption) (*FindAllResponse, error)
	FindAllCategory(ctx context.Context, in *FindAllRequest, opts ...client.CallOption) (*FindAllResponse, error)
}

type catagoryService struct {
	c    client.Client
	name string
}

func NewCatagoryService(name string, c client.Client) CatagoryService {
	return &catagoryService{
		c:    c,
		name: name,
	}
}

func (c *catagoryService) CreateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CreateCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catagory.CreateCategory", in)
	out := new(CreateCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catagoryService) UpdateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*UpdateCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catagory.UpdateCategory", in)
	out := new(UpdateCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catagoryService) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...client.CallOption) (*DeleteCatoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catagory.DeleteCategory", in)
	out := new(DeleteCatoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catagoryService) FindCategoryByName(ctx context.Context, in *FindByNameRequest, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catagory.FindCategoryByName", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catagoryService) FindCategoryByID(ctx context.Context, in *FindByIdRequest, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catagory.FindCategoryByID", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catagoryService) FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, opts ...client.CallOption) (*FindAllResponse, error) {
	req := c.c.NewRequest(c.name, "Catagory.FindCategoryByLevel", in)
	out := new(FindAllResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catagoryService) FindCategoryByParent(ctx context.Context, in *FindByParentRequest, opts ...client.CallOption) (*FindAllResponse, error) {
	req := c.c.NewRequest(c.name, "Catagory.FindCategoryByParent", in)
	out := new(FindAllResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catagoryService) FindAllCategory(ctx context.Context, in *FindAllRequest, opts ...client.CallOption) (*FindAllResponse, error) {
	req := c.c.NewRequest(c.name, "Catagory.FindAllCategory", in)
	out := new(FindAllResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Catagory service

type CatagoryHandler interface {
	CreateCategory(context.Context, *CategoryRequest, *CreateCategoryResponse) error
	UpdateCategory(context.Context, *CategoryRequest, *UpdateCategoryResponse) error
	DeleteCategory(context.Context, *DeleteCategoryRequest, *DeleteCatoryResponse) error
	FindCategoryByName(context.Context, *FindByNameRequest, *CategoryResponse) error
	FindCategoryByID(context.Context, *FindByIdRequest, *CategoryResponse) error
	FindCategoryByLevel(context.Context, *FindByLevelRequest, *FindAllResponse) error
	FindCategoryByParent(context.Context, *FindByParentRequest, *FindAllResponse) error
	FindAllCategory(context.Context, *FindAllRequest, *FindAllResponse) error
}

func RegisterCatagoryHandler(s server.Server, hdlr CatagoryHandler, opts ...server.HandlerOption) error {
	type catagory interface {
		CreateCategory(ctx context.Context, in *CategoryRequest, out *CreateCategoryResponse) error
		UpdateCategory(ctx context.Context, in *CategoryRequest, out *UpdateCategoryResponse) error
		DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, out *DeleteCatoryResponse) error
		FindCategoryByName(ctx context.Context, in *FindByNameRequest, out *CategoryResponse) error
		FindCategoryByID(ctx context.Context, in *FindByIdRequest, out *CategoryResponse) error
		FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, out *FindAllResponse) error
		FindCategoryByParent(ctx context.Context, in *FindByParentRequest, out *FindAllResponse) error
		FindAllCategory(ctx context.Context, in *FindAllRequest, out *FindAllResponse) error
	}
	type Catagory struct {
		catagory
	}
	h := &catagoryHandler{hdlr}
	return s.Handle(s.NewHandler(&Catagory{h}, opts...))
}

type catagoryHandler struct {
	CatagoryHandler
}

func (h *catagoryHandler) CreateCategory(ctx context.Context, in *CategoryRequest, out *CreateCategoryResponse) error {
	return h.CatagoryHandler.CreateCategory(ctx, in, out)
}

func (h *catagoryHandler) UpdateCategory(ctx context.Context, in *CategoryRequest, out *UpdateCategoryResponse) error {
	return h.CatagoryHandler.UpdateCategory(ctx, in, out)
}

func (h *catagoryHandler) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, out *DeleteCatoryResponse) error {
	return h.CatagoryHandler.DeleteCategory(ctx, in, out)
}

func (h *catagoryHandler) FindCategoryByName(ctx context.Context, in *FindByNameRequest, out *CategoryResponse) error {
	return h.CatagoryHandler.FindCategoryByName(ctx, in, out)
}

func (h *catagoryHandler) FindCategoryByID(ctx context.Context, in *FindByIdRequest, out *CategoryResponse) error {
	return h.CatagoryHandler.FindCategoryByID(ctx, in, out)
}

func (h *catagoryHandler) FindCategoryByLevel(ctx context.Context, in *FindByLevelRequest, out *FindAllResponse) error {
	return h.CatagoryHandler.FindCategoryByLevel(ctx, in, out)
}

func (h *catagoryHandler) FindCategoryByParent(ctx context.Context, in *FindByParentRequest, out *FindAllResponse) error {
	return h.CatagoryHandler.FindCategoryByParent(ctx, in, out)
}

func (h *catagoryHandler) FindAllCategory(ctx context.Context, in *FindAllRequest, out *FindAllResponse) error {
	return h.CatagoryHandler.FindAllCategory(ctx, in, out)
}
