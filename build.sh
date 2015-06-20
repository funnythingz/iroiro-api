#!/bin/sh

# iroiro-api
APP=iroiro
REPO=tutum.co/funnythingz/$APP-api

GOOS=linux gom build -o $APP
tar cvfz $APP.tar.gz ./$APP ./db/database.toml ./config/config.toml
cat $APP.tar.gz | docker import - $REPO
docker push $REPO
rm ./$APP ./$APP.tar.gz
echo done

# migration
APP=iroiro-api-migration
REPO=tutum.co/funnythingz/$APP
GOOS=linux gom build -o $APP migrate/migration.go
tar cvfz $APP.tar.gz ./$APP ./db/database.toml
cat $APP.tar.gz | docker import - $REPO
docker push $REPO
rm ./$APP ./$APP.tar.gz
echo migration done
