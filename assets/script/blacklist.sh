#!/bin/bash
set -o nounset
set -o xtrace

git grep -n 'trace.StringAttribute(.*, string(' "*.go" "**/*.go"
git grep -n 'trace.StringAttribute(.*, fmt.Sprintf(".*%s.*"' "*.go" "**/*.go"
git grep -n -P '^\t+go (?!slog\.Recover\()' "*.go" "**/*.go"
git grep -n -P '^\t+(_, )*(err|ok) := .*\)$' "*.go" "**/*.go"
git grep -l '<-ctx.Done():' "*.go" "**/*.go"|xargs pcregrep -n -M '<-.*\.?ctx\.Done\(\):\n\t+(?!.*err :=.* ([^\s]*\.)?ctx\.Err\()'
