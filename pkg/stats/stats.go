package stats

import (
	"github.com/Ilhom5005/bank/v2/pkg/types"
	)

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFAIL {
			continue
		}
		sum += payment.Amount
	}

	return types.Money(sum)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, payment := range payments {
		if payment.Status == types.StatusFAIL {
			continue
		}
		if category == payment.Category {
			sum += int(payment.Amount)
		}

	}
	
	return types.Money(sum)
}
