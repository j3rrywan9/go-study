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
func (de *Engine) Address() string {
	return de.address
}

// Setter
func (de *Engine) SetAddress(addr string) {
	de.address = addr
}
