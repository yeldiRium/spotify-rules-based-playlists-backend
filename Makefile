BINARY_NAME := "spotify-rules-based-playlists-backend"

qa: analyse test

analyse:
	@go vet ./...

test:
	@go test -cover ./...

coverage:
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/cover.out ./...
	@go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
	@open ./coverage/cover.html

clean:
	@rm -rf build/

build: qa clean
	$(eval VERSION=$(shell git tag --points-at HEAD))
	$(eval VERSION=$(or $(VERSION), (version unavailable)))

	@GOOS=linux GOARCH=amd64 go build -ldflags="-X 'github.com/thenativeweb/$(BINARY_NAME)/version.Version=$(VERSION)'" -o "./build/$(BINARY_NAME)-linux-amd64"

build-docker: build
	$(eval VERSION=$(shell git tag --points-at HEAD))
	$(eval IMAGE_VERSION=$(or $(VERSION), latest))
	$(eval VERSION=$(or $(VERSION), (version unavailable)))

	docker build --build-arg version="$(VERSION)" -t "thenativeweb/$(BINARY_NAME):latest" -t "thenativeweb/$(BINARY_NAME):$(IMAGE_VERSION)" .

.PHONY: analyse build build-docker clean coverage qa test
