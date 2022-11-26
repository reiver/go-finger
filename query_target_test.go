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
			Query: finger.CreateQueryUser("dariush"),
			Expected: finger.CreateTarget("dariush"),
		},



		{
			Query: finger.CreateQueryUserHost("dariush", "example.com"),
			Expected: finger.CreateTarget("dariush@example.com"),
		},
		{
			Query: finger.CreateQueryUserHostPort("dariush", "example.com", 1971),
			Expected: finger.CreateTarget("dariush@example.com:1971"),
		},



		{
			Query: finger.CreateQueryUserHosts("dariush", "example.com", "something.social"),
			Expected: finger.CreateTarget("dariush@example.com@something.social"),
		},
		{
			Query: finger.AssembleQueryUserAddresses(
				finger.CreateUser("dariush"),
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddressHost("something.social"),
			),
			Expected: finger.CreateTarget("dariush@example.com:1971@something.social"),
		},
		{
			Query: finger.AssembleQueryUserAddresses(
				finger.CreateUser("dariush"),
				finger.CreateAddressHost("example.com"),
				finger.CreateAddress("something.social", 1234),
			),
			Expected: finger.CreateTarget("dariush@example.com@something.social:1234"),
		},
		{
			Query: finger.AssembleQueryUserAddresses(
				finger.CreateUser("dariush"),
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddress("something.social", 1234),
			),
			Expected: finger.CreateTarget("dariush@example.com:1971@something.social:1234"),
		},



		{
			Query: finger.CreateQueryHost("example.com"),
			Expected: finger.CreateTarget("@example.com"),
		},
		{
			Query: finger.CreateQueryHostPort("example.com", 1971),
			Expected: finger.CreateTarget("@example.com:1971"),
		},
		{
			Query: finger.CreateQueryHosts("example.com", "something.social"),
			Expected: finger.CreateTarget("@example.com@something.social"),
		},
		{
			Query: finger.AssembleQueryAddresses(
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddressHost("something.social"),
			),
			Expected: finger.CreateTarget("@example.com:1971@something.social"),
		},
		{
			Query: finger.AssembleQueryAddresses(
				finger.CreateAddressHost("example.com"),
				finger.CreateAddress("something.social", 1234),
			),
			Expected: finger.CreateTarget("@example.com@something.social:1234"),
		},
		{
			Query: finger.AssembleQueryAddresses(
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddress("something.social", 1234),
			),
			Expected: finger.CreateTarget("@example.com:1971@something.social:1234"),
		},



		{
			Query: finger.CreateQueryHost("once"),
			Expected: finger.CreateTarget("@once"),
		},
		{
			Query: finger.CreateQueryHosts("once", "twice"),
			Expected: finger.CreateTarget("@once@twice"),
		},
		{
			Query: finger.CreateQueryHosts("once", "twice", "thrice"),
			Expected: finger.CreateTarget("@once@twice@thrice"),
		},
		{
			Query: finger.CreateQueryHosts("once", "twice", "thrice", "fource"),
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
