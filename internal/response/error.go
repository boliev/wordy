package response

// Error response struct
type Error struct {
	Error string `json:"error"`
}

// NewError Error response struct constructor
func NewError(error string) *Error {
	return &Error{
		Error: error,
	}
}
