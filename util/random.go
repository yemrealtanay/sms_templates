package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUserID() int32 {
	return RandomInt(1, 10)
}

func RandomName() string {
	return RandomString(6)
}

func RandomSubject() string {
	return RandomString(15)
}

func RandomSentence() string {
	return fmt.Sprintf("%s %s %s %s %s ", RandomString(5), RandomString(6), RandomString(4), RandomString(8), RandomString(7))
}

func RandomBool() bool {
	number1 := RandomInt(1, 10)
	number2 := RandomInt(1, 10)
	if number1 > number2 {
		return true
	} else {
		return false
	}
}

func RandomKey() string {
	return fmt.Sprintf("%s_%s_%s", RandomString(5), RandomString(6), RandomString(4))
}
