FROM ubuntu:20.04 as build-ui

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

ARG VERSION
RUN echo "${VERSION}" > "./version.txt"

COPY web/package.json ./
COPY web/package-lock.json ./

RUN npm install

COPY web/fonts ./fonts
COPY web/pic ./pic
COPY web/public ./public
COPY web/src ./src

RUN VERSION="${VERSION}" npm run build

FROM ubuntu:20.04 as build-svc

RUN apt update \
 && apt install -y --no-install-recommends \
      ca-certificates \
      openssl \
      build-essential \
      curl \
      git \
      wget \
 && apt clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ARG GO_VERSION=1.17
RUN export FILENAME="go${GO_VERSION}.linux-amd64.tar.gz" \
 && export URL="https://golang.org/dl/${FILENAME}" \
 && wget "${URL}" -O "${FILENAME}" \
 && tar xzvf "${FILENAME}" -C "/usr/local"
ENV PATH "${PATH}:/usr/local/go/bin"

WORKDIR "/build"

COPY ./go.* ./
RUN go mod download

COPY ./configs ./configs
COPY ./development ./development
COPY ./internal ./internal
COPY ./logs ./logs
COPY ./tests ./tests
COPY ./main.go ./
COPY ./Makefile ./

ARG VERSION
RUN VERSION=$VERSION make build

ARG VERSION
RUN echo "${VERSION}" > "./version.txt"

FROM ubuntu:20.04 as runtime

RUN apt update \
 && apt install -y --no-install-recommends \
      ca-certificates \
      gettext-base \
      supervisor \
      nginx \
      curl \
      vim \
 && apt clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENV HOME "/delivery"
WORKDIR "${HOME}"

COPY --from=build-ui "/build/version.txt" "${HOME}/version.txt"

ENV PATH="${PATH}:${HOME}/bin"
RUN mkdir -p "${HOME}/bin"

COPY ./infra/*.sh "${HOME}/bin"

COPY --from=build-ui "/build/certificates" "/etc/nginx/certificates/"
COPY --from=build-ui "/build/dhparam.pem" "/etc/nginx/dhparam.pem"
COPY "./infra/templates/" "${HOME}/templates"
RUN groupadd "nginx" && useradd -g "nginx" -d "/usr/share/nginx/html" "nginx" \
 && rm -f "/usr/share/nginx/html/index.html" \
 && echo "pong" > "/usr/share/nginx/html/ping.txt" \
 && chown -Rv nginx:nginx "/usr/share/nginx/html"
EXPOSE 80 443

COPY --from=build-ui "/build/dist" "/usr/share/nginx/html"

COPY --from=build-svc "/build/bin/delivery-linux-amd64" "${HOME}/bin/delivery"
RUN chmod +x "${HOME}/bin/delivery"

HEALTHCHECK --interval=60s --timeout=5s --start-period=30s --retries=3 \
  CMD curl -fk "https://localhost/ping.txt" || exit 1

ENTRYPOINT ["/delivery/bin/entrypoint.sh"]