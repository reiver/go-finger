package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestRequest_Sentence(t *testing.T) {

	tests := []struct{
		FingerRequest finger.Request
		Expected string
	}{
		{
			FingerRequest: finger.EmptyRequest(),
			Expected: "",
		},



		{
			FingerRequest: finger.CreateRequestSwitch("/W"),
			Expected: "/W",
		},
		{
			FingerRequest: finger.CreateRequestSwitch("/PULL"),
			Expected: "/PULL",
		},



		{
			FingerRequest: finger.CreateRequestTarget("joeblow"),
			Expected: "joeblow",
		},
		{
			FingerRequest: finger.CreateRequestTarget("dariush"),
			Expected: "dariush",
		},



		{
			FingerRequest: finger.CreateRequest("/W", "joeblow"),
			Expected: "/W joeblow",
		},
		{
			FingerRequest: finger.CreateRequest("/W", "dariush"),
			Expected: "/W dariush",
		},



		{
			FingerRequest: finger.CreateRequest("/PULL", "joeblow"),
			Expected: "/PULL joeblow",
		},
		{
			FingerRequest: finger.CreateRequest("/PULL", "dariush"),
			Expected: "/PULL dariush",
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual string = test.FingerRequest.Sentence()

		if expected != actual {
			t.Errorf("For test #%d, the actual finger-protocol request (string) is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
