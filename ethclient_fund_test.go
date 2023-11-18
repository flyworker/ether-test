package main // Use the appropriate package name

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"math/big"
	"os"
	"testing"
)

// TestConnectToTestnet tests the ability to connect to the testnet

func TestTransferEth(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	senderPrivateKey := os.Getenv("SENDER_PRIVATE_KEY")
	if senderPrivateKey == "" {
		t.Fatal("No private key found in .env file")
	}

	recipientAddress := "0x96216849c49358B10257cb55b28eA603c874b05E" // Replace with recipient's address
	amount := big.NewInt(100000000000)                               // 0.0000001 ETH in Wei

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(senderPrivateKey)
	if err != nil {
		t.Fatalf("Failed to parse private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Fatalf("Failed to get nonce: %v", err)
	}

	gasLimit := uint64(21000) // Standard limit for a transfer
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatalf("Failed to suggest gas price: %v", err)
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(recipientAddress), amount, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		t.Fatalf("Failed to get network ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatalf("Failed to send transaction: %v", err)
	}

	t.Logf("Transaction sent: %s", signedTx.Hash().Hex())
}

func TestWriteMessageToContract(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1)) // Use the correct chain ID
	if err != nil {
		t.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contractAddress := common.HexToAddress("0xBEE4684EA19D09ae89038B09381c548e3202AaaA")
	msgContract, err := contract.NewMessageContract(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to instantiate the contract: %v", err)
	}

	tx, err := msgContract.WriteMessage(auth, "Hello, Ethereum!")
	if err != nil {
		t.Fatalf("Failed to send transaction: %v", err)
	}

	t.Logf("Transaction sent! Tx Hash: %s", tx.Hash().Hex())
}
