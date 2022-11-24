package finger_test

import (
	"github.com/reiver/go-finger"

	"strings"
	"testing"
)

func TestRequest_WriteTo(t *testing.T) {

	tests := []struct{
		FingerRequest finger.Request
		Expected string
	}{
		{
			FingerRequest: finger.Request{},
			Expected: "\r\n",
		},



		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("W"),
			},
			Expected: "/W\r\n",
		},
		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("PULL"),
			},
			Expected: "/PULL\r\n",
		},



		{
			FingerRequest: finger.Request{
				Target: finger.SomeTarget("joeblow"),
			},
			Expected: "joeblow\r\n",
		},
		{
			FingerRequest: finger.Request{
				Target: finger.SomeTarget("dariush"),
			},
			Expected: "dariush\r\n",
		},



		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("W"),
				Target: finger.SomeTarget("joeblow"),
			},
			Expected: "/W joeblow\r\n",
		},
		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("W"),
				Target: finger.SomeTarget("dariush"),
			},
			Expected: "/W dariush\r\n",
		},



		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("PULL"),
				Target: finger.SomeTarget("joeblow"),
			},
			Expected: "/PULL joeblow\r\n",
		},
		{
			FingerRequest: finger.Request{
				Switch: finger.SomeSwitch("PULL"),
				Target: finger.SomeTarget("dariush"),
			},
			Expected: "/PULL dariush\r\n",
		},
	}

	for testNumber, test := range tests {

		var storage strings.Builder
		{
			n, err := test.FingerRequest.WriteTo(&storage)

			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %q", err, err)
				continue
			}

			if expected, actual := int64(len(test.Expected)), n; expected != actual {
				t.Errorf("For test #%d, the actual number of bytes written is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			var expected string = test.Expected
			var actual   string = storage.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual finger-protocol request (string) is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}
