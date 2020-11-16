package engine

import (
	"net/http"
)

type HttpSession struct {
	Client http.Client
}

type Engine struct {
	address string
	Session HttpSession
}

// Getter
func (e *Engine) Address() string {
	return this.address
}

// Setter
func (e *Engine) SetAddress(address string) {
	e.address = address
}
