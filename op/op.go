package op

import (
	"fmt"
	"github.com/swicherwich/gork/ds"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Reducer[T Number] interface {
	doReduce(acc *T, a T)
}

type ReduceFunc[T Number] func(*T, T)

func (reducer ReduceFunc[T]) doReduce(acc *T, a T) {
	reducer(acc, a)
}

func Reduce[T Number](ds ds.Dataset, col string, acc T, reducer ReduceFunc[T]) T {
	for i := 0; i < len(ds.Data); i++ {
		reducer.doReduce(&acc, ds.Data[i][col].(T))
	}
	return acc
}

type Mapper[T any] interface {
	doMap(a T) T
}

type MapFunc[T any] func(a T) T

func (mapper MapFunc[T]) doMap(a T) T {
	return mapper(a)
}

func Map[T any](ds *ds.Dataset, col string, mapper MapFunc[T]) {
	for i := 0; i < len(ds.Data); i++ {
		ds.Data[i][col] = mapper.doMap(ds.Data[i][col].(T))
	}
}

type filterI[T any] interface {
	doFilter(a T) bool
}

type FilterFunc[T any] func(a T) bool

func (filter FilterFunc[T]) doFilter(a T) bool {
	return filter(a)
}

func Filter[T any](d *ds.Dataset, col string, filter FilterFunc[T]) {
	filteredData := make([]map[string]any, 0)

	for i := 0; i < len(d.Data); i++ {
		if filter(d.Data[i][col].(T)) {
			filteredData = append(filteredData, d.Data[i])
		}
	}
	d.Data = filteredData
}

func Group[T Number](d *ds.Dataset, gByCol []string, gDataCol string, reducer ReduceFunc[T]) {
	groups := make(map[string][]map[string]any)

	for _, row := range d.Data {
		var groupKey string
		for _, col := range gByCol {
			groupKey += fmt.Sprintf("%v:", row[col])
		}

		if _, ok := groups[groupKey]; !ok {
			groups[groupKey] = make([]map[string]any, 0)
		}

		groups[groupKey] = append(groups[groupKey], row)
	}

	groupedData := make([]map[string]any, 0)

	for _, groupData := range groups {
		group := make(map[string]any)
		for _, col := range gByCol {
			group[col] = groupData[0][col]
		}

		var total T
		for _, row := range groupData {
			reducer.doReduce(&total, row[gDataCol].(T))
		}
		group[gDataCol] = total

		groupedData = append(groupedData, group)
	}

	headers := make([]string, 0)
	for _, v := range gByCol {
		headers = append(headers, v)
	}
	headers = append(headers, gDataCol)

	d.Headers = headers
	d.Data = groupedData
}
