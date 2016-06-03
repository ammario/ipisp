package ipisp

import (
	"net"
	"time"
)

//Response contains a response from Cymru
type Response struct {
	IP        net.IP
	ASN       ASN
	Country   *Country
	Registry  string
	Range     *net.IPNet
	Allocated time.Time
	Name      *Name
}
