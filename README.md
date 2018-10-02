# Skatepark API - GO #

[![Build Status](http://jenkins.derekpedersen.com/buildStatus/icon?job=the-cluster/skatepark-api-go)](http://jenkins.derekpedersen.com/job/the-cluster/job/skatepark-api-go)

My skatepark api written in golang. It exposes all of the skateparks that I have been too.

## Dependency Management ##

- https://github.com/golang/dep
- go get -u github.com/golang/dep/cmd/dep
- dep init
- dep ensure
- dep ensure -add github.com/foo/bar

## Logging ##

- https://github.com/jeanphorn/log4go
- dep ensure -add github.com/jeanphorn/log4go

## Docker ##

The API runs within a docker container.

### Build ###

- docker build ./ -t skatepark-api-go

### Execute ###

- docker run -d --rm -it -p 8080:8080 --name=skatepark-api-go-container -t skatepark-api-go

### GCR ###

- docker tag api-go us.gcr.io/${PROJECT_ID}/skatepark-api-go:latest
- gcloud docker -- push us.gcr.io/${PROJECT_ID}/skatepark-api-go:latest