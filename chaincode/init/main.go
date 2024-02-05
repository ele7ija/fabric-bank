package main

import (
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/chaincode"
)

type SmartContract struct {
	contractapi.Contract
}

type Bank struct {
    BankId uuid.UUID
    CentralLocation string
    FoundingYear int
    PIB string
    Users []User
}

type User struct {
    UserId uuid.UUID
    Name string
    LastName string
    Email string
    Receipts []Account
}

type Account struct {
    AccountId uuid.UUID
    Amount float64
    Currency string
    Cards []Card
}

type Card struct {
    CardId uuid.UUID
}


func (s *SmartContract) Init (Apistub shim.ChaincodeStubInterface) peer.Response {
    banks := []Bank{

    }

    for _, bank := range banks {
        bankJson, _ := json.Marshal(bank)
        Apistub.PutState(bank.BankId.String(), bankJson)
    }


    return shim.Success(nil)
}

func (s *SmartContract) Invoke (Apistub shim.ChaincodeStubInterface) peer.Response {
    function, args := Apistub.GetFunctionAndParameters()

    switch function {
        case "createBank":
            return s.createBank(Apistub, args)
        case "createAccounts":
            return s.createAccounts(Apistub, args)
        case "depositMoney":
            return s.depositMoney(Apistub, args)
        case "withdrawMoney":
            return s.withdrawMoney(Apistub, args)
        case "transferMoney":
            return s.transferMoney(Apistub, args)
        default:
            return shim.Error("Invalid function name.")
    }

}

func (s *SmartContract) createBank(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
    // Implementirati...
    return shim.Success(nil)
}

func (s *SmartContract) createAccounts(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
    // Implementirati...
    return shim.Success(nil)
}

func (s *SmartContract) depositMoney(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
    // Implementirati...
    return shim.Success(nil)
}


func (s *SmartContract) withdrawMoney(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
    // Implementirati...
    return shim.Success(nil)
}

func (s *SmartContract) transferMoney(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
    // Implementirati...
    return shim.Success(nil)
}

func main() {
    smartContract := new(SmartContract)

    cc, err := contractapi.NewChaincode(smartContract)

    if err != nil {
        log.Panicf("Error creating chaincode: %v", err)
    }

    if err := cc.Start(); err != nil {
        log.Panicf("Error starting chaincode: %v", err)
    }
}
