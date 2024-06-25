/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"local/Validate-Identity/chaincode"
)

func main() {
	identityChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating identity-validate-basic chaincode: %v", err)
	}

	if err := identityChaincode.Start(); err != nil {
		log.Panicf("Error starting identity-validate-basic chaincode: %v", err)
	}
}
