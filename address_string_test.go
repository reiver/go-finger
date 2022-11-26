package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestAddress_String(t *testing.T) {

	tests := []struct{
		Address finger.Address
		Expected string
	}{
		{
			Address: finger.SomeAddress("example.com", 1971),
			Expected: "example.com:1971",
		},



		{
			Address: finger.SomeAddressHost("example.com"),
			Expected: "example.com",
		},



		{
			Address: finger.SomeAddressPort(1971),
			Expected: ":1971",
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual   string = test.Address.String()

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}
