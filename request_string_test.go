package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestRequest_String(t *testing.T) {

	tests := []struct{
		FingerRequest finger.Request
		Expected string
	}{
		{
			FingerRequest: finger.EmptyRequest(),
			Expected: "\r\n",
		},



		{
			FingerRequest: finger.CreateRequestSwitch("/W"),
			Expected: "/W\r\n",
		},
		{
			FingerRequest: finger.CreateRequestSwitch("/PULL"),
			Expected: "/PULL\r\n",
		},



		{
			FingerRequest: finger.CreateRequestTarget("joeblow"),
			Expected: "joeblow\r\n",
		},
		{
			FingerRequest: finger.CreateRequestTarget("dariush"),
			Expected: "dariush\r\n",
		},



		{
			FingerRequest: finger.CreateRequest("/W", "joeblow"),
			Expected: "/W joeblow\r\n",
		},
		{
			FingerRequest: finger.CreateRequest("/W", "dariush"),
			Expected: "/W dariush\r\n",
		},



		{
			FingerRequest: finger.CreateRequest("/PULL", "joeblow"),
			Expected: "/PULL joeblow\r\n",
		},
		{
			FingerRequest: finger.CreateRequest("/PULL", "dariush"),
			Expected: "/PULL dariush\r\n",
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual string = test.FingerRequest.String()

		if expected != actual {
			t.Errorf("For test #%d, the actual finger-protocol request (string) is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
