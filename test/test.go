package test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"nft-hook-sdk-go/util"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"

	contract "nft-hook-sdk-go/contract/eth"
)

var password = "123456"

func TestGenerateNewAddress(t *testing.T) {
	addr, pk, _ := util.GenerateNewAddress()

	fmt.Println(" addr: ", addr)
	fmt.Println(" pk: ", pk)
}

func TestGenerateKeyfile(t *testing.T) {

	addr, keyfile, _ := util.GenerateKeyfile(password, "./keystore")

	fmt.Println(" addr: ", addr)
	fmt.Println(" keyfile: ", keyfile)

	keyJson, err := os.ReadFile(keyfile)
	if err != nil {
		log.Fatalf("Failed to read keystore file: %v", err)
	}

	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		log.Fatalf("Failed to decrypt key: %v", err)
	}

	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	fmt.Printf("Private key: %s\n", privateKeyHex)
	fmt.Printf("Address: %s\n", key.Address.Hex())
}

func TestDeployContract(t *testing.T) {

	data, err := os.ReadFile("secrets.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var secrets map[string]string
	err = json.Unmarshal(data, &secrets)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	sdk, err := contract.NewContractSDK("https://sepolia.infura.io/v3/"+secrets["INFURA_KEY"], "../keystore/UTC--2024-08-07T09-45-30.086279000Z--ed240f3831f6bce7b892cfa53d0d84cd27f2db6d", password, big.NewInt(1))
	if err != nil {
		log.Fatalf("Failed to create SDK: %v", err)
	}

	abiPath := "path/to/your/contract.abi"
	binPath := "path/to/your/contract.bin"

	constructorArgs := []interface{}{
		"MyNFT", // name
		"MNFT",  // symbol
	}

	address, tx, _, err := sdk.DeployContract(abiPath, binPath, constructorArgs...)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed! Address: %s\nTransaction hash: %s\n", address.Hex(), tx.Hash().Hex())
}
