package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestParseRequest(t *testing.T) {

	tests := []struct{
		RequestString string
		Expected finger.Request
	}{
		{
			RequestString: "",
			Expected: finger.EmptyRequest(),
		},









		{
			RequestString: "joeblow",
			Expected: finger.AssembleRequestTarget(finger.CreateTarget("joeblow")),
		},
		{
			RequestString: "joeblow@example.com",
			Expected: finger.AssembleRequestTarget(finger.CreateTarget("joeblow@example.com")),
		},
		{
			RequestString: "joeblow@example.com@something.dom",
			Expected: finger.AssembleRequestTarget(finger.CreateTarget("joeblow@example.com@something.dom")),
		},









		{
			RequestString: "/W",
			Expected: finger.AssembleRequestSwitch(finger.CreateSwitch("/W")),
		},
		{
			RequestString: "/W joeblow",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow")),
		},
		{
			RequestString: "/W  joeblow",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow")),
		},
		{
			RequestString: "/W   joeblow",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow")),
		},
		{
			RequestString: "/W    joeblow",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow")),
		},
		{
			RequestString: "/W joeblow@example.com",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow@example.com")),
		},
		{
			RequestString: "/W  joeblow@example.com",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow@example.com")),
		},
		{
			RequestString: "/W   joeblow@example.com",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow@example.com")),
		},
		{
			RequestString: "/W    joeblow@example.com",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow@example.com")),
		},
		{
			RequestString: "/W joeblow@example.com@something.dom",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow@example.com@something.dom")),
		},
		{
			RequestString: "/W  joeblow@example.com@something.dom",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow@example.com@something.dom")),
		},
		{
			RequestString: "/W   joeblow@example.com@something.dom",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow@example.com@something.dom")),
		},
		{
			RequestString: "/W    joeblow@example.com@something.dom",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow@example.com@something.dom")),
		},



		{
			RequestString: "\xEF\xBB\xBF"+"/W joeblow@example.com@something.dom",
			Expected: finger.AssembleRequest(finger.CreateSwitch("/W"), finger.CreateTarget("joeblow@example.com@something.dom")),
		},
	}

	for testNumber, test := range tests {

		actualRequest, err := finger.ParseRequest(test.RequestString)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("REQUEST-STRING: %q", test.RequestString)
			continue
		}

		{
			expected := test.Expected
			actual   := actualRequest

			if expected != actual {
				t.Errorf("For test #%d, the actual value for the finger.Requset is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("REQUEST-STRING: %q", test.RequestString)
				continue
			}
		}
	}
}
