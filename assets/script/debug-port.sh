#!/bin/bash
set -o errexit
set -o xtrace
set -o nounset

# Wait for the connecting to remote host
until ping -c 1 -W 1 $REMOTE_HOST;do sleep 1s;done

ssh -y -i ~/.ssh/id_rsa -R $REMOTE_PORT:localhost:22 $REMOTE_USER@$REMOTE_HOST 'sleep 1h'
