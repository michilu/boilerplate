#!/bin/bash
set -o nounset
set -o errexit
set -o xtrace

(cd $1
 ls -tr1 | grep -v -e "\.json$" | tail -n +$2 | xargs -n1 rm -rf
)
