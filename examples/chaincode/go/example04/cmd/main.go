/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"

	"github.com/hyperledger/fabric-gm/core/chaincode/shim"
	"github.com/hyperledger/fabric-gm/examples/chaincode/go/example04"
)

func main() {
	err := shim.Start(new(example04.SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
