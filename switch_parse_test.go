package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestParseSwitch(t *testing.T) {

	tests := []struct{
		SwitchString string
		Expected finger.Switch
	}{
		{
			SwitchString:                 "/W",
			Expected: finger.CreateSwitch("/W"),
		},



		{
			SwitchString:                 "/PULL",
			Expected: finger.CreateSwitch("/PULL"),
		},



		{
			SwitchString:                 "/once/twice/thrice/fource",
			Expected: finger.CreateSwitch("/once/twice/thrice/fource"),
		},



		{
			SwitchString:                 "/",
			Expected: finger.CreateSwitch("/"),
		},
	}

	for testNumber, test := range tests {

		var expected finger.Switch = test.Expected

		actual, err := finger.ParseSwitch(test.SwitchString)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected." , testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
