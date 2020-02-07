#!/bin/bash
set -o errexit
set -o xtrace

name=$4
set -o nounset

service=$1
host=$2
package=$3

mkdir -p /home/root
[ -w /home/root ]\
  || (echo Permission denied ; exit 1)

cd /home/root

curl -sf "http://$host/$package"\
  |tar xzvf -

[ -n "$name" ] &&\
  curl -sf "http://$host/$name.tar.gz"\
  |tar xzvf -

cp -a etc/systemd/system/debug-port.service /etc/systemd/system/.
cp -a etc/systemd/system/$service.service /etc/systemd/system/.

systemctl daemon-reload
systemctl enable debug-port
systemctl enable $service
systemctl restart $service

( type tree > /dev/null 2>&1 ) && tree
./$service version

systemctl restart debug-port
