package piscine

func CountIf(f func(string) bool, tab []string) int {

	count := 0
	for _, s := range tab {
		if f(s) == true {
			count++
		}
	}
	return count
}
