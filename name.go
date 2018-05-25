package ipisp

import (
	"regexp"
	"strings"
)

// Name contains an IPISP ISP name
type Name struct {
	Raw   string
	Short string
	Long  string
}

var stripASRegex = regexp.MustCompile(`(-AS$)`)

// stringAS strips the AS portion of a name.
func stripAS(raw string) string {
	return stripASRegex.ReplaceAllString(raw, "")
}

// ParseName returns a pointer to a new name
func ParseName(raw string) Name {
	n := Name{
		Raw:   raw,
		Short: raw,
		Long:  raw,
	}

	tokens := strings.Split(raw, " - ")
	switch {
	case len(tokens) == 0:
		// Just use raw name.
	case len(tokens) == 1:
		n.Short = tokens[0]
	case len(tokens) == 2:
		n.Short = tokens[0]
		n.Long = tokens[1]
	}

	n.Short = stripAS(n.Short)
	n.Long = stripAS(n.Long)
	return n
}

// String returns a human readable representation of the name.
func (n Name) String() string {
	return n.Long
}
