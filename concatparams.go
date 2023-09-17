package piscine

func ConcatParams(args []string) string {
	str := ""
	for i := range args {
		if len(str)-1 != i {
			str += args[i] + "\n"
		} else {
			str += args[i]
		}
	}
	return str
}
