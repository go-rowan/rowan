package csv

type options struct {
	comma rune
}

func defaultOptions() options {
	return options{
		comma: 0,
	}
}

type Option func(*options)

func WithDelimiter(r rune) Option {
	return func(o *options) {
		o.comma = r
	}
}
