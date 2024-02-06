peer lifecycle chaincode package chaincode.tar.gz --path $(pwd) --lang golang --label 0.1.0
peer lifecycle chaincode install chaincode.tar.gz
peer lifecycle chaincode approveformyorg --channelID secondChannel --name base --version 0.1.0 --package-id 83908be9db1a0257f9011f8a2ad7f094f23d748c91d323973247bd450b43783d --sequence 1 --waitForEvent # SEQUENCE HAS TO CHANGE EACH TIME
peer lifecycle chaincode approveformyorg --channelID myChannel --name base --version 0.1.0 --package-id 83908be9db1a0257f9011f8a2ad7f094f23d748c91d323973247bd450b43783d --sequence 1 --waitForEvent --channel-config-policy /home/aleksa/Desktop/fabric-bank/fabric-samples/test-network/configtx/configtx.yaml # SEQUENCE HAS TO CHANGE EACH TIME


