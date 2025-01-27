package lending

import "github.com/gin-gonic/gin"

func ApproveBooksByAdmin(ctx *gin.Context, approveBooks []ApproveBooksRequest) error {
	for _, approveBook := range approveBooks {
		tx := db.WithContext(ctx).Model(&LendBook{}).Where("id = ?", approveBook.LendBookIDs).Updates(LendBook{Status: "Approved"})
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}
