package ipisp

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	// ErrUnassigned is returned if an IP address is not assigned.
	// See https://en.wikipedia.org/wiki/Reserved_IP_addresses for some example ranges.
	ErrUnassigned = errors.New("address is unassigned")
)




// LookupIP performs a lookup using the DNS client.
// Please use the BulkClient if you're looking up more than
// one address at a time. The service may ban you, otherwise.
func LookupIP(ip string) (*Response, error) {

}

// LookupASN performs a lookup using the DNS client.
// Please use the BulkClient if you're looking up more than
// one address at a time. The service may ban you, otherwise.
func LookupASN(asn string) (*Response, error) {

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
