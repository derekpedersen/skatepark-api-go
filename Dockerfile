# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/derekpedersen/skatepark-api-go
RUN ls -l /go/src/github.com/derekpedersen/skatepark-api-go

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get -u github.com/golang/dep/cmd/dep
RUN cd /go/src/github.com/derekpedersen/skatepark-api-go && make build

# Run the command by default when the container starts.
RUN ls -l /go/src/github.com/derekpedersen/skatepark-api-go/bin
ENTRYPOINT cd /go/src/github.com/derekpedersen/skatepark-api-go && ./bin/skatepark-api-go

# Document that the service listens on port 8080.
EXPOSE 8080
EXPOSE 8000