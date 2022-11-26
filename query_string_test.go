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
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
			},
			Expected: "dariush",
		},



		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
				},
			},
			Expected: "dariush@example.com",
		},
		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
				},
			},
			Expected: "dariush@example.com:1971",
		},



		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddressHost("something.social"),
				},
			},
			Expected: "dariush@example.com@something.social",
		},
		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddressHost("something.social"),
				},
			},
			Expected: "dariush@example.com:1971@something.social",
		},
		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddress("something.social", 1234),
				},
			},
			Expected: "dariush@example.com@something.social:1234",
		},
		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddress("something.social", 1234),
				},
			},
			Expected: "dariush@example.com:1971@something.social:1234",
		},



		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
				},
			},
			Expected: "@example.com",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
				},
			},
			Expected: "@example.com:1971",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddressHost("something.social"),
				},
			},
			Expected: "@example.com@something.social",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddressHost("something.social"),
				},
			},
			Expected: "@example.com:1971@something.social",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddress("something.social", 1234),
				},
			},
			Expected: "@example.com@something.social:1234",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddress("something.social", 1234),
				},
			},
			Expected: "@example.com:1971@something.social:1234",
		},



		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("once"),
				},
			},
			Expected: "@once",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("once"),
					finger.SomeAddressHost("twice"),
				},
			},
			Expected: "@once@twice",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("once"),
					finger.SomeAddressHost("twice"),
					finger.SomeAddressHost("thrice"),
				},
			},
			Expected: "@once@twice@thrice",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("once"),
					finger.SomeAddressHost("twice"),
					finger.SomeAddressHost("thrice"),
					finger.SomeAddressHost("fource"),
				},
			},
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
