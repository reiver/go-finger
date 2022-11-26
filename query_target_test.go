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
			Expected: finger.NoTarget(),
		},



		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
			},
			Expected: finger.SomeTarget("dariush"),
		},



		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
				},
			},
			Expected: finger.SomeTarget("dariush@example.com"),
		},
		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
				},
			},
			Expected: finger.SomeTarget("dariush@example.com:1971"),
		},



		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddressHost("something.social"),
				},
			},
			Expected: finger.SomeTarget("dariush@example.com@something.social"),
		},
		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddressHost("something.social"),
				},
			},
			Expected: finger.SomeTarget("dariush@example.com:1971@something.social"),
		},
		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddress("something.social", 1234),
				},
			},
			Expected: finger.SomeTarget("dariush@example.com@something.social:1234"),
		},
		{
			Query: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddress("something.social", 1234),
				},
			},
			Expected: finger.SomeTarget("dariush@example.com:1971@something.social:1234"),
		},



		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
				},
			},
			Expected: finger.SomeTarget("@example.com"),
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
				},
			},
			Expected: finger.SomeTarget("@example.com:1971"),
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddressHost("something.social"),
				},
			},
			Expected: finger.SomeTarget("@example.com@something.social"),
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddressHost("something.social"),
				},
			},
			Expected: finger.SomeTarget("@example.com:1971@something.social"),
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddress("something.social", 1234),
				},
			},
			Expected: finger.SomeTarget("@example.com@something.social:1234"),
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddress("something.social", 1234),
				},
			},
			Expected: finger.SomeTarget("@example.com:1971@something.social:1234"),
		},



		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("once"),
				},
			},
			Expected: finger.SomeTarget("@once"),
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("once"),
					finger.SomeAddressHost("twice"),
				},
			},
			Expected: finger.SomeTarget("@once@twice"),
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("once"),
					finger.SomeAddressHost("twice"),
					finger.SomeAddressHost("thrice"),
				},
			},
			Expected: finger.SomeTarget("@once@twice@thrice"),
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
			Expected: finger.SomeTarget("@once@twice@thrice@fource"),
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
