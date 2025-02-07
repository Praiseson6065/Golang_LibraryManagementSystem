package events

type BookForPayload struct {
	BookId    string `json:"bookId"`
	Quantity  int    `json:"quantity"`
	Available bool   `json:"avaliable"`
}

type EventUserSignedUpPayload struct {
	UserId string `json:"userId"`
}

type EventCartCheckedOutPayload struct {
	LendCartItem     []BookForPayload `json:"lendItems"`
	PurchaseCartItem []BookForPayload `json:"purchaseItems"`
}
