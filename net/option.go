package net

// WithEnv new env option
func WithEnv(env string) Option {
	return newFuncOption(func(o *option) {
		o.env = env
	})
}

// Option config option
type Option interface {
	apply(*option)
}

type option struct {
	env string
}

func decodeOpts(opts []Option) option {
	op := defaultOption()
	for _, opt := range opts {
		opt.apply(&op)
	}
	return op
}

func defaultOption() option {
	return option{
		env: "HOST_ADDRESS",
	}
}

type funcOption struct {
	f func(*option)
}

func (fo *funcOption) apply(do *option) {
	fo.f(do)
}

func newFuncOption(f func(*option)) *funcOption {
	return &funcOption{
		f: f,
	}
}
