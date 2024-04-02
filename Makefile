export GIT_COMMIT_SHA = $(shell git rev-parse HEAD)

dependencies:
	go mod tidy && \
	go mod download

mocks:
	.tools/scripts/mocks.sh

test: mocks
	go test ./... -covermode=count -v -coverprofile cp.out && \
	go tool cover -html=cp.out -o cp.html && \
	go tool cover -func=cp.out

swagger:
	rm -f .docs/swagger/swagger.json && \
	go generate && \
	swagger validate .docs/swagger/swagger.json

swagger-view: swagger
	swagger serve .docs/swagger/swagger.json

set-version:
	./.helm/set-version.sh

build: dependencies
	rm -rf .bin && \
	cd cmd && \
	go build -o ../.bin/skatepark-api-go

run: build
	.bin/skatepark-api-go

docker: build
	docker build ./ -t skatepark-api-go:latest --no-cache

publish:
	docker tag skatepark-api-go us.gcr.io/sleipnir/skatepark-api-go:${GIT_COMMIT_SHA}
	gcloud docker -- push us.gcr.io/sleipnir/skatepark-api-go:${GIT_COMMIT_SHA}

deploy: set-version
	helm upgrade skatepark-api .helm

secret:
	kubectl create -f .kubernetes/secret.yaml

kubernetes: build test docker publish deploy