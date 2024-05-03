export GIT_COMMIT_SHA = $(shell git rev-parse HEAD)

dependencies:
	go mod tidy && \
	go mod download

test:
	go test ./... -covermode=count -v -coverprofile cp.out && \
	go tool cover -html=cp.out -o cp.html && \
	go tool cover -func=cp.out

swagger:
	rm -f .docs/swagger/swagger.json && \
	go generate && \
	swagger validate .docs/swagger/swagger.json

swagger-view:
	swagger serve .docs/swagger/swagger.json

set-version:
	./.helm/set-version.sh

build: dependencies
	rm -rf .bin && \
	cd cmd && \
	go build -o ../.bin/skatepark-api-go

run: 
	.bin/skatepark-api-go

docker: 
	docker build ./ -t skatepark-api:latest --no-cache

publish:
	docker tag derekpedersen/skatepark-api:${GIT_COMMIT_SHA}
	docker push derekpedersen/skatepark-api:${GIT_COMMIT_SHA}

deploy: set-version
	helm upgrade skatepark-api .helm --install

secret:
	kubectl apply -f .kubernetes/secret.yaml

kubernetes: build test docker publish deploy