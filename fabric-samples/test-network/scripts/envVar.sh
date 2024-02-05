#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

# imports
. scripts/utils.sh

export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export PEER0_ORG1_CA=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export PEER1_ORG1_CA=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt
export PEER2_ORG1_CA=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer2.org1.example.com/tls/ca.crt
export PEER3_ORG1_CA=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer3.org1.example.com/tls/ca.crt

export PEER0_ORG2_CA=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export PEER1_ORG2_CA=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer1.org2.example.com/tls/ca.crt
export PEER2_ORG2_CA=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer2.org2.example.com/tls/ca.crt
export PEER3_ORG2_CA=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer3.org2.example.com/tls/ca.crt

export PEER0_ORG3_CA=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
export PEER1_ORG3_CA=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer1.org3.example.com/tls/ca.crt
export PEER2_ORG3_CA=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer2.org3.example.com/tls/ca.crt
export PEER3_ORG3_CA=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer3.org3.example.com/tls/ca.crt

export PEER0_ORG4_CA=${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt
export PEER1_ORG4_CA=${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer1.org4.example.com/tls/ca.crt
export PEER2_ORG4_CA=${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer2.org4.example.com/tls/ca.crt
export PEER3_ORG4_CA=${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer3.org4.example.com/tls/ca.crt

# Set environment variables for the peer org
setGlobals() {
  #DODATO - POCETAK
  local PEER_NUMBER=${2:-0}  # Ukoliko $2 nije poslato, postavi na 0
  #DODATO - KRAJ

  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  infoln "Using organization ${USING_ORG} peer ${PEER_NUMBER}"
  if [ $USING_ORG -eq 1 ]; then
    export CORE_PEER_LOCALMSPID="Org1MSP"
    #DODATO - POCETAK
    peer_ca_variable="PEER${PEER_NUMBER}_ORG1_CA"
    #DODATO - KRAJ
    export CORE_PEER_TLS_ROOTCERT_FILE="${!peer_ca_variable}" #izmenjeno
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    #export CORE_PEER_ADDRESS=localhost:7051 zamenjeno
    case $PEER_NUMBER in
        0) export CORE_PEER_ADDRESS="localhost:7051" ;;
        1) export CORE_PEER_ADDRESS="localhost:7052" ;;
        2) export CORE_PEER_ADDRESS="localhost:7053" ;;
        3) export CORE_PEER_ADDRESS="localhost:7055" ;;
        *) echo "Invalid PEER_NUMBER" ;;
    esac
  elif [ $USING_ORG -eq 2 ]; then
    export CORE_PEER_LOCALMSPID="Org2MSP"
    #DODATO - POCETAK
    peer_ca_variable="PEER${PEER_NUMBER}_ORG2_CA"
    #DODATO - KRAJ
    #export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_TLS_ROOTCERT_FILE="${!peer_ca_variable}" #zamenilo liniju iznad
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    #export CORE_PEER_ADDRESS=localhost:9051
    case $PEER_NUMBER in
        0) export CORE_PEER_ADDRESS="localhost:9051" ;;
        1) export CORE_PEER_ADDRESS="localhost:9052" ;;
        2) export CORE_PEER_ADDRESS="localhost:9053" ;;
        3) export CORE_PEER_ADDRESS="localhost:9055" ;;
        *) echo "Invalid PEER_NUMBER" ;;
    esac
  elif [ $USING_ORG -eq 3 ]; then
    export CORE_PEER_LOCALMSPID="Org3MSP"
    #DODATO - POCETAK
    peer_ca_variable="PEER${PEER_NUMBER}_ORG3_CA"
    #DODATO - KRAJ
    #export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
    export CORE_PEER_TLS_ROOTCERT_FILE="${!peer_ca_variable}" #zamenilo liniju iznad
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
    case $PEER_NUMBER in
        0) export CORE_PEER_ADDRESS="localhost:11051" ;;
        1) export CORE_PEER_ADDRESS="localhost:11052" ;;
        2) export CORE_PEER_ADDRESS="localhost:11053" ;;
        3) export CORE_PEER_ADDRESS="localhost:11055" ;;
        *) echo "Invalid PEER_NUMBER" ;;
    esac
  elif [ $USING_ORG -eq 4 ]; then
    export CORE_PEER_LOCALMSPID="Org4MSP"
    #DODATO - POCETAK
    peer_ca_variable="PEER${PEER_NUMBER}_ORG4_CA"
    #DODATO - KRAJ
    #export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG4_CA
    export CORE_PEER_TLS_ROOTCERT_FILE="${!peer_ca_variable}" #zamenilo liniju iznad
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org4.example.com/users/Admin@org4.example.com/msp
    case $PEER_NUMBER in
        0) export CORE_PEER_ADDRESS="localhost:13051" ;;
        1) export CORE_PEER_ADDRESS="localhost:13052" ;;
        2) export CORE_PEER_ADDRESS="localhost:13053" ;;
        3) export CORE_PEER_ADDRESS="localhost:13055" ;;
        *) echo "Invalid PEER_NUMBER" ;;
    esac
  else
    errorln "ORG Unknown"
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

