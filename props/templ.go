package props

func FirstOrDefault[T any](def T, props ...T) T {
	if len(props) > 0 {
		return props[0]
	}
	return def
}
