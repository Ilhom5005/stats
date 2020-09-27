package stats

import (
	"github.com/Ilhom5005/bank/pkg/typesv2"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment {
{
	ID: 1,
	Amount: 10_000,
	Category: "jacket",
},
{
	ID: 2,
	Amount: 20_000,
	Category: "jeans",
},
{
	ID: 3,
	Amount: 30_000,
	Category: "T-shirt",
},
	}
fmt.Println(Avg(payments))

//Output:
//20000

}
