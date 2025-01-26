package payments

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

func createCheckoutSession(paymentRequest *PaymentRequest) (string, error) {

	stripe.Key = stripeKey

	var lineItems []*stripe.CheckoutSessionLineItemParams
	for _, book := range paymentRequest.Books {
		lineItem := &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("inr"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name:   stripe.String(book.Title),
					Images: []*string{stripe.String(book.ImageURL)},
				},
				UnitAmount: stripe.Int64(int64(book.Price)),
			},
			Quantity: stripe.Int64(int64(book.Quantity)),
		}
		lineItems = append(lineItems, lineItem)
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:         stripe.String("https://example.com/success"),
		CancelURL:          stripe.String("https://example.com/cancel"),
	}

	session, err := session.New(params)

	if err != nil {
		return "", err
	}

	return session.ID, nil
}
