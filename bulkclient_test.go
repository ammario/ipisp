package ipisp

import (
	"context"
	"net"
	"strings"
	"testing"
)

func TestWhoisClient(t *testing.T) {
	t.Parallel()
	client, err := DialBulkClient(context.Background())
	if err != nil {
		t.Fatalf("dial bulk client: %v", err)
	}
	defer client.Close()

	t.Run("IPs", func(t *testing.T) {
		// These are well-known IPs.
		resp, err := client.LookupIPs(
			net.ParseIP("1.1.1.1"),
			net.ParseIP("4.2.2.2"),
		)
		if err != nil {
			t.Fatalf("lookup IPs: %v", err)
		}
		if resp[0].ASN != ASN(13335) {
			t.Errorf("unexpected ASN %v", resp[0].ASN)
		}
		if !strings.Contains(resp[0].ISPName, "CLOUDFLARE") {
			t.Errorf("unexpected owner %v", resp[0].ISPName)
		}
		if !strings.Contains(resp[1].ISPName, "LEVEL3") {
			t.Errorf("unexpected owner %v", resp[1].ISPName)
		}
		t.Logf("%+v", resp)
	})

	t.Run("ASNs", func(t *testing.T) {
		// These are well-known IPs.
		resp, err := client.LookupASNs(
			13335,
			15169,
		)
		if err != nil {
			t.Fatalf("lookup ASNs: %v", err)
		}
		if resp[0].ASN != ASN(13335) {
			t.Errorf("unexpected ASN %v", resp[0].ASN)
		}
		if !strings.Contains(resp[0].ISPName, "CLOUDFLARE") {
			t.Errorf("unexpected owner %v", resp[0].ISPName)
		}
		if !strings.Contains(resp[1].ISPName, "GOOGLE") {
			t.Errorf("unexpected owner %v", resp[1].ISPName)
		}
		t.Logf("%+v", resp)
	})

}
