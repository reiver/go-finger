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
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
					},
				},
			},
			Expected: "dariush@example.com",
		},
		{
			Query: finger.Query{
				UserName: finger.SomeUserName("dariush"),
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
						Port: finger.SomePort(1971),
					},
				},
			},
			Expected: "dariush@example.com:1971",
		},



		{
			Query: finger.Query{
				UserName: finger.SomeUserName("dariush"),
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
					},
					finger.Address{
						Host: finger.SomeHost("something.social"),
					},
				},
			},
			Expected: "dariush@example.com@something.social",
		},
		{
			Query: finger.Query{
				UserName: finger.SomeUserName("dariush"),
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
						Port: finger.SomePort(1971),
					},
					finger.Address{
						Host: finger.SomeHost("something.social"),
					},
				},
			},
			Expected: "dariush@example.com:1971@something.social",
		},
		{
			Query: finger.Query{
				UserName: finger.SomeUserName("dariush"),
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
					},
					finger.Address{
						Host: finger.SomeHost("something.social"),
						Port: finger.SomePort(1234),
					},
				},
			},
			Expected: "dariush@example.com@something.social:1234",
		},
		{
			Query: finger.Query{
				UserName: finger.SomeUserName("dariush"),
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
						Port: finger.SomePort(1971),
					},
					finger.Address{
						Host: finger.SomeHost("something.social"),
						Port: finger.SomePort(1234),
					},
				},
			},
			Expected: "dariush@example.com:1971@something.social:1234",
		},



		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
					},
				},
			},
			Expected: "@example.com",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
						Port: finger.SomePort(1971),
					},
				},
			},
			Expected: "@example.com:1971",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
					},
					finger.Address{
						Host: finger.SomeHost("something.social"),
					},
				},
			},
			Expected: "@example.com@something.social",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
						Port: finger.SomePort(1971),
					},
					finger.Address{
						Host: finger.SomeHost("something.social"),
					},
				},
			},
			Expected: "@example.com:1971@something.social",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
					},
					finger.Address{
						Host: finger.SomeHost("something.social"),
						Port: finger.SomePort(1234),
					},
				},
			},
			Expected: "@example.com@something.social:1234",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("example.com"),
						Port: finger.SomePort(1971),
					},
					finger.Address{
						Host: finger.SomeHost("something.social"),
						Port: finger.SomePort(1234),
					},
				},
			},
			Expected: "@example.com:1971@something.social:1234",
		},



		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("once"),
					},
				},
			},
			Expected: "@once",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("once"),
					},
					finger.Address{
						Host: finger.SomeHost("twice"),
					},
				},
			},
			Expected: "@once@twice",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("once"),
					},
					finger.Address{
						Host: finger.SomeHost("twice"),
					},
					finger.Address{
						Host: finger.SomeHost("thrice"),
					},
				},
			},
			Expected: "@once@twice@thrice",
		},
		{
			Query: finger.Query{
				Addresses: []finger.Address{
					finger.Address{
						Host: finger.SomeHost("once"),
					},
					finger.Address{
						Host: finger.SomeHost("twice"),
					},
					finger.Address{
						Host: finger.SomeHost("thrice"),
					},
					finger.Address{
						Host: finger.SomeHost("fource"),
					},
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
