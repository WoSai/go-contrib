package validation

type (
	Notification interface {
		AddError(reason string)
		Err() error
	}
)
