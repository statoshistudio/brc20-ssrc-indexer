package evm

// for demo

func main() {
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("error casting public key to ECDSA")
	// }

	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// auth := bind.NewKeyedTransactor(privateKey)
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)     // in wei
	// auth.GasLimit = uint64(300000) // in units
	// auth.GasPrice = gasPrice

	// address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	// instance, err := stake.NewAbi(address, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// key := [32]byte{}
	// value := [32]byte{}
	// copy(key[:], []byte("foo"))
	// copy(value[:], []byte("bar"))

	// tx, err := instance.Stake(auth, address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("tx sent: %s", tx.Hash().Hex())

	// result, err := instance.Items(&bind.CallOpts{}, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(result[:])) // "bar"
}
