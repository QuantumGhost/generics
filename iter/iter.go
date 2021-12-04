package iter

const (
	ErrNoMore sentinel = "no more item"
)

type sentinel string

func (s sentinel) Error() string {
	return string(s)
}


type Iterator[T any] interface {
	Next() (T, error)
}


type iterFunc [T any] func () (T, error)

func (f iterFunc[T]) Next() (T, error) {
	return f()
}

func Slice[T any](s []T) Iterator[T] {
	length := len(s)
	i := 0
	return iterFunc[T](func () (T, error) {
		if i < length {
			res := s[i]
			i++
			return res, nil
		}
		var res T
		return res, ErrNoMore
	})
}

type Pair[K any, V any] struct {
	Key K
	Value V
}

func Map[K comparable, V any](m map[K]V) Iterator[Pair[K, V]] {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys =append(keys, k)
	}
	keyIdx := 0
	length := len(keys)
	return iterFunc[Pair[K, V]](func () (Pair[K, V], error) {
		for i := keyIdx; i<length; i++ {
			key := keys[i]
			value, ok := m[key]
			if !ok {
				continue
			}
			keyIdx = i + 1
			return Pair[K, V]{Key: key, Value: value}, nil
		}
		return Pair[K, V]{}, ErrNoMore
	})
}
