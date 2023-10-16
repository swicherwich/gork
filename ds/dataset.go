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

func (ds Dataset) SaveCsv(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	err = writeRecord(w, ds.Headers)
	if err != nil {
		removeErr := os.Remove(path)
		return fmt.Errorf("error occurred while wring to file: %s, %s", err, removeErr)
	}
	w.Flush()

	flushSize := 100
	var recCounter int
	for _, record := range ds.Data {
		row := convertMapToSlice(ds.Headers, record)
		err := writeRecord(w, row)
		if err != nil {
			removeErr := os.Remove(path)
			return fmt.Errorf("error occurred while wring to file: %s, %s", err, removeErr)
		}
		if recCounter%flushSize == 0 {
			w.Flush()
		}
	}
	return nil
}

func writeRecord(w *csv.Writer, record []string) error {
	err := w.Write(record)
	if err != nil {
		return err
	}
	return nil
}

func convertMapToSlice(headers []string, record map[string]any) []string {
	slice := make([]string, 0)
	for _, header := range headers {
		slice = append(slice, fmt.Sprintf("%v", record[header]))
	}
	return slice
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
