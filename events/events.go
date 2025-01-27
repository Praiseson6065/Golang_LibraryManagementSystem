package events

import (
	"github.com/gin-gonic/gin"
)

type EventType string

const (
	EventUserSignedUp          EventType = "user.signed.up"
	EventUserLoggedIn          EventType = "user.logged.in"
	EventUserLoggedOut         EventType = "user.logged.out"
	EventCartCheckedOut        EventType = "cart.checked.out"
	EventBookValidationSuccess EventType = "book.validation.success"
	EventBookValidationFailed  EventType = "book.validation.failed"
	EventPaymentMade           EventType = "payment.made"
	EventBookLent              EventType = "book.lent"
	EventBookReturned          EventType = "book.returned"
	EventReviewAdded           EventType = "review.added"
)

type Event struct {
	Type    EventType
	Context *gin.Context
	Payload interface{}
}
