package events

type EventUserCartCreationResponse struct {
	Error error `json:"error"`
}

type EventCartCheckedOutResponse struct {
	LendCartItem     []BookForPayload `json:"lendItems"`
	PurchaseCartItem []BookForPayload `json:"purchaseItems"`
	Success          bool             `json:"success"`
	Error            error            `json:"error"`
}
