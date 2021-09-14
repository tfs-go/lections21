LINT_VERSION = v1.37.0
DOCKER_LINT_IMAGE = golangci/golangci-lint:$(LINT_VERSION)
DOCKER_LINT_FLAGS = --rm -v $(CURDIR):/app -w/app -e GOGC=30
DOCKER_LINT_PARAMS = golangci-lint run -v -c .golangci.yml ./...

lint:
	docker run $(DOCKER_LINT_FLAGS) $(DOCKER_LINT_IMAGE) $(DOCKER_LINT_PARAMS)