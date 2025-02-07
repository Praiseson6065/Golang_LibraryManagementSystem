package events

type EventUserSignedUpResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}
