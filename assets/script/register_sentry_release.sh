#!/bin/sh
set -o nounset
set -o errexit
set -o xtrace

VERSION=$1
HASH=`git rev-parse HEAD`

curl -X POST \
  -H "Authorization: Bearer ${SENTRY_AUTH_TOKEN}" \
  -H "Content-Type: application/json" \
  -d "{\"version\":\"$VERSION\",\"ref\":\"$HASH\",\"projects\":[\"$SENTRY_PROJECT\"]}" \
  https://sentry.io/api/0/organizations/${SENTRY_ORG}/releases/
