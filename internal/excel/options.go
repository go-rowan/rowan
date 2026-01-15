package excel

type options struct {
	rangeA1 string
}

type Option func(*options)

func WithRange(rangeA1 string) Option {
	return func(o *options) {
		o.rangeA1 = rangeA1
	}
}
