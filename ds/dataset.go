package ds

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type Dataset struct {
	Headers []string
	Types   map[string]reflect.Kind
	Data    []map[string]any
}

func (ds Dataset) String() string {
	var s string
	for i, v := range ds.Headers {
		s += fmt.Sprintf("| <'%s'>[index=%d][type=%s] ", v, i, ds.Types[v])
	}
	s += fmt.Sprintf("|\n")

	for i := 0; i < len(ds.Data); i++ {
		for _, v := range ds.Headers {
			s += fmt.Sprintf("| %v", ds.Data[i][v])
		}
		s += fmt.Sprintf("| ")
		s += fmt.Sprintf("\n")
	}

	return s
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
	fmt.Println(headers)

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
