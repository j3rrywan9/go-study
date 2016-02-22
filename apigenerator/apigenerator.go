package apigenerator

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/j3rrywan9/engine"
)

type ApiGenerator struct {
}

func (ag *ApiGenerator) LoadSchemasFromEngine(engine *engine.Engine, path string) string {
	url := "http://" + engine.Address() + path
	client := new(http.Client)
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(robots)
}
