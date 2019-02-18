package ipisp

import (
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testClient(c Client, t *testing.T) {
	t.Run("ASN", func(t *testing.T) {
		asn := ASN(15169)
		resp, err := c.LookupASN(asn)
		require.Nil(t, err)
		assert.Equal(t, ASN(15169), resp.ASN)
		assert.Contains(t, resp.Name.Long, "Google")
		assert.Equal(t, "US", resp.Country)
		assert.Equal(t, "ARIN", resp.Registry)
		assert.Nil(t, resp.Range)
		expCreated, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2000-03-30 00:00:00 +0000 UTC")
		require.NoError(t, err)
		assert.Equal(t, expCreated, resp.AllocatedAt)
	})

	t.Run("IP", func(t *testing.T) {
		t.Run("Normal IP", func(t *testing.T) {
			ip := net.ParseIP("8.8.8.8")
			resp, err := c.LookupIP(ip)
			require.NoError(t, err)
			assert.Equal(t, ASN(15169), resp.ASN)
			assert.Contains(t, resp.Name.Long, "Google")
			assert.Equal(t, "US", resp.Country)
			assert.Equal(t, "ARIN", resp.Registry)
			_, expRange, err := net.ParseCIDR("8.8.8.0/24")
			require.NoError(t, err)
			assert.Equal(t, expRange, resp.Range)
		})

		t.Run("Raw IP", func(t *testing.T) {
			ip := net.IP{8, 8, 8, 8}
			resp, err := c.LookupIP(ip)
			require.NoError(t, err)

			assert.Equal(t, ASN(15169), resp.ASN)
			assert.Contains(t, resp.Name.Long, "Google")
			assert.Equal(t, "US", resp.Country)
			assert.Equal(t, "ARIN", resp.Registry)
			_, expRange, err := net.ParseCIDR("8.8.8.0/24")
			require.NoError(t, err)
			assert.Equal(t, expRange, resp.Range)
		})

		t.Run("Multiple ASNs", func(t *testing.T) {
			t.Skip("This test is too flakey at the moment because the IP changes owners.")

			// TODO: find an IP address which changes hands less frequently.
			ip := net.ParseIP("103.235.224.237") // See #6
			resp, err := c.LookupIP(ip)
			require.NoError(t, err)
			assert.Equal(t, ASN(23724), resp.ASN)
		})
	})
}

func Test_parseASNs(t *testing.T) {
	type args struct {
		asnList string
	}
	tests := []struct {
		name    string
		args    args
		want    []ASN
		wantErr bool
	}{
		{"single", args{"1010"}, []ASN{1010}, false},
		{"double", args{"1010 1000"}, []ASN{1010, 1000}, false},
		{"none", args{""}, nil, true},
		{"bad", args{"hello"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseASNs(tt.args.asnList)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseASNs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseASNs() = %v, want %v", got, tt.want)
			}
		})
	}
}
