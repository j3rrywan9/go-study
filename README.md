export GOPATH=/c/Users/jwang/go

go build engine/engine.go

cd engine

go install

cd ../blackbox/

go run main.go
