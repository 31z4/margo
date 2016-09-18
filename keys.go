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
	if value, err := c.storage.Get(ctx.Key); err == nil {
		return ctx.OK(value)
	}
	return ctx.NotFound()
}

func (c *KeysController) List(ctx *app.ListKeysContext) error {
	return ctx.OK(c.storage.Keys())
}

func (c *KeysController) Remove(ctx *app.RemoveKeysContext) error {
	c.storage.Remove(ctx.Key)
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
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	return ctx.OK([]byte(""))
}
