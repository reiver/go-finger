package finger_test

import (
	"github.com/reiver/go-finger"

	"testing"
)

func TestQuery_String(t *testing.T) {

	tests := []struct{
		Query finger.Query
		Expected string
	}{
		{
			Query: finger.Query{},
			Expected: "",
		},



		{
			Query: finger.Query{
				UserName: finger.SomeUserName("dariush"),
			},
			Expected: "dariush",
		},
		{
			Query: finger.Query{
				UserName: finger.SomeUserName("dariush"),
				Hosts: []finger.Host{
					finger.SomeHost("example.com"),
				},
			},
			Expected: "dariush@example.com",
		},
		{
			Query: finger.Query{
				UserName: finger.SomeUserName("dariush"),
				Hosts: []finger.Host{
					finger.SomeHost("example.com"),
					finger.SomeHost("something.social"),
				},
			},
			Expected: "dariush@example.com@something.social",
		},



		{
			Query: finger.Query{
				Hosts: []finger.Host{
					finger.SomeHost("example.com"),
				},
			},
			Expected: "@example.com",
		},
		{
			Query: finger.Query{
				Hosts: []finger.Host{
					finger.SomeHost("example.com"),
					finger.SomeHost("something.social"),
				},
			},
			Expected: "@example.com@something.social",
		},



		{
			Query: finger.Query{
				Hosts: []finger.Host{
					finger.SomeHost("once"),
				},
			},
			Expected: "@once",
		},
		{
			Query: finger.Query{
				Hosts: []finger.Host{
					finger.SomeHost("once"),
					finger.SomeHost("twice"),
				},
			},
			Expected: "@once@twice",
		},
		{
			Query: finger.Query{
				Hosts: []finger.Host{
					finger.SomeHost("once"),
					finger.SomeHost("twice"),
					finger.SomeHost("thrice"),
				},
			},
			Expected: "@once@twice@thrice",
		},
		{
			Query: finger.Query{
				Hosts: []finger.Host{
					finger.SomeHost("once"),
					finger.SomeHost("twice"),
					finger.SomeHost("thrice"),
					finger.SomeHost("fource"),
				},
			},
			Expected: "@once@twice@thrice@fource",
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual   string = test.Query.String()

		if expected != actual {
			t.Errorf("For test #%d, the actual query string value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}
