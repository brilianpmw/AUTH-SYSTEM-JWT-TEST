run-http: 
	go run cmd/app.go

build:
	go build -v -o cmd/app.go

test:
	@go test -v -race ./...
