package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var (
	SECRET = []byte("my_secret")
)

const (
	CHANNEL        = "mychannel"
	USERNAME_CLAIM = "username"
	BANK_CLAIM     = "bank"
)

func Json(jsonData []byte) (string, error) {
	var prettyJSON bytes.Buffer

	err := json.Indent(&prettyJSON, jsonData, "", "\t")
	if err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}

func PopulateWallet(wallet *gateway.Wallet, org string) error {
	basePath := filepath.Join(
		"..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		org+".example.com",
		"users",
		"User1@"+org+".example.com",
		"msp",
	)

	certPath := filepath.Join(basePath, "signcerts", "cert.pem")
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))

	if err != nil {
		return err
	}

	keystorePath := filepath.Join(basePath, "keystore")
	files, err := ioutil.ReadDir(keystorePath)

	if err != nil {
		return err
	}

	if len(files) != 1 {
		return errors.New("there must be a secret key file in the keystore path")
	}

	keyPath := filepath.Join(keystorePath, files[0].Name())
	key, err := ioutil.ReadFile(keyPath)

	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(strings.Title(org+"MSP"), string(cert), string(key))

	err = wallet.Put("User1@"+org+".example.com", identity)

	if err != nil {
		return err
	}

	return nil
}

func ValidateOrPopulateWallet(wallet *gateway.Wallet, org string) error {
	if !wallet.Exists("User1@" + org + ".example.com") {
		err := PopulateWallet(wallet, org)
		if err != nil {
			return err
		}
	}

	log.Println("Validated/Populated the wallet.")
	return nil
}

func GetGateway(wallet *gateway.Wallet, org string) (*gateway.Gateway, error) {
	connectionPath := filepath.Join(
		"..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		org+".example.com",
		"connection-"+org+".json",
	)

	gateway, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(connectionPath))),
		gateway.WithIdentity(wallet, "User1@"+org+".example.com"),
	)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the gateway.")
	return gateway, nil
}

func GetContract(wallet *gateway.Wallet, chaincodeName string, org string, channel string) (*gateway.Contract, error) {
	err := ValidateOrPopulateWallet(wallet, org)
	if err != nil {
		return nil, err
	}
	gateway, err := GetGateway(wallet, org)
	if err != nil {
		return nil, err
	}
	defer gateway.Close()

	network, err := gateway.GetNetwork(channel)
	log.Println("Got network.")

	if err != nil {
		return nil, err
	}

	contract := network.GetContract(chaincodeName)
	log.Println("Got contract.")

	return contract, nil
}

type Claims map[string]interface{}

func (c Claims) Get(item string) string {
	if _, ok := c[item]; !ok {
		return ""
	}
	return c[item].(string)
}

