package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type SmartContract struct {
	contractapi.Contract
}

type Bank struct {
	BankId          uuid.UUID
	Name            string
	CentralLocation string
	FoundingYear    int
	PIB             string
	Users           []User
}

type User struct {
	UserId   uuid.UUID
	Name     string
	LastName string
	Email    string
	Receipts []Account
}

type Account struct {
	AccountId uuid.UUID
	Amount    float64
	Currency  string
	Cards     []Card
}

type Card struct {
	CardId uuid.UUID
}

func (s *SmartContract) Init(Apistub shim.ChaincodeStubInterface) peer.Response {
	accounts := []Account{
		{
			AccountId: uuid.New(),
			Amount:    20000,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    500,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    1000,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    5500,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    2300,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    2400,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    60,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    1400,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    1500,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    3500,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    4230,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    100,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    35000,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    600,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    2300,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    3600,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    900,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    420,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    4350,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    30,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    400,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    100,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    2300,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
		{
			AccountId: uuid.New(),
			Amount:    3400,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.New()}, Card{CardId: uuid.New()}},
		},
	}

	users := []User{
		{
			UserId:   uuid.New(),
			Name:     "Nikola",
			LastName: "Malinovic",
			Email:    "nmalinovic@gmail.com",
			Receipts: []Account{accounts[0], accounts[1]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Igor",
			LastName: "Tot",
			Email:    "itot@gmail.com",
			Receipts: []Account{accounts[2], accounts[3]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Jelena",
			LastName: "Petrovic",
			Email:    "jpetrovic@gmail.com",
			Receipts: []Account{accounts[4], accounts[5]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Petar",
			LastName: "Djukic",
			Email:    "pdjukic@gmail.com",
			Receipts: []Account{accounts[6], accounts[7]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Aleksandar",
			LastName: "Vukovic",
			Email:    "nmalinovic@gmail.com",
			Receipts: []Account{accounts[8], accounts[9]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Nenad",
			LastName: "Obradovic",
			Email:    "nobradovic@gmail.com",
			Receipts: []Account{accounts[10], accounts[11]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Ognjen",
			LastName: "Zalis",
			Email:    "ozalis@gmail.com",
			Receipts: []Account{accounts[12], accounts[13]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Milan",
			LastName: "Kovacevic",
			Email:    "mkovacevic@gmail.com",
			Receipts: []Account{accounts[14], accounts[15]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Ksenija",
			LastName: "Jovancevic",
			Email:    "kjovancevic@gmail.com",
			Receipts: []Account{accounts[16], accounts[17]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Milica",
			LastName: "Simovic",
			Email:    "msimovic@gmail.com",
			Receipts: []Account{accounts[18], accounts[19]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Luka",
			LastName: "Dragovic",
			Email:    "ldragovic@gmail.com",
			Receipts: []Account{accounts[20], accounts[21]},
		},
		{
			UserId:   uuid.New(),
			Name:     "Vanja",
			LastName: "Tanovic",
			Email:    "vtanovic@gmail.com",
			Receipts: []Account{accounts[22], accounts[23]},
		},
	}

	banks := []Bank{
		{
			BankId:          uuid.New(),
			Name:            "Banka1",
			CentralLocation: "Beograd",
			FoundingYear:    1952,
			PIB:             "123456789",
			Users:           []User{users[0], users[1], users[2]},
		},
		{
			BankId:          uuid.New(),
			Name:            "Banka2",
			CentralLocation: "Novi Sad",
			FoundingYear:    1963,
			PIB:             "234567891",
			Users:           []User{users[3], users[4], users[5]},
		},
		{
			BankId:          uuid.New(),
			Name:            "Banka3",
			CentralLocation: "Nis",
			FoundingYear:    1972,
			PIB:             "345678912",
			Users:           []User{users[6], users[7], users[8]},
		},
		{
			BankId:          uuid.New(),
			Name:            "Banka4",
			CentralLocation: "Subotica",
			FoundingYear:    1981,
			PIB:             "456789123",
			Users:           []User{users[9], users[10], users[11]},
		},
	}

	for _, bank := range banks {
		bankJson, _ := json.Marshal(bank)
		Apistub.PutState(bank.BankId.String(), bankJson)
	}

	return shim.Success(nil)
}

func (s *SmartContract) Invoke(Apistub shim.ChaincodeStubInterface) peer.Response {
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
