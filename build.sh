#!/bin/sh
GOOS=linux gom build -o iroiro
tar cvfz iroiro.tar.gz ./iroiro
mv ./iroiro.tar.gz ~/vagrant/coreos-vagrant/
rm ./iroiro
echo done
