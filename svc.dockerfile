FROM ubuntu:20.04

ARG VERSION
ARG GO_VERSION=1.18.3

RUN apt update \
 && apt install -y --no-install-recommends \
      ca-certificates \
      openssl \
      build-essential \
      curl \
      git \
      wget \
 && apt clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN export FILENAME="go${GO_VERSION}.linux-amd64.tar.gz" \
 && export URL="https://golang.org/dl/${FILENAME}" \
 && wget "${URL}" -O "${FILENAME}" \
 && tar xzvf "${FILENAME}" -C "/usr/local"
ENV PATH "${PATH}:/usr/local/go/bin"

WORKDIR "/build"

COPY ./go.* ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./vendor ./vendor
COPY ./configs ./configs
COPY ./development ./development
COPY ./internal ./internal
COPY ./logs ./logs
COPY ./tests ./tests
COPY ./main.go ./
COPY ./Makefile ./

RUN VERSION=$VERSION make build

RUN echo "${VERSION}" > "./version.txt"