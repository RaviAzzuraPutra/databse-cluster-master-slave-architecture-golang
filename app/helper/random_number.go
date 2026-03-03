package helper

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GenerateRandomNumber() string {

	min := 100000
	max := 999999

	rangeVal := big.NewInt(int64(max - min + 1))

	number, _ := rand.Int(rand.Reader, rangeVal)

	result := number.Int64() + int64(min)

	return strconv.FormatInt(result, 10)

}
