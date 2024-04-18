package main

import (
	"math"
	"math/big"
)

func SwapOrder(sequences []*big.Int) []*big.Int {
	swapped := make([]*big.Int, len(sequences))
	for i, j := len(sequences)-1, 0; i >= 0; i, j = i-1, j+1 {
		swapped[j] = sequences[i]
	}

	return swapped
}

func LargestMultiple(numbers *big.Int, sequenceLen int) []*big.Int {
	var (
		largestSequence = []*big.Int{}
		largestMultiple = big.NewInt(0)
	)

	for i := 0; numbers.Cmp(big.NewInt(0)) == 1; i++ {
		sequenceNumbers := []*big.Int{}
		multipleSequences := big.NewInt(1)
		calcNumbers := big.NewInt(0).Set(numbers)

		for j := 0; j < sequenceLen; j++ {
			indexedNumb := big.NewInt(0).Mod(calcNumbers, big.NewInt(10))
			if indexedNumb.Cmp(big.NewInt(0)) == 0 {
				break
			}

			sequenceNumbers = append(sequenceNumbers, indexedNumb)
			multipleSequences = multipleSequences.Mul(multipleSequences, indexedNumb)
			calcNumbers = calcNumbers.Sub(calcNumbers, sequenceNumbers[j]).Div(calcNumbers, big.NewInt(10))
		}

		if len(sequenceNumbers) < sequenceLen {
			numbers = numbers.Div(numbers, big.NewInt(int64(math.Pow10(len(sequenceNumbers)+1))))
			continue
		}

		if multipleSequences.Cmp(largestMultiple) == 1 {
			largestMultiple = multipleSequences
			largestSequence = sequenceNumbers
		}

		numbers = numbers.Div(numbers, big.NewInt(10))
	}

	return SwapOrder(largestSequence)
}
