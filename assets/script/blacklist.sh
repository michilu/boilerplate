#!/bin/bash
set -o nounset
set -o xtrace

git grep -n 'trace.StringAttribute(.*, string(' "**/*.go"
git grep -n 'trace.StringAttribute(.*, fmt.Sprintf(".*%s.*"' "**/*.go"
git grep -n -P '^\t+go (?!slog\.Recover\()' "**/*.go"
git grep -n -P '^\t+(_, )*err := .*\)$' "**/*.go"
