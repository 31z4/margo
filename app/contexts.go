//************************************************************************//
// API "Margo API": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/31z4/margo/design
// --out=$(GOPATH)/src/github.com/31z4/margo
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// GetKeysContext provides the keys get action context.
type GetKeysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Key string
}

// NewGetKeysContext parses the incoming request URL and body, performs validations and creates the
// context used by the keys controller get action.
func NewGetKeysContext(ctx context.Context, service *goa.Service) (*GetKeysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetKeysContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramKey := req.Params["key"]
	if len(paramKey) > 0 {
		rawKey := paramKey[0]
		rctx.Key = rawKey
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetKeysContext) OK(r interface{}) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GetKeysContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetKeysContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// GetElementKeysContext provides the keys getElement action context.
type GetElementKeysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Element string
	Key     string
}

// NewGetElementKeysContext parses the incoming request URL and body, performs validations and creates the
// context used by the keys controller getElement action.
func NewGetElementKeysContext(ctx context.Context, service *goa.Service) (*GetElementKeysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetElementKeysContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramElement := req.Params["element"]
	if len(paramElement) > 0 {
		rawElement := paramElement[0]
		rctx.Element = rawElement
	}
	paramKey := req.Params["key"]
	if len(paramKey) > 0 {
		rawKey := paramKey[0]
		rctx.Key = rawKey
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetElementKeysContext) OK(r interface{}) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GetElementKeysContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetElementKeysContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// ListKeysContext provides the keys list action context.
type ListKeysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListKeysContext parses the incoming request URL and body, performs validations and creates the
// context used by the keys controller list action.
func NewListKeysContext(ctx context.Context, service *goa.Service) (*ListKeysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListKeysContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListKeysContext) OK(r []string) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// RemoveKeysContext provides the keys remove action context.
type RemoveKeysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Key string
}

// NewRemoveKeysContext parses the incoming request URL and body, performs validations and creates the
// context used by the keys controller remove action.
func NewRemoveKeysContext(ctx context.Context, service *goa.Service) (*RemoveKeysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RemoveKeysContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramKey := req.Params["key"]
	if len(paramKey) > 0 {
		rawKey := paramKey[0]
		rctx.Key = rawKey
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *RemoveKeysContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *RemoveKeysContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RemoveKeysContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// SetKeysContext provides the keys set action context.
type SetKeysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Key     string
	Payload SetKeysPayload
}

// NewSetKeysContext parses the incoming request URL and body, performs validations and creates the
// context used by the keys controller set action.
func NewSetKeysContext(ctx context.Context, service *goa.Service) (*SetKeysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := SetKeysContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramKey := req.Params["key"]
	if len(paramKey) > 0 {
		rawKey := paramKey[0]
		rctx.Key = rawKey
	}
	return &rctx, err
}

// SetKeysPayload is the keys set action payload.
type SetKeysPayload interface{}

// OK sends a HTTP response with status code 200.
func (ctx *SetKeysContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *SetKeysContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// UpdateKeysContext provides the keys update action context.
type UpdateKeysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Key     string
	Payload UpdateKeysPayload
}

// NewUpdateKeysContext parses the incoming request URL and body, performs validations and creates the
// context used by the keys controller update action.
func NewUpdateKeysContext(ctx context.Context, service *goa.Service) (*UpdateKeysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateKeysContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramKey := req.Params["key"]
	if len(paramKey) > 0 {
		rawKey := paramKey[0]
		rctx.Key = rawKey
	}
	return &rctx, err
}

// UpdateKeysPayload is the keys update action payload.
type UpdateKeysPayload interface{}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateKeysContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateKeysContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateKeysContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}
