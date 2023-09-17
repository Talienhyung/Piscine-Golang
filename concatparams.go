package piscine

func ConcatParams(args []string) string {
	str := ""
	for i := range args {
		if len(args)-1 != i {
			str += args[i] + "\n"
		} else {
			str += args[i]
		}
	}
	return str
}
