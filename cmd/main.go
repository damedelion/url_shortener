package main

import (
	"fmt"
	"math/rand"
	"slices"
)

const CharsCount int64 = 63
const ShortURLLength int64 = 10

func main() {
	var in string
	fmt.Scanf("%s", &in)

	variationsCount := PowInt64(CharsCount, ShortURLLength)

	id := rand.Int63n(variationsCount)

	res, _ := ToBase63(id)
	fmt.Println(res)
}

func ToBase63(id int64) (string, error) {
	res := make([]byte, ShortURLLength)
	remainder := byte('0')

	for i := int64(0); i < ShortURLLength; i++ {
		if id > 0 {
			remainder = byte(id % CharsCount)
			var err error
			remainder, err = Base63Table(remainder)
			if err != nil {
				return "", fmt.Errorf("bad id")
			}
			id /= CharsCount
		} else {
			remainder = byte('0')
		}

		res[i] = remainder
	}

	slices.Reverse(res)

	return string(res), nil

}

func Base63Table(num byte) (byte, error) {
	if num >= 0 && num < 10 {
		return byte('0') + num, nil
	} else if num >= 10 && num < 36 {
		inc := byte(num - 10)
		return byte('a') + inc, nil
	} else if num >= 36 && num < 62 {
		inc := byte(num - 36)
		return byte('A') + inc, nil
	} else if num == 62 {
		return byte('_'), nil
	} else {
		return 0, fmt.Errorf("num is > 63")
	}
}

func PowInt64(x, y int64) int64 {
	res := int64(1)
	for i := int64(0); i < y; i++ {
		res *= int64(x)
	}

	return res
}
