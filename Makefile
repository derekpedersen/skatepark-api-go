run:
	dep ensure
	go build -o bin/skatepark-api-go
	./bin/skatepark-api-go

go-build:
	dep ensure
	go build -o bin/skatepark-api-go

test:
	go test ./... -v

build:
	dep ensure
	go build -o bin/skatepark-api-go
	docker build ./ -t skatepark-api-go

publish:
	docker tag skatepark-api-go us.gcr.io/${PROJECT_ID}/skatepark-api-go:latest
	gcloud docker -- push us.gcr.io/${PROJECT_ID}/skatepark-api-go:latest

deploy:
	sed -i -e 's/%PROJECT_ID%/${PROJECT_ID}/g' ./kubernetes-deployment.yaml
	kubectl delete deployment skatepark-api-go-deployment
	kubectl create -f ./kubernetes-deployment.yaml
	kubectl apply -f ./kubernetes-service.yaml

kubernetes: build test publish deploy