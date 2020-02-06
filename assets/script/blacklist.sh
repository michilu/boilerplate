#!/bin/bash
set -o nounset
set -o errexit
set -o xtrace

git grep -n "= proto.Marshal(" "**/*.go"
