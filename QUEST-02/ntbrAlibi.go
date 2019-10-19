package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	s := 1
	if n < 0 {
		s = -1
		z01.PrintRune('-')
	}

	if n != 0 {
		f := (n / 10) * s
		if f != 0 {
			PrintNbr(f)
		}
		k := ((n % 10) * s) + '0'
		z01.PrintRune(rune(k))

	} else {
		z01.PrintRune('0')
	}
}

func PrintNbr444(n int) {
	last := false
	if n == -9223372036854775808 {
		n = -922337203685477580
		last = true
	}
	if n < 0 {
		n = -1 * n
		z01.PrintRune('-')
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	var arr [36]int
	div := n
	rem := 0
	i := 0
	for div > 0 {
		rem = div % 10
		div /= 10
		arr[i] = rem
		i++
	}
	for j := i - 1; j >= 0; j-- {
		c := 0
		char := '0'
		for c < arr[j] {
			char++
			c++
		}
		z01.PrintRune(char)
	}
	if last {
		z01.PrintRune('8')
	}
}
  
func PrintNbr123(n int) {
  // count number of digits in n (ex: 123 -> 3 digits
  count := 0
  for c := n; c != 0; count++ {
    c = c / 10
  }

  if count == 0 {
    z01.PrintRune('0')
  } else if n != 0 {
    elem := n
    if n < 0 && n != -9223372036854775808 {
      // convert negatives number to positives & display '-'
      elem = -n
      z01.PrintRune('-')
    } else if n == -9223372036854775808 {
      elem = -n - 1
      z01.PrintRune('-')
    }
    for i := count; i > 0; i-- {
      power := 1
      for d := i -1; d > 0; d-- {
        power = power * 10
      } 
      // treat special limit of 9223372036854775808
      if n == -9223372036854775808 && power == 1 {
        z01.PrintRune(rune(((elem / power) % 10) + 49))
      } else if !(n == -9223372036854775808 && power == 1) {
        z01.PrintRune(rune(((elem / power) % 10) + 48))
      }
    } 
  }
}

func main(){
	PrintNbr123(-12365765756776)
	

}