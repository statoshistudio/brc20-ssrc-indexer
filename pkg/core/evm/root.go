package evm

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	registry "github.com/ByteGum/go-icms/pkg/core/chain/evm/abis/registry"
	stake "github.com/ByteGum/go-icms/pkg/core/chain/evm/abis/stake"
	token "github.com/ByteGum/go-icms/pkg/core/chain/evm/abis/token"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ToHexAddress(address string) common.Address {
	return common.HexToAddress(address)
}
func StakeContract(rpc_url string, contractAddress string) (*stake.Stake, *ethclient.Client, common.Address, error) {
	address := common.HexToAddress(contractAddress)
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, nil, address, err
	}

	instance, err := stake.NewStake(address, client)
	if err != nil {
		return nil, nil, address, err
	}
	return instance, client, address, err
}

func TokenContract(rpc_url string, contractAddress string) (*token.IcmToken, error) {
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)

	instance, err := token.NewIcmToken(address, client)
	if err != nil {
		return nil, err
	}
	return instance, err
}

func RegistryContract(rpc_url string, contractAddress string) (*registry.Abi, error) {
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)

	instance, err := registry.NewAbi(address, client)
	if err != nil {
		return nil, err
	}
	return instance, err
}

func AuthOption(privateKey string, client ethclient.Client) *bind.TransactOpts {
	pKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := pKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(pKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth
}
