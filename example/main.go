package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ammario/ipisp/v2"
)

func main() {
	resp, err := ipisp.LookupIP(context.Background(), net.ParseIP("4.2.2.2"))
	if err != nil {
		log.Fatalf("lookup 4.2.2.2: %v", err)
	}
	fmt.Printf("resolved IP 4.2.2.2: %+v\n", resp)

	resp, err = ipisp.LookupASN(ipisp.ASN(666))
	if err != nil {
		log.Fatalf("lookup ASN 666: %v", err)
	}
	fmt.Printf("resolved ASN 666: %+v\n", resp)
}
