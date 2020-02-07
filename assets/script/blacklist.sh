#!/bin/bash
set -o nounset
set -o xtrace

git grep -n "= proto.Marshal(" "**/*.go"
git grep -n "trace.StringAttribute.*string(" "**/*.go"
