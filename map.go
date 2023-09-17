package piscine

func Map(f func(int) bool, a []int) []bool {
	var mapi []bool
	for _, i := range a {
		mapi = append(mapi, f(i))
	}
	return mapi
}
