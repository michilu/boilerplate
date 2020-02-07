#!/bin/bash
set -o nounset
set -o errexit
#set -o xtrace

ip=$(/sbin/ifconfig|grep "inet "|grep -v 127.0.0.1|grep -v " 169.254."|awk '{print $2}')

( type jq > /dev/null 2>&1 ) || (
  echo "require jq. https://stedolan.github.io/jq/download/"
  exit 1
  )

service=$(basename $(pwd))
package_name="$service-*.linux-arm.tar.gz"
package=$(ls -1t $package_name|head -n 1)
[ -f "$package" ] || (
  echo "must be download '$package_name' from https://github.com/$(basename $(dirname $(pwd)))/$service/branches"
  exit 1
  )

rm -rf provision
mkdir provision
cp -a\
  $package\
  assets/script/create_user.sh\
  assets/script/install.sh\
  provision
cmd="curl -sf http://$ip:8000/install.sh | sh /dev/stdin $service $ip:8000 $package"
( type pbcopy > /dev/null 2>&1 ) && echo $cmd|pbcopy
echo
echo $cmd
echo
(cd provision && python3 -m http.server || python2 -m SimpleHTTPServer)
