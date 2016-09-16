//************************************************************************//
// API "Margo API": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// KeysController is the controller interface for the Keys actions.
type KeysController interface {
	goa.Muxer
	Get(*GetKeysContext) error
	List(*ListKeysContext) error
	Remove(*RemoveKeysContext) error
	Set(*SetKeysContext) error
	Update(*UpdateKeysContext) error
}

// MountKeysController "mounts" a Keys resource controller on the given service.
func MountKeysController(service *goa.Service, ctrl KeysController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetKeysContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	service.Mux.Handle("GET", "/keys/:key", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "Keys", "action", "Get", "route", "GET /keys/:key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListKeysContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/keys", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Keys", "action", "List", "route", "GET /keys")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRemoveKeysContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Remove(rctx)
	}
	service.Mux.Handle("DELETE", "/keys/:key", ctrl.MuxHandler("Remove", h, nil))
	service.LogInfo("mount", "ctrl", "Keys", "action", "Remove", "route", "DELETE /keys/:key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSetKeysContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(SetKeysPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Set(rctx)
	}
	service.Mux.Handle("PUT", "/keys/:key", ctrl.MuxHandler("Set", h, unmarshalSetKeysPayload))
	service.LogInfo("mount", "ctrl", "Keys", "action", "Set", "route", "PUT /keys/:key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateKeysContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(UpdateKeysPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PATCH", "/keys/:key", ctrl.MuxHandler("Update", h, unmarshalUpdateKeysPayload))
	service.LogInfo("mount", "ctrl", "Keys", "action", "Update", "route", "PATCH /keys/:key")
}

// unmarshalSetKeysPayload unmarshals the request body into the context request data Payload field.
func unmarshalSetKeysPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	var payload SetKeysPayload
	if err := service.DecodeRequest(req, &payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload
	return nil
}

// unmarshalUpdateKeysPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateKeysPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	var payload UpdateKeysPayload
	if err := service.DecodeRequest(req, &payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload
	return nil
}
