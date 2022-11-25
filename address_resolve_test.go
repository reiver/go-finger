package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestAddress_Resolve(t *testing.T) {

	tests := []struct{
		Address finger.Address
		Expected string
	}{
		{
			Address: finger.Address{
				Host: finger.SomeHost("example.com"),
				Port: finger.SomePort(1971),
			},
			Expected: "example.com:1971",
		},



		{
			Address: finger.Address{
				Host: finger.SomeHost("example.com"),
			},
			Expected: "example.com:79",
		},



		{
			Address: finger.Address{
				Port: finger.SomePort(1971),
			},
			Expected: "127.0.0.1:1971",
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual   string = test.Address.Resolve()

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}
