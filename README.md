Usage:

```go
func main() {
	dataset := ds.Dataset{
		Headers: []string{"col1", "col2"},
		Types: map[string]reflect.Kind{
			"col1": reflect.String, "col2": reflect.Int,
		},
		Data: []map[string]any{
			{"col1": "rec1", "col2": 11},
			{"col1": "rec2", "col2": 22},
			{"col1": "rec3", "col2": 33},
		},
	}
	doubler := func(a int) int { return a * 2 }

	fmt.Println(dataset.String())
	op.Map[int](&dataset, "col2", doubler)
	fmt.Println(dataset.String())

	fmt.Println(op.Reduce(dataset, "col2", 0, func(acc *int, a int) { *acc += a }))
}
```