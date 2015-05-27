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
func (this *Engine) Address() string {
	return this.address
}

// Setter
func (this *Engine) SetAddress(address string) {
	this.address = address
}
