package piscine

func Compare(a, b string) int {
	lenea := len(a) - 1
	leneb := len(b) - 1
	lettre := a[0]
	lettre2 := b[0]
	lettre3 := a[lenea]
	lettre4 := b[leneb]
	if lettre == lettre2 && lettre3 == lettre4 {
		return 0
	}
	if lettre == lettre2 && lettre3 != lettre4 {
		return 1
	}
	if lettre != lettre2 && lettre3 == lettre4 {
		return -1
	}
	return 0
}
