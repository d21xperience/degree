package utils

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DeployResult struct {
	Address common.Address
	TxHash  common.Hash
}

func DeployFromABIBytecode(abiJSON, bytecodeHex string) (*DeployResult, error) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return nil, err
	}

	privKey, _ := crypto.HexToECDSA(os.Getenv("SUPERADMIN_PRIVATE_KEY"))
	pubKeyInterface := privKey.Public()
	pubKey, ok := pubKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Gagal konversi public key")
	}
	from := crypto.PubkeyToAddress(*pubKey)

	chainID, _ := client.NetworkID(context.Background())
	nonce, _ := client.PendingNonceAt(context.Background(), from)
	gasPrice, _ := client.SuggestGasPrice(context.Background())

	auth, _ := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 5_000_00
	auth.GasPrice = gasPrice

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	bytecode := common.FromHex(bytecodeHex)

	addr, tx, _, err := bind.DeployContract(auth, parsedABI, bytecode, client)
	if err != nil {
		return nil, err
	}

	return &DeployResult{Address: addr, TxHash: tx.Hash()}, nil
}