# Set environment variables for use in the CLI container 
setGlobalsCLI() {
  setGlobals $1 $2 #izmenjeno, dodato $2
  #DODATO - POCETAK
  local PEER_NUMBER=${2:-0}  # Ukoliko $2 nije poslato, postavi na 0
  #DODATO - KRAJ

  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  if [ $USING_ORG -eq 1 ]; then
    case $PEER_NUMBER in
        0) export CORE_PEER_ADDRESS=peer0.org1.example.com:7051 ;;
        1) export CORE_PEER_ADDRESS=peer1.org1.example.com:7052 ;;
        2) export CORE_PEER_ADDRESS=peer2.org1.example.com:7053 ;;
        3) export CORE_PEER_ADDRESS=peer3.org1.example.com:7055 ;;
        *) echo "Invalid PEER_NUMBER" ;;
    esac
    #export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
  elif [ $USING_ORG -eq 2 ]; then
    case $PEER_NUMBER in
        0) export CORE_PEER_ADDRESS=peer0.org2.example.com:9051 ;;
        1) export CORE_PEER_ADDRESS=peer1.org2.example.com:9052 ;;
        2) export CORE_PEER_ADDRESS=peer2.org2.example.com:9053 ;;
        3) export CORE_PEER_ADDRESS=peer3.org2.example.com:9055 ;;
        *) echo "Invalid PEER_NUMBER" ;;
    esac
    #export CORE_PEER_ADDRESS=peer0.org2.example.com:9051
  elif [ $USING_ORG -eq 3 ]; then
    case $PEER_NUMBER in
        0) export CORE_PEER_ADDRESS=peer0.org3.example.com:11051 ;;
        1) export CORE_PEER_ADDRESS=peer1.org3.example.com:11052 ;;
        2) export CORE_PEER_ADDRESS=peer2.org3.example.com:11053 ;;
        3) export CORE_PEER_ADDRESS=peer3.org3.example.com:11055 ;;
        *) echo "Invalid PEER_NUMBER" ;;
    esac
    #export CORE_PEER_ADDRESS=peer0.org3.example.com:11051
  elif [ $USING_ORG -eq 4 ]; then
    case $PEER_NUMBER in
        0) export CORE_PEER_ADDRESS=peer0.org4.example.com:13051 ;;
        1) export CORE_PEER_ADDRESS=peer1.org4.example.com:13052 ;;
        2) export CORE_PEER_ADDRESS=peer2.org4.example.com:13053 ;;
        3) export CORE_PEER_ADDRESS=peer3.org4.example.com:13055 ;;
        *) echo "Invalid PEER_NUMBER" ;;
    esac
    #export CORE_PEER_ADDRESS=peer0.org4.example.com:13051
  else
    errorln "ORG Unknown"
  fi
}

# TODO: IZMENITI
# parsePeerConnectionParameters $@
# Helper function that sets the peer connection parameters for a chaincode
# operation
parsePeerConnectionParameters() {
  PEER_CONN_PARMS=""
  PEERS=""
  while [ "$#" -gt 0 ]; do
    for ((i=0; i<4; i++)); do
      setGlobals $1 $i
      PEER="peer$i.org$1"
      ## Set peer addresses
      PEERS="$PEERS $PEER"
      PEER_CONN_PARMS="$PEER_CONN_PARMS --peerAddresses $CORE_PEER_ADDRESS"
      ## Set path to TLS certificate
      TLSINFO=$(eval echo "--tlsRootCertFiles \$PEER${i}_ORG$1_CA")
      PEER_CONN_PARMS="$PEER_CONN_PARMS $TLSINFO"
    done
    # shift by one to get to the next organization
    shift
  done
  # remove leading space for output
  PEERS="$(echo -e "$PEERS" | sed -e 's/^[[:space:]]*//')"
}

verifyResult() {
  if [ $1 -ne 0 ]; then
    fatalln "$2"
  fi
}
