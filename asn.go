package ipisp

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

//ASN contains an Autonomous Systems Number
type ASN int

//ParseASN parses a string like AS2341 into an ASN
func ParseASN(asn string) (ASN, error) {
	//Make case insensitive
	asn = strings.ToUpper(asn)
	if len(asn) > 2 && asn[:2] == "AS" {
		asn = asn[2:]
	}

	nn, err := strconv.Atoi(asn)
	return ASN(nn), errors.Wrap(err, "failed to conv into to string")
}

//String implements fmt.Stringer
func (a ASN) String() string {
	return "AS" + strconv.Itoa(int(a))
}