func checkLoggedIn(r *http.Request) (Claims, error) {
	var tokenStr string
	fmt.Println("Got auth: ", r.Header["Authorization"])
	if tokenArr, ok := r.Header["Authorization"]; !ok {
		return nil, fmt.Errorf("token isn't provided")
	} else {
		tokenStr = strings.Split(tokenArr[0], " ")[1]
	}
	token := strings.Split(tokenStr, ".")
	if len(token) != 3 {
		return nil, fmt.Errorf("bad token")
	}
	mac := hmac.New(sha256.New, SECRET)
	mac.Write([]byte(fmt.Sprintf("%s.%s", token[0], token[1])))
	sig, _ := base64.RawURLEncoding.DecodeString(token[2])
	if !hmac.Equal(sig, mac.Sum(nil)) {
		return nil, fmt.Errorf("bad hash")
	}
	claimsByte, err := base64.RawURLEncoding.DecodeString(token[1])
	if err != nil {
		return nil, err
	}

	var claims Claims
	err = json.Unmarshal(claimsByte, &claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

type loginResponse struct {
	Token string `json:"token"`
}

func getOrgFromBank(bank string) string {
	switch bank {
	case "bank1":
		return "org1"
	case "bank2":
		return "org2"
	case "bank3":
		return "org3"
	default:
		return "org4"
	}
}

type User struct {
	UserId   string
	Password string
	Name     string
	LastName string
	Email    string
	Receipts []string
}

type ProfileAccount struct {
	Id       string  `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type Profile struct {
	Id       string           `json:"id"`
	Email    string           `json:"email"`
	Accounts []ProfileAccount `json:"accounts"`
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

func main() {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}
	chaincodeName := "basic"

	http.HandleFunc("/api/checkhealth", func(w http.ResponseWriter, r *http.Request) {
		username, err := checkLoggedIn(r)
		if err != nil {
			fmt.Fprintf(w, "Hello, World Random user")
		} else {
			fmt.Fprintf(w, "Hello, World %s", username)
		}

	})

	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		bankId := r.Form.Get("bank")
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if bankId == "" || username == "" || password == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You have to supply bank, username and password")
			return
		}
		fmt.Printf("username: %s, pass: %s, bank: %s\n", username, password, bankId)

		org := getOrgFromBank(bankId)
		fmt.Println(org, chaincodeName)
		contract, err := GetContract(wallet, chaincodeName, org, CHANNEL)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		result, err := contract.EvaluateTransaction("GetUser", username)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, "User not found: %s", err)
			return
		}
		user := &User{}
		err = json.Unmarshal(result, user)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		if user.Password != password {
			w.WriteHeader(403)
			w.Write([]byte("Bad password"))
			return
		}

		var c Claims = map[string]interface{}{}
		c[USERNAME_CLAIM] = username
		c[BANK_CLAIM] = bankId
		headerEnc, _ := json.Marshal(map[string]string{"typ": "JWT", "alg": "HS256"})
		claimsEnc, _ := json.Marshal(c)
		jwtStr := fmt.Sprintf(
			"%s.%s",
			base64.RawURLEncoding.EncodeToString(headerEnc),
			base64.RawURLEncoding.EncodeToString(claimsEnc),
		)

		// Sign with sha 256
		mac := hmac.New(sha256.New, SECRET)
		mac.Write([]byte(jwtStr))

		token := fmt.Sprintf("%s.%s", jwtStr, base64.RawURLEncoding.EncodeToString(mac.Sum(nil)))
		res, err := json.Marshal(&loginResponse{Token: token})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(res)
	})

	http.HandleFunc("/api/initledger", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}
		channel := r.Form.Get("channel")
		org := r.Form.Get("org")
		contract, err := GetContract(wallet, chaincodeName, org, channel)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		_, err = contract.SubmitTransaction("InitLedger")
		log.Println("Transaction submitted.")
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		log.Println("Ledger initialized.")
		fmt.Fprintf(w, "Success")
	})

	http.HandleFunc("/api/get/users", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}
		channel := r.Form.Get("channel")
		org := r.Form.Get("org")

		contract, err := GetContract(wallet, chaincodeName, org, channel)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		result, err := contract.EvaluateTransaction("GetAllUsers")
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}

		json, err := Json(result)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}

		log.Println("Users queried.")
		fmt.Fprintf(w, json)
	})

	http.HandleFunc("/api/get/banks", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}
		channel := r.Form.Get("channel")
		org := r.Form.Get("org")

		contract, err := GetContract(wallet, chaincodeName, org, channel)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		result, err := contract.EvaluateTransaction("GetAllBanks")
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}

		json, err := Json(result)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}

		log.Println("Banks queried.")
		fmt.Fprintf(w, json)
	})

	http.HandleFunc("/api/get/accounts", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}
		channel := r.Form.Get("channel")
		org := r.Form.Get("org")

		contract, err := GetContract(wallet, chaincodeName, org, channel)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		result, err := contract.EvaluateTransaction("GetAllAccounts")
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}

		json, err := Json(result)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}

		log.Println("Accounts queried.")
		fmt.Fprintf(w, json)
	})

	http.HandleFunc("/api/deposit", func(w http.ResponseWriter, r *http.Request) {
		claims, err := checkLoggedIn(r)
		if err != nil {
			w.WriteHeader(403)
			fmt.Fprintf(w, "You have to be authorized")
			return
		}
		fmt.Printf("User %s is depositing money.\n", claims.Get(USERNAME_CLAIM))
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}
		accountId := r.Form.Get("accountId")
		if accountId == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You have to supply accountId param")
			return
		}
		amount := r.Form.Get("amount")
		if amount == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You have to supply amount param")
			return
		}
		contract, err := GetContract(wallet, chaincodeName, getOrgFromBank(claims.Get(BANK_CLAIM)), CHANNEL)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		username := claims.Get(USERNAME_CLAIM)
		result, err := contract.EvaluateTransaction("GetUser", username)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, "User not found: %s", err)
			return
		}
		user := &User{}
		err = json.Unmarshal(result, user)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		flag := false
		for _, account := range user.Receipts {
			if account == accountId {
				flag = true
				break
			}
		}
		if !flag {
			w.WriteHeader(403)
			fmt.Fprintf(w, "Account: %s does not belong to you, user %s!", accountId, username)
			return
		}

		_, err = contract.SubmitTransaction("DepositMoney", accountId, amount)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		fmt.Fprintf(w, "Successfully deposited money.")
	})

	http.HandleFunc("/api/withdraw", func(w http.ResponseWriter, r *http.Request) {
		claims, err := checkLoggedIn(r)
		if err != nil {
			w.WriteHeader(403)
			fmt.Fprintf(w, "You have to be authorized")
			return
		}
		fmt.Printf("User %s is depositing money.\n", claims.Get(USERNAME_CLAIM))
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}
		accountId := r.Form.Get("accountId")
		if accountId == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You have to supply accountId param")
			return
		}
		amount := r.Form.Get("amount")
		if amount == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You have to supply amount param")
			return
		}
		currency := r.Form.Get("currency")
		if currency == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You have to supply currency param")
			return
		}
		contract, err := GetContract(wallet, chaincodeName, getOrgFromBank(claims.Get(BANK_CLAIM)), CHANNEL)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		username := claims.Get(USERNAME_CLAIM)
		result, err := contract.EvaluateTransaction("GetUser", username)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, "User not found: %s", err)
			return
		}
		user := &User{}
		err = json.Unmarshal(result, user)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		flag := false
		for _, account := range user.Receipts {
			if account == accountId {
				flag = true
				break
			}
		}
		if !flag {
			w.WriteHeader(403)
			fmt.Fprintf(w, "Account: %s does not belong to you, user %s!", accountId, username)
			return
		}

		_, err = contract.SubmitTransaction("WithdrawMoney", accountId, amount, currency)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		fmt.Fprintf(w, "Successfully withdrew money.")
	})

	http.HandleFunc("/api/transfer", func(w http.ResponseWriter, r *http.Request) {
		claims, err := checkLoggedIn(r)
		if err != nil {
			w.WriteHeader(403)
			fmt.Fprintf(w, "You have to be authorized")
			return
		}
		fmt.Printf("User %s is transferring money.\n", claims.Get(USERNAME_CLAIM))
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		}
		accountFrom := r.Form.Get("accountFrom")
		if accountFrom == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You have to supply accountFrom param")
			return
		}
		accountTo := r.Form.Get("accountTo")
		if accountTo == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You have to supply accountTo param")
			return
		}
		amount := r.Form.Get("amount")
		if amount == "" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "You have to supply amount param")
			return
		}
		convert := r.Form.Get("convert")
		if convert != "true" {
			convert = "false"
		}
		contract, err := GetContract(wallet, chaincodeName, getOrgFromBank(claims.Get(BANK_CLAIM)), CHANNEL)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		username := claims.Get(USERNAME_CLAIM)
		result, err := contract.EvaluateTransaction("GetUser", username)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, "User not found: %s", err)
			return
		}
		user := &User{}
		err = json.Unmarshal(result, user)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		flagFrom := false
		flagTo := false
		for _, account := range user.Receipts {
			if account == accountFrom {
				flagFrom = true
			}
			if account == accountTo {
				flagTo = true
			}
			if flagFrom && flagTo {
				break
			}
		}
		if !flagFrom || !flagTo {
			w.WriteHeader(403)
			fmt.Fprintf(w, "Account %s or %s does not belong to you, user %s!", accountFrom, accountTo, username)
			return
		}

		_, err = contract.SubmitTransaction("TransferMoney", accountFrom, accountTo, amount, convert)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		fmt.Fprintf(w, "Successfully transferred %s from account %s to %s.", amount, accountFrom, accountTo)
	})

	http.HandleFunc("/api/profile", func(w http.ResponseWriter, r *http.Request) {
		claims, err := checkLoggedIn(r)
		if err != nil {
			w.WriteHeader(403)
			fmt.Fprintf(w, "You have to be authorized")
			return
		}
		contract, err := GetContract(wallet, chaincodeName, getOrgFromBank(claims.Get(BANK_CLAIM)), CHANNEL)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		username := claims.Get(USERNAME_CLAIM)
		result, err := contract.EvaluateTransaction("GetUser", username)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, "User not found: %s", err)
			return
		}
		user := &User{}
		err = json.Unmarshal(result, user)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		profile := &Profile{
			Id:       user.UserId,
			Email:    user.Email,
			Accounts: make([]ProfileAccount, 0, len(user.Receipts)),
		}
		for _, account := range user.Receipts {
			result, err := contract.EvaluateTransaction("GetAccount", account)
			if err != nil {
				w.WriteHeader(404)
				fmt.Fprintf(w, "User not found: %s", err)
				return
			}
			acc := &Account{}
			err = json.Unmarshal(result, acc)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				return
			}
			profile.Accounts = append(profile.Accounts, ProfileAccount{
				Id:       acc.AccountId,
				Amount:   acc.Amount,
				Currency: acc.Currency,
			})
		}
		res, err := json.MarshalIndent(profile, "", "\t")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}

		w.Write(res)
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
