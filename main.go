package main

import (
	"fmt"
	"github.com/swicherwich/gork/internal/pkg/dataset"
	"reflect"
)

func main() {
	ds := dataset.Dataset{
		Headers: []string{"col1", "col2"},
		Types: map[string]reflect.Kind{
			"col1": reflect.String, "col2": reflect.Int,
		},
		Data: []map[string]any{
			map[string]any{"col1": "rec1", "col2": 11},
			map[string]any{"col1": "rec2", "col2": 22},
		},
	}
	fmt.Println(ds.String())
}
