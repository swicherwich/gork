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

	for i := 0; i < len(ds.Data); i++ {
		for _, v := range ds.Headers {
			s += fmt.Sprintf("| %v", ds.Data[i][v])
		}
		s += fmt.Sprintf("| ")
		s += fmt.Sprintf("\n")
	}

	return s
}
