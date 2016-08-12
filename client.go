package ipisp

import (
	"net"
)

type Client interface {
	LookupIPs([]net.IP) ([]Response, error)
	LookupIP(net.IP) (*Response, error)
	LookupASNs([]ASN) ([]Response, error)
	LookupASN(ASN) (*Response, error)
	Close() error
}
