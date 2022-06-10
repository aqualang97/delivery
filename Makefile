SHELL := /bin/bash

export GIT_HASH ?= $(shell git log -n 1 2> /dev/null | head -n 1 | sed -e 's/^commit //' | head -c 8)
export VERSION ?= v1-$(GIT_HASH)

IMAGE_LOCAL_NAME ?= delivery
IMAGE_NAME ?=

TAG_NAME ?= latest

TARGET_PLATFORM := linux
PLATFORMS ?= $(TARGET_PLATFORM)
BUILD_NAME := delivery
BUILD_DIR := ./bin
LDFLAGS = -ldflags "-X $(PACKAGE_NAME)/pkg/util.Version=$(VERSION)"
#GCFLAGS = -gcflags="all=-N -l"
GCFLAGS =

.PHONY: build
build: $(PLATFORMS)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	GO111MODULE=on GOOS=$@ GOARCH=amd64 CGO_ENABLED=0 go build ${LDFLAGS} ${GCFLAGS} \
	  -o "${BUILD_DIR}/${BUILD_NAME}-$@-amd64" main.go

.PHONY: start
start:
	echo starting... && \
	  npm run serve

.PHONY: docker-build
docker-build:
	# For future ci caching
	docker build \
	  -f "./svc.dockerfile" \
	  -t "${IMAGE_LOCAL_NAME}:${TAG_NAME}-build-svc" \
	  --build-arg VERSION="$(VERSION)" \
	  --cache-from "${IMAGE_LOCAL_NAME}:${TAG_NAME}-build-svc" \
	  .
	docker build \
	  -f "./ui.dockerfile" \
	  -t "${IMAGE_LOCAL_NAME}:${TAG_NAME}-build-ui" \
	  --build-arg VERSION="$(VERSION)" \
	  --cache-from "${IMAGE_LOCAL_NAME}:${TAG_NAME}-build-ui" \
	  .
	docker build \
	  -f "./Dockerfile" \
	  -t "${IMAGE_LOCAL_NAME}:${TAG_NAME}" \
	  --build-arg VERSION="$(VERSION)" \
	  --build-arg IMAGE_UI="${IMAGE_LOCAL_NAME}:${TAG_NAME}-build-ui" \
	  --build-arg IMAGE_SVC="${IMAGE_LOCAL_NAME}:${TAG_NAME}-build-svc" \
	  --cache-from "${IMAGE_LOCAL_NAME}:${TAG_NAME}" \
	  .

.PHONY: docker-run
docker-run: docker-build
docker-run:
	docker run -it --rm -p 8443:443 "${IMAGE_LOCAL_NAME}:${TAG_NAME}"

.PHONY: docker-rebuild
docker-rebuild:
  # For future ci caching
	docker build \
    -f "./Dockerfile" \
    -t "${IMAGE_LOCAL_NAME}-build-ui" \
    --target build-ui \
    --no-cache \
    .
	docker build \
    -f "./Dockerfile" \
    -t "${IMAGE_LOCAL_NAME}" \
    --no-cache \
    .

.PHONY: docker-deploy
docker-deploy:
	docker tag ${IMAGE_LOCAL_NAME} ${IMAGE_NAME}:latest
	docker push ${IMAGE_NAME}:latest