#!/bin/bash
set -ev

# install chaincode for channelsales1
docker exec cli peer chaincode install -n music-cc -v 1.0 -p chaincode/go
sleep 1
# instantiate chaincode for channelsales1
docker exec cli peer chaincode instantiate -o orderer.acornpub.com:7050 -C channelsales1 -n music-cc -v 1.0 -c '{"Args":[""]}' -P "OR ('Sales1Org.member','CustomerOrg.member')"
sleep 10
# invoke chaincode for channelsales1
docker exec cli peer chaincode invoke -o orderer.acornpub.com:7050 -C channelsales1 -n music-cc -c '{"function":"initWallet","Args":[""]}'
docker exec cli peer chaincode invoke -o orderer.acornpub.com:7050 -C channelsales1 -n music-cc -c '{"function":"setMusic","Args":["Fabric", "Hyper", "20", "1Q2W3E4R"]}'
sleep 3
# query chaincode for channelsales1
docker exec cli peer chaincode query -o orderer.acornpub.com:7050 -C channelsales1 -n music-cc -c '{"function":"getAllMusic","Args":[""]}'