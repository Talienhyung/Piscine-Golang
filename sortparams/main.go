package main

func sortArgs(args []string) {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if compareStrings(args[i], args[j]) > 0 {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}

func compareStrings(a, b string) int {
	i := 0
	for i < len(a) && i < len(b) {
		if a[i] != b[i] {
			return int(a[i]) - int(b[i])
		}
		i++
	}
	return len(a) - len(b)
}
