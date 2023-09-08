package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}
	for i, c := range base {
		for j, d := range base {
			if i != j && c == d {
				return 0
			}
		}
	}
	for _, c := range base {
		if c == '+' || c == '-' {
			return 0
		}
	}
	baseMap := make(map[rune]int)
	for i, c := range base {
		baseMap[c] = i
	}
	result := 0
	multiplier := 1
	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])
		if val, ok := baseMap[c]; ok {
			result += val * multiplier
			multiplier *= len(base)
		} else {
			return 0
		}
	}
	return result
}
