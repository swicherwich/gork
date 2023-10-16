package ds

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func TestDataset_String(t *testing.T) {
	dataset := getDataset()
	res := dataset.String()
	want := `col1              |col2
---------------   |---------------
rec1              |11
rec2              |22
rec3              |33
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

func TestParseValue(t *testing.T) {
	val, kind := parseValue("str")
	if val != "str" || kind != reflect.String {
		t.Fatalf("parseValue(\"str\") = [val=%s, kind=%s], want = [val=%s, kind=%s]", val, kind, "str", reflect.String)
	}
	val, kind = parseValue("1")
	if val != 1 || kind != reflect.Int {
		t.Fatalf("parseValue(1) = [val=%s, kind=%s], want = [val=%d, kind=%s]", val, kind, 1, reflect.Int)
	}
	val, kind = parseValue("true")
	if val != true || kind != reflect.Bool {
		t.Fatalf("parseValue(1) = [val=%s, kind=%s], want = [val=%v, kind=%s]", val, kind, true, reflect.Bool)
	}
}

func createTempFile(t *testing.T) *os.File {
	file, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatal("Error creating temporary file:", err)
	}
	return file
}

func removeTempFile(t *testing.T, file *os.File) {
	if err := os.Remove(file.Name()); err != nil {
		t.Fatal("Error removing temporary file:", err)
	}
}

func TestSaveCsv(t *testing.T) {
	ds := Dataset{
		Headers: []string{"Name", "Age"},
		Data: []map[string]any{
			{"Name": "Alice", "Age": 30},
			{"Name": "Bob", "Age": 25},
		},
	}

	file := createTempFile(t)
	defer removeTempFile(t, file)

	err := ds.SaveCsv(file.Name())
	if err != nil {
		t.Fatal("SaveCsv returned an error:", err)
	}

	file.Close()
	csvFile, err := os.Open(file.Name())
	if err != nil {
		t.Fatal("Error opening the saved CSV file:", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatal("Error reading the saved CSV file:", err)
	}

	expected := [][]string{
		{"Name", "Age"},
		{"Alice", "30"},
		{"Bob", "25"},
	}

	if len(records) != len(expected) {
		t.Fatalf("Expected %d records, but got %d", len(expected), len(records))
	}

	for i, record := range records {
		for j, value := range record {
			if value != expected[i][j] {
				t.Errorf("Expected value %s at [%d][%d], but got %s", expected[i][j], i, j, value)
			}
		}
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
