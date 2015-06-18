export GOPATH=/c/Users/jwang/go

go build engine/engine.go

cd engine

go install

cd ../blackbox/

go run main.go

Revel
=====
go get github.com/revel/revel

go get github.com/revel/cmd/revel

revel help

revel new godashboard

mv godashboard github.com/j3rrywan9

revel run github.com/j3rrywan9/godashboard
