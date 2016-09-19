package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Margo API", func() {
	Title("Margo API")
	Description("Rudimentary yet scalable in-memory cache.")
	Version("0.0.1")
	Contact(func() {
		Name("Elisey Zanko")
		Email("elisey.zanko@gmail.com")
		URL("https://github.com/31z4/margo")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/31z4/margo/blob/master/LICENSE")
	})
	Scheme("http")
	Host("localhost:8080")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("keys", func() {
	BasePath("/keys")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all keys.")
		Response(OK, ArrayOf(String), func() {
			Media("application/json")
		})
	})

	Action("get", func() {
		Routing(
			GET("/:key"),
		)
		Description("Get the value of a key.")
		Params(func() {
			Param("key", String)
		})
		Response(OK, Any, func() {
			Media("application/json")
		})
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("getElement", func() {
		Routing(
			GET("/:key/:element"),
		)
		Description("Get the element of the list or dict value stored at key.")
		Params(func() {
			Param("key", String)
		})
		Params(func() {
			Param("element", String)
		})
		Response(OK, Any, func() {
			Media("application/json")
		})
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("set", func() {
		Routing(
			PUT("/:key"),
		)
		Description("Set the value of a key.")
		Payload(Any)
		Params(func() {
			Param("key", String)
		})
		Response(OK)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PATCH("/:key"),
		)
		Description("Update the value of a key.")
		Payload(Any)
		Params(func() {
			Param("key", String)
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("remove", func() {
		Routing(
			DELETE("/:key"),
		)
		Description("Delete a key.")
		Params(func() {
			Param("key", String)
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})
