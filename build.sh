#!/bin/sh

# iroiro-api
app=iroiro
repo=tutum.co/funnythingz/$app-api

GOOS=linux gom build -o $app
tar cvfz $app.tar.gz ./$app ./db/database.toml
cat $app.tar.gz | docker import - $repo
docker push $repo
rm ./$app ./$app.tar.gz
echo done

# migration
GOOS=linux gom build migrate/migration.go
tar cvfz migration.tar.gz ./migration ./db/database.toml
rm ./migration ./migration.tar.gz
echo migration done
