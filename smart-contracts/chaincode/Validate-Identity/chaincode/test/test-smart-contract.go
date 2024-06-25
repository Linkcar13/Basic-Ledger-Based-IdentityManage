package chaincode_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	"local/Validate-Identity/chaincode"
	"local/Validate-Identity/chaincode/mocks"
	"github.com/stretchr/testify/require"
)

//go:generate counterfeiter -o mocks/transaction.go -fake-name TransactionContext . transactionContext
type transactionContext interface {
	contractapi.TransactionContextInterface
}

//go:generate counterfeiter -o mocks/chaincodestub.go -fake-name ChaincodeStub . chaincodeStub
type chaincodeStub interface {
	shim.ChaincodeStubInterface
}

//go:generate counterfeiter -o mocks/statequeryiterator.go -fake-name StateQueryIterator . stateQueryIterator
type stateQueryIterator interface {
	shim.StateQueryIteratorInterface
}

func TestInitLedger(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	identityManage := chaincode.SmartContract{}
	err := identityManage.InitLedger(transactionContext)
	require.NoError(t, err)

	chaincodeStub.PutStateReturns(fmt.Errorf("failed inserting key"))
	err = identityManage.InitLedger(transactionContext)
	require.EqualError(t, err, "failed to put to world state. failed inserting key")
}

func TestCreateIdentity(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	identityManage := chaincode.SmartContract{}
	err := identityManage.CreateIdentity(transactionContext, "", "" , "", "" , "", "", "", "", "", "")
	require.NoError(t, err)

	chaincodeStub.GetStateReturns([]byte{}, nil)
	err = identityManage.CreateIdentity(transactionContext, "201720953", "" , "", "" , "", "", "", "", "", "")
	require.EqualError(t, err, "the identity already exists")

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve identity"))
	err = identityManage.CreateIdentity(transactionContext, "201720953", "" , "", "" , "", "", "", "", "", "")
	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")
}

func TestReadIdentity(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	expectedAsset := &chaincode.Identity{ID: "201720953"}
	bytes, err := json.Marshal(expectedAsset)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	identityManage := chaincode.SmartContract{}
	identity, err := identityManage.ReadIdentity(transactionContext, "")
	require.NoError(t, err)
	require.Equal(t, expectedAsset, identity)

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve identity"))
	_, err = identityManage.ReadIdentity(transactionContext, "")
	require.EqualError(t, err, "failed to read from world state: unable to retrieve identity")

	chaincodeStub.GetStateReturns(nil, nil)
	identity, err = identityManage.ReadIdentity(transactionContext, "201720953")
	require.EqualError(t, err, "the identity does not exist")
	require.Nil(t, identity)
}

/* func TestUpdateAsset(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	expectedAsset := &chaincode.Asset{ID: "asset1"}
	bytes, err := json.Marshal(expectedAsset)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	identityManage := chaincode.SmartContract{}
	err = identityManage.UpdateAsset(transactionContext, "", "", 0, "", 0)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(nil, nil)
	err = identityManage.UpdateAsset(transactionContext, "asset1", "", 0, "", 0)
	require.EqualError(t, err, "the asset asset1 does not exist")

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
	err = identityManage.UpdateAsset(transactionContext, "asset1", "", 0, "", 0)
	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")
} */

func TestDeleteIdentity(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	identity := &chaincode.Identity{ID: "201720953"}
	bytes, err := json.Marshal(identity)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	chaincodeStub.DelStateReturns(nil)
	identityManage := chaincode.SmartContract{}
	err = identityManage.DeleteIdentity(transactionContext, "")
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(nil, nil)
	err = identityManage.DeleteIdentity(transactionContext, "201720953")
	require.EqualError(t, err, "the identity provide does not exist")

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve identity"))
	err = identityManage.DeleteIdentity(transactionContext, "")
	require.EqualError(t, err, "failed to read from world state: unable to retrieve identity")
}

/* func TestTransferAsset(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	asset := &chaincode.Asset{ID: "asset1"}
	bytes, err := json.Marshal(asset)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	identityManage := chaincode.SmartContract{}
	_, err = identityManage.TransferAsset(transactionContext, "", "")
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
	_, err = identityManage.TransferAsset(transactionContext, "", "")
	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")
} */

func TestGetAllIdentities(t *testing.T) {
	identity := &chaincode.Identity{ID: "201720953"}
	bytes, err := json.Marshal(identity)
	require.NoError(t, err)

	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturnsOnCall(0, true)
	iterator.HasNextReturnsOnCall(1, false)
	iterator.NextReturns(&queryresult.KV{Value: bytes}, nil)

	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateByRangeReturns(iterator, nil)
	identityManage := &chaincode.SmartContract{}
	identities, err := identityManage.GetAllIdentities(transactionContext)
	require.NoError(t, err)
	require.Equal(t, []*chaincode.Identity{identity}, identities)

	iterator.HasNextReturns(true)
	iterator.NextReturns(nil, fmt.Errorf("failed retrieving next item"))
	identities, err = identityManage.GetAllIdentities(transactionContext)
	require.EqualError(t, err, "failed retrieving next item")
	require.Nil(t, identities)

	chaincodeStub.GetStateByRangeReturns(nil, fmt.Errorf("failed retrieving all identities"))
	identities, err = identityManage.GetAllIdentities(transactionContext)
	require.EqualError(t, err, "failed retrieving all identities")
	require.Nil(t, identities)
}
