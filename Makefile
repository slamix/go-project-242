build:
	go build cmd/hexlet-path-size/main.go

lint:
	golangci-lint run .