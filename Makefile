export GIT_COMMIT_SHA = $(shell git rev-parse HEAD)

test:
	.tools/scripts/test.sh

swagger:
	# .tools/scripts/swagger.sh

swagger-view: swagger
	swagger serve .docs/swagger/swagger.json

build:
	.tools/scripts/build.sh

run: build
	./bin/skatepark-api-go

docker: build
	docker build ./ -t skatepark-api-go:latest --no-cache

publish:
	docker tag skatepark-api-go us.gcr.io/${GCLOUD_PROJECT_ID}/skatepark-api-go:${GIT_COMMIT_SHA}
	gcloud docker -- push us.gcr.io/${GCLOUD_PROJECT_ID}/skatepark-api-go:${GIT_COMMIT_SHA}

deploy:
	helm upgrade skatepark-api .helm

secret:
	kubectl create -f .kubernetes/secret.yaml

kubernetes: build test docker publish deploy