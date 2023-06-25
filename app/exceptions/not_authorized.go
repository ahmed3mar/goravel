package exceptions

type NotAuthorized struct {
}

func (r *NotAuthorized) Error() string {
	return "Not Authorized"
}

func (r *NotAuthorized) StatusCode() int {
	return 401
}
