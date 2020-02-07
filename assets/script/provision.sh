#!/bin/bash
set -o nounset
set -o errexit
#set -o xtrace

ip=$(/sbin/ifconfig|grep "inet "|grep -v 127.0.0.1|grep -v " 169.254."|awk '{print $2}')

( type jq > /dev/null 2>&1 ) || (
  echo "require jq. https://stedolan.github.io/jq/download/"
  exit 1
  )

[ -r assets/.ssh/id_rsa ]\
  || (if [ -f .ansiblevaultkey ]; then
    pipenv run ansible-vault decrypt\
    --vault-password-file=.ansiblevaultkey\
    --output=assets/.ssh/id_rsa\
    assets/.ssh/id_rsa.encrypted
  else
    pipenv run ansible-vault decrypt\
    --output=assets/.ssh/id_rsa\
    assets/.ssh/id_rsa.encrypted
  fi;)

service=$(basename $(pwd))
package_name="$service-*.linux-arm.tar.gz"
package=$(ls -1t $package_name|head -n 1)
[ -f "$package" ] || (
  echo "must be download '$package_name' from https://github.com/$(basename $(dirname $(pwd)))/$service/branches"
  exit 1
  )

name="$(date -u '+%Y%m%d-%H%M%S')""-""$(cat /dev/urandom | base64 | tr -d "+/" | fold -w 6 | head -n 1)"
port="5$(cat /dev/urandom | base64 | tr -d "+/[A-z]" | fold -w 4 |head -n 1)"
sed -i '' "s|^REMOTE_PORT=.*|REMOTE_PORT=$port|" assets/debian/home/root/debug-port.env

tar --file=$name.tar --create\
 config.toml\
 ;
tar --file=$name.tar --append --directory=assets\
 .ssh/id_rsa\
 ;
tar --file=$name.tar --append --directory=assets/debian/home/root\
 debug-port.env\
 ;
tar --file=$name.tar --list --verbose
gzip $name.tar

rm -rf provision
mkdir provision
cp -a\
  $name.tar.gz\
  $package\
  assets/script/create_user.sh\
  assets/script/install.sh\
  provision
cmd="curl -sf http://$ip:8000/install.sh | sh /dev/stdin $service $ip:8000 $package $name"
( type pbcopy > /dev/null 2>&1 ) && echo $cmd|pbcopy
echo
echo $cmd
echo
echo "curl -sf http://$ip:8000/create_user.sh | sh /dev/stdin $(whoami)"
echo
(cd provision && python3 -m http.server || python2 -m SimpleHTTPServer)
