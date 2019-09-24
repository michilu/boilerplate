#!/bin/bash
set -o nounset
set -o errexit
set -o xtrace

assetslib=assets/lib
rocksdb=${assetslib}/rocksdb
version=$1

mkdir -p ${rocksdb}
curl -sL "https://github.com/facebook/rocksdb/archive/v${version}.tar.gz" \
| tar zx --strip=1 -C ${rocksdb}
( cd ${rocksdb} && make shared_lib )
cp -pd ${rocksdb}/librocksdb\.* ${assetslib}
