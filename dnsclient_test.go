package ipisp

import (
	"context"
	"net"
	"strings"
	"testing"
)

func TestLookupIP(t *testing.T) {
	tests := []struct {
		name    string
		ip      net.IP
		wantISP string
		wantErr bool
	}{
		{"google", net.ParseIP("8.8.8.8"), "GOOGLE", false},
		{"fail", net.ParseIP("0.0.1.1"), "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LookupIP(context.Background(), tt.ip)
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("LookupIP() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			t.Logf("error: %v, response: %+v", err, got)
			if !strings.Contains(got.ISPName, tt.wantISP) {
				t.Errorf("LookupIP() got = %v, want %v", got, tt.wantISP)
			}
		})
	}
}
