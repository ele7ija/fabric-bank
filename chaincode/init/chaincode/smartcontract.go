package chaincode

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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
	Password string
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

type UserWithAccounts struct {
	User     User
	Accounts []Account
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
			Password: "user1",
			Name:     "Nikola",
			LastName: "Malinovic",
			Email:    "nmalinovic@gmail.com",
			Receipts: []string{accounts[0].AccountId, accounts[1].AccountId},
		},
		{
			UserId:   "user2",
			Password: "user2",
			Name:     "Igor",
			LastName: "Tot",
			Email:    "itot@gmail.com",
			Receipts: []string{accounts[2].AccountId, accounts[3].AccountId},
		},
		{
			UserId:   "user3",
			Password: "user3",
			Name:     "Jelena",
			LastName: "Petrovic",
			Email:    "jpetrovic@gmail.com",
			Receipts: []string{accounts[4].AccountId, accounts[5].AccountId},
		},
		{
			UserId:   "user4",
			Password: "user4",
			Name:     "Petar",
			LastName: "Djukic",
			Email:    "pdjukic@gmail.com",
			Receipts: []string{accounts[6].AccountId, accounts[7].AccountId},
		},
		{
			UserId:   "user5",
			Password: "user5",
			Name:     "Nikolina",
			LastName: "Vukovic",
			Email:    "nvukovic@gmail.com",
			Receipts: []string{accounts[8].AccountId, accounts[9].AccountId},
		},
		{

			UserId:   "user6",
			Password: "user6",
			Name:     "Nenad",
			LastName: "Obradovic",
			Email:    "nobradovic@gmail.com",
			Receipts: []string{accounts[10].AccountId, accounts[11].AccountId},
		},
		{
			UserId:   "user7",
			Password: "user7",
			Name:     "Ognjen",
			LastName: "Zalis",
			Email:    "ozalis@gmail.com",
			Receipts: []string{accounts[12].AccountId, accounts[13].AccountId},
		},
		{
			UserId:   "user8",
			Password: "user8",
			Name:     "Milan",
			LastName: "Kovacevic",
			Email:    "mkovacevic@gmail.com",
			Receipts: []string{accounts[14].AccountId, accounts[15].AccountId},
		},
		{
			UserId:   "user9",
			Password: "user9",
			Name:     "Ksenija",
			LastName: "Jovancevic",
			Email:    "kjovancevic@gmail.com",
			Receipts: []string{accounts[16].AccountId, accounts[17].AccountId},
		},
		{
			UserId:   "user10",
			Password: "user10",
			Name:     "Milica",
			LastName: "Simovic",
			Email:    "msimovic@gmail.com",
			Receipts: []string{accounts[18].AccountId, accounts[19].AccountId},
		},
		{
			UserId:   "user11",
			Password: "user11",
			Name:     "Luka",
			LastName: "Dragovic",
			Email:    "ldragovic@gmail.com",
			Receipts: []string{accounts[20].AccountId, accounts[21].AccountId},
		},
		{
			UserId:   "user12",
			Password: "user12",
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

	return nil
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func (s *SmartContract) GetAllBanks(ctx contractapi.TransactionContextInterface) ([]*Bank, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("bank1", "bank99999")
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

func (s *SmartContract) GetAllUsers(ctx contractapi.TransactionContextInterface) ([]*User, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("user1", "user99999")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var users []*User
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var user User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (s *SmartContract) GetAllAccounts(ctx contractapi.TransactionContextInterface) ([]*Account, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("account1", "account99999")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var accounts []*Account
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var account Account
		err = json.Unmarshal(queryResponse.Value, &account)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}

func (s *SmartContract) GetAccount(ctx contractapi.TransactionContextInterface, id string) (*Account, error) {
	exist, err := s.AssetExists(ctx, id)
	if !exist {
		return nil, fmt.Errorf("Account not found")
	}
	accountJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	var account Account
	err = json.Unmarshal(accountJSON, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *SmartContract) WithdrawMoney(ctx contractapi.TransactionContextInterface, accountId string, amount float64) error {
	exists, err := s.AssetExists(ctx, accountId)
	if !exists {
		return fmt.Errorf("Account does not exist")
	}
	accountJSON, err := ctx.GetStub().GetState(accountId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	var account Account
	err = json.Unmarshal(accountJSON, &account)
	if err != nil {
		return err
	}
	if account.Amount < amount {
		return fmt.Errorf("Account does not have that amount of money")
	}
	account.Amount -= amount
	accountJSON, err = json.Marshal(account)
	if err != nil {
		return err
	}
	ctx.GetStub().PutState(account.AccountId, accountJSON)
	return nil
}

func (s *SmartContract) DepositMoney(ctx contractapi.TransactionContextInterface, accountId string, amount float64) error {
	exists, err := s.AssetExists(ctx, accountId)
	if !exists {
		return fmt.Errorf("Account does not exist")
	}
	accountJSON, err := ctx.GetStub().GetState(accountId)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	var account Account
	err = json.Unmarshal(accountJSON, &account)
	if err != nil {
		return err
	}
	account.Amount += amount
	accountJSON, err = json.Marshal(&account)
	if err != nil {
		return err
	}
	ctx.GetStub().PutState(account.AccountId, accountJSON)
	return nil
}

func (s *SmartContract) TransferMoney(ctx contractapi.TransactionContextInterface, accountFromId string, accountToId string, amount float64) error {
	accountFrom, err := s.GetAccount(ctx, accountFromId)
	if err != nil {
		return err
	}
	accountTo, err := s.GetAccount(ctx, accountToId)
	if err != nil {
		return err
	}
	if accountFrom.Amount < amount {
		return fmt.Errorf("Account does not have that amount of money")
	}
	amount_to := amount
	if accountFrom.Currency != accountTo.Currency {
		amount_to, err = TransferAmountToCurrency(amount, accountFrom.Currency, accountTo.Currency)
		if err != nil {
			return fmt.Errorf("Invalid currency")
		}
	}
	accountFrom.Amount -= amount
	accountTo.Amount += amount_to
	err = putAccount(ctx, accountFrom)
	if err != nil {
		return err
	}
	err = putAccount(ctx, accountTo)
	if err != nil {
		return err
	}
	return nil
}

func (s *SmartContract) CreateUser(ctx contractapi.TransactionContextInterface, name string, lastName string, email string, bankId string) error {
	users, err := s.GetAllUsers(ctx)
	if err != nil {
		return err
	}
	max := -1
	for _, user := range users {
		num, err := extractNumber("user", user.UserId)
		if err != nil {
			return err
		}
		if num > max {
			max = num
		}
	}
	bank, err := GetBank(ctx, bankId)
	if err != nil {
		return err
	}
	for _, userId := range bank.Users {
		user, err := s.GetUser(ctx, userId)
		if err != nil {
			return nil
		}
		if user.Email == email {
			return fmt.Errorf("Email already assigned to user in this bank")
		}
	}
	new_user := User{
		UserId:   fmt.Sprintf("user%d", max+1),
		Name:     name,
		LastName: lastName,
		Email:    email,
		Receipts: []string{},
	}
	err = putUser(ctx, &new_user)
	if err != nil {
		return err
	}
	bank.Users = append(bank.Users, new_user.UserId)
	err = putBank(ctx, bank)
	if err != nil {
		return err
	}
	return nil
}

func (s *SmartContract) CreateAccount(ctx contractapi.TransactionContextInterface, userId string, currency string) error {
	accounts, err := s.GetAllAccounts(ctx)
	if err != nil {
		return err
	}
	max_card := -1
	max_account := -1
	for _, account := range accounts {
		num_acc, err := extractNumber("account", account.AccountId)
		if err != nil {
			return err
		}
		if num_acc > max_account {
			max_account = num_acc
		}
		for _, card := range account.Cards {
			num, err := extractNumber("card", card.CardId)
			if err != nil {
				return err
			}
			if num > max_card {
				max_card = num
			}
		}
	}
	_, ok := makeMap()[currency]
	if !ok {
		fmt.Errorf("Currency %s not found", currency)
	}
	new_account := Account{
		AccountId: fmt.Sprintf("account%d", max_account+1),
		Currency:  currency,
		Amount:    0,
		Cards:     []Card{Card{CardId: fmt.Sprintf("card%d", max_card+1)}, Card{fmt.Sprintf("card%d", max_card+2)}},
	}
	user, err := s.GetUser(ctx, userId)
	if err != nil {
		return nil
	}
	user.Receipts = append(user.Receipts, new_account.AccountId)
	err = putAccount(ctx, &new_account)
	if err != nil {
		return err
	}
	err = putUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func extractNumber(prefix string, s string) (int, error) {
	s = strings.TrimPrefix(s, prefix)
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func putAccount(ctx contractapi.TransactionContextInterface, account *Account) error {
	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}
	ctx.GetStub().PutState(account.AccountId, accountJSON)
	return nil
}

func putBank(ctx contractapi.TransactionContextInterface, bank *Bank) error {
	objJSON, err := json.Marshal(bank)
	if err != nil {
		return err
	}
	ctx.GetStub().PutState(bank.BankId, objJSON)
	return nil
}

func putUser(ctx contractapi.TransactionContextInterface, user *User) error {
	objJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	ctx.GetStub().PutState(user.UserId, objJSON)
	return nil
}

func GetBank(ctx contractapi.TransactionContextInterface, id string) (*Bank, error) {
	bankJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if bankJSON == nil {
		return nil, fmt.Errorf("Bank not found")
	}
	var bank Bank
	err = json.Unmarshal(bankJSON, &bank)
	if err != nil {
		return nil, err
	}
	return &bank, nil
}

func (s *SmartContract) GetUser(ctx contractapi.TransactionContextInterface, id string) (*User, error) {
	userJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userJSON == nil {
		return nil, fmt.Errorf("User not found")
	}
	var user User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func TransferAmountToCurrency(amount float64, from string, to string) (float64, error) {
	currencies := makeMap()
	valueFrom, ok := currencies[from]
	if !ok {
		return -1, fmt.Errorf("Invalid currency")
	}
	valueTo, ok := currencies[to]
	if !ok {
		return -1, fmt.Errorf("Invalid currency")
	}
	return (valueFrom / valueTo) * amount, nil
}

func makeMap() map[string]float64 {
	currencies := make(map[string]float64)
	currencies["RSD"] = 1
	currencies["EUR"] = 117.1672
	currencies["AUD"] = 71.0492
	currencies["CAD"] = 80.6770
	currencies["CNY"] = 15.1375
	currencies["CZK"] = 4.6982
	currencies["DKK"] = 15.7035
	currencies["HUF"] = 30.3173
	currencies["INR"] = 1.3114
	currencies["JPY"] = 73.5559
	currencies["KWD"] = 353.4456
	currencies["NOK"] = 10.2950
	return currencies
}

func (s *SmartContract) QueryUsers(ctx contractapi.TransactionContextInterface, name string, lastName string, email string, minReceiptsCount int) ([]*User, error) {
	queryString := fmt.Sprintf(`{
        "selector": {
			"UserId": {"$regex": "^user"},
            "Name": {"$regex": "%s"},
            "LastName": {"$regex": "%s"},
            "Email": {"$regex": "%s"}
        }
    }`, name, lastName, email, minReceiptsCount)
	queryResultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer queryResultsIterator.Close()
	filteredUsers := []*User{}
	for queryResultsIterator.HasNext() {
		queryResult, err := queryResultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("error retrieving next query result: %v", err)
		}
		var user User
		if err := json.Unmarshal(queryResult.Value, &user); err != nil {
			return nil, fmt.Errorf("error unmarshalling user: %v", err)
		}
		if len(user.Receipts) >= minReceiptsCount {
			filteredUsers = append(filteredUsers, &user)
		}
	}

	return filteredUsers, nil
}

func (s *SmartContract) GetUsersWithMoreResources(ctx contractapi.TransactionContextInterface, amount float64, currency string) ([]*UserWithAccounts, error) {
	currencies := makeMap()
	_, ok := currencies[currency]
	if !ok {
		return nil, fmt.Errorf("Invalid currency")
	}
	accountsQuery := fmt.Sprintf(`{
		"selector": {
			"Currency": "%s",
			"Amount": {
				"$gte": %f
			}
		}
	}`, currency, amount)
	accountsResult, err := ctx.GetStub().GetQueryResult(accountsQuery)
	if err != nil {
		return nil, err
	}
	defer accountsResult.Close()
	var accounts []Account
	for accountsResult.HasNext() {
		accountResult, err := accountsResult.Next()
		if err != nil {
			return nil, err
		}
		var account Account
		if err := json.Unmarshal(accountResult.Value, &account); err != nil {
			return nil, fmt.Errorf("error unmarshalling account: %v", err)
		}
		accounts = append(accounts, account)
	}

	accountConditions := []string{}
	for _, acc := range accounts {
		condition := fmt.Sprintf(`{"Receipts": {"$elemMatch": {"$eq": "%s"}}}`, acc.AccountId)
		accountConditions = append(accountConditions, condition)
	}
	usersQuery := fmt.Sprintf(`
	{
		"selector": {
			"$and": [
				{
					"$or": [
						%s
					]
				},
				{
					"UserId": {"$regex": "^user"}
				}
			]
		}
	}`, strings.Join(accountConditions, "\n"))
	usersQueryIterator, err := ctx.GetStub().GetQueryResult(usersQuery)
	if err != nil {
		return nil, fmt.Errorf("error executing users query: %v", err)
	}
	defer usersQueryIterator.Close()

	var users []User
	for usersQueryIterator.HasNext() {
		queryResult, err := usersQueryIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("error retrieving next user query result: %v", err)
		}

		var user User
		if err := json.Unmarshal(queryResult.Value, &user); err != nil {
			return nil, fmt.Errorf("error unmarshalling user: %v", err)
		}
		users = append(users, user)
	}

	retVal := []*UserWithAccounts{}
	for _, user := range users {
		userWithAccounts := UserWithAccounts{
			User:     user,
			Accounts: []Account{},
		}
		for _, accId := range user.Receipts {
			for _, account := range accounts {
				if account.AccountId == accId {
					userWithAccounts.Accounts = append(userWithAccounts.Accounts, account)
				}
			}
		}
	}
	return retVal, nil
}
