package magicfinger_test

import (
	"github.com/reiver/go-finger/magic"

	"testing"
)

func TestQuoteSentence(t *testing.T) {

	tests := []struct{
		Sentence string
		Expected string
	}{
		{
			Sentence:"",
			Expected:"{}",
		},



		{
			Sentence: "ONCE",
			Expected:"{ONCE}",
		},
		{
			Sentence: "ONCE TWICE",
			Expected:"{ONCE TWICE}",
		},
		{
			Sentence: "ONCE TWICE THRICE",
			Expected:"{ONCE TWICE THRICE}",
		},
		{
			Sentence: "ONCE TWICE THRICE FOURCE",
			Expected:"{ONCE TWICE THRICE FOURCE}",
		},



		{
			Sentence:`\`,
			Expected:`{\\}`,
		},
		{
			Sentence:`\\`,
			Expected:`{\\\\}`,
		},
		{
			Sentence:`\\\`,
			Expected:`{\\\\\\}`,
		},
		{
			Sentence:`\\\\`,
			Expected:`{\\\\\\\\}`,
		},
		{
			Sentence:`\\\\\`,
			Expected:`{\\\\\\\\\\}`,
		},



		{
			Sentence:`{`,
			Expected:`{\{}`,
		},
		{
			Sentence:`{{`,
			Expected:`{\{\{}`,
		},
		{
			Sentence:`{{{`,
			Expected:`{\{\{\{}`,
		},
		{
			Sentence:`{{{{`,
			Expected:`{\{\{\{\{}`,
		},



		{
			Sentence:`}`,
			Expected:`{\}}`,
		},
		{
			Sentence:`}}`,
			Expected:`{\}\}}`,
		},
		{
			Sentence:`}}}`,
			Expected:`{\}\}\}}`,
		},
		{
			Sentence:`}}}}`,
			Expected:`{\}\}\}\}}`,
		},



		{
			Sentence: `/GET joeblow/something.txt`,
			Expected:`{/GET joeblow/something.txt}`,
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual   string = magicfinger.QuoteSentence(test.Sentence)

		if expected != actual {
			t.Errorf("For test #%d, the actual value for the quoted-sentence is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}
