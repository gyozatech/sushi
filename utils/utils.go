package utils

import (
	"encoding/json"
	"math/rand"
	"strings"
)

// RandomString generates random string of given lenght
func RandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	str := make([]rune, length)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
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
