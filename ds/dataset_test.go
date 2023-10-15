package ds

import (
	"reflect"
	"testing"
)

func TestDataset_String(t *testing.T) {
	dataset := getDataset()
	res := dataset.String()
	want := `| <'col1'>[index=0][type=string] | <'col2'>[index=1][type=int] |
| rec1| 11| 
| rec2| 22| 
| rec3| 33| 
`
	if want != res {
		t.Fatalf("dataset.String() = \n%s, want = \n%s", res, want)
	}
}

func TestDatasetFromCsv(t *testing.T) {
	dataset := DatasetFromCsv("testdata/data.csv")

	if reflect.DeepEqual(getDataset(), dataset) {
		t.Fatalf("DatasetFromCsv(\"testdata/data.csv\") = \n%s, want = \n%s", dataset, getDataset())
	}
}

func getDataset() Dataset {
	return Dataset{
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
