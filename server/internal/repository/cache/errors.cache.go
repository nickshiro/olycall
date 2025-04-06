package cache

type ConflictError struct {
	Field string
}

func (e ConflictError) Error() string {
	return e.Field
}
