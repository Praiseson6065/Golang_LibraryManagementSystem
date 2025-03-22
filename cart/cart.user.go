package cart

import (
	"LibManMicroServ/events"
	"LibManMicroServ/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckOutCart
// @Summary Check Out Cart
// @Description CheckOutCart
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /checkout [post]
func CheckOutCart(eventBus *events.EventBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := middleware.GetUserID(ctx)
		getCart, err := GetActiveCart(ctx, userId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		responses, err := EventValidateBooks(ctx, eventBus, getCart.LendItems, getCart.PurchaseItems)

		if len(responses) > 0 {
			_, err = responses[0].(events.EventCartCheckedOutResponse).Success, responses[0].(events.EventCartCheckedOutResponse).Error
			lendCartItem := responses[0].(events.EventCartCheckedOutResponse).LendCartItem
			purchaseCartItem := responses[0].(events.EventCartCheckedOutResponse).PurchaseCartItem
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"data": gin.H{
						"lendCartItem":     lendCartItem,
						"purchaseCartItem": purchaseCartItem,
					},
				})
				return
			}

		}

	}
}

// AddItemToCart godoc
// @Summary Add Item To Cart
// @Description AddItemToCart
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param item body LendCartItem true "Item to add"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /addItemToCart [post]
func AddItemToCart(eventBus *events.EventBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := middleware.GetUserID(ctx)

		var lendItem LendCartItem
		var purchaseItem PurchaseCartItem

		if err := ctx.ShouldBindBodyWithJSON(&lendItem); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := ctx.ShouldBindBodyWithJSON(&purchaseItem); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		responses, err := EventValidateBooks(ctx, eventBus, []LendCartItem{lendItem}, []PurchaseCartItem{purchaseItem})
		if len(responses) > 0 {
			_, err = responses[0].(events.EventCartCheckedOutResponse).Success, responses[0].(events.EventCartCheckedOutResponse).Error
			lendCartItem := responses[0].(events.EventCartCheckedOutResponse).LendCartItem
			purchaseCartItem := responses[0].(events.EventCartCheckedOutResponse).PurchaseCartItem
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"data": gin.H{
						"lendCartItem":     lendCartItem,
						"purchaseCartItem": purchaseCartItem,
					},
				})
				return
			}
		}

		err = AddLendItem(ctx, &lendItem)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = AddPurchaseItem(ctx, &purchaseItem)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		getCart, err := GetActiveCart(ctx, userId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, getCart)
	}

}
