package simple_http

func compose(middleware []HandlerFunc) HandlerFunc {

	return func(ctx *Context, next NextFunc) {

		index := -1

		var dispatch func(i int) func()

		dispatch = func(i int) func() {

			if i <= index {
				panic("next() called multiple times")
			}

			index = i

			if i == len(middleware) {
				return next
			}

			fn := middleware[i]

			return func() {
				fn(ctx, dispatch(i+1))
			}
		}

		dispatch(0)()
	}
}
