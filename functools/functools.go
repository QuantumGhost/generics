package temp

func Map[T any, U any](data []T, f func(T) U) []U {
	out := make([]U, 0, len(data))
	for _, v := range data {
		out = append(out, f(v))
	}
	return out
}

func MapErr[T any, U any](data []T, f func (T) (U, error)) ([]U, error) {
	out := make([]U, 0, len(data))
	for _, t := range data {
		u, err := f(t)
		if err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, nil
}

func Filter[T any](data []T, filterFn func (T) bool) []T {
	filtered := make([]T, 0, len(data))
	for _, v := range data {
		if filterFn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Reduce[T any](data []T, initial T, combiner func (T, T) T) T {
	value := initial
	for _, elem := range data {
		value = combiner(value, elem)
	}
	return value
}
