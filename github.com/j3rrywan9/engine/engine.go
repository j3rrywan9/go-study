package engine

import (
	"net/http"
)

const testAddress string = "172.16.100.79"

type HttpSession struct {
	client http.Client
}

type Engine struct {
	address string
	session HttpSession
}

func (de *Engine) Address() string {
	de.address = testAddress
	return de.address
}
