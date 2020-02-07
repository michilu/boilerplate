#!/bin/bash
set -o errexit
set -o xtrace
set -o nounset

USER=$1
id $USER\
  && userdel -r $USER
useradd -m -s /bin/bash -g operator --password $USER $USER

mkdir -p /home/${USER}/.ssh
cd /home/${USER}

touch .ssh/authorized_keys
curl https://github.com/${USER}.keys > .ssh/authorized_keys

chmod 700 .ssh
chmod 600 .ssh/authorized_keys
chown -R ${USER}:operator .ssh
