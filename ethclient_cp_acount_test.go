package main // Use the appropriate package name

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flyworker/ether-test/ubi"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"
)

func TestChangeBeneficiary(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	defer client.Close()

	senderPrivateKey := os.Getenv("SENDER_PRIVATE_KEY")
	if senderPrivateKey == "" {
		t.Fatal("No private key found in .env file")
	}

	privateKey, err := crypto.HexToECDSA(senderPrivateKey)
	if err != nil {
		t.Fatalf("Failed to parse private key: %v", err)
	}

	networkID, err := client.NetworkID(ctx)
	if err != nil {
		t.Fatalf("Failed to get network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	if err != nil {
		t.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contractAddress := common.HexToAddress(os.Getenv("CP_ACCOUNT_CONTRACT_ADDRESS"))
	if contractAddress.Hex() == "0x0000000000000000000000000000000000000000" {
		t.Fatal("Invalid or no contract address provided")
	}

	cpAccount, err := ubi.NewCpaccount(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to instantiate the Cpaccount contract: %v", err)
	}

	// Example: Change beneficiary
	newBeneficiary := common.HexToAddress("0xE53AEd6DEA9e44116D4551a93eEeE28bC8684916") // Replace with actual address
	newQuota := big.NewInt(1000)                                                        // Example quota
	newExpiration := big.NewInt(12345678)                                               // Example expiration

	tx, err := cpAccount.ChangeBeneficiary(auth, newBeneficiary, newQuota, newExpiration)
	if err != nil {
		t.Fatalf("Failed to execute ChangeBeneficiary: %v", err)
	}

	t.Logf("ChangeBeneficiary transaction sent! Tx Hash: %s", tx.Hash().Hex())
}

func TestReadBeneficiary(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(os.Getenv("CP_ACCOUNT_CONTRACT_ADDRESS"))
	if contractAddress.Hex() == "0x0000000000000000000000000000000000000000" {
		t.Fatal("Invalid or no contract address provided")
	}

	cpAccount, err := ubi.NewCpaccount(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to instantiate the Cpaccount contract: %v", err)
	}

	beneficiaryDetails, err := cpAccount.Beneficiary(&bind.CallOpts{})
	if err != nil {
		t.Fatalf("Failed to retrieve beneficiary details: %v", err)
	}

	t.Logf("Beneficiary Address: %s", beneficiaryDetails.BeneficiaryAddress.Hex())
	t.Logf("Quota: %s", beneficiaryDetails.Quota.String())
	t.Logf("Expiration: %s", beneficiaryDetails.Expiration.String())
}

func TestSubmitUBIProof(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	defer client.Close()

	senderPrivateKey := os.Getenv("SENDER_PRIVATE_KEY")
	if senderPrivateKey == "" {
		t.Fatal("No private key found in .env file")
	}

	privateKey, err := crypto.HexToECDSA(senderPrivateKey)
	if err != nil {
		t.Fatalf("Failed to parse private key: %v", err)
	}

	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		t.Fatalf("Failed to get network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	if err != nil {
		t.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contractAddress := common.HexToAddress(os.Getenv("CP_ACCOUNT_CONTRACT_ADDRESS"))
	if contractAddress.Hex() == "0x0000000000000000000000000000000000000000" {
		t.Fatal("Invalid or no contract address provided")
	}

	cpAccount, err := ubi.NewCpaccount(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to instantiate the Cpaccount contract: %v", err)
	}

	// Example values for submitUBIProof - replace these with actual values
	taskId := "exampleTaskId"
	taskType := uint8(0) // Example task type
	proof := "exampleProofString"

	tx, err := cpAccount.SubmitUBIProof(auth, taskId, taskType, proof)
	if err != nil {
		t.Fatalf("Failed to execute submitUBIProof: %v", err)
	}

	// Optionally, wait for the transaction to be confirmed and then perform additional checks
	// ...

	t.Logf("submitUBIProof transaction sent! Tx Hash: %s", tx.Hash().Hex())
}

func TestReadTransactionInput(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	defer client.Close()

	// Use the transaction hash of the submitUBIProof transaction
	txHash := common.HexToHash("0x294fbddde7a3c0162a22c51e6bd1fb91565f16a660713c3df4e60023d1346ea6")

	// Retrieve the transaction by its hash
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		t.Fatalf("Failed to retrieve transaction: %v", err)
	}
	if isPending {
		t.Log("Transaction is still pending")
	}

	inputData := tx.Data()
	methodID := hex.EncodeToString(inputData[:4])
	log.Printf("Method ID: %s", methodID)

	// Read the ABI from a file
	abiFile, err := ioutil.ReadFile("ubi/CPAccount.abi") // Replace with the actual path
	if err != nil {
		t.Fatalf("Failed to read ABI file: %v", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		t.Fatalf("Failed to parse ABI: %v", err)
	}

	method, exists := parsedABI.Methods["submitUBIProof"]
	if !exists {
		t.Fatal("Method submitUBIProof not found")
	}

	decodedData, err := method.Inputs.Unpack(inputData[4:])
	if err != nil {
		t.Fatalf("Failed to decode input data: %v", err)
	}

	for i, input := range decodedData {
		t.Logf("Input %d: %v", i, input)
	}
}
