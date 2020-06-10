#!/bin/bash

echo "Generate genesis.block"

FABRIC_CFG_PATH=$PWD ./bin/configtxgen -profile DataTransaction -outputBlock ./artifacts/orderer.genesis.block

echo "Generate channel.tx"

FABRIC_CFG_PATH=$PWD ./bin/configtxgen -profile DataTransaction -outputCreateChannelTx ./artifacts/datatransaction.channel.tx -channelID datatransaction

echo "Generate org1.datatransaction.anchors.tx"

FABRIC_CFG_PATH=$PWD ./bin/configtxgen -profile DataTransaction -outputAnchorPeersUpdate ./artifacts/org1.datatransaction.anchors.tx -channelID datatransaction -asOrg Org1
