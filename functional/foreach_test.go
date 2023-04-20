package functional

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestForEach(t *testing.T) {
	testName := "TestForEach"

	type DB struct {
		Type    string
		Version string
	}

	var fetchDBFromCode Function[string, DB] = func(dbCode string) (*DB, error) {
		split := strings.Split(dbCode, "-")
		if len(split) != 2 {
			return nil, fmt.Errorf("invalid DB version provided %s", dbCode)
		}
		return &DB{Type: split[0], Version: split[1]}, nil
	}

	testCases := []struct {
		description   string
		input         []string
		expectedSlice []DB
		expectedErr   error
	}{
		{
			description:   "Null input slice",
			input:         nil,
			expectedSlice: []DB{},
			expectedErr:   nil,
		},
		{
			description:   "Empty input slice",
			input:         []string{},
			expectedSlice: []DB{},
			expectedErr:   nil,
		},
		{
			description:   "Input slice with invalid version",
			input:         []string{"mysql-v8.5.6", "postgresv5.3.1", "postgres-v4.2.4", "mysql-v11.2.3"},
			expectedSlice: nil,
			expectedErr:   fmt.Errorf("invalid DB version provided postgresv5.3.1"),
		},
		{
			description: "Input slice with all valid versions",
			input:       []string{"mysql-v8.5.6", "postgres-v5.3.1", "postgres-v4.2.4", "mysql-v11.2.3"},
			expectedSlice: []DB{
				{
					Type:    "mysql",
					Version: "v8.5.6",
				},
				{
					Type:    "postgres",
					Version: "v5.3.1",
				},
				{
					Type:    "postgres",
					Version: "v4.2.4",
				},
				{
					Type:    "mysql",
					Version: "v11.2.3",
				},
			},
			expectedErr: nil,
		},
	}
	for _, testCase := range testCases {
		actualSlice, actualErr := ForEach(testCase.input, fetchDBFromCode)
		if !reflect.DeepEqual(testCase.expectedSlice, actualSlice) {
			t.Errorf("%s failed (%s), expected slice %+v got %+v", testName, testCase.description, testCase.expectedSlice, actualSlice)
		}
		if actualErr != nil && testCase.expectedErr.Error() != actualErr.Error() {
			t.Errorf("%s failed (%s), expected err %s got %s", testName, testCase.description, testCase.expectedErr, actualErr)
		}
	}
}
