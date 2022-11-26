package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestQuery_ClientParameters(t *testing.T) {

	tests := []struct{
		QueryString string
		ExpectedAddress finger.Address
		ExpectedTarget  finger.Target
	}{
		{
			QueryString: "",
			ExpectedAddress: finger.DefaultAddress(),
			ExpectedTarget: finger.EmptyTarget(),
		},



		{
			QueryString: "dariush",
			ExpectedAddress: finger.DefaultAddress(),
			ExpectedTarget: finger.SomeTarget("dariush"),
		},



		{
			QueryString: "dariush@example.com",
			ExpectedAddress: finger.SomeAddressHost("example.com"),
			ExpectedTarget: finger.SomeTarget("dariush"),
		},
		{
			QueryString: "dariush@example.com:1971",
			ExpectedAddress: finger.SomeAddress("example.com", 1971),
			ExpectedTarget: finger.SomeTarget("dariush"),
		},



		{
			QueryString: "dariush@example.com@something.social",
			ExpectedAddress: finger.SomeAddressHost("something.social"),
			ExpectedTarget: finger.SomeTarget("dariush@example.com"),
		},
		{
			QueryString: "dariush@example.com:1971@something.social",
			ExpectedAddress: finger.SomeAddressHost("something.social"),
			ExpectedTarget: finger.SomeTarget("dariush@example.com:1971"),
		},
		{
			QueryString: "dariush@example.com@something.social:1234",
			ExpectedAddress: finger.SomeAddress("something.social", 1234),
			ExpectedTarget: finger.SomeTarget("dariush@example.com"),
		},
		{
			QueryString: "dariush@example.com:1971@something.social:1234",
			ExpectedAddress: finger.SomeAddress("something.social", 1234),
			ExpectedTarget: finger.SomeTarget("dariush@example.com:1971"),
		},



		{
			QueryString: "@example.com",
			ExpectedAddress: finger.SomeAddressHost("example.com"),
			ExpectedTarget: finger.EmptyTarget(),
		},
		{
			QueryString: "@example.com:1971",
			ExpectedAddress: finger.SomeAddress("example.com", 1971),
			ExpectedTarget: finger.EmptyTarget(),
		},
		{
			QueryString: "@example.com@something.social",
			ExpectedAddress: finger.SomeAddressHost("something.social"),
			ExpectedTarget: finger.SomeTarget("@example.com"),
		},
		{
			QueryString: "@example.com:1971@something.social",
			ExpectedAddress: finger.SomeAddressHost("something.social"),
			ExpectedTarget: finger.SomeTarget("@example.com:1971"),
		},
		{
			QueryString: "@example.com@something.social:1234",
			ExpectedAddress: finger.SomeAddress("something.social", 1234),
			ExpectedTarget: finger.SomeTarget("@example.com"),
		},
		{
			QueryString: "@example.com:1971@something.social:1234",
			ExpectedAddress: finger.SomeAddress("something.social", 1234),
			ExpectedTarget: finger.SomeTarget("@example.com:1971"),
		},



		{
			QueryString: "@once",
			ExpectedAddress: finger.SomeAddressHost("once"),
			ExpectedTarget: finger.EmptyTarget(),
		},
		{
			QueryString: "@once@twice",
			ExpectedAddress: finger.SomeAddressHost("twice"),
			ExpectedTarget: finger.SomeTarget("@once"),
		},
		{
			QueryString: "@once@twice@thrice",
			ExpectedAddress: finger.SomeAddressHost("thrice"),
			ExpectedTarget: finger.SomeTarget("@once@twice"),
		},
		{
			QueryString: "@once@twice@thrice@fource",
			ExpectedAddress: finger.SomeAddressHost("fource"),
			ExpectedTarget: finger.SomeTarget("@once@twice@thrice"),
		},
	}

	for testNumber, test := range tests {

		var query finger.Query
		{
			var err error

			query, err = finger.ParseQuery(test.QueryString)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				continue
			}
		}

		expectedAddress := test.ExpectedAddress
		expectedTarget  := test.ExpectedTarget
		actualAddress, actualQuery := query.ClientParameters()

		{
			expected := expectedAddress
			actual   := actualAddress

			if expected != actual {
				t.Errorf("For test #%d, the actual address value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}

		{
			var expected finger.Target = expectedTarget
			var actual   finger.Target = actualQuery.Target()

			if expected != actual {
				t.Errorf("For test #%d, the actual target is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}
	}
}
