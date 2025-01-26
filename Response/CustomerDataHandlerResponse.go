package Response

import "Vizinhos_Back_End/Entity"

type CustomerDataHandlerResponse struct {
	Orders    []Entity.Order          `json:"orders"`
	Addresses []Entity.StoreOrAddress `json:"addresses"`
}
