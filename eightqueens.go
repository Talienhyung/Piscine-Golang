package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	printTableau(15863724)
	printTableau(16837425)
	printTableau(17468253)
	printTableau(17582463)
	printTableau(24683175)
	printTableau(25713864)
	printTableau(25741863)
	printTableau(26174835)
	printTableau(26831475)
	printTableau(27368514)
	printTableau(27581463)
	printTableau(28613574)
	printTableau(31758246)
	printTableau(35281746)
	printTableau(35286471)
	printTableau(35714286)
	printTableau(35841726)
	printTableau(36258174)
	printTableau(36271485)
	printTableau(36275184)
	printTableau(36418572)
	printTableau(36428571)
	printTableau(36814752)
	printTableau(36815724)
	printTableau(36824175)
	printTableau(37285146)
	printTableau(37286415)
	printTableau(38471625)
	printTableau(41582736)
	printTableau(41586372)
	printTableau(42586137)
	printTableau(42736815)
	printTableau(42736851)
	printTableau(42751863)
	printTableau(42857136)
	printTableau(42861357)
	printTableau(46152837)
	printTableau(46827135)
	printTableau(46831752)
	printTableau(47185263)
	printTableau(47382516)
	printTableau(47526138)
	printTableau(47531682)
	printTableau(48136275)
	printTableau(48157263)
	printTableau(48531726)
	printTableau(51468273)
	printTableau(51842736)
	printTableau(51863724)
	printTableau(52468317)
	printTableau(52473861)
	printTableau(52617483)
	printTableau(52814736)
	printTableau(53168247)
	printTableau(53172864)
	printTableau(53847162)
	printTableau(57138642)
	printTableau(57142863)
	printTableau(57248136)
	printTableau(57263148)
	printTableau(57263184)
	printTableau(57413862)
	printTableau(58413627)
	printTableau(58417263)
	printTableau(61528374)
	printTableau(62713584)
	printTableau(62714853)
	printTableau(63175824)
	printTableau(63184275)
	printTableau(63185247)
	printTableau(63571428)
	printTableau(63581427)
	printTableau(63724815)
	printTableau(63728514)
	printTableau(63741825)
	printTableau(64158273)
	printTableau(64285713)
	printTableau(64713528)
	printTableau(64718253)
	printTableau(68241753)
	printTableau(71386425)
	printTableau(72418536)
	printTableau(72631485)
	printTableau(73168524)
	printTableau(73825164)
	printTableau(74258136)
	printTableau(74286135)
	printTableau(75316824)
	printTableau(82417536)
	printTableau(82531746)
	printTableau(83162574)
	printTableau(84136275)
	// for i := 15863724; i < 87654321; i++ {
	// 	tab := transformation(i)
	// 	if valide(tab) && verification(tab) {
	// 		printTableau(tab)
	// 	}
	// }
}

func transformation(n int) []int {
	digits := make([]int, 0)
	t2 := make([]int, 0)
	for n > 0 {
		p := n % 10
		digits = append(digits, p)
		n = n / 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		t2 = append(t2, digits[i])
	}
	return t2
}

// func valide(tab []int) bool {
// 	for i := range tab {
// 		if tab[i] == 0 || tab[i] == 9 {
// 			return false
// 		}
// 	}
// 	for i, c := range tab {
// 		for j, d := range tab {
// 			if i != j && c == d {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func validationManger(i int, j int, s []int) string {
// 	if i+j < 8 {
// 		if s[i+j] == s[i]+j || s[i+j] == s[i]-j {
// 			return "false"
// 		}
// 	} else if i-j >= 0 {
// 		if s[i-j] == s[i]+j || s[i-j] == s[i]-j {
// 			return "false"
// 		}
// 	}
// 	return "true"
// }

// func verification(tab []int) bool {
// 	for i := 0; i < 8; i++ {
// 		for j := 1; j < 8; j++ {
// 			if validationManger(i, j, tab) == "false" {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func printTableau(n int) {
	tab := transformation(n)
	for i := 0; i < len(tab); i++ {
		z01.PrintRune(rune(tab[i] + '0'))
	}
	z01.PrintRune('\n')
}
