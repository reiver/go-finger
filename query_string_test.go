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
			Query: finger.EmptyQuery(),
			Expected: "",
		},



		{
			Query: finger.CreateQueryUser("dariush"),
			Expected: "dariush",
		},



		{
			Query: finger.CreateQueryUserPath("dariush", "/path/to/file.ext"),
			Expected: "dariush/path/to/file.ext",
		},



		{
			Query: finger.CreateQueryUserHost("dariush", "example.com"),
			Expected: "dariush@example.com",
		},
		{
			Query: finger.CreateQueryUserHostPort("dariush", "example.com", 1971),
			Expected: "dariush@example.com:1971",
		},



		{
			Query: finger.CreateQueryUserPathHost("dariush", "/path/to/file.ext", "example.com"),
			Expected: "dariush/path/to/file.ext@example.com",
		},
		{
			Query: finger.CreateQueryUserPathHostPort("dariush", "/path/to/file.ext", "example.com", 1971),
			Expected: "dariush/path/to/file.ext@example.com:1971",
		},



		{
			Query: finger.CreateQueryUserHosts("dariush", "example.com", "something.social"),
			Expected: "dariush@example.com@something.social",
		},
		{
			Query: finger.AssembleQueryUserAddresses(
				finger.CreateUser("dariush"),
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddressHost("something.social"),
			),
			Expected: "dariush@example.com:1971@something.social",
		},
		{
			Query: finger.AssembleQueryUserAddresses(
				finger.CreateUser("dariush"),
				finger.CreateAddressHost("example.com"),
				finger.CreateAddress("something.social", 1234),
			),
			Expected: "dariush@example.com@something.social:1234",
		},
		{
			Query: finger.AssembleQueryUserAddresses(
				finger.CreateUser("dariush"),
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddress("something.social", 1234),
			),
			Expected: "dariush@example.com:1971@something.social:1234",
		},



		{
			Query: finger.CreateQueryHost("example.com"),
			Expected: "@example.com",
		},
		{
			Query: finger.CreateQueryHostPort("example.com", 1971),
			Expected: "@example.com:1971",
		},
		{
			Query: finger.CreateQueryHosts("example.com", "something.social"),
			Expected: "@example.com@something.social",
		},
		{
			Query: finger.AssembleQueryAddresses(
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddressHost("something.social"),
			),
			Expected: "@example.com:1971@something.social",
		},
		{
			Query: finger.AssembleQueryAddresses(
				finger.CreateAddressHost("example.com"),
				finger.CreateAddress("something.social", 1234),
			),
			Expected: "@example.com@something.social:1234",
		},
		{
			Query: finger.AssembleQueryAddresses(
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddress("something.social", 1234),
			),
			Expected: "@example.com:1971@something.social:1234",
		},



		{
			Query: finger.CreateQueryHosts("once"),
			Expected: "@once",
		},
		{
			Query: finger.CreateQueryHosts("once", "twice"),
			Expected: "@once@twice",
		},
		{
			Query: finger.CreateQueryHosts("once", "twice", "thrice"),
			Expected: "@once@twice@thrice",
		},
		{
			Query: finger.CreateQueryHosts("once", "twice", "thrice", "fource"),
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
