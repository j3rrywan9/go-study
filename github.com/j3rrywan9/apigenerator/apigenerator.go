package apigenerator

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/j3rrywan9/engine"
)

type ApiGenerator struct {
}

func (ag *ApiGenerator) LoadSchemaFromEngine(engine *engine.Engine, path string) {
	url := "http://" + engine.Address() + path
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
