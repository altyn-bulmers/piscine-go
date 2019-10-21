package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	} else {
		for i := 0; i < nb/2; i++ {
			if i*i == nb {
				return i
			}
		}
	}
	return 0
}

/*
func Sqrt(nb int) int{

	result := 1
	nb := 16

	if result*10%10 != 0{
		return 0
	} else {
		for i:=0; i<20; i++ {
			result -= (result*result - nb) / (2*result)

		}
		return result
	}



}
*/
