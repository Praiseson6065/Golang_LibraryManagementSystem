package books

import (
	"LibManMicroServ/events"

	"github.com/gin-gonic/gin"
)

func Router(eventBus *events.EventBus, r *gin.Engine) {

	booksEventsQueue, responses := eventBus.Subscribe(string(events.EventBooksValidation))
	go func() {
		for event := range booksEventsQueue {
			payload := event.Payload.(events.EventCartCheckedOutPayload)
			c := event.Context

			responses <- validateCheckOutCart(c, payload)
		}
	}()

	books := r.Group("/books")
	{
		books.GET("/", getAllBooks())
		books.GET("/:id", getOneBook())
	}
}

func AdminRouter(r *gin.Engine) {
	books := r.Group("/admin/books")
	{
		books.POST("/", bookAdd())
		books.PUT("/", bookUpdate())
		books.DELETE("/:id", bookDelete())
	}
}
