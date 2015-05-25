#!/bin/sh

# iroiro-api
GOOS=linux gom build -o iroiro
tar cvfz iroiro.tar.gz ./iroiro ./db/database.toml
mv ./iroiro.tar.gz ~/vagrant/iroiro/
rm ./iroiro
echo done

# migration
GOOS=linux gom build migrate/migration.go
tar cvfz migration.tar.gz ./migration ./db/database.toml
mv ./migration.tar.gz ~/vagrant/iroiro/
rm ./migration
echo migration done
