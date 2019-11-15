#!/bin/bash

# clean out existing mocks
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
