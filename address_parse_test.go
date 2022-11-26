package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestParseAddress(t *testing.T) {

	tests := []struct{
		AddressString string
		Expected finger.Address
	}{
		{
			AddressString: "",
			Expected: finger.Address{},
		},



		{
			AddressString: "example.com",
			Expected: finger.SomeAddressHost("example.com"),
		},



		{
			AddressString: "once.com",
			Expected: finger.SomeAddressHost("once.com"),
		},
		{
			AddressString: "twice.net",
			Expected: finger.SomeAddressHost("twice.net"),
		},
		{
			AddressString: "thrice.org",
			Expected: finger.SomeAddressHost("thrice.org"),
		},
		{
			AddressString: "fource.dev",
			Expected: finger.SomeAddressHost("fource.dev"),
		},



		{
			AddressString: "once.com:79",
			Expected: finger.SomeAddress("once.com", 79),
		},
		{
			AddressString: "twice.net:1079",
			Expected: finger.SomeAddress("twice.net", 1079),
		},
		{
			AddressString: "thrice.org:1971",
			Expected: finger.SomeAddress("thrice.org", 1971),
		},
		{
			AddressString: "fource.dev:7979",
			Expected: finger.SomeAddress("fource.dev", 7979),
		},
	}

	for testNumber, test := range tests {

		expected := test.Expected

		actual, err := finger.ParseAddress(test.AddressString)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}

		if expected != actual {
			t.Errorf("For test #%d, the actual finger-protocol address is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
