test:
	go test -v

build:
	CGO_ENABLED=0 go build

lint:
	go vet
