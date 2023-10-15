package ds

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Dataset struct {
	Headers []string
	Types   map[string]reflect.Kind
	Data    []map[string]any
}

func (ds Dataset) String() string {
	var buffer strings.Builder

	w := tabwriter.NewWriter(&buffer, 1, 0, 3, ' ', tabwriter.Debug)

	headerRow := strings.Join(ds.Headers, "\t")
	fmt.Fprintln(w, headerRow)

	separator := make([]string, len(ds.Headers))
	for i := range separator {
		separator[i] = strings.Repeat("-", 15)
	}
	separatorRow := strings.Join(separator, "\t")
	fmt.Fprintln(w, separatorRow)

	for _, row := range ds.Data {
		var rowValues []string
		for _, header := range ds.Headers {
			value := fmt.Sprintf("%v", row[header])
			rowValues = append(rowValues, value)
		}
		rowStr := strings.Join(rowValues, "\t")
		fmt.Fprintln(w, rowStr)
	}

	w.Flush()

	return buffer.String()
}

func DatasetFromCsv(path string) *Dataset {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Error occurred while opening csv file: %s", err.Error()))
	}

	r := csv.NewReader(file)

	headers, err := r.Read()
	if err != nil {
		panic(fmt.Sprintf("Error reading csv file: %s", err.Error()))
	}

	ds := &Dataset{}
	ds.Headers = headers
	ds.Types = make(map[string]reflect.Kind)
	ds.Data = make([]map[string]any, 0)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		recMap := make(map[string]any)
		for i, v := range ds.Headers {
			val, kind := parseValue(record[i])
			recMap[v] = val
			ds.Types[v] = kind
		}
		ds.Data = append(ds.Data, recMap)
	}

	return ds
}

func parseValue(val any) (v any, kind reflect.Kind) {
	s := val.(string)
	i, err := strconv.Atoi(s)
	if err == nil {
		return i, reflect.Int
	}
	b, err := strconv.ParseBool(s)
	if err == nil {
		return b, reflect.Bool
	}
	return s, reflect.String
}
