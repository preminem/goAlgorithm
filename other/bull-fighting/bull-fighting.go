package bull_fighting

import "errors"

func BullFighting(a, b string) (int, error) {
	if len(a) != 5 || len(b) != 5 {
		return 0, errors.New("parameter error")
	}
	var arrayA, arrayB [5]int
	err := stringToArray(a, &arrayA)
	if err != nil {
		return 0, err
	}
	err = stringToArray(b, &arrayB)
	if err != nil {
		return 0, err
	}
	var nonBullA, nonBullB [2]int
	//determine if a and b are bull
	isBullA, isBullB := bull(arrayA, &nonBullA), bull(arrayB, &nonBullB)
	if isBullA && !isBullB {
		return 1, nil
	} else if !isBullA && isBullB {
		return -1, nil
	} else if !isBullA && !isBullB {
		maxA, maxB := max(arrayA), max(arrayB)
		if maxA > maxB {
			return 1, nil
		} else if maxB > maxA {
			return -1, nil
		} else {
			return 0, nil
		}
	} else if (nonBullA[0]+nonBullA[1])%10 == (nonBullB[0]+nonBullB[1])%10 {
		return 0, nil
	} else if (nonBullA[0]+nonBullA[1])%10 == 0 || (nonBullA[0]+nonBullA[1])%10 > (nonBullB[0]+nonBullB[1])%10 {
		return 1, nil
	} else {
		return -1, nil
	}
}

func stringToArray(str string, arr *[5]int) error {
	for i, s := range str {
		var value int
		if s == 'A' {
			value = 1
		} else if s == 'J' || s == 'Q' || s == 'K' || s == 'T' {
			value = 10
		} else if s < 50 || s > 57 {
			return errors.New("parameter error")
		} else {
			value = int(s - '0')
		}
		arr[i] = value
	}
	return nil
}

func bull(array [5]int, nonBull *[2]int) bool {
	isBull := false
I:
	for i := 0; i < len(array)-2; i++ {
		for j := i + 1; j < len(array)-1; j++ {
			for k := j + 1; k < len(array); k++ {
				if (array[i]+array[j]+array[k])%10 == 0 {
					isBull = true
					temp := make([]int, 0)
					for l, v := range array {
						if l != i && l != j && l != k {
							temp = append(temp, v)
						}
					}
					if (temp[0]+temp[1])%10 == 0 {
						nonBull[0], nonBull[1] = temp[0], temp[1]
						break I
					} else if (temp[0]+temp[1])%10 > (nonBull[0]+nonBull[1])%10 {
						nonBull[0], nonBull[1] = temp[0], temp[1]
					}
				}
			}
		}
	}
	return isBull
}

func max(arr [5]int) int {
	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if maxVal > arr[i] {
			maxVal = arr[i]
		}
	}
	return maxVal
}
