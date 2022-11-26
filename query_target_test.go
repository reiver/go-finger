package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestQuery_Target(t *testing.T) {

	tests := []struct{
		Query finger.Query
		Expected finger.Target
	}{
		{
			Query: finger.Query{},
			Expected: finger.EmptyTarget(),
		},



		{
			Query: finger.SomeQueryUser("dariush"),
			Expected: finger.CreateTarget("dariush"),
		},



		{
			Query: finger.SomeQueryUserHost("dariush", "example.com"),
			Expected: finger.CreateTarget("dariush@example.com"),
		},
		{
			Query: finger.SomeQueryUserHostPort("dariush", "example.com", 1971),
			Expected: finger.CreateTarget("dariush@example.com:1971"),
		},



		{
			Query: finger.SomeQueryUserHosts("dariush", "example.com", "something.social"),
			Expected: finger.CreateTarget("dariush@example.com@something.social"),
		},
		{
			Query: finger.SomeQueryUserAddresses(
				"dariush",
				finger.SomeAddress("example.com", 1971),
				finger.SomeAddressHost("something.social"),
			),
			Expected: finger.CreateTarget("dariush@example.com:1971@something.social"),
		},
		{
			Query: finger.SomeQueryUserAddresses(
				"dariush",
				finger.SomeAddressHost("example.com"),
				finger.SomeAddress("something.social", 1234),
			),
			Expected: finger.CreateTarget("dariush@example.com@something.social:1234"),
		},
		{
			Query: finger.SomeQueryUserAddresses(
				"dariush",
				finger.SomeAddress("example.com", 1971),
				finger.SomeAddress("something.social", 1234),
			),
			Expected: finger.CreateTarget("dariush@example.com:1971@something.social:1234"),
		},



		{
			Query: finger.SomeQueryHost("example.com"),
			Expected: finger.CreateTarget("@example.com"),
		},
		{
			Query: finger.SomeQueryHostPort("example.com", 1971),
			Expected: finger.CreateTarget("@example.com:1971"),
		},
		{
			Query: finger.SomeQueryHosts("example.com", "something.social"),
			Expected: finger.CreateTarget("@example.com@something.social"),
		},
		{
			Query: finger.SomeQueryAddresses(
				finger.SomeAddress("example.com", 1971),
				finger.SomeAddressHost("something.social"),
			),
			Expected: finger.CreateTarget("@example.com:1971@something.social"),
		},
		{
			Query: finger.SomeQueryAddresses(
				finger.SomeAddressHost("example.com"),
				finger.SomeAddress("something.social", 1234),
			),
			Expected: finger.CreateTarget("@example.com@something.social:1234"),
		},
		{
			Query: finger.SomeQueryAddresses(
				finger.SomeAddress("example.com", 1971),
				finger.SomeAddress("something.social", 1234),
			),
			Expected: finger.CreateTarget("@example.com:1971@something.social:1234"),
		},



		{
			Query: finger.SomeQueryHost("once"),
			Expected: finger.CreateTarget("@once"),
		},
		{
			Query: finger.SomeQueryHosts("once", "twice"),
			Expected: finger.CreateTarget("@once@twice"),
		},
		{
			Query: finger.SomeQueryHosts("once", "twice", "thrice"),
			Expected: finger.CreateTarget("@once@twice@thrice"),
		},
		{
			Query: finger.SomeQueryHosts("once", "twice", "thrice", "fource"),
			Expected: finger.CreateTarget("@once@twice@thrice@fource"),
		},
	}

	for testNumber, test := range tests {

		var expected finger.Target = test.Expected
		var actual   finger.Target = test.Query.Target()

		if expected != actual {
			t.Errorf("For test #%d, the actual query string value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}

	}
}
