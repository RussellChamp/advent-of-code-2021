package filters

func ByLength(size int) func(string) bool {
	return func(val string) bool { return len(val) == size }
}
