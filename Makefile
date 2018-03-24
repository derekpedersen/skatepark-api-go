main:
	dep ensure
	go build -o bin/skatepark-api-go

test:
	go test ./... -v