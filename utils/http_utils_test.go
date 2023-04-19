package utils

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestValidateIPv6(t *testing.T) {
	ipv6Addresses := []string{
		"2001:4860:4860::8888",
		"2001:4860:4860::8844",
		"2345:0425:2CA1:0000:0000:0567:5673:23b5",
		"2345:0425:2CA1:0000:0000:0567::",
		"2345:0425:2CA1:0:0:0567:5673:23b5",
		"2345:0425:2CA1::0000:0567:5673:23b5",
		"2345:0425:2CA1::::0567:5673:23b5",
		"FF01::1",
		"FF01:0:0:0:0:0:0:1",
		"FF01:0000:0000:0000:0000:0000:0000:0001",
		"2001:db8:1234:ffff:ffff:ffff:ffff:ffff",
		"2001:2f:ffff:ffff:ffff:ffff:ffff:ffff",
		"2001::ffff:ffff:ffff:ffff:ffff:ffff",
		"100::ffff:ffff:ffff:ffff",
		"64:ff9b:1:ffff:ffff:ffff:255.255.255.255",
		"::1",
		"2001::",
		"2001:20::",
		"2001:db8::",
		"::",
		"::ffff:0:255.255.255.255",
		"fe80::ffff:ffff:ffff:ffff",
		"::ffff:0:255.255.255.255",
	}
	ipv6PrefixedAddresses := []string{
		"2001:AACD:3B:23::/64",
		"0000::/3",
		"2001:db8:a::/64",
		"::/0",
		"::1/128",
		"::ffff:0:0/96",
		"::ffff:0:0:0/96",
		"64:ff9b::/96",
		"64:ff9b:1::/48",
		"100::/64",
		"2001:0000::/32",
		"fc00::/7",
		"2001:4860:4860::8888/12",
		"2001:4860:4860::8844/154",
		"2345:0425:2CA1:0000:0000:0567:5673:23b5/56",
		"2345:0425:2CA1:0000:0000:0567::/1",
		"2345:0425:2CA1:0:0:0567:5673:23b5/254",
		"2345:0425:2CA1::0000:0567:5673:23b5/155",
		"2345:0425:2CA1::::0567:5673:23b5/145",
	}
	ipv4Addresses := []string{
		"127.0.0.1",
		"77.54.11.1/32",
		"192.168.0.1/24",
		"172.15.0.0/16",
	}

	for _, address := range ipv6Addresses {
		if !ValidateIPv6(address, false) {
			t.Errorf("Address %s should be considered as an v6 IP address\n", address)
		}
	}
	for _, address := range ipv6PrefixedAddresses {
		if !ValidateIPv6(address, true) {
			t.Errorf("Address %s should be considered as a prefixed v6 IP address\n", address)
		}
	}
	for _, address := range ipv4Addresses {
		if ValidateIPv6(address, false) || ValidateIPv6(address, true) {
			t.Errorf("Address %s should be considered as an v4 IP address\n", address)
		}
	}

}

func TestValidateIPv4(t *testing.T) {
	ipv4Addresses := []string{
		"0.0.0.0",
		"0.255.255.255",
		"10.0.0.0",
		"10.255.255.255",
		"100.64.0.0",
		"100.127.255.255",
		"127.0.0.0",
		"127.255.255.255",
		"8.8.8.8",
		"1.1.1.1",
	}
	ipv4PrefixedAddresses := []string{
		"0.0.0.0/8",
		"10.0.0.0/8",
		"100.64.0.0/10",
		"127.0.0.0/8",
		"169.254.0.0/16",
		"198.51.100.0/24",
		"255.255.255.255/32",
		"240.0.0.0/4",
	}
	ipv6Addresses := []string{
		"2001:AACD:3B:23::/64",
		"0000::/3",
		"2001:db8:a::/64",
		"::/0",
		"::1/128",
		"::ffff:0:0/96",
		"100::ffff:ffff:ffff:ffff",
		"64:ff9b:1:ffff:ffff:ffff:255.255.255.255",
		"::1",
		"2001::",
		"2001:20::",
		"2001:db8::",
	}
	for _, address := range ipv4Addresses {
		if !ValidateIPv4(address, false) {
			t.Errorf("Address %s should be considered as an v4 IP address\n", address)
		}
	}
	for _, address := range ipv4PrefixedAddresses {
		if !ValidateIPv4(address, true) {
			t.Errorf("Address %s should be considered as a prefixed v4 IP address\n", address)
		}
	}
	for _, address := range ipv6Addresses {
		if ValidateIPv4(address, false) || ValidateIPv4(address, true) {
			t.Errorf("Address %s should be considered as an v6 IP address\n", address)
		}
	}
}

