ARG IMAGE_UI
ARG IMAGE_SVC
FROM ${IMAGE_UI} as build-ui
FROM ${IMAGE_SVC} as build-svc

FROM ubuntu:20.04

RUN apt update \
 && apt install -y --no-install-recommends \
      ca-certificates \
      gettext-base \
      supervisor \
      nginx \
      curl \
      wget \
      vim \
    && apt clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENV HOME "/delivery"
WORKDIR "${HOME}"

COPY --from=build-ui "/build/version.txt" "${HOME}/version.txt"

ENV PATH="${PATH}:${HOME}/bin"
RUN mkdir -p "${HOME}/bin"

RUN wget https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh \
      && chmod +x "wait-for-it.sh"
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