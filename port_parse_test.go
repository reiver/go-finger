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
			Expected: finger.NoPort(),
		},



		{
			PortString: "0",
			Expected: finger.SomePort(0),
		},
		{
			PortString: "1",
			Expected: finger.SomePort(1),
		},
		{
			PortString: "2",
			Expected: finger.SomePort(2),
		},
		{
			PortString: "3",
			Expected: finger.SomePort(3),
		},
		{
			PortString: "4",
			Expected: finger.SomePort(4),
		},
		{
			PortString: "5",
			Expected: finger.SomePort(5),
		},
		{
			PortString: "6",
			Expected: finger.SomePort(6),
		},
		{
			PortString: "7",
			Expected: finger.SomePort(7),
		},
		{
			PortString: "8",
			Expected: finger.SomePort(8),
		},
		{
			PortString: "9",
			Expected: finger.SomePort(9),
		},
		{
			PortString: "10",
			Expected: finger.SomePort(10),
		},
		{
			PortString: "11",
			Expected: finger.SomePort(11),
		},
		{
			PortString: "12",
			Expected: finger.SomePort(12),
		},
		{
			PortString: "13",
			Expected: finger.SomePort(13),
		},



		{
			PortString: "74",
			Expected: finger.SomePort(74),
		},
		{
			PortString: "75",
			Expected: finger.SomePort(75),
		},
		{
			PortString: "76",
			Expected: finger.SomePort(76),
		},
		{
			PortString: "77",
			Expected: finger.SomePort(77),
		},
		{
			PortString: "78",
			Expected: finger.SomePort(78),
		},
		{
			PortString: "79",
			Expected: finger.SomePort(79),
		},



		{
			PortString: "1971",
			Expected: finger.SomePort(1971),
		},



		{
			PortString: "123",
			Expected: finger.SomePort(123),
		},
		{
			PortString: "1234",
			Expected: finger.SomePort(1234),
		},
		{
			PortString: "12345",
			Expected: finger.SomePort(12345),
		},



		{
			PortString: "65535",
			Expected: finger.SomePort(65535),
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
