package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"math/big"
	"crypto/ecdsa"

	"github.com/ev3go/ev3dev"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("########################################################################")
	//-->Create an IPC based RPC connection to a remote node (here our RPI)
	conn, err := ethclient.Dial("http://192.168.1.38:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	} else {
		fmt.Println("Sucess! you are connected to the Ethereum Network")
	}

	//-->Print block header
	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	} else {
		fmt.Println("The block header is:",header.Number.String())
	}

	//-->Get the contract address from hex string
	addr := common.HexToAddress("0x72f8a3ad080ed0101af41e30b381d7b004b7adea")

	//-->Bind to an already deployed contract and print coolNumber
	speedMonitor, err := NewEV3SpeedMonitor(addr, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate coolNumber contract: %v", err)
	} else {
		speed, _ := speedMonitor.Speed(&bind.CallOpts{})
		//device, _ := monitor.DeviceId(&bind.CallOpts{})
		fmt.Println("########")
		fmt.Println("The last speed is:", speed)
		//fmt.Println("The EV3 device id is:", device)
	}

	//-->Bind to an account
	privateKey, err := crypto.HexToECDSA("2c88557feb65fa683290492e7a91b67dae8876260dba0972f98b4a5578c942d9")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("########")
	//fmt.Println("Address of account:", fromAddress)

	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pending Nonce:", nonce)

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("SuggestGasPrice:", gasPrice)

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(3))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	//-->Write in the smart contract
	metric := make(chan uint32) //create metric channel for concurrent goroutines
	go monitorDeviceRun(metric)
	m1 := <- metric //receive value from metric channel
	fmt.Println("########")
	fmt.Println("The new speed is:", m1)
	speedMonitor.Update(auth, m1)
}

func monitorDeviceRun(metric chan uint32) {
	var result uint32
	// Get the handle for the large motor on outA.
	outA, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find motor on outA: %v", err)
	}
	// get max speed
	maxSpeed := outA.MaxSpeed()
	// set speed
	outA.SetSpeedSetpoint(50 * maxSpeed / 100)
	// set time to run, here 10seconds
	outA.SetTimeSetpoint(time.Duration(10) * time.Second)
	// start
	outA.Command("run-timed")
	// while #duration seconds have not passed
	for i := 0; i < 10; i += 5 {
		// get some metric from motor, e.g current speed
		value, _ := outA.Speed()
		result = uint32(value)
		time.Sleep(5 * time.Second)
	}
	metric <- result //send value to metric channel
}
