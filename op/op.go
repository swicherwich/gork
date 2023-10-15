package op

import (
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
