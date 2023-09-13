package piscine

func StrRev(s string) string {
	var v string
	for i := len(s) - 1; i >= 0; i-- {
		v += string(s[i] + '0')
	}
	return v
}
