package iprange_test

import (
	"net"
	"testing"
	"time"

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

	assert("0.0.0.0/0,::0", "127.0.0.1", true)
	assert("0.0.0.0/0,::/0", "::1", true)
}

func TestNewWhenFail(t *testing.T) {
	_, err := iprange.New("invalid as /CIDR")
	if err == nil {
		t.Error("Expacted an error, but got nil")
	}
}

type ConnMock struct {
	addr *net.TCPAddr
}

func (_ *ConnMock) Read([]byte) (int, error)           { return 0, nil }
func (_ *ConnMock) Write([]byte) (int, error)          { return 0, nil }
func (_ *ConnMock) Close() error                       { return nil }
func (_ *ConnMock) LocalAddr() net.Addr                { return nil }
func (c *ConnMock) RemoteAddr() net.Addr               { return c.addr }
func (c *ConnMock) SetDeadline(t time.Time) error      { return nil }
func (c *ConnMock) SetReadDeadline(t time.Time) error  { return nil }
func (c *ConnMock) SetWriteDeadline(t time.Time) error { return nil }

func NewConnMock(ip string) *ConnMock {
	return &ConnMock{
		addr: &net.TCPAddr{IP: net.ParseIP(ip)},
	}
}

var _ net.Conn = &ConnMock{}

func TestIncludeConn(t *testing.T) {
	assert := func(ra, ip string, res bool) {
		r, err := iprange.New(ra)
		if err != nil {
			t.Error(err)
		}
		conn := NewConnMock(ip)
		if r.InlucdeConn(conn) != res {
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

	assert("0.0.0.0/0,::0", "127.0.0.1", true)
	assert("0.0.0.0/0,::/0", "::1", true)
}
