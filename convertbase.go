package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return PrintNbrBase(AtoiBase(nbr, baseFrom), baseTo)
}
