package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Create an IPC based RPC connection to a remote node (here our RPI)
	conn, err := ethclient.Dial("http://192.168.1.110:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	} else {
		fmt.Println("Sucess! you are connected to the Ethereum Network")
	}

	//Get the contract address from hex string
	addr := common.HexToAddress("0xB202CCCd3b66F63f567D40d6e22C1C1BAB627824")

	// Bind to an already deployed contract and print randomNumber
	contract, err := NewRandomNumberContract(addr, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate randomNumber contract: %v", err)
	} else {
		randomNumber, _ := contract.RandomNumber(&bind.CallOpts{})
		fmt.Println("The randomNumber is:", randomNumber)
	}

	//Print block header
	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	} else {
		fmt.Println("The block header is:",header.Number.String())
	}
}
