package ipisp

import (
	"strconv"
	"strings"
)

//ASN contains an Autonomous Systems Number
type ASN int

//ParseASN parses a string like AS2341 into an ASN
func ParseASN(asn string) ASN {
	//Make case insensitive
	asn = strings.ToUpper(asn)
	if len(asn) > 2 {
		asn = strings.TrimPrefix(asn, "AS")
		nn, _ := strconv.Atoi(asn)
		return ASN(nn)
	}
	return 0
}

//String implements fmt.Stringer
func (a ASN) String() string {
	return "AS" + strconv.Itoa(int(a))
}
