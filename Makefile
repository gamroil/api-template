BINARY_NAME = api-template
IMAGE_NAME = api-template

build:
	@go build -o $(BINARY_NAME) -ldflags="-X main.Commit=$(git rev-parse HEAD)" .

build-image:
	@docker build -t $(IMAGE_NAME) .

run-image:
	@docker run -p 8080:8080 $(IMAGE_NAME)

lint:
	@golangci-lint run --config golangci.yaml

lint-fix:
	@golangci-lint run --config golangci.yaml --fix
