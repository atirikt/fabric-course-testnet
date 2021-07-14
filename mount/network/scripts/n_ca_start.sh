#!/bin/bash
cd ~/mount/network/docker

docker-compose -f docker-compose-test-net.yaml -f docker-compose-ca.yaml -f docker-compose-couch.yaml down
docker volume prune -f

sudo rm -fr ~/mount/network/organizations/ordererOrganizations/*
sudo rm -fr ~/mount/network/organizations/peerOrganizations/*
sudo rm -fr ~/mount/network/system-genesis-block/*

docker-compose -f docker-compose-ca.yaml up -d
sleep 10

cd ..
./organizations/fabric-ca/registerEnroll.sh

export FABRIC_CFG_PATH=$PWD/configtx/
configtxgen -profile TwoOrgsApplicationGenesis -channelID system-channel -outputBlock ./system-genesis-block/genesis.block 

