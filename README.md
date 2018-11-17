# Skatepark API - GO #

[![Build Status](https://jenkins.derekpedersen.com/buildStatus/icon?job=derekpedersen/skatepark-api-go/master&style=plastic&.png)](https://jenkins.derekpedersen.com/job/derekpedersen/job/skatepark-api-go/job/master/)
[![Coverage Status](https://coveralls.io/repos/github/derekpedersen/skatepark-api-go/badge.png?branch=master)](https://coveralls.io/github/derekpedersen/skatepark-api-go)

My skatepark api written in golang. It exposes all of the skateparks that I have been too.

## Dependency Management ##

- https://github.com/golang/dep
- go get -u github.com/golang/dep/cmd/dep
- dep init
- dep ensure
- dep ensure -add github.com/foo/bar

## Logging ##

- https://github.com/sirupsen/logrus
- dep ensure -add github.com/sirupsen/logrus

## Testing ##

### gomock ###

https://github.com/golang/mock

```bash
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

## Docker ##

The API runs within a docker container.

### Build ###

- docker build ./ -t skatepark-api-go

### Execute ###

- docker run -d --rm -it -p 8080:8080 --name=skatepark-api-go-container -t skatepark-api-go

### GCR ###

- docker tag api-go us.gcr.io/${PROJECT_ID}/skatepark-api-go:latest
- gcloud docker -- push us.gcr.io/${PROJECT_ID}/skatepark-api-go:latest