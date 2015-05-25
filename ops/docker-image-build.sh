#!/bin/sh

cat share/iroiro.tar.gz | docker import - tutum.co/funnythingz/iroiro
docker push tutum.co/funnythingz/iroiro

cat share/migration.tar.gz | docker import - tutum.co/funnythingz/iroiro-migration
docker push tutum.co/funnythingz/iroiro-migration
