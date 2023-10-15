package ds

import (
	"fmt"
	"reflect"
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

	for _, v := range ds.Data {
		for _, val := range v {
			s += fmt.Sprintf("| %v", val)
		}
		s += fmt.Sprintf("| ")
		s += fmt.Sprintf("\n")
	}

	return s
}
