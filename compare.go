package piscine

func Compare(a, b string) int {
	lena := 0
	for range a {
		lena++
	}
	lenb := 0
	for range b {
		lenb++
	}
	lenea := lena - 1
	leneb := lenb - 1
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
	return 1
}
