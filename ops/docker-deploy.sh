#!/bin/sh

docker pull tutum.co/funnythingz/iroiro
docker pull tutum.co/funnythingz/iroiro-migration
docker run -dit -name iroiro-migration tutum.co/funnythingz/iroiro-migration /migration production migrate
docker run -dit -name iroiro -p 80:8000 tutum.co/funnythingz/iroiro /iroiro production
