Usage example:

```go
package main

import (
	"fmt"
	"github.com/swicherwich/gork/ds"
	"github.com/swicherwich/gork/op"
	"time"
)

func main() {
	dataset := ds.DatasetFromCsv("example_data.csv")

	op.Map(dataset, "time", func(a string) string {
		t, _ := time.Parse("2006-01-02T15:04:05", a)
		return t.Format("01-2006")
	})
	op.Filter(dataset, "time", func(a string) bool { return a == "10-2023" })
	op.Group(dataset, []string{"time", "mcc"}, "amount", func(acc *int, a int) { *acc += a })
	res := op.Reduce(*dataset, "amount", 0, func(i *int, i2 int) {
		*i += i2
	})
	fmt.Println("Total amount spent:", res) // -505

	dataset.SaveCsv("res.csv")

}
```

Example data:
```csv
time,mcc,amount
2023-10-14T17:11:08,Grocery,-53
2023-10-14T16:33:47,Grocery,-123
2023-10-14T15:35:38,Grocery,-61
2023-10-14T13:36:23,Fast-food,-68
2023-10-14T11:48:53,Digital goods,-200
2023-09-13T21:53:14,Transfer,-20
2023-09-13T16:25:31,Digital goods,-280
2023-09-13T15:15:31,Fast-food,-62
2023-09-13T10:39:55,Grocery,-121
```

Result:
```csv
time,mcc,amount
10-2023,Fast-food,-68
10-2023,Digital goods,-200
10-2023,Grocery,-237
```