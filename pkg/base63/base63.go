package base63

import (
	"fmt"
	"slices"
)

const CharsCount int64 = 63

func ToBase63(id int64, resLength int) (string, error) {
	res := make([]byte, resLength)
	remainder := byte('0')

	for i := 0; i < resLength; i++ {
		if id > 0 {
			remainder = byte(id % CharsCount)
			var err error
			remainder, err = base63Table(remainder)
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

func base63Table(num byte) (byte, error) {
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
