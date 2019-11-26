package controller

import "strings"

// Hello is hello controller.
type Hello struct {
	rawstring string
}

// NewHello returns new Hello.
func NewHello() *Hello {
	return &Hello{rawstring: "hello"}
}

// HelloGetResponse is the response data.
type HelloGetResponse struct {
	data string
}

// Get returns HelloGetResponse that has repeated `hello` string.
func (h *Hello) Get(count uint32) *HelloGetResponse {
	resp := strings.Repeat(h.rawstring, int(count))

	return &HelloGetResponse{data: resp}
}
