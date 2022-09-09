export GIT_COMMIT_SHA = $(shell git rev-parse HEAD)

mocks:
	.tools/scripts/mocks.sh

test: mocks
	.tools/scripts/test.sh

swagger:
	.tools/scripts/swagger.sh

swagger-view: swagger
	swagger serve .docs/swagger/swagger.json

set-version:
	./.helm/set-version.sh

build:
	.tools/scripts/build.sh

run: build
	./bin/skatepark-api-go

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