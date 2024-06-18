package exception


func IsNotFound(err error) bool {
	if e, ok := err.(*ApiException); ok {
		if e.BusinessCode == 404 {
			return true
		}
	}
	return false
}

func NewRecordNotFound(format string, a ...any) *ApiException {
	return NewApiException(RecordNotFound, format, a...)
}