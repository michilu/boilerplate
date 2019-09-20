#!/bin/bash
set -o nounset
set -o errexit
set -o xtrace

rocksdb=lib/rocksdb

GO111MODULE=on \
CGO_CFLAGS="-I$PWD/${rocksdb}/include" \
CGO_LDFLAGS="-L$PWD/${rocksdb} -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd -ldl" \
go get -v -u github.com/tecbot/gorocksdb
