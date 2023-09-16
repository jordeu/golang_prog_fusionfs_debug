all: get build

release:
	CGO_ENABLED=0 GOARCH=amd64 go build -o fusionfs_debug --ldflags '-extldflags "-static"' fusionfs_debug.go

get:
	go get -t ./...

build:
	go build -o fusionfs_debug fusionfs_debug.go
