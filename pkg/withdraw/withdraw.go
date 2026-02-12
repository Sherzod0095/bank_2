package withdraw

import "github.com/Sherzod0095/bank_2/pkg/types"

const withdrwaLimit = 20_000_00

func Withdraw(card *types.Card, amount types.Money) {
	if amount < 0 {
		return
	}

	if amount > withdrwaLimit {
		return
	}

	if !card.Active {
		return
	}

	if card.Balance < amount {
		return
	}

	card.Balance -= amount
}
