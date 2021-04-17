// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	cardNumber := os.Getenv("CARD_NUMBER")
	cardHolder := os.Getenv("CARD_HOLDER")
	expMonth := os.Getenv("EXP_MONTH")
	expYear := os.Getenv("EXP_YEAR")
	userID := os.Getenv("USER_ID")

	formula.Formula{
		CardNumber: cardNumber,
		CardHolder: cardHolder,
		ExpMonth:   expMonth,
		ExpYear:    expYear,
		UserID:     userID,
	}.Run(os.Stdout)
}
