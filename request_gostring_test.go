package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestRequest_GoString(t *testing.T) {

	tests := []struct{
		FingerRequest finger.Request
		Expected string
	}{
		{
			FingerRequest: finger.Request{},
			Expected: `finger.Request{}`,
		},



		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("W"),
			},
			Expected: `finger.Request{ Switch: finger.SomeSwitch("W") }`,
		},
		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("PULL"),
			},
			Expected: `finger.Request{ Switch: finger.SomeSwitch("PULL") }`,
		},



		{
			FingerRequest: finger.Request{
				Target: finger.SomeTarget("joeblow"),
			},
			Expected: `finger.Request{ Target: finger.SomeTarget("joeblow") }`,
		},
		{
			FingerRequest: finger.Request{
				Target: finger.SomeTarget("dariush"),
			},
			Expected: `finger.Request{ Target: finger.SomeTarget("dariush") }`,
		},



		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("W"),
				Target: finger.SomeTarget("joeblow"),
			},
			Expected: `finger.Request{ Switch: finger.SomeSwitch("W"), Target: finger.SomeTarget("joeblow") }`,
		},
		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("W"),
				Target: finger.SomeTarget("dariush"),
			},
			Expected: `finger.Request{ Switch: finger.SomeSwitch("W"), Target: finger.SomeTarget("dariush") }`,
		},



		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("PULL"),
				Target: finger.SomeTarget("joeblow"),
			},
			Expected: `finger.Request{ Switch: finger.SomeSwitch("PULL"), Target: finger.SomeTarget("joeblow") }`,
		},
		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("PULL"),
				Target: finger.SomeTarget("dariush"),
			},
			Expected: `finger.Request{ Switch: finger.SomeSwitch("PULL"), Target: finger.SomeTarget("dariush") }`,
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual string = test.FingerRequest.GoString()

		if expected != actual {
			t.Errorf("For test #%d, the actual finger-protocol request (string) is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
