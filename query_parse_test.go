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
			Expected: finger.SomeQueryUser("dariush"),
		},



		{
			QueryString: "@example.com",
			Expected: finger.SomeQueryHost("example.com"),
		},
		{
			QueryString: "@example.com@something.social",
			Expected: finger.SomeQueryHosts("example.com", "something.social"),
		},



		{
			QueryString: "@example.com:1971",
			Expected: finger.SomeQueryHostPort("example.com", 1971),
		},
		{
			QueryString: "@example.com:1971@something.social",
			Expected: finger.SomeQueryAddresses(
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddressHost("something.social"),
			),
		},
		{
			QueryString: "@example.com@something.social:79",
			Expected: finger.SomeQueryAddresses(
				finger.CreateAddressHost("example.com"),
				finger.CreateAddress("something.social", 79),
			),
		},
		{
			QueryString: "@example.com:1971@something.social:79",
			Expected: finger.SomeQueryAddresses(
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddress("something.social", 79),
			),
		},



		{
			QueryString: "dariush@example.com",
			Expected: finger.SomeQueryUserHost("dariush", "example.com"),
		},
		{
			QueryString: "dariush@example.com@something.social",
			Expected: finger.SomeQueryUserHosts("dariush", "example.com", "something.social"),
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
