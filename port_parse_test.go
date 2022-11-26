package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestParsePort(t *testing.T) {

	tests := []struct{
		PortString string
		Expected finger.Port
	}{
		{
			PortString: "",
			Expected: finger.EmptyPort(),
		},



		{
			PortString: "0",
			Expected: finger.CreatePort(0),
		},
		{
			PortString: "1",
			Expected: finger.CreatePort(1),
		},
		{
			PortString: "2",
			Expected: finger.CreatePort(2),
		},
		{
			PortString: "3",
			Expected: finger.CreatePort(3),
		},
		{
			PortString: "4",
			Expected: finger.CreatePort(4),
		},
		{
			PortString: "5",
			Expected: finger.CreatePort(5),
		},
		{
			PortString: "6",
			Expected: finger.CreatePort(6),
		},
		{
			PortString: "7",
			Expected: finger.CreatePort(7),
		},
		{
			PortString: "8",
			Expected: finger.CreatePort(8),
		},
		{
			PortString: "9",
			Expected: finger.CreatePort(9),
		},
		{
			PortString: "10",
			Expected: finger.CreatePort(10),
		},
		{
			PortString: "11",
			Expected: finger.CreatePort(11),
		},
		{
			PortString: "12",
			Expected: finger.CreatePort(12),
		},
		{
			PortString: "13",
			Expected: finger.CreatePort(13),
		},



		{
			PortString: "74",
			Expected: finger.CreatePort(74),
		},
		{
			PortString: "75",
			Expected: finger.CreatePort(75),
		},
		{
			PortString: "76",
			Expected: finger.CreatePort(76),
		},
		{
			PortString: "77",
			Expected: finger.CreatePort(77),
		},
		{
			PortString: "78",
			Expected: finger.CreatePort(78),
		},
		{
			PortString: "79",
			Expected: finger.CreatePort(79),
		},



		{
			PortString: "1971",
			Expected: finger.CreatePort(1971),
		},



		{
			PortString: "123",
			Expected: finger.CreatePort(123),
		},
		{
			PortString: "1234",
			Expected: finger.CreatePort(1234),
		},
		{
			PortString: "12345",
			Expected: finger.CreatePort(12345),
		},



		{
			PortString: "65535",
			Expected: finger.CreatePort(65535),
		},
	}

	for testNumber, test := range tests {

		var expected finger.Port = test.Expected

		actual, err := finger.ParsePort(test.PortString)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		if expected != actual {
			t.Errorf("For test #%d, the actual value for the finger-protocol port is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
