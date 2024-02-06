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
			AccountId: uuid.Must(uuid.Parse("5113068e-7162-4378-8318-3cafa26ac0eb")),
			Amount:    20000,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("91fd3109-9e32-4e6c-945c-7dbe9b1b539e"))}, Card{CardId: uuid.Must(uuid.Parse("166d3c32-67d1-4fbb-aa90-fe7a00ba3fc0"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("db6a4cb9-4e7f-4b26-b3c2-5f9ebdba888d")),
			Amount:    500,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("7ed9b8b0-9ef6-479c-bab6-446f22aa8820"))}, Card{CardId: uuid.Must(uuid.Parse("114d937a-f8c4-4583-b10a-152e4d260048"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("f3febd2b-9eb6-41a9-a747-2feb1065f45a")),
			Amount:    1000,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("26b44b2f-4e0f-49f8-8d9d-e393a3027abb"))}, Card{CardId: uuid.Must(uuid.Parse("29bcb510-8456-4894-abd1-57d8a7ce5244"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("dc6bbcf3-e170-40e3-bb6f-d506725fc4a0")),
			Amount:    5500,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("ffa8bdd1-64d8-40f8-a964-64935045dbc9"))}, Card{CardId: uuid.Must(uuid.Parse("aa1892d2-7094-4bcc-b39f-04a6168283b3"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("709fce20-2897-4a74-93e7-6adb25bbb116")),
			Amount:    2300,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("ae805b1b-18f2-4b7d-8df0-182008d0674a"))}, Card{CardId: uuid.Must(uuid.Parse("e9f85639-9467-4369-b409-1c098b1b87cc"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("bddafeb8-1909-48b6-8d66-f0844edee433")),
			Amount:    2400,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("8ba5ed48-34fb-43af-b156-a57ae6bb6d90"))}, Card{CardId: uuid.Must(uuid.Parse("9ee8c2e5-5a51-4e77-9a18-f2fa9ea562d8"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("26384418-14fb-4ebe-b404-5015f3c275b1")),
			Amount:    60,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("b668ca2f-d0ad-47ab-be26-926f55995ba7"))}, Card{CardId: uuid.Must(uuid.Parse("d07e331c-6c6f-4631-9f75-fb44c0e19c61"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("1702f106-d3f5-48e2-b4c1-71589cde0c05")),
			Amount:    1400,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("9e3cca5a-39bb-484a-88e9-fd0fe0fc0702"))}, Card{CardId: uuid.Must(uuid.Parse("bf9f68e6-1809-4050-a1ee-8ca7176d18f6"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("5a2b4b75-c5e3-4e28-a53e-5dc9d4166125")),
			Amount:    1500,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("5a322c49-8baf-4418-9775-812fe3faea9f"))}, Card{CardId: uuid.Must(uuid.Parse("b2feb921-9630-4c1f-9751-e3a9302ac435"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("6844ce91-3e8f-4cdc-8180-5d3456d8dc6f")),
			Amount:    3500,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("92a0ccbd-3e4f-401e-8a56-16b147e12b73"))}, Card{CardId: uuid.Must(uuid.Parse("b275072b-a3e5-4775-925e-886d213a5978"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("6208aa56-426c-42cd-afa8-cbd5aad56053")),
			Amount:    4230,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("a9fb80ed-de77-4ede-a84a-61919e0393ba"))}, Card{CardId: uuid.Must(uuid.Parse("f6c96081-cd4b-413c-bb96-ad274ac41fec"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("683af87d-da90-45a6-87ce-19a068402f38")),
			Amount:    100,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("58dd9770-62f3-4065-824e-68bd00fb6c2d"))}, Card{CardId: uuid.Must(uuid.Parse("9cb20b7a-d17e-4aec-8fd1-d89bdc35c016"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("012fc710-2a80-4f21-acbf-27b787c77602")),
			Amount:    35000,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("67d08b8c-b080-4697-becc-760f451bc1df"))}, Card{CardId: uuid.Must(uuid.Parse("24e25448-66f9-4743-a7e6-63ab2be98836"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("098b7094-96db-42e2-b6fa-1c4284907f1f")),
			Amount:    600,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("d4abb341-1f96-4bb7-8df3-499a54db3fe0"))}, Card{CardId: uuid.Must(uuid.Parse("2fb65ae7-80a2-4af4-a626-f3461906f938"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("4326dc23-6df7-41dc-b1fa-85c06e2225e1")),
			Amount:    2300,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("08f32860-60bf-4153-bbe7-87816e6865a4"))}, Card{CardId: uuid.Must(uuid.Parse("3beb9c22-01b4-4ac5-b1d9-56035d84b10b"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("139c07f7-541e-49aa-8626-240c20860b30")),
			Amount:    3600,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("5dec08fe-3079-410d-b5cd-5e17194869d7"))}, Card{CardId: uuid.Must(uuid.Parse("23c37413-b575-473a-8005-843a8af5ea21"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("3c04c2fd-72b4-42ff-a395-fa987bbfccbd")),
			Amount:    900,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("6b79133d-48fb-4eb3-992d-00ace89df9c3"))}, Card{CardId: uuid.Must(uuid.Parse("eff7b5d3-acf6-415a-8b53-8bdeb2499f26"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("2fac8a4d-fd3d-4ff8-bff0-2565d8e91033")),
			Amount:    420,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("1f36f8ec-d2eb-4725-9e08-45bcd48b013c"))}, Card{CardId: uuid.Must(uuid.Parse("81a158d7-779f-4e81-9383-9483e4877dcc"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("bb8b4944-827c-4e2e-b78e-b9ef3d39eab7")),
			Amount:    4350,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("c85b170d-38ef-4e73-b149-30a9683148cc"))}, Card{CardId: uuid.Must(uuid.Parse("68cf8ec1-0af4-4d8a-9946-baa5dea234bd"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("dd8fe886-f3de-46be-9743-5343985e007d")),
			Amount:    30,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("945405d3-fc1f-4ebe-9358-54550db72c1d"))}, Card{CardId: uuid.Must(uuid.Parse("9f812212-9e42-43c6-ba08-8a2fe4ba90a5"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("e93dee0d-37db-4e5d-a757-fb7f330cf2c0")),
			Amount:    400,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("f9d2d85a-3b08-40d0-86a4-5ce6b4234e7f"))}, Card{CardId: uuid.Must(uuid.Parse("34faa9be-e6c3-4445-9972-aea5036497e1"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("123fe85e-097e-4e6c-8fdd-9328671dcdca")),
			Amount:    100,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("5ac9ee77-c744-4e7c-b297-7e2002c260d5"))}, Card{CardId: uuid.Must(uuid.Parse("41a50d3d-28ad-4612-97dc-e60edba9ddd3"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("b35eec43-f6b3-4acc-85d5-029811d528c1")),
			Amount:    2300,
			Currency:  "EUR",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("0544de79-6d53-4b2a-a793-99fdd75e1579"))}, Card{CardId: uuid.Must(uuid.Parse("7bcc9922-525c-4be7-b22e-290bf82984c0"))}},
		},
		{
			AccountId: uuid.Must(uuid.Parse("40f46ecb-8d9e-40ec-a83d-9b815c421316")),
			Amount:    3400,
			Currency:  "RSD",
			Cards:     []Card{Card{CardId: uuid.Must(uuid.Parse("5168e0e1-52e2-4f6b-a0a4-1e370e50a36f"))}, Card{CardId: uuid.Must(uuid.Parse("fc4ab8c6-5ae6-4e2b-904a-aa72a440dea5"))}},
		},
	}

	users := []User{
		{
			UserId:   uuid.Must(uuid.Parse("0b6e2461-430b-4efe-bfa9-33eb92bde240")),
			Name:     "Nikola",
			LastName: "Malinovic",
			Email:    "nmalinovic@gmail.com",
			Receipts: []Account{accounts[0], accounts[1]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("1ead71bb-c5ea-41f3-882c-cfac6d50af85")),
			Name:     "Igor",
			LastName: "Tot",
			Email:    "itot@gmail.com",
			Receipts: []Account{accounts[2], accounts[3]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("195873d6-c2ac-47cf-9b9a-37c8db61b8cb")),
			Name:     "Jelena",
			LastName: "Petrovic",
			Email:    "jpetrovic@gmail.com",
			Receipts: []Account{accounts[4], accounts[5]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("f5ca9b23-6fac-48b2-8aae-7ea8c5ed926b")),
			Name:     "Petar",
			LastName: "Djukic",
			Email:    "pdjukic@gmail.com",
			Receipts: []Account{accounts[6], accounts[7]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("e963f1e2-285d-4398-a9db-44d9895ecc52")),
			Name:     "Aleksandar",
			LastName: "Vukovic",
			Email:    "nmalinovic@gmail.com",
			Receipts: []Account{accounts[8], accounts[9]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("2716153b-9814-40bd-8bcc-a7aa8c9b5236")),
			Name:     "Nenad",
			LastName: "Obradovic",
			Email:    "nobradovic@gmail.com",
			Receipts: []Account{accounts[10], accounts[11]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("9683fac0-74a5-4146-adaa-8d76dd2e42dd")),
			Name:     "Ognjen",
			LastName: "Zalis",
			Email:    "ozalis@gmail.com",
			Receipts: []Account{accounts[12], accounts[13]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("b774b07a-a917-4a02-8d20-d2a0bac8e7a4")),
			Name:     "Milan",
			LastName: "Kovacevic",
			Email:    "mkovacevic@gmail.com",
			Receipts: []Account{accounts[14], accounts[15]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("d2642b7c-2e2d-45b9-be91-a7bfc03cad01")),
			Name:     "Ksenija",
			LastName: "Jovancevic",
			Email:    "kjovancevic@gmail.com",
			Receipts: []Account{accounts[16], accounts[17]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("fff40cff-1df7-4ba7-8081-f18f220c0e9f")),
			Name:     "Milica",
			LastName: "Simovic",
			Email:    "msimovic@gmail.com",
			Receipts: []Account{accounts[18], accounts[19]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("86fc36db-2f9a-4043-b832-9180f99d03fa")),
			Name:     "Luka",
			LastName: "Dragovic",
			Email:    "ldragovic@gmail.com",
			Receipts: []Account{accounts[20], accounts[21]},
		},
		{
			UserId:   uuid.Must(uuid.Parse("c2f18f19-37fb-4fee-9aec-dc6648866108")),
			Name:     "Vanja",
			LastName: "Tanovic",
			Email:    "vtanovic@gmail.com",
			Receipts: []Account{accounts[22], accounts[23]},
		},
	}

	banks := []Bank{
		{
			BankId:          uuid.Must(uuid.Parse("8c54fb1d-b6f8-4c33-ab5f-103c5942b94c")),
			Name:            "Banka1",
			CentralLocation: "Beograd",
			FoundingYear:    1952,
			PIB:             "123456789",
			Users:           []User{users[0], users[1], users[2]},
		},
		{
			BankId:          uuid.Must(uuid.Parse("16377c57-88ff-45db-b80c-7a02940ddf51")),
			Name:            "Banka2",
			CentralLocation: "Novi Sad",
			FoundingYear:    1963,
			PIB:             "234567891",
			Users:           []User{users[3], users[4], users[5]},
		},
		{
			BankId:          uuid.Must(uuid.Parse("0d60c79a-9231-4b18-a28d-626d7f5e3fa5")),
			Name:            "Banka3",
			CentralLocation: "Nis",
			FoundingYear:    1972,
			PIB:             "345678912",
			Users:           []User{users[6], users[7], users[8]},
		},
		{
			BankId:          uuid.Must(uuid.Parse("bf1cc056-e18a-4f1e-972d-258f0df53676")),
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
