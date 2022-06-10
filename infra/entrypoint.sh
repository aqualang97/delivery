#!/usr/bin/env bash
set -e

export VERSION="$(cat "${HOME}/version.txt")"
SUBST="\$HOME \$VERSION \$DELIVERY_API_ENDPOINT"

envsubst "${SUBST}" <"${HOME}/templates/nginx.conf" >"/etc/nginx/nginx.conf"
envsubst "${SUBST}" <"${HOME}/templates/supervisord.conf" >"${HOME}/supervisord.conf"
envsubst "${SUBST}" <"${HOME}/templates/endpoint.js" >"/usr/share/nginx/html/endpoint.js"

if [ $# -gt 0 ]; then
  exec "$@"
else
  ./wait-for-it.sh -t 15 "${DB_LISTEN}" -- echo "DB is up"
  /delivery/bin/delivery create
  supervisord --nodaemon --configuration "${HOME}/supervisord.conf"
fi
