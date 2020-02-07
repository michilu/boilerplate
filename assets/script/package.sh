#!/bin/bash
set -o nounset
set -o errexit
set -o xtrace

[ -d package ] && rm -rf package

name=$(basename $PWD)
mkdir -p package/$name
base=package/$name/$name
tar --file=$base.tar --create\
 assets/daemon/loop\
 assets/script/debug-port.sh\
 ;
tar --file=$base.tar --append --directory=assets/debian/home/root\
 debug-port.env\
 ;
tar --file=$base.tar --list --verbose
[ "$(tar --file=$base.tar --list|sort)" == "$(cat assets/script/package.txt)" ]
for i in $(find assets/gox -type f -name "linux-*" -or -type f -name "darwin-*"); do
 suffix=$(echo $i|sed -e "s/assets\/gox\///" -e "s/\//./g")
 target=package/$name-$suffix
 cp $i $base
 cp $base.tar $target.tar
 tar --file=$target.tar --append --directory=package/$name\
  $name\
  ;
 echo -e "\n"$target.tar
 tar --file=$target.tar --list --verbose
 gzip $target.tar
done

rm -rf package/$name
