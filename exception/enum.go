package exception

type Code int

const (
	RecordNotFound Code = 404
	AuthenticationFailed Code = 5000
	RequestTimeout Code = 5001
)