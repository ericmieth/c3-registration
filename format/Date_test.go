package format

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	cases := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), "17. November 2009, 20:34"},
	}

	for _, c := range cases {
		got := Date(c.input)
		if got != c.want {
			t.Errorf("Date(%q) ==\n%q (is)\n%q (should)", c.input, got, c.want)
		}
	}
}
