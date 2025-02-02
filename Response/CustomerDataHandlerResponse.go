package Response

import "Vizinhos_Back_End/Entity"

type CustomerDataHandlerResponse struct {
	Orders  []Entity.Order        `json:"orders"`
	Address Entity.StoreOrAddress `json:"address"`
}
