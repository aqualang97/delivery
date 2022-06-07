

FROM ubuntu:20.04

ARG VERSION

RUN apt update \
 && apt install -y --no-install-recommends \
      ca-certificates \
      openssl \
      python \
      build-essential \
      curl \
      wget \
 && curl -fsSL https://deb.nodesource.com/setup_14.x | bash - \
 && apt-get install -y nodejs \
 && apt clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN npm install -g npm

WORKDIR "/build"

RUN mkdir -p "./certificates" \
 && openssl req \
      -x509 \
      -newkey rsa:2048 \
      -keyout "./certificates/key.pem" \
      -out "./certificates/cert.pem" \
      -days 3650 \
      -nodes \
      -subj "/CN=localhost" \
 && openssl dhparam -out "./dhparam.pem" 2048

RUN echo "${VERSION}" > "./version.txt"

COPY web/package.json ./
COPY web/package-lock.json ./

RUN npm install

COPY web/fonts ./fonts
COPY web/pic ./pic
COPY web/public ./public
COPY web/src ./src

RUN VERSION="${VERSION}" npm run build