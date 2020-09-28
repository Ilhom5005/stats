package stats

import (
	"github.com/Ilhom5005/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment {
{
	ID: 1,
	Amount: 10_000,
	Category: "jacket",
	Status: "FAIL",
},
{
	ID: 2,
	Amount: 20_000,
	Category: "jeans",
	Status: "OK",
},
{
	ID: 3,
	Amount: 30_000,
	Category: "T-shirt",
	Status: "OK",
},
	}
fmt.Println(Avg(payments))

//Output:
//20000

}
