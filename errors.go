package errors

// New represents a string error
type New string

// Error returns the error string
func (n New) Error() string {
	return string(n)
}
