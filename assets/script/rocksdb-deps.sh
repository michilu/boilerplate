#!/bin/bash
set -o nounset
set -o errexit
set -o xtrace

apt-get update && apt-get install --no-install-recommends -y\
 libbz2-dev\
 liblz4-dev\
 libsnappy-dev\
 libzstd-dev\
 zlib1g-dev\
;
