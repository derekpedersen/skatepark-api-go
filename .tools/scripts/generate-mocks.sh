
# clean out existing mocks
rm -rf mock

# repository mocks
mockgen -source=repository/skateparks.go -destination=mock/mock_skateparks_repository.go -package=mock

# controller mocks
mockgen -source=controller/skateparks.go -destination=mock/mock_skateparks_controller.go -package=mock

# service mocks
mockgen -source=service/skateparks.go -destination=mock/mock_skateparks_service.go -package=mock
mockgen -source=vendor/github.com/derekpedersen/imgur-go/service/album.go -destination=mock/mock_album_service.go -package=mock