package piscine

func ConcatParams(args []string) string {
	str := ""
	for i := range args {
		str = args[i] + "\n"
	}
	return str
}
