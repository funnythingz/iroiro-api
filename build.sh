#!/bin/sh
GOOS=linux gom build -o iroiro
tar cvfz iroiro.tar.gz ./iroiro ./database.toml
mv ./iroiro.tar.gz ~/vagrant/iroiro/
rm ./iroiro
echo done
