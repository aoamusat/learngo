package design

import (
	"fmt"

	. "goa.design/goa/v3/dsl"
)

const HTTP_PORT = 6040
const GRPC_PORT = 6041

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for multiplying numbers, a Goa teaser")
	Server("calc", func() {
		Host("localhost", func() {
			URI(fmt.Sprintf("http://localhost:%d", HTTP_PORT))
			URI(fmt.Sprintf("grpc://localhost:%d", GRPC_PORT))
		})
	})
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("multiply", func() {
		Payload(func() {
			Field(1, "a", Int64, "Left operand")
			Field(2, "b", Int64, "Right operand")
			Required("a", "b")
		})

		Result(Int64)

		HTTP(func() {
			GET("/multiply/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
