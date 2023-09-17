package piscine

func IsSorted(f func(int, int) int, arr []int) bool {
	cesce := true
	decresce := true
	for i := 1; i < Lent4(arr); i++ {
		if !(f(arr[i-1], arr[i]) >= 0) {
			cesce = false
		}
	}
	for i := 1; i < Lent4(arr); i++ {
		if !(f(arr[i-1], arr[i]) <= 0) {
			decresce = false
		}
	}
	return cesce || decresce
}

func Lent4(d []int) int {
	return len(d)
}
