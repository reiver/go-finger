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
			Expected: finger.Address{
				Host: finger.SomeHost("example.com"),
			},
		},



		{
			AddressString: "once.com",
			Expected: finger.Address{
				Host: finger.SomeHost("once.com"),
			},
		},
		{
			AddressString: "twice.net",
			Expected: finger.Address{
				Host: finger.SomeHost("twice.net"),
			},
		},
		{
			AddressString: "thrice.org",
			Expected: finger.Address{
				Host: finger.SomeHost("thrice.org"),
			},
		},
		{
			AddressString: "fource.dev",
			Expected: finger.Address{
				Host: finger.SomeHost("fource.dev"),
			},
		},



		{
			AddressString: "once.com:79",
			Expected: finger.Address{
				Host: finger.SomeHost("once.com"),
				Port: finger.SomePort(79),
			},
		},
		{
			AddressString: "twice.net:1079",
			Expected: finger.Address{
				Host: finger.SomeHost("twice.net"),
				Port: finger.SomePort(1079),
			},
		},
		{
			AddressString: "thrice.org:1971",
			Expected: finger.Address{
				Host: finger.SomeHost("thrice.org"),
				Port: finger.SomePort(1971),
			},
		},
		{
			AddressString: "fource.dev:7979",
			Expected: finger.Address{
				Host: finger.SomeHost("fource.dev"),
				Port: finger.SomePort(7979),
			},
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