func TestIsEqual(t *testing.T) {
	testName := "TestIsEqual"

	type testCase struct {
		description string
		inputA      interface{}
		inputB      interface{}
		expected    bool
	}

	testCases := []testCase{
		{
			description: "Comparing an integer and an int64",
			inputA:      3,
			inputB:      int64(3),
			expected:    true,
		},
		{
			description: "Comparing an unit16 and an int64",
			inputA:      uint16(3),
			inputB:      int64(3),
			expected:    true,
		},
		{
			description: "Comparing an integer in string format and an int64",
			inputA:      "3",
			inputB:      int64(3),
			expected:    true,
		},
		{
			description: "Comparing a pointer to a integer in string format and a pointer to an int64",
			inputA:      PointerTo("3"),
			inputB:      PointerTo(int64(3)),
			expected:    true,
		},
		{
			description: "Comparing a pointer to a float64 and a pointer to an int32",
			inputA:      PointerTo(float64(3)),
			inputB:      PointerTo(int32(3)),
			expected:    true,
		},
		{
			description: "Comparing a string and a pointer to a string",
			inputA:      "hello",
			inputB:      PointerTo("hello"),
			expected:    true,
		},
		{
			description: "Comparing two nil interfaces",
			inputA:      nil,
			inputB:      nil,
			expected:    true,
		},
		{
			description: "Comparing two errors",
			inputA:      fmt.Errorf("error!"),
			inputB:      errors.New("error!"),
			expected:    true,
		},
		{
			description: "Comparing a bool and a bool in string format",
			inputA:      true,
			inputB:      "true",
			expected:    true,
		},
		{
			description: "Comparing a pointer to a bool and a pointer to a bool in string format",
			inputA:      PointerTo(true),
			inputB:      PointerTo("true"),
			expected:    true,
		},
		{
			description: "Comparing a float in string format and a float",
			inputA:      "3.4",
			inputB:      float32(3.4),
			expected:    true,
		},
		{
			description: "Comparing a pointer to a float in string format and a pointer to a float",
			inputA:      PointerTo("3.4"),
			inputB:      PointerTo(float32(3.4)),
			expected:    true,
		},
		{
			description: "Comparing a slice of integers in string format and a slice of integers",
			inputA:      []string{"1", "2", "3"},
			inputB:      []int{1, 2, 3},
			expected:    true,
		},
		{
			description: "Comparing a slice of integers in string format and a slice of pointers to integers",
			inputA:      []string{"1", "2", "3"},
			inputB:      []*int{PointerTo(1), PointerTo(2), PointerTo(3)},
			expected:    true,
		},
		{
			description: "Comparing a map of pointers of integers in string format and a map of integers",
			inputA:      map[string]*string{"1": PointerTo("1"), "2": PointerTo("2")},
			inputB:      map[int]int{1: 1, 2: 2},
			expected:    true,
		},
		{
			description: "Comparing two different strings",
			inputA:      "hello",
			inputB:      "h1",
			expected:    false,
		},
		{
			description: "Comparing two different integers",
			inputA:      3,
			inputB:      2,
			expected:    false,
		},
		{
			description: "Comparing two different booleans",
			inputA:      true,
			inputB:      false,
			expected:    false,
		},
		{
			description: "Comparing two different slices",
			inputA:      []int{1, 2, 3},
			inputB:      []int{1, 2},
			expected:    false,
		},
	}
	for _, test := range testCases {
		if actual := IsEqual(test.inputA, test.inputB); actual != test.expected {
			t.Errorf("%s failed (%s): expected %t, found %t", testName, test.description, test.expected, actual)
		}
	}
}

func TestEnrichContextWithValue(t *testing.T) {
	testName := "TestEnrichContextWithValue"

	type Book struct {
		Title  string
		Author string
	}

	ctx := context.TODO()
	ctx = EnrichContextWithValue(ctx, "name", "Alessandro")
	ctx = EnrichContextWithValue(ctx, "age", 35)
	ctx = EnrichContextWithValue(ctx, "subscribed", true)
	ctx = EnrichContextWithValue(ctx, "book", Book{Title: "Jurassic Park", Author: "Michael Chricton"})

	var name *string = FetchContextValue[string](ctx, "name")
	var age *int = FetchContextValue[int](ctx, "age")
	var subscribed *bool = FetchContextValue[bool](ctx, "subscribed")
	var book *Book = FetchContextValue[Book](ctx, "book")
	var unexisting = FetchContextValue[string](ctx, "unexisting")

	if name == nil || *name != "Alessandro" {
		t.Errorf("%s failed: expected 'name' %s", testName, "Alessandro")
	}
	if age == nil || *age != 35 {
		t.Errorf("%s failed: expected 'age' %d", testName, 35)
	}
	if subscribed == nil || !*subscribed {
		t.Errorf("%s failed: expected 'age' %t", testName, true)
	}
	if book == nil || *book != (Book{Title: "Jurassic Park", Author: "Michael Chricton"}) {
		t.Errorf("%s failed: expected 'book' %+v", testName, Book{Title: "Jurassic Park", Author: "Michael Chricton"})
	}
	if unexisting != nil {
		t.Errorf("%s failed: expected 'unesisting' to be nil. Got %v", testName, unexisting)
	}
}
