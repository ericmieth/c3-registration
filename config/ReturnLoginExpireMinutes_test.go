package config

import (
	"testing"
	"time"
)

func TestReturnLoginExpireMinutes(t *testing.T) {
	cases := []struct {
		want time.Duration
	}{
		{30},
	}
	for _, c := range cases {
		got := ReturnLoginExpireMinutes()
		if got != c.want {
			t.Errorf("ReturnLoginExpireTime() ==\n%q (is)\n%q (should)", got, c.want)
		}

	}
}
