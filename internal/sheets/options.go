package sheets

type options struct {
	isURL   bool
	rangeA1 string
}

type Option func(*options)

func WithSheetsURL() Option {
	return func(o *options) {
		o.isURL = true
	}
}

func WithRange(rangeA1 string) Option {
	return func(o *options) {
		o.rangeA1 = rangeA1
	}
}
