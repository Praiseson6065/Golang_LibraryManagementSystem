package events

type EventType string

const (
	EventUserLoggedIn          EventType = "user_logged_in"
	EventUserLoggedOut         EventType = "user_logged_out"
	EventCartCheckedOut        EventType = "cart_checked_out"
	EventBookValidationSuccess EventType = "book_validation_success"
	EventBookValidationFailed  EventType = "book_validation_failed"
	EventPaymentMade           EventType = "payment_made"
	EventBookLent              EventType = "book_lent"
	EventBookReturned          EventType = "book_returned"
	EventReviewAdded           EventType = "review_added"
)

type Event struct {
	Type    EventType
	Payload interface{}
}
