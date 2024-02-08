# fabric-bank

Projekat iz predmeta PDASP - 2023/2024 - grupa G6.

Autori:
- Aleksa Bajat (E2 114/2023)
- Aleksandar Radišić (E2 17/2023)
- Bojan Popržen (E2 4/2022)

## Pokretanje

Pokretanje ima dva dela:
1. Podizanje mreze i instalacija chaincode-a
2. Klijentska interakcija sa chaincode-om

### 1. Podizanje mreze
```
# Preuzmi potrebne Hyperledger fabric docker slike i preuzmi binarne fajlove i
# prebaci ih u ./fabric-samples/bin i ./fabric-samples/config.
./install-fabric.sh --fabric-version 2.2.6

# Spusti prethodno podignutu mrezu i podigni novu sa 4 organizacije sa po 4
# peer-a.
cd ./fabric-samples/test-network
./network.sh down
./network.sh up

# Kreiranje kanala.
./network.sh createChannel

# Deployment chaincode-a.
./network.sh deployCC -ccn basic -ccp ../../chaincode/init/ -ccl go
```

### 2a. Interakcija pomocu SDK-a (custom klijenta) kroz REST API

```bash
# Pokrenuti REST API server
cd ../../app
go run main.go
```

Primer pozivanja REST endpoint-a:

```bash
$ curl "localhost:8080/api/login?username=user1&password=user1&bank=bank1"                                                                                                              
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYW5rIjoiYmFuazEiLCJ1c2VybmFtZSI6InVzZXIxIn0.e9RVH1J5AgZ2MlEiK5gQfon9UOPDVhiToghj4TAvCo0"}

$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYW5rIjoiYmFuazEiLCJ1c2VybmFtZSI6InVzZXIxIn0.e9RVH1J5AgZ2MlEiK5gQfon9UOPDVhiToghj4TAvCo0" "localhost:8080/api/profile"
{
        "id": "user1",
        "email": "nmalinovic@gmail.com",
        "accounts": [
                {
                        "id": "account1",
                        "amount": 17000,
                        "currency": "RSD"
                },
                {
                        "id": "account2",
                        "amount": 517.0696235806608,
                        "currency": "EUR"
                }
        ]
}

$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYW5rIjoiYmFuazEiLCJ1c2VybmFtZSI6InVzZXIxIn0.e9RVH1J5AgZ2MlEiK5gQfon9UOPDVhiToghj4TAvCo0" "localhost:8080/api/deposit?amount=1000&accountId=account1&currency=RSD"
Successfully deposited money.

$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYW5rIjoiYmFuazEiLCJ1c2VybmFtZSI6InVzZXIxIn0.e9RVH1J5AgZ2MlEiK5gQfon9UOPDVhiToghj4TAvCo0" "localhost:8080/api/profile"
{
        "id": "user1",
        "email": "nmalinovic@gmail.com",
        "accounts": [
                {
                        "id": "account1",
                        "amount": 18000,
                        "currency": "RSD"
                },
                {
                        "id": "account2",
                        "amount": 517.0696235806608,
                        "currency": "EUR"
                }
        ]
}

$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYW5rIjoiYmFuazEiLCJ1c2VybmFtZSI6InVzZXIxIn0.e9RVH1J5AgZ2MlEiK5gQfon9UOPDVhiToghj4TAvCo0" "localhost:8080/api/transfer?amount=2000&accountFrom=account1&accountTo=account2"

$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYW5rIjoiYmFuazEiLCJ1c2VybmFtZSI6InVzZXIxIn0.e9RVH1J5AgZ2MlEiK5gQfon9UOPDVhiToghj4TAvCo0" "localhost:8080/api/profile"
{
        "id": "user1",
        "email": "nmalinovic@gmail.com",
        "accounts": [
                {
                        "id": "account1",
                        "amount": 16000,
                        "currency": "RSD"
                },
                {
                        "id": "account2",
                        "amount": 534.1392471613217,
                        "currency": "EUR"
                }
        ]
}

```

### 2b. Interakcija pomocu Hyperledger alata

Koristiti Go 1.14!

```bash
sudo apt-get install bison
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.14 --binary
gvm use go1.14
```

```bash
# Preuzmi potrebne Hyperledger fabric docker slike i preuzmi binarne fajlove i
# prebaci ih u ./fabric-samples/bin i ./fabric-samples/config.
./install-fabric.sh --fabric-version 2.2.6

# Spusti prethodno podignutu mrezu i podigni novu sa 4 organizacije sa po 4
# peer-a.
cd ./fabric-samples/test-network
./network.sh down
./network.sh up

# Kreiranje kanala.
./network.sh createChannel

# Deployment chaincode-a.
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go

# Invoke chaincode-a.
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" --peerAddresses localhost:13051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'

peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAssets"]}'
# Returns:
# [{"ID":"asset1","color":"blue","size":5,"owner":"Tomoko","appraisedValue":300},{"ID":"asset2","color":"red","size":5,"owner":"Brad","appraisedValue":400},{"ID":"asset3","color":"green","size":10,"owner":"Jin Soo","appraisedValue":500},{"ID":"asset4","color":"yellow","size":10,"owner":"Max","appraisedValue":600},{"ID":"asset5","color":"black","size":15,"owner":"Adriana","appraisedValue":700},{"ID":"asset6","color":"white","size":15,"owner":"Michel","appraisedValue":800}]

#Uzimanje novca sa racuna
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" --peerAddresses localhost:13051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt" -c '{"function":"WithdrawMoney","Args":["account1", "18000"]}'

peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAccounts"]}'

#Dodavanje novca na racun
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" --peerAddresses localhost:13051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt" -c '{"function":"DepositMoney","Args":["account1", "1000"]}'

peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAccounts"]}'

#Transfer novca
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" --peerAddresses localhost:13051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt" -c '{"function":"TransferMoney","Args":["account1", "account10", "1000"]}'

peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAccounts"]}'

#Novi korisnik
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" --peerAddresses localhost:13051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt" -c '{"function":"CreateUser","Args":["Aleksandar", "Radisic", "aradisic@gmail.com", "bank1"]}'

peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllUsers"]}'
peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllBanks"]}'

#Novi racun za korisnika
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt" --peerAddresses localhost:13051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org4.example.com/peers/peer0.org4.example.com/tls/ca.crt" -c '{"function":"CreateAccount","Args":["user13", "RSD"]}'

peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAccounts"]}'
peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllUsers"]}'

peer chaincode query -C mychannel -n basic -c '{"Args":["QueryUsers", "Nikola", "", "", "1"]}'
# Returns:
# [{"UserId":"user1","Password":"user1","Name":"Nikola","LastName":"Malinovic","Email":"nmalinovic@gmail.com","Receipts":["account1","account2"]}]


peer chaincode query -C mychannel -n basic -c '{"Args":["GetUsersWithMoreResources", "20000", "RSD"]}'
# Returns:
# [{"User":{"UserId":"user1","Password":"user1","Name":"Nikola","LastName":"Malinovic","Email":"nmalinovic@gmail.com","Receipts":["account1","account2"]},"Accounts":[{"AccountId":"account1","Amount":20000,"Currency":"RSD","Cards":[{"CardId":"card1"},{"CardId":"card2"}]}]},{"User":{"UserId":"user7","Password":"user7","Name":"Ognjen","LastName":"Zalis","Email":"ozalis@gmail.com","Receipts":["account13","account14"]},"Accounts":[{"AccountId":"account13","Amount":35000,"Currency":"RSD","Cards":[{"CardId":"card25"},{"CardId":"card26"}]}]}]


```


