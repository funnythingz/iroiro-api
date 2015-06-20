#!/bin/sh

# API
APP=iroiro-api
docker rm -f $APP
docker pull tutum.co/funnythingz/$APP
docker run -dit --name iroiro -p 80:8000 tutum.co/funnythingz/$APP /iroiro production

# migration
docker rm -f iroiro-migration
docker pull tutum.co/funnythingz/iroiro-migration
docker run -dit --name iroiro-migration tutum.co/funnythingz/iroiro-migration /migration production migrate
