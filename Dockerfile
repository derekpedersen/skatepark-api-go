# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM debian:stretch-slim

# Install wget and install/updates certificates
RUN apt-get update \
 && apt-get install -y -q --no-install-recommends \
    ca-certificates \
 && apt-get clean \
 && rm -r /var/lib/apt/lists/*

# Copy the local package files to the container's workspace.
COPY ./.bin/skatepark-api-go /go/bin/

# Copy the json repository files
COPY ./.db/json/ /repository/json/

# Copy the swagger files
COPY ./.docs/swagger/* ./.docs/swagger/

# Run the command by default when the container starts.
CMD /go/bin/skatepark-api-go

# Document that the service listens on port 8080.
EXPOSE 8080 3000