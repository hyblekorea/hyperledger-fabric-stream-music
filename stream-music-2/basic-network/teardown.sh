#!/bin/bash
set -ev

# Shut down the Docker containers for the system tests.
docker-compose -f docker-compose-ca.yaml stop
docker-compose stop
docker-compose -f docker-compose-ca.yaml kill && docker-compose -f docker-compose-ca.yaml down --volumes --remove-orphans
docker-compose -f kill && docker-compose down --volumes --remove-orphans

# remove sdk wallet
rm -rf $GOPATH/src/stream-music-2/application/wallet/

# Your system is now clean
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)

# remove chaincode docker images
docker rmi -f $(docker images dev-* -q)