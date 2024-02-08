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

func checkLoggedIn(r *http.Request) (string, error) {
	var tokenStr string
	fmt.Println("Got auth: ", r.Header["Authorization"])
	if tokenArr, ok := r.Header["Authorization"]; !ok {
		return "", fmt.Errorf("token isn't provided")
	} else {
		tokenStr = strings.Split(tokenArr[0], " ")[1]
	}
	token := strings.Split(tokenStr, ".")
	if len(token) != 3 {
		return "", fmt.Errorf("bad token")
	}
	mac := hmac.New(sha256.New, SECRET)
	mac.Write([]byte(fmt.Sprintf("%s.%s", token[0], token[1])))
	sig, _ := base64.RawURLEncoding.DecodeString(token[2])
	if !hmac.Equal(sig, mac.Sum(nil)) {
		return "", fmt.Errorf("bad hash")
	}
	claimsByte, err := base64.RawURLEncoding.DecodeString(token[1])
	if err != nil {
		return "", err
	}

	var claims Claims
	err = json.Unmarshal(claimsByte, &claims)
	if err != nil {
		return "", err
	}

	return claims["username"].(string), nil
}

type loginResponse struct {
	Token string `json:"token"`
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
		username := r.Form.Get("username")
		fmt.Println("username:", username)
		// password := r.Form.Get("password")
		// TODO check exists
		// Encode header and claims
		var c Claims = map[string]interface{}{}
		c["username"] = username
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

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
