package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestAddress_GoString(t *testing.T) {

	tests := []struct{
		Address finger.Address
		Expected string
	}{
		{
			Address:   finger.CreateAddress("example.com", 1971),
			Expected: `finger.CreateAddress("example.com", 1971)`,
		},



		{
			Address:   finger.CreateAddressHost("example.com"),
			Expected: `finger.CreateAddressHost("example.com")`,
		},



		{
			Address:   finger.CreateAddressPort(1971),
			Expected: `finger.CreateAddressPort(1971)`,
		},



		{
			Address:   finger.EmptyAddress(),
			Expected: `finger.EmptyAddress()`,
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual   string = test.Address.GoString()

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}
