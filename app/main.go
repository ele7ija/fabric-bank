package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
    "bytes"
    "encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
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
		org + ".example.com",
		"users",
		"User1@" + org + ".example.com",
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
        return errors.New("There must be a secret key file in the keystore path")
    }

    keyPath := filepath.Join(keystorePath, files[0].Name())
    key, err := ioutil.ReadFile(keyPath)

    if err != nil {
        return err
    }

    identity := gateway.NewX509Identity(strings.Title(org + "MSP"), string(cert), string(key))

    err = wallet.Put("User1@" + org + ".example.com", identity)

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
		org + ".example.com",
        "connection-" + org + ".json",
    )

    gateway, err := gateway.Connect(
        gateway.WithConfig(config.FromFile(filepath.Clean(connectionPath))),
        gateway.WithIdentity(wallet, "User1@" + org + ".example.com"),
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

func main(){
    os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
    wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}
    chaincodeName := "basic"

    http.HandleFunc("/api/checkhealth", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello, World!")
    })

    http.HandleFunc("/api/initledger", func(w http.ResponseWriter, r *http.Request){
        err := r.ParseForm()
        if err != nil {
            log.Println(err)
            fmt.Fprintf(w, "Error: %s", err)
        }
        channel := r.Form.Get("channel")
        org := r.Form.Get("org")
        contract, err := GetContract(wallet, chaincodeName, org, channel)
        if err != nil{
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

    http.HandleFunc("/api/get/users", func(w http.ResponseWriter, r *http.Request){
        err := r.ParseForm()
        if err != nil {
            log.Println(err)
            fmt.Fprintf(w, "Error: %s", err)
        }
        channel := r.Form.Get("channel")
        org := r.Form.Get("org")

        contract, err := GetContract(wallet, chaincodeName, org, channel)
        if err != nil{
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

    http.HandleFunc("/api/get/banks", func(w http.ResponseWriter, r *http.Request){
        err := r.ParseForm()
        if err != nil {
            log.Println(err)
            fmt.Fprintf(w, "Error: %s", err)
        }
        channel := r.Form.Get("channel")
        org := r.Form.Get("org")

        contract, err := GetContract(wallet, chaincodeName, org, channel)
        if err != nil{
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

    http.HandleFunc("/api/get/accounts", func(w http.ResponseWriter, r *http.Request){
        err := r.ParseForm()
        if err != nil {
            log.Println(err)
            fmt.Fprintf(w, "Error: %s", err)
        }
        channel := r.Form.Get("channel")
        org := r.Form.Get("org")

        contract, err := GetContract(wallet, chaincodeName, org, channel)
        if err != nil{
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

