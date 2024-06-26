/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"local/Identity-Magnament/chaincode"
)

func main() {
	identityChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating identity-manage-basic chaincode: %v", err)
	}

	if err := identityChaincode.Start(); err != nil {
		log.Panicf("Error starting identity-manage-basic chaincode: %v", err)
	}
}
