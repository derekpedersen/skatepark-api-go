export GIT_COMMIT_SHA = $(shell git rev-parse HEAD)

mocks:
	rm -fr mock
	mkdir mock
	# repositories
	mockgen -source=repository/skateparks.go -destination=mock/mock_skateparks_repository.go -package=mock
	# services
	mockgen -source=service/skateparks.go -destination=mock/mock_skateparks_service.go -package=mock
	mockgen -source=service/health.go -destination=mock/mock_health_service.go -package=mock
	# controllres
	mockgen -source=controller/skateparks.go -destination=mock/mock_skateparks_controller.go -package=mock
	mockgen -source=controller/health.go -destination=mock/mock_health_controller.go -package=mock
	# external deps
	mockgen -source=vendor/github.com/derekpedersen/imgur-go/service/album.go -destination=mock/mock_album_service.go -package=mock

test: mocks
	go test ./... -covermode=count -v -coverprofile cp.out
	go get github.com/t-yuki/gocover-cobertura
	go tool cover -html=cp.out -o cp.html && gocover-cobertura < cp.out > cp.xml
	go tool cover -func=cp.out

swagger:
	rm -f .docs/swagger/swagger.json
	go generate
	swagger validate .docs/swagger/swagger.json

swagger-view: swagger
	swagger serve .docs/swagger/swagger.json

build:
	rm -rf bin
	dep ensure
	go build -o bin/skatepark-api-go

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

kubernetes: build test docker publish deploy