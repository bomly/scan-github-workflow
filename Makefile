GITHUB_REPOSITORY ?= scan-gomod
GITHUB_SHA ?= $(shell git rev-parse --short HEAD)
IMAGE ?= ${GITHUB_REPOSITORY}:${GITHUB_SHA}

test:
	go test ./internal/... -count 1 -race

test-coverage:
	go test -coverprofile=coverage.txt ./internal/... -count 1 -race

vet:
	go vet ./...

cover:
	go test ./... -count 1 -race -cover

gen:
	go generate ./...

plantuml-docker:
	docker run -v $(shell pwd)/docs:/docs -w /docs ghcr.io/plantuml/plantuml *.pu

plantuml:
	plantuml docs/*.pu

build:
	docker buildx build -t ${IMAGE} .
