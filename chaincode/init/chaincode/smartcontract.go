package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}
type Bank struct {
	BankId          string
	Name            string
	CentralLocation string
	FoundingYear    int
	PIB             string
	Users           []string
}

type User struct {
	UserId   string
	Name     string
	LastName string
	Email    string
	Receipts []string
}

type Account struct {
	AccountId string
	Amount    float64
	Currency  string
	Cards     []Card
}

type Card struct {
	CardId string
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	ID             string `json:"ID"`
	Color          string `json:"color"`
	Size           int    `json:"size"`
	Owner          string `json:"owner"`
	AppraisedValue int    `json:"appraisedValue"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	accounts := []Account{
		{
			AccountId: "account1",
			Amount:    20000,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card1"}, Card{CardId: "card2"}},
		},
		{
			AccountId: "account2",
			Amount:    500,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card3"}, Card{CardId: "card4"}},
		},
		{
			AccountId: "account3",
			Amount:    1000,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card5"}, Card{CardId: "card6"}},
		},
		{
			AccountId: "account4",
			Amount:    5500,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card7"}, Card{CardId: "card8"}},
		},
		{
			AccountId: "account5",
			Amount:    2300,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card9"}, Card{CardId: "card10"}},
		},
		{
			AccountId: "account6",
			Amount:    2400,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card11"}, Card{CardId: "card12"}},
		},
		{
			AccountId: "account7",
			Amount:    60,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card13"}, Card{CardId: "card14"}},
		},
		{
			AccountId: "account8",
			Amount:    1400,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card15"}, Card{CardId: "card16"}},
		},
		{
			AccountId: "account9",
			Amount:    1500,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card17"}, Card{CardId: "card18"}},
		},
		{
			AccountId: "account10",
			Amount:    3500,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card19"}, Card{CardId: "card20"}},
		},
		{
			AccountId: "account11",
			Amount:    4230,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card21"}, Card{CardId: "card22"}},
		},
		{
			AccountId: "account12",
			Amount:    100,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card23"}, Card{CardId: "card24"}},
		},
		{
			AccountId: "account13",
			Amount:    35000,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card25"}, Card{CardId: "card26"}},
		},
		{
			AccountId: "account14",
			Amount:    600,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card27"}, Card{CardId: "card28"}},
		},
		{
			AccountId: "account15",
			Amount:    2300,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card29"}, Card{CardId: "card30"}},
		},
		{
			AccountId: "account16",
			Amount:    3600,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card31"}, Card{CardId: "card32"}},
		},
		{
			AccountId: "account17",
			Amount:    900,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card33"}, Card{CardId: "card34"}},
		},
		{
			AccountId: "account18",
			Amount:    420,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card35"}, Card{CardId: "card36"}},
		},
		{
			AccountId: "account19",
			Amount:    4350,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card37"}, Card{CardId: "card38"}},
		},
		{
			AccountId: "account20",
			Amount:    30,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card39"}, Card{CardId: "card40"}},
		},
		{
			AccountId: "account21",
			Amount:    400,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card41"}, Card{CardId: "card42"}},
		},
		{
			AccountId: "account22",
			Amount:    100,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card43"}, Card{CardId: "card44"}},
		},
		{
			AccountId: "account23",
			Amount:    2300,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: "card45"}, Card{CardId: "card46"}},
		},
		{
			AccountId: "account24",
			Amount:    3400,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: "card47"}, Card{CardId: "card48"}},
		},
	}
	users := []User{
		{
			UserId:   "user1",
			Name:     "Nikola",
			LastName: "Malinovic",
			Email:    "nmalinovic@gmail.com",
			Receipts: []string{accounts[0].AccountId, accounts[1].AccountId},
		},
		{
			UserId:   "user2",
			Name:     "Igor",
			LastName: "Tot",
			Email:    "itot@gmail.com",
			Receipts: []string{accounts[2].AccountId, accounts[3].AccountId},
		},
		{
			UserId:   "user3",
			Name:     "Jelena",
			LastName: "Petrovic",
			Email:    "jpetrovic@gmail.com",
			Receipts: []string{accounts[4].AccountId, accounts[5].AccountId},
		},
		{
			UserId:   "user4",
			Name:     "Petar",
			LastName: "Djukic",
			Email:    "pdjukic@gmail.com",
			Receipts: []string{accounts[6].AccountId, accounts[7].AccountId},
		},
		{
			UserId:   "user5",
			Name:     "Aleksandar",
			LastName: "Vukovic",
			Email:    "nmalinovic@gmail.com",
			Receipts: []string{accounts[8].AccountId, accounts[9].AccountId},
		},
		{

			UserId:   "user6",
			Name:     "Nenad",
			LastName: "Obradovic",
			Email:    "nobradovic@gmail.com",
			Receipts: []string{accounts[10].AccountId, accounts[11].AccountId},
		},
		{
			UserId:   "user7",
			Name:     "Ognjen",
			LastName: "Zalis",
			Email:    "ozalis@gmail.com",
			Receipts: []string{accounts[12].AccountId, accounts[13].AccountId},
		},
		{
			UserId:   "user8",
			Name:     "Milan",
			LastName: "Kovacevic",
			Email:    "mkovacevic@gmail.com",
			Receipts: []string{accounts[14].AccountId, accounts[15].AccountId},
		},
		{
			UserId:   "user9",
			Name:     "Ksenija",
			LastName: "Jovancevic",
			Email:    "kjovancevic@gmail.com",
			Receipts: []string{accounts[16].AccountId, accounts[17].AccountId},
		},
		{
			UserId:   "user10",
			Name:     "Milica",
			LastName: "Simovic",
			Email:    "msimovic@gmail.com",
			Receipts: []string{accounts[18].AccountId, accounts[19].AccountId},
		},
		{
			UserId:   "user11",
			Name:     "Luka",
			LastName: "Dragovic",
			Email:    "ldragovic@gmail.com",
			Receipts: []string{accounts[20].AccountId, accounts[21].AccountId},
		},
		{
			UserId:   "user12",
			Name:     "Vanja",
			LastName: "Tanovic",
			Email:    "vtanovic@gmail.com",
			Receipts: []string{accounts[22].AccountId, accounts[23].AccountId},
		},
	}
	banks := []Bank{
		{
			BankId:          "bank1",
			Name:            "Banka1",
			CentralLocation: "Beograd",
			FoundingYear:    1952,
			PIB:             "123456789",
			Users:           []string{users[0].UserId, users[1].UserId, users[2].UserId},
		},
		{
			BankId:          "bank2",
			Name:            "Banka2",
			CentralLocation: "Novi Sad",
			FoundingYear:    1963,
			PIB:             "234567891",
			Users:           []string{users[3].UserId, users[4].UserId, users[5].UserId},
		},
		{

			BankId:          "bank3",
			Name:            "Banka3",
			CentralLocation: "Nis",
			FoundingYear:    1972,
			PIB:             "345678912",
			Users:           []string{users[6].UserId, users[7].UserId, users[8].UserId},
		},

		{

			BankId:          "bank4",
			Name:            "Banka4",
			CentralLocation: "Subotica",
			FoundingYear:    1981,
			PIB:             "456789123",
			Users:           []string{users[9].UserId, users[10].UserId, users[11].UserId},
		},
	}

	for _, bank := range banks {
		bankJSON, err := json.Marshal(bank)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(bank.BankId, bankJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	for _, user := range users {
		userJSON, err := json.Marshal(user)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(user.UserId, userJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	for _, account := range accounts {
		accountJSON, err := json.Marshal(account)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(account.AccountId, accountJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	assets := []Asset{
		{ID: "asset1", Color: "blue", Size: 5, Owner: "Tomoko", AppraisedValue: 300},
		{ID: "asset2", Color: "red", Size: 5, Owner: "Brad", AppraisedValue: 400},
		{ID: "asset3", Color: "green", Size: 10, Owner: "Jin Soo", AppraisedValue: 500},
		{ID: "asset4", Color: "yellow", Size: 10, Owner: "Max", AppraisedValue: 600},
		{ID: "asset5", Color: "black", Size: 15, Owner: "Adriana", AppraisedValue: 700},
		{ID: "asset6", Color: "white", Size: 15, Owner: "Michel", AppraisedValue: 800},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Owner = newOwner
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (s *SmartContract) GetAllBanks(ctx contractapi.TransactionContextInterface) ([]*Bank, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var banks []*Bank
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var bank Bank
		err = json.Unmarshal(queryResponse.Value, &bank)
		if err != nil {
			return nil, err
		}
		banks = append(banks, &bank)
	}

	return banks, nil
}
