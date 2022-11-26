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
			FingerRequest: finger.SomeRequestSwitch("W"),
			Expected:     `finger.SomeRequestSwitch("W")`,
		},
		{
			FingerRequest: finger.SomeRequestSwitch("PULL"),
			Expected:     `finger.SomeRequestSwitch("PULL")`,
		},



		{
			FingerRequest: finger.SomeRequestTarget("joeblow"),
			Expected:     `finger.SomeRequestTarget("joeblow")`,
		},
		{
			FingerRequest: finger.SomeRequestTarget("dariush"),
			Expected:     `finger.SomeRequestTarget("dariush")`,
		},



		{
			FingerRequest: finger.SomeRequestTarget("joeblow@example.com"),
			Expected:     `finger.SomeRequestTarget("joeblow@example.com")`,
		},
		{
			FingerRequest: finger.SomeRequestTarget("dariush@changelog.ca"),
			Expected:     `finger.SomeRequestTarget("dariush@changelog.ca")`,
		},



		{
			FingerRequest: finger.SomeRequestTarget("joeblow@example.com@something.social"),
			Expected:     `finger.SomeRequestTarget("joeblow@example.com@something.social")`,
		},
		{
			FingerRequest: finger.SomeRequestTarget("dariush@changelog.ca@example.dev"),
			Expected:     `finger.SomeRequestTarget("dariush@changelog.ca@example.dev")`,
		},



		{
			FingerRequest: finger.SomeRequest("W", "joeblow"),
			Expected:     `finger.SomeRequest("W", "joeblow")`,
		},
		{
			FingerRequest: finger.SomeRequest("W", "dariush"),
			Expected:     `finger.SomeRequest("W", "dariush")`,
		},



		{
			FingerRequest: finger.SomeRequest("PULL", "joeblow"),
			Expected:     `finger.SomeRequest("PULL", "joeblow")`,
		},
		{
			FingerRequest: finger.SomeRequest("PULL", "dariush"),
			Expected:     `finger.SomeRequest("PULL", "dariush")`,
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
