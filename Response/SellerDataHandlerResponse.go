package Response

import "Vizinhos_Back_End/Entity"

type SellerDataHandlerResponse struct {
	StoreAddress Entity.StoreOrAddress `json:"storeAddress"`
	Products     []Entity.Product      `json:"products"`
}
