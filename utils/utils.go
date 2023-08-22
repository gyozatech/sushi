package utils

import (
	"encoding/json"
	"math/rand"
	"strings"
)

// RandomString generates random string of given lenght
func RandomString(length int) string {
	return RandomStringFrom("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", length)
}

// RandomPassword generate a random password given the desired length
func RandomPassword(length int) string {
	return RandomStringFrom("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{};:',.<>?/", length)
}

// RandomStringFrom generate a random string given the set of characters and the desired length
func RandomStringFrom(charset string, length int) string {
	var characters = []rune(charset)
	str := make([]rune, length)
	for i := range str {
		str[i] = characters[mrand.Intn(len(characters))]
	}
	return string(str)
}

// IsEqual compares two values regardless of their types for testing purpose
func IsEqual(a, b interface{}) bool {
	A, _ := json.Marshal(a)
	B, _ := json.Marshal(b)

	return strings.ReplaceAll(string(A), "\"", "") == strings.ReplaceAll(string(B), "\"", "")
}

// CollectResults returns a slice collecting the results of a function:
// for example: CollectResults(json.Marshal(a))[0] will fetch the first result and discard the error
func CollectResults(results ...interface{}) []interface{} {
	return results
}
