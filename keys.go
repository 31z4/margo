package main

import (
	"github.com/31z4/margo/app"
	"github.com/goadesign/goa"
	"github.com/31z4/margo/storage"
)

type KeysController struct {
	*goa.Controller
	storage *storage.Storage
}

func NewKeysController(service *goa.Service) *KeysController {
	return &KeysController{
		Controller: service.NewController("KeysController"),
		storage: storage.New(),
	}
}

func (c *KeysController) Get(ctx *app.GetKeysContext) error {
	value, err := c.storage.Get(ctx.Key)
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	}
	return ctx.OK(value)
}

func (c *KeysController) GetElement(ctx *app.GetElementKeysContext) error {
	value, err := c.storage.GetElement(ctx.Key, ctx.Element)
	if err != nil {
		switch e := err.(type) {
		case *storage.NotFoundError:
			return ctx.NotFound(goa.ErrNotFound(e))
		case *storage.TypeError:
			return ctx.BadRequest(goa.ErrBadRequest(e))
		}
	}
	return ctx.OK(value)
}

func (c *KeysController) List(ctx *app.ListKeysContext) error {
	return ctx.OK(c.storage.Keys())
}

func (c *KeysController) Remove(ctx *app.RemoveKeysContext) error {
	if err := c.storage.Remove(ctx.Key); err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	}
	return ctx.OK([]byte(""))
}

func (c *KeysController) Set(ctx *app.SetKeysContext) error {
	if err := c.storage.Set(ctx.Key, ctx.Payload); err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	return ctx.OK([]byte(""))
}

func (c *KeysController) Update(ctx *app.UpdateKeysContext) error {
	if err := c.storage.Update(ctx.Key, ctx.Payload); err != nil {
		switch e := err.(type) {
		case *storage.NotFoundError:
			return ctx.NotFound(goa.ErrNotFound(e))
		case *storage.TypeError:
			return ctx.BadRequest(goa.ErrBadRequest(e))
		}
	}
	return ctx.OK([]byte(""))
}
