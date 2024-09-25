package token

import(
	"testing"
)

func TestLookupIdent(t *testing.T) {
	tt := []struct {
		input string
		want TokenType
	}{
		{"func", FUNCTION},
		{"myVar", IDENTIFIER},
	}

	for _, tc := range tt {
		got := LookupIdent(tc.input)
		 if tc.want != got {
			t.Fatalf("For input %s, Test fail - got= %s, want= %s", tc.input, got, tc.want)
		 }
	}
}