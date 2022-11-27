package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestSwitch_GoString(t *testing.T) {

	tests := []struct{
		Switch finger.Switch
		Expected string
	}{
		{
			Switch:    finger.CreateSwitch("/W"),
			Expected: `finger.CreateSwitch("/W")`,
		},



		{
			Switch:    finger.CreateSwitch("/PULL"),
			Expected: `finger.CreateSwitch("/PULL")`,
		},



		{
			Switch:    finger.CreateSwitch("/once/twice/thrice/fource"),
			Expected: `finger.CreateSwitch("/once/twice/thrice/fource")`,
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual   string = test.Switch.GoString()

		if expected != actual {
			t.Errorf("For test #%d, the actual value from the finger.Switch.GoString() method is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}
