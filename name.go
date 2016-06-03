package ipisp

import "strings"

//Name contains an IPISP name
type Name struct {
	Raw string
}

//NewName returns a pointer to a new name
func NewName(raw string) *Name {
	return &Name{
		Raw: strings.TrimSpace(raw),
	}
}

//Short attempts to provide the short name of an ISP
func (n *Name) Short() string {
	nameTokens := strings.Split(n.Raw, "-")
	if len(nameTokens) < 2 {
		return n.Raw
	}
	return strings.TrimSpace(nameTokens[0])
}

//Long attempts to provide the log name of an ISP
func (n *Name) Long() string {
	nameTokens := strings.Split(n.Raw, "-")
	if len(nameTokens) < 2 {
		return n.Raw
	}
	return strings.TrimSpace(nameTokens[1])
}
