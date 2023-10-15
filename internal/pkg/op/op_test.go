package op

import (
	"github.com/swicherwich/gork/internal/pkg/ds"
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	dataset := getDataset()
	res := Reduce(dataset, "col2", 0, func(acc *int, a int) { *acc += a })

	if res != 66 {
		t.Fatalf(`Reduce(dataset, "col2", 0, func(acc *int, a int) { *acc += a }) = %d, want 66`, res)
	}
}

func TestMap(t *testing.T) {
	dataset := getDataset()

	col := "col2"
	Map[int](&dataset, col, func(a int) int { return a * 2 })
	want := []int{22, 44, 66}

	for i := 0; i < len(dataset.Data); i++ {
		res := dataset.Data[i][col]
		if res != want[i] {
			t.Fatalf(`Map[int](&tmpDataset, col, func(a int) int { return a * 2 }) = %d, want %d`, res, want)
		}
	}
}

func getDataset() ds.Dataset {
	return ds.Dataset{
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
}
