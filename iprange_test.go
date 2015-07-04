package iprange_test

import (
	"testing"

	"github.com/pocke/go-iprange"
)

func TestIncludeStr(t *testing.T) {
	assert := func(ra, ip string, res bool) {
		r, err := iprange.New(ra)
		if err != nil {
			t.Error(err)
		}
		if r.IncludeStr(ip) != res {
			t.Errorf("Expected %q, but got %q", res, !res)
		}
	}
	assert("192.168.0.1", "192.168.0.1", true)
	assert("192.168.0.1", "192.168.0.2", false)

	assert("192.168.0.0/24", "192.168.0.1", true)
	assert("192.168.0.0/24", "192.168.0.2", true)
	assert("192.168.0.0/24", "192.168.1.2", false)

	assert("192.168.0.0/24,172.0.0.0/16,192.168.1.1", "192.168.0.1", true)
	assert("192.168.0.0/24,172.0.0.0/16,192.168.1.1", "192.168.1.1", true)
	assert("192.168.0.0/24,172.0.0.0/16,192.168.1.1", "172.0.10.11", true)

	assert("2001:0db8:bd05:01d2:288a:1fc0:0001:10ee", "2001:0db8:bd05:01d2:288a:1fc0:0001:10ee", true)
	assert("2001:0db8:bd05:01d2:288a:1fc0:0001:10ee", "192.168.0.1", false)
}
