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
			FingerRequest: finger.EmptyRequest(),
			Expected:     `finger.EmptyRequest()`,
		},



		{
			FingerRequest: finger.CreateRequestSwitch("/W"),
			Expected:     `finger.CreateRequestSwitch("/W")`,
		},
		{
			FingerRequest: finger.CreateRequestSwitch("/PULL"),
			Expected:     `finger.CreateRequestSwitch("/PULL")`,
		},



		{
			FingerRequest: finger.CreateRequestTarget("joeblow"),
			Expected:     `finger.CreateRequestTarget("joeblow")`,
		},
		{
			FingerRequest: finger.CreateRequestTarget("dariush"),
			Expected:     `finger.CreateRequestTarget("dariush")`,
		},



		{
			FingerRequest: finger.CreateRequestTarget("joeblow@example.com"),
			Expected:     `finger.CreateRequestTarget("joeblow@example.com")`,
		},
		{
			FingerRequest: finger.CreateRequestTarget("dariush@reiver.link"),
			Expected:     `finger.CreateRequestTarget("dariush@reiver.link")`,
		},



		{
			FingerRequest: finger.CreateRequestTarget("joeblow@example.com@something.social"),
			Expected:     `finger.CreateRequestTarget("joeblow@example.com@something.social")`,
		},
		{
			FingerRequest: finger.CreateRequestTarget("dariush@reiver.link@example.dev"),
			Expected:     `finger.CreateRequestTarget("dariush@reiver.link@example.dev")`,
		},



		{
			FingerRequest: finger.CreateRequest("/W", "joeblow"),
			Expected:     `finger.CreateRequest("/W", "joeblow")`,
		},
		{
			FingerRequest: finger.CreateRequest("/W", "dariush"),
			Expected:     `finger.CreateRequest("/W", "dariush")`,
		},



		{
			FingerRequest: finger.CreateRequest("/PULL", "joeblow"),
			Expected:     `finger.CreateRequest("/PULL", "joeblow")`,
		},
		{
			FingerRequest: finger.CreateRequest("/PULL", "dariush"),
			Expected:     `finger.CreateRequest("/PULL", "dariush")`,
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
