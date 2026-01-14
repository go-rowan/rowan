package sheets

type Error struct {
	Kind ErrorKind
	Err  error
}

type ErrorKind string

const (
	ErrAPIDisabled  ErrorKind = "api_disabled"
	ErrQuotaProject ErrorKind = "quota_project_missing"
	ErrUnauthorized ErrorKind = "unauthorized"
	ErrUnknown      ErrorKind = "unknown"
)

func (e *Error) Error() string {
	switch e.Kind {
	case ErrAPIDisabled:
		return "rowan: Google Sheets API is not enabled"
	case ErrQuotaProject:
		return "rowan: Google Sheets requires a quota project"
	case ErrUnauthorized:
		return "rowan: not authorized to access this spreadsheet"
	default:
		return "rowan: failed to read Google Sheets"
	}
}

func (e *Error) Unwrap() error {
	return e.Err
}
