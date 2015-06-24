#!/bin/sh

APP=iroiro-api
docker rm -f $APP
docker pull tutum.co/funnythingz/$APP
docker run -dit --name $APP -p 80:8000 tutum.co/funnythingz/$APP /iroiro production

APP=iroiro-api-migration
docker rm -f $APP
docker pull tutum.co/funnythingz/$APP
docker run -dit --name $APP tutum.co/funnythingz/$APP /$APP production migrate
