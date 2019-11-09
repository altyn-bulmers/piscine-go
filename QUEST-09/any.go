package piscine

func Any(f func(string) bool, arr []string) bool {

	for _, s := range arr {
		if f(s) == true {
			return true
		}
	}
	return false

}
