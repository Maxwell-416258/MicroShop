

.PHONY: proto
proto:
	protoc -I ./ --go_out=./ --micro_out=./ ./*.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o product-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t product-service:latest
