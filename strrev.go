package piscine

func StrRev(s string) string {
	var v string
	for _, c := range s {
		v = string(c) + v
	}
	return s
}
