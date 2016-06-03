package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ammario/ipisp"
)

func main() {
	client, err := ipisp.NewClient()

	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	fmt.Println("Successfully made client. Looking up ip 4.2.2.2 and AS666")

	resp, err := client.LookupIP(net.ParseIP("4.2.2.2"))

	if err != nil {
		log.Fatalf("Error looking up 4.2.2.2: %v", err)
	}

	fmt.Printf("Resolved IP 4.2.2.2\n")
	fmt.Printf("IP: %v\n", resp.IP)
	fmt.Printf("ASN: %v\n", resp.ASN)
	fmt.Printf("Range: %v\n", resp.Range)
	fmt.Printf("Country: %v\n", resp.Country.Name)
	fmt.Printf("Registry: %v\n", resp.Registry)
	fmt.Printf("ISP: %v\n", resp.Name.Raw)

	fmt.Print("\n\nResolved AS666\n")
	resp, err = client.LookupASN(ipisp.ASN(666))

	fmt.Printf("ASN: %v\n", resp.ASN)
	fmt.Printf("Country: %v\n", resp.Country.Name)
	fmt.Printf("Registry: %v\n", resp.Registry)
	fmt.Printf("ISP: %v\n", resp.Name.Raw)

	client.Close()
}
