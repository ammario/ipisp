package ipisp

import (
	"net"
	"strings"

	"github.com/pkg/errors"
)

//Client is a lookup client
type Client interface {
	LookupIPs([]net.IP) ([]Response, error)
	LookupIP(net.IP) (*Response, error)
	LookupASNs([]ASN) ([]Response, error)
	LookupASN(ASN) (*Response, error)
	Close() error
}

// parseASNs parses an ASN list like "1024 1111 11202".
// If it doesn't return an error, the returned slice has at least one entry.
func parseASNs(asnList string) ([]ASN, error) {
	tokens := strings.Split(strings.TrimSpace(asnList), " ")
	if len(tokens) == 0 {
		return nil, errors.New("no ASNs")
	}

	asns := make([]ASN, len(tokens))

	for i, tok := range tokens {
		asn, err := ParseASN(tok)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse asn")
		}
		asns[i] = ASN(asn)
	}

	return asns, nil
}
