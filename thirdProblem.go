package main

func DigitLetterComb(digits int) int {
	var (
		prevComb, comb, prevDigit int = 1, 1, 0
		isDup                     bool
	)

	for i := 0; digits > 0; i++ {
		digit := digits % 10
		digits /= 10

		digitLetter := (digit * 10) + prevDigit
		prevDigit = digit

		if i == 0 {
			digitLetter = digit
			continue
		}

		if digitLetter > 26 {
			isDup = false
			continue
		}

		if isDup {
			prevComb, comb = comb, comb+prevComb
		}

		if !isDup {
			prevComb, comb = comb, comb*2
			isDup = true
		}
	}

	return comb
}
