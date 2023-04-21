package internal

// sliceAt returns the value at index 'idx' or 'fallback' if out of range
func sliceAt[T any](lst []T, idx int, fallback T) T {
	if idx < 0 || len(lst) <= idx {
		return fallback
	}

	return lst[idx]
}
