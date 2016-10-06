package subscription

import (
	"testing"
)

func TestIsUniversityMail(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"abc12def@studserv.uni-leipzig.de", true},
		{"ab12defg@studserv.uni-leipzig.de", true},
	}

	for _, c := range cases {
		got := isUniversityMail(c.input)
		if got != c.want {
			t.Errorf("isUniversityMail(%q) ==\n%q (is)\n%q (should)", c.input, got, c.want)
		}
	}
}
