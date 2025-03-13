package common

func DerefOrEmpty[T any](val *T) T {
	if val == nil {
		var empty T
		return empty
	}

	return *val
}
