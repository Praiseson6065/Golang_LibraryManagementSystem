package handlers

import (
	"fmt"

	"strconv"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
)

func ProfilePage(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, err := middlewares.CookieGetData(cookie, c)

	if err != nil {
		return c.JSON(fiber.Map{
			"msg": "Not Logged In",
		})
	} else {
		return c.JSON(fiber.Map{
			"msg":   "Logged In",
			"name":  claims["name"],
			"email": claims["email"],
		})
	}

}
func UserRequestedBooks(c *fiber.Ctx) error {
	db, err := database.DbGormConnect()
	if err != nil {
		return c.JSON(err)
	}
	var ReqBook models.UserRequestedBooks
	db.AutoMigrate(&models.UserRequestedBooks{})
	c.BodyParser(&ReqBook)
	ReqBook.RequestStatus = false
	err = db.Create(&ReqBook).Error
	if err != nil {
		return c.JSON(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(true)
}
func RequestedBooks(c *fiber.Ctx) error {
	userId, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return c.JSON(err)
	}
	db, err := database.DbGormConnect()
	if err != nil {
		return c.JSON(err)
	}
	var ReqBooks []models.UserRequestedBooks

	db.Find(&ReqBooks).Where("user_id=?", userId)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(ReqBooks)
}
func BookReviewByUser(c *fiber.Ctx) error {
	var bkRvw models.BookReviews
	UserId, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return c.JSON(err)
	}
	c.BodyParser(&bkRvw)
	db, err := database.DbGormConnect()
	if err != nil {
		return c.JSON(err)
	}
	bkRvw.UserId = UserId
	User, err := models.GetUser(UserId)
	if err != nil {
		return c.JSON(err)
	}
	bkRvw.UserName = User.Name
	db.AutoMigrate(&models.BookReviews{})
	err = db.Create(&bkRvw).Error
	if err != nil {
		return c.JSON(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(true)
}
func DeleteReviewByUser(c *fiber.Ctx) error {
	UserId, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return c.JSON(err)
	}
	BookId, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return c.JSON(err)
	}

	db, err := database.DbGormConnect()
	if err != nil {
		return c.JSON(err)
	}
	Query := fmt.Sprintf("Delete from book_reviews where book_id=%d and user_id=%d ;", BookId, UserId)
	err = db.Exec(Query).Error
	if err != nil {
		return c.JSON(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(true)
}
func UpdateReviewByUser(c *fiber.Ctx) error {
	UserId, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return c.JSON(err)
	}
	var bkRvw models.BookReviews
	c.BodyParser(&bkRvw)
	UpdatedBookReview, err := models.GetReviewByUserBookId(UserId, bkRvw.BookId)
	if err != nil {
		return c.JSON(err)
	}
	UpdatedBookReview.Review = bkRvw.Review
	db, err := database.DbGormConnect()
	if err != nil {
		return c.JSON(err)
	}
	db.Save(&UpdatedBookReview)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(true)
}
func UserBookDetails(c *fiber.Ctx) error {
	var UserBookDetails models.UserBookDetails
	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	bookId, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}
	ApprovedBooks, err := models.GetUserApprovedBooks(userid)
	if err != nil {
		return err
	}
	for _, i := range ApprovedBooks {
		if i.ID == bookId {
			UserBookDetails.Approve = true
			break
		}
	}
	IssuedBooks, err := models.GetIssuedBooks(userid)
	if err != nil {
		return err
	}
	for _, i := range IssuedBooks {
		if i.ID == bookId {
			UserBookDetails.Issued = true
			break
		}
	}
	Cartbooks, err := models.GetCartBooksByUserID(userid)
	if err != nil {
		return err
	}
	for _, i := range Cartbooks {
		if i.ID == bookId {
			UserBookDetails.Cart = true
			break
		}
	}

	return c.JSON(UserBookDetails)
}
func UserRemoveIssueBook(c *fiber.Ctx) error {
	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	bookid, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}
	_, err = models.RemoveIssueBookByUser(userid, bookid)
	if err != nil {
		return err
	}
	return c.JSON(true)
}
func GetUserCart(c *fiber.Ctx) error {
	Userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	var CBooks []models.Book
	CBooks, err = models.GetCartBooksByUserID(Userid)
	if err != nil {
		return err
	}
	return c.JSON(CBooks)
}
func GetPurchaseCart(c *fiber.Ctx) error {
	Userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	var cartPurchaseBooks []models.CartPurchaseBooks
	cartPurchaseBooks, err = models.UserPurchaseCart(Userid)
	if err != nil {
		return err
	}
	var PurchaseCart []models.PurchaseCart
	for _, i := range cartPurchaseBooks {
		book, err := models.GetBookById(i.Book)
		if err != nil {
			return err
		}
		var Pc models.PurchaseCart
		Pc.Book = book
		Pc.PurchaseDetails = i
		PurchaseCart = append(PurchaseCart, Pc)
	}
	return c.JSON(PurchaseCart)

}
func RemoveFromUserPurCart(c *fiber.Ctx) error {
	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	bookid, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}
	_, err = models.RemovefromPurchaseCart(userid, bookid)
	if err != nil {
		return err
	}
	return c.JSON(true)
}
func PurchaseBook(c *fiber.Ctx) error {
	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	_, err = models.PurchaseBooks(userid)
	if err != nil {
		return err
	}
	return c.JSON(true)

}
