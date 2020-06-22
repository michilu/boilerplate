#!/bin/bash
set -o errexit
set -o xtrace
set -o nounset

target='assets/gcp/bq'
[ -d ${target} ] || mkdir -p ${target}
target_json="${target}/transfer_config.json"
bq --format=prettyjson ls   --transfer_config --transfer_location='us'\
  |jq '.|sort_by(.name)'\
  > ${target_json}
[ "$(jq -r '.[].state' < ${target_json})" = "SUCCEEDED" ]
