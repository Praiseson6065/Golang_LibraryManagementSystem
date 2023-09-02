package handlers

import (
	"fmt"
	"log"

	"strconv"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

func PaymentHandler(c *fiber.Ctx) error {
	Userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}

	data, err := models.GetPurchaseCartbyId(Userid)
	if err != nil {
		return err
	}
	Price := 0
	for _, i := range data {
		Price += ((int(i.Book.Price)) * (i.PurchaseDetails.Quantity))
	}

	stripe.Key = config.EnvConfigs().STRIPE_key
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(Price * 100)),
		Currency: stripe.String(string(stripe.CurrencyINR)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)
	if err != nil {
		return err
	}
	var payment models.Pi
	payment.Id = pi.ID
	payment.ClientSecret = pi.ClientSecret

	return c.JSON(payment)

}

func PaymentConfirm(c *fiber.Ctx) error {

	var data models.Pi
	c.BodyParser(&data)
	fmt.Println(data)
	stripe.Key = config.EnvConfigs().STRIPE_key
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
		ReturnURL:     stripe.String("/profile.html"),
	}
	pi, _ := paymentintent.Confirm(
		data.Id,
		params,
	)

	if pi.ClientSecret == data.ClientSecret {
		return c.JSON(pi.NextAction.RedirectToURL.URL)
	}
	return c.JSON(false)
}
