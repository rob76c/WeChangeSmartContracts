package main

import (
	"/chaincode"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {

	transactionLedgerChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})

	if err != nil {
		log.Panicf("Error creating transactionLedger chaincode: %v", err)
	}

	if err := transactionLedgerChaincode.Start(); err != nil {
		log.Panicf("Error starting transactionLedger chaincode: %v", err)
	}
}
