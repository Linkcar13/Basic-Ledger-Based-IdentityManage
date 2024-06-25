package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

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



/* // CreateIdentity issues a new identity to the world state with given details.
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
} */

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

// ValidateIdentity validates the provided identity's expiration date and updates its state
// ValidateIdentity validates the provided identity's expiration date and updates its state
func (s *SmartContract) ValidateIdentity(ctx contractapi.TransactionContextInterface, id string) (string, error) {
    // Leer la identidad existente del estado mundial
    storedIdentity, err := s.ReadIdentity(ctx, id)
    if err != nil {
        return "", fmt.Errorf("failed to read identity: %v", err)
    }

    // Parsear la fecha de expiración
    expiryDate, err := time.Parse("2006-01-02", storedIdentity.Fecha_de_Expiracion) // Assuming the date is in YYYY-MM-DD format
    if err != nil {
        return "", fmt.Errorf("failed to parse expiration date: %v", err)
    }

    // Obtener la fecha actual
    currentDate := time.Now()

    // Validar la fecha de expiración
    var statusMessage string
    if currentDate.After(expiryDate) {
        storedIdentity.Estado = "Inactivo"
        statusMessage = "La identidad es inválida"
    } else {
        storedIdentity.Estado = "Activo"
        statusMessage = "La identidad es válida"
    }

    // Actualizar la identidad en el estado mundial
    updatedIdentityJSON, err := json.Marshal(storedIdentity)
    if err != nil {
        return "", fmt.Errorf("failed to marshal updated identity: %v", err)
    }

    err = ctx.GetStub().PutState(id, updatedIdentityJSON)
    if err != nil {
        return "", fmt.Errorf("failed to update identity state: %v", err)
    }

    err = ctx.GetStub().SetEvent("IdentityStatusValidated", []byte(statusMessage))
    if err != nil {
        return "", fmt.Errorf("failed to set event: %v", err)
    }

    return statusMessage, nil
}


/* func (s *SmartContract) DeleteIdentity(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.IdentityExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the identity %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
} */

/* // IdentityExists returns true when asset with given ID exists in world state
func (s *SmartContract) IdentityExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	identityJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read identity from world state: %v", err)
	}

	return identityJSON != nil, nil
}



// GetAllAssets returns all identities found in world state
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
} */
