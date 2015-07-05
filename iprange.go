package iprange

import (
	"net"
	"strings"
)

type Range struct {
	allows []*net.IPNet
}

func New(ipStr string) (*Range, error) {
	IPs := strings.Split(ipStr, ",")
	r := &Range{
		allows: make([]*net.IPNet, 0, len(IPs)),
	}

	for _, i := range IPs {
		if !strings.Contains(i, "/") {
			if strings.Contains(i, ".") { // IPv4
				i += "/32"
			} else { // IPv6
				i += "/128"
			}
		}

		_, mask, err := net.ParseCIDR(i)
		if err != nil {
			return nil, err
		}
		r.allows = append(r.allows, mask)
	}

	return r, nil
}

func (r *Range) IncludeStr(addr string) bool {
	return r.Include(net.ParseIP(addr))
}

func (r *Range) Include(addr net.IP) bool {
	for _, m := range r.allows {
		if m.Contains(addr) {
			return true
		}
	}

	return false
}

func (r *Range) InlucdeConn(conn net.Conn) bool {
	addr, _ := conn.RemoteAddr().(*net.TCPAddr)
	return r.Include(addr.IP)
}
