package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestQuery_String(t *testing.T) {

	tests := []struct{
		Query finger.Query
		Expected string
	}{
		{
			Query: finger.Query{},
			Expected: "",
		},



		{
			Query: finger.SomeQueryUser("dariush"),
			Expected: "dariush",
		},



		{
			Query: finger.SomeQueryUserHost("dariush", "example.com"),
			Expected: "dariush@example.com",
		},
		{
			Query: finger.SomeQueryUserHostPort("dariush", "example.com", 1971),
			Expected: "dariush@example.com:1971",
		},



		{
			Query: finger.SomeQueryUserHosts("dariush", "example.com", "something.social"),
			Expected: "dariush@example.com@something.social",
		},
		{
			Query: finger.SomeQueryUserAddresses(
				"dariush",
				finger.SomeAddress("example.com", 1971),
				finger.SomeAddressHost("something.social"),
			),
			Expected: "dariush@example.com:1971@something.social",
		},
		{
			Query: finger.SomeQueryUserAddresses(
				"dariush",
				finger.SomeAddressHost("example.com"),
				finger.SomeAddress("something.social", 1234),
			),
			Expected: "dariush@example.com@something.social:1234",
		},
		{
			Query: finger.SomeQueryUserAddresses(
				"dariush",
				finger.SomeAddress("example.com", 1971),
				finger.SomeAddress("something.social", 1234),
			),
			Expected: "dariush@example.com:1971@something.social:1234",
		},



		{
			Query: finger.SomeQueryHost("example.com"),
			Expected: "@example.com",
		},
		{
			Query: finger.SomeQueryHostPort("example.com", 1971),
			Expected: "@example.com:1971",
		},
		{
			Query: finger.SomeQueryHosts("example.com", "something.social"),
			Expected: "@example.com@something.social",
		},
		{
			Query: finger.SomeQueryAddresses(
				finger.SomeAddress("example.com", 1971),
				finger.SomeAddressHost("something.social"),
			),
			Expected: "@example.com:1971@something.social",
		},
		{
			Query: finger.SomeQueryAddresses(
				finger.SomeAddressHost("example.com"),
				finger.SomeAddress("something.social", 1234),
			),
			Expected: "@example.com@something.social:1234",
		},
		{
			Query: finger.SomeQueryAddresses(
				finger.SomeAddress("example.com", 1971),
				finger.SomeAddress("something.social", 1234),
			),
			Expected: "@example.com:1971@something.social:1234",
		},



		{
			Query: finger.SomeQueryHosts("once"),
			Expected: "@once",
		},
		{
			Query: finger.SomeQueryHosts("once", "twice"),
			Expected: "@once@twice",
		},
		{
			Query: finger.SomeQueryHosts("once", "twice", "thrice"),
			Expected: "@once@twice@thrice",
		},
		{
			Query: finger.SomeQueryHosts("once", "twice", "thrice", "fource"),
			Expected: "@once@twice@thrice@fource",
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual   string = test.Query.String()

		if expected != actual {
			t.Errorf("For test #%d, the actual query string value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}
