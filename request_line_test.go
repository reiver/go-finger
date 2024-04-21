package finger_test

import (
	"github.com/reiver/go-finger"

	"reflect"
	"io"
	"strings"

	"testing"
)

func TestReadRequestLine_one(t *testing.T) {

	var request1 string = "/W joeblow\r\n"

	var inputRequestLine string = request1

	var reader io.Reader = strings.NewReader(inputRequestLine)

	actualRequestLine, err := finger.ReadRequestLine(reader)
	if nil != err {
		t.Errorf("Did not expect an error but actually got one.")
		t.Logf("ERROR: (%T) %q", err, err)
		return
	}

	{
		var expected string = request1[:len(request1)-2]
		var actual   string = actualRequestLine

		if expected != actual {
			t.Errorf("The actual request-line is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}
}

func TestReadRequestLine_two(t *testing.T) {

	var request1 string = "/W joeblow\r\n"
	var request2 string = "darius@reiver.link\r\n"

	var inputRequestLine string = request1 + request2

	var reader io.Reader = strings.NewReader(inputRequestLine)

	{
		actualRequestLine, err := finger.ReadRequestLine(reader)
		if nil != err {
			t.Errorf("Did not expect an error but actually got one.")
			t.Logf("ERROR: (%T) %q", err, err)
			return
		}

		{
			var expected string = request1[:len(request1)-2]
			var actual   string = actualRequestLine

			if expected != actual {
				t.Errorf("The actual request-line is not what was expected.")
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				return
			}
		}
	}

	{
		actualRequestLine, err := finger.ReadRequestLine(reader)
		if nil != err {
			t.Errorf("Did not expect an error but actually got one.")
			t.Logf("ERROR: (%T) %q", err, err)
			return
		}

		{
			var expected string = request2[:len(request2)-2]
			var actual   string = actualRequestLine

			if expected != actual {
				t.Errorf("The actual request-line is not what was expected.")
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				return
			}
		}
	}
}

func TestReadRequestLine(t *testing.T) {

	tests := []struct{
		Input string
		Expected []string
	}{
		{
			Input:
				"\r\n",
			Expected: []string{
				"",
			},
		},
		{
			Input:
				"v\r\n",
			Expected: []string{
				"v",
			},
		},
		{
			Input:
				"joeblow\r\n",
			Expected: []string{
				"joeblow",
			},
		},
		{
			Input:
				"joeblow@example.com\r\n",
			Expected: []string{
				"joeblow@example.com",
			},
		},
		{
			Input:
				"joeblow@example.com@reiver.link\r\n",
			Expected: []string{
				"joeblow@example.com@reiver.link",
			},
		},
		{
			Input:
				"/W\r\n",
			Expected: []string{
				"/W",
			},
		},
		{
			Input:
				"/W joeblow\r\n",
			Expected: []string{
				"/W joeblow",
			},
		},
		{
			Input:
				"/W joeblow@example.com\r\n",
			Expected: []string{
				"/W joeblow@example.com",
			},
		},
		{
			Input:
				"/W joeblow@example.com@reiver.link\r\n",
			Expected: []string{
				"/W joeblow@example.com@reiver.link",
			},
		},
		{
			Input:
				"/PULL\r\n",
			Expected: []string{
				"/PULL",
			},
		},



		{
			Input:
				"/GET  /path/to/file.txt HTTP/1.1\r\n",
			Expected: []string{
				"/GET  /path/to/file.txt HTTP/1.1",
			},
		},



		{
			Input:
				"\r\n"+
				"once\r\n"+
				"twice\r\n"+
				"thrice\r\n"+
				"fource\r\n"+
				"\r\n"+
				"/W once\r\n"+
				"/W twice\r\n"+
				"/W thrice\r\n"+
				"/W fource\r\n"+
				"\r\n"+
				"once@example.com\r\n"+
				"twice@example.com\r\n"+
				"thrice@example.com\r\n"+
				"fource@example.com\r\n"+
				"\r\n"+
				"/W once@example.com\r\n"+
				"/W twice@example.com\r\n"+
				"/W thrice@example.com\r\n"+
				"/W fource@example.com\r\n",
			Expected: []string{
				"",
				"once",
				"twice",
				"thrice",
				"fource",
				"",
				"/W once",
				"/W twice",
				"/W thrice",
				"/W fource",
				"",
				"once@example.com",
				"twice@example.com",
				"thrice@example.com",
				"fource@example.com",
				"",
				"/W once@example.com",
				"/W twice@example.com",
				"/W thrice@example.com",
				"/W fource@example.com",
			},
		},
	}

	testloop: for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Input)

		var actual []string

		for {
			requestline, err := finger.ReadRequestLine(reader)
			if io.EOF == err {
				break
			}
			if nil != err && io.EOF != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %q", err, err)
				t.Logf("EXPECTED-LINES: %#v", test.Expected)
				t.Logf("ACTUAL-LINES:   %#v", actual)
				t.Logf("INPUT: %q", test.Input)
	/////////////////////// CONTINUE
				continue testloop
			}

			actual = append(actual, requestline)
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual read request-lines is not what was expected.", testNumber)
				t.Logf("EXPECTED-LINES: %#v", expected)
				t.Logf("ACTUAL-LINES:   %#v", actual)
				t.Logf("INPUT: %q", test.Input)
				continue
			}
		}
	}
}
