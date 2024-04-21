package finger_test

import (
	"github.com/reiver/go-finger"

	"reflect"

	"testing"
)

func TestParseRFC1288Query(t *testing.T) {

	tests := []struct{
		QueryString string
		Expected finger.Query
	}{
		{
			QueryString: "",
			Expected: finger.EmptyQuery(),
		},



		{
			QueryString: "dariush",
			Expected: finger.CreateQueryActor("dariush"),
		},



		{
			QueryString: "@example.com",
			Expected: finger.CreateQueryHost("example.com"),
		},
		{
			QueryString: "@example.com@something.social",
			Expected: finger.CreateQueryHosts("example.com", "something.social"),
		},



		{
			QueryString: "@example.com:1971",
			Expected: finger.CreateQueryHostPort("example.com", 1971),
		},
		{
			QueryString: "@example.com:1971@something.social",
			Expected: finger.AssembleQueryAddresses(
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddressHost("something.social"),
			),
		},
		{
			QueryString: "@example.com@something.social:79",
			Expected: finger.AssembleQueryAddresses(
				finger.CreateAddressHost("example.com"),
				finger.CreateAddress("something.social", 79),
			),
		},
		{
			QueryString: "@example.com:1971@something.social:79",
			Expected: finger.AssembleQueryAddresses(
				finger.CreateAddress("example.com", 1971),
				finger.CreateAddress("something.social", 79),
			),
		},



		{
			QueryString: "dariush@example.com",
			Expected: finger.CreateQueryActorHost("dariush", "example.com"),
		},
		{
			QueryString: "dariush@example.com@something.social",
			Expected: finger.CreateQueryActorHosts("dariush", "example.com", "something.social"),
		},
	}

	for testNumber, test := range tests {

		var expected finger.Query = test.Expected

		var actual finger.Query
		var err error
		actual, err = finger.ParseRFC1288Query(test.QueryString)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			t.Logf("QUERY-STRING: %q", test.QueryString)
			continue
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual query is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("QUERY-STRING: %q", test.QueryString)
			continue
		}
	}
}
