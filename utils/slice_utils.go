package utils

import "reflect"

// Contains searches if an element is present in a slice
func Contains[T comparable](slice []T, element T) bool {
	if len(slice) == 0 {
		return false
	}
	for _, s := range slice {
		if element == s {
			return true
		}
	}
	return false
}

// TruncateSlice truncates a slice to a maximum given length
func TruncateSlice[T any](slice []T, maxLength int) []T {
	if len(slice) > maxLength {
		return slice[0:maxLength]
	}
	return slice
}

// AppendUnique appends elements to a slice with the guarantee of uniqueneness
func AppendUnique[T comparable](slice []T, elements ...T) []T {
	for _, el := range elements {
		if !Contains(slice, el) {
			slice = append(slice, el)
		}
	}
	return slice
}

// RemoveByIndex removes the element specified by the index from a given slice
func RemoveByIndex[T comparable](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveElement removes element from the list
func RemoveElement[T comparable](list []T, element T) []T {
	newList := []T{}
	for _, e := range list {
		if e != element {
			newList = append(newList, e)
		}
	}
	return newList
}

// RemoveEmptyValueInSlice Function to empty value in a slice
func RemoveEmptyValueInSlice[T comparable](slice []T) []T {
	var result []T
	for _, value := range slice {
		if !reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
			result = append(result, value)
		}
	}
	return result
}

// FindOne filters a slice finding one element given a condition
func FindOne[T any](slice []T, where func(t T) bool) *T {
	for _, item := range slice {
		if where(item) {
			return &item
		}
	}
	return nil
}

// FindMany filters a slice finding all the element which verify a given condition
func FindMany[T any](slice []T, where func(t T) bool) []T {
	result := []T{}
	for _, item := range slice {
		if where(item) {
			result = append(result, item)
		}
	}
	return result
}
