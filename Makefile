export GIT_COMMIT_SHA = $(shell git rev-parse HEAD)

test:
	.tools/scripts/test.sh

swagger:
	.tools/scripts/swagger.sh

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
	sed -e 's/%GCLOUD_PROJECT_ID%/${GCLOUD_PROJECT_ID}/g' -e 's/%GIT_COMMIT_SHA%/${GIT_COMMIT_SHA}/g' .kubernetes/deployment.yaml > deployment.sed.yaml
	kubectl apply -f ./deployment.sed.yaml
	kubectl apply -f .kubernetes/service.yaml

secret:
	kubectl create -f .kubernetes/secret.yaml

kubernetes: build test docker publish