#!/bin/bash
set -o nounset
set -o xtrace

git grep -n "trace.StringAttribute(.*, string(" "**/*.go"
git grep -n 'trace.StringAttribute(.*, fmt.Sprintf(".*%s.*"' "**/*.go"
