export GIT_COMMIT_SHA = $(shell git rev-parse HEAD)

mocks:
	rm -fr mock
	mkdir mock
	mockgen -source=repository/skateparks.go -destination=mock/mock_skateparks_repository.go -package=mock
	mockgen -source=service/skateparks.go -destination=mock/mock_skateparks_service.go -package=mock
	mockgen -source=vendor/github.com/derekpedersen/imgur-go/service/album.go -destination=mock/mock_album_service.go -package=mock

test: mocks
	go test ./... -v -coverprofile cp.out
	go tool cover -html=cp.out

build:
	dep ensure
	go build -o bin/skatepark-api-go

run: build
	./bin/skatepark-api-go

docker:
	docker build ./ -t skatepark-api-go

publish:
	docker tag skatepark-api-go us.gcr.io/${GCLOUD_PROJECT_ID}/skatepark-api-go:${GIT_COMMIT_SHA}
	gcloud docker -- push us.gcr.io/${GCLOUD_PROJECT_ID}/skatepark-api-go:${GIT_COMMIT_SHA}

deploy:
	sed -e 's/%GCLOUD_PROJECT_ID%/${GCLOUD_PROJECT_ID}/g' -e 's/%GIT_COMMIT_SHA%/${GIT_COMMIT_SHA}/g' ./kubernetes-deployment.yaml > deployment.sed.yaml
	kubectl apply -f ./deployment.sed.yaml
	kubectl apply -f ./kubernetes-service.yaml

secret:
	kubectl create -f ./kubernetes-secret.yaml

kubernetes: build test docker publish deploy