package finger_test

import (
	"github.com/reiver/go-finger"

	"reflect"

	"testing"
)

func TestParseQuery(t *testing.T) {

	tests := []struct{
		QueryString string
		Expected finger.Query
	}{
		{
			QueryString: "",
			Expected: finger.Query{},
		},



		{
			QueryString: "dariush",
			Expected: finger.Query{
				User: finger.SomeUser("dariush"),
			},
		},



		{
			QueryString: "@example.com",
			Expected: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
				},
			},
		},
		{
			QueryString: "@example.com@something.social",
			Expected: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddressHost("something.social"),
				},
			},
		},



		{
			QueryString: "@example.com:1971",
			Expected: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
				},
			},
		},
		{
			QueryString: "@example.com:1971@something.social",
			Expected: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddressHost("something.social"),
				},
			},
		},
		{
			QueryString: "@example.com@something.social:79",
			Expected: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddress("something.social", 79),
				},
			},
		},
		{
			QueryString: "@example.com:1971@something.social:79",
			Expected: finger.Query{
				Addresses: []finger.Address{
					finger.SomeAddress("example.com", 1971),
					finger.SomeAddress("something.social", 79),
				},
			},
		},



		{
			QueryString: "dariush@example.com",
			Expected: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
				},
			},
		},
		{
			QueryString: "dariush@example.com@something.social",
			Expected: finger.Query{
				User: finger.SomeUser("dariush"),
				Addresses: []finger.Address{
					finger.SomeAddressHost("example.com"),
					finger.SomeAddressHost("something.social"),
				},
			},
		},
	}

	for testNumber, test := range tests {

		var expected finger.Query = test.Expected

		var actual finger.Query
		var err error
		actual, err = finger.ParseQuery(test.QueryString)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual query is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
