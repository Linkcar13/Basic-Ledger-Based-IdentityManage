package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Identity
type SmartContract struct {
	contractapi.Contract
}

// Identity describes basic details of what makes up a simple identity
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Identity struct {
	Apellidos string `json:"Apellidos"`
	Direccion string `json:"Direccion"`
	Estado string `json:"Estado"`
	Facultad string `json:"Facultad"`
	Fecha_de_Emision string `json:"Fecha_de_Emision"`
	Fecha_de_Expiracion string `json:"Fecha_de_Expiracion"`
	ID string `json:"ID"`
	Nombres string `json:"Nombres"`
	Semestre string `json:"Semestre"`
	Tipo_de_Sangre string `json:"Tipo_de_Sangre"`
	Universidad string `json:"Universidad"`
}

// InitLedger adds a base set of identities to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	identities := []Identity{
		{ID: "201720953", Nombres: "Carlos Alejandro", Apellidos: "Tester", Facultad: "Sistemas", Tipo_de_Sangre: "A+", Semestre: "9no", Universidad: "Escuela Politecnica Nacional", Fecha_de_Emision: "2024-02-12", Fecha_de_Expiracion: "2024-08-28", Estado:"Activo"},
		{ID: "201820856", Nombres: "Leonardo Mijail", Apellidos: "Tester", Facultad: "Civil", Tipo_de_Sangre: "O+", Semestre: "9no", Universidad: "Pontificia Universidad Catolica", Fecha_de_Emision: "2024-02-12", Fecha_de_Expiracion: "2024-08-28", Estado:"Activo"},
	}

	for _, identity := range identities {
		assetJSON, err := json.Marshal(identity)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(identity.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateIdentity issues a new identity to the world state with given details.
func (s *SmartContract) CreateIdentity(ctx contractapi.TransactionContextInterface, id string, nombres string, apellidos string, facultad string, tips string, semestre string, universidad string, datemi string, datexp string, estado string) error {
	exists, err := s.IdentityExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	identity := Identity{
		ID:		id, 
		Nombres:		nombres, 
		Apellidos:		apellidos, 
		Facultad:		facultad, 
		Tipo_de_Sangre:		tips, 
		Semestre:		semestre, 
		Universidad:		universidad, 
		Fecha_de_Emision:		datemi, 
		Fecha_de_Expiracion:		datexp, 
		Estado:		estado,
	}
	identityJSON, err := json.Marshal(identity)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(id, identityJSON)
  	if err != nil {
    	return err
  	}
  return ctx.GetStub().SetEvent("IdentityIssued", identityJSON)
}

// ReadIdentity returns the asset stored in the world state with given id.
func (s *SmartContract) ReadIdentity(ctx contractapi.TransactionContextInterface, id string) (*Identity, error) {
	identityJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read identity: %v", err)
	}
	if identityJSON == nil {
		return nil, fmt.Errorf("the identity %s does not exist", id)
	}

	var identity Identity
	err = json.Unmarshal(identityJSON, &identity)
	if err != nil {
		return nil, err
	}

	return &identity, nil
}

/* // UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
} */

// DeleteIdentity deletes an given identity from the world state.
func (s *SmartContract) DeleteIdentity(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.IdentityExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the identity %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// IdentityExists returns true when asset with given ID exists in world state
func (s *SmartContract) IdentityExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	identityJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read identity from world state: %v", err)
	}

	return identityJSON != nil, nil
}

/* // VerifiyIdentity updates the owner field of asset with given id in world state, and returns the old owner.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) (string, error) {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return "", err
	}

	oldOwner := asset.Owner
	asset.Owner = newOwner

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return oldOwner, nil
} */

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllIdentities(ctx contractapi.TransactionContextInterface) ([]*Identity, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var identities []*Identity
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var identity Identity
		err = json.Unmarshal(queryResponse.Value, &identity)
		if err != nil {
			return nil, err
		}
		identities = append(identities, &identity)
	}

	return identities, nil
}
