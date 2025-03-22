package cart

import (
	"LibManMicroServ/middleware"

	"github.com/gin-gonic/gin"
)

func CreateUserCart(ctx *gin.Context, UserId string) error {
	var cart Cart
	cart.UserID = UserId
	tx := db.WithContext(ctx).Create(&cart)

	if tx.Error != nil {

		return tx.Error

	}
	return nil
}

func GetActiveCart(ctx *gin.Context, UserId string) (*Cart, error) {
	var cart Cart
	tx := db.WithContext(ctx).Preload("LendItems", "status = ?", "pending").Preload("PurchaseItems", "status = ?", "pending").Where("user_id = ?", UserId).First(&cart)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &cart, nil
}

func GetCart(ctx *gin.Context, UserId string) (*Cart, error) {
	var cart Cart
	tx := db.WithContext(ctx).Preload("LendItems").Preload("PurchaseItems").Where("user_id = ?", UserId).First(&cart)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &cart, nil
}
func CheckOutCartbyUser(ctx *gin.Context) (Cart, error) {

	userId := middleware.GetUserID(ctx)
	var cart Cart
	if err := db.WithContext(ctx).Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return cart, err
	}

	cart.LendItems = nil
	cart.PurchaseItems = nil
	cart.Status = "active"

	if err := db.WithContext(ctx).Save(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil

}

func AddLendItem(ctx *gin.Context, lc *LendCartItem) error {
	tx := db.WithContext(ctx).Create(&lc)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func AddPurchaseItem(ctx *gin.Context, pc *PurchaseCartItem) error {
	tx := db.WithContext(ctx).Create(&pc)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
