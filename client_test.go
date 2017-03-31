package ipisp

import (
	"testing"
	"time"

	"net"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testClient(c Client, t *testing.T) {
	t.Run("ASN", func(t *testing.T) {
		asn := ASN(15169)
		resp, err := c.LookupASN(asn)
		require.Nil(t, err)
		assert.Equal(t, ASN(15169), resp.ASN)
		assert.Equal(t, "Google Inc., US", resp.Name.Long)
		assert.Equal(t, "US", resp.Country)
		assert.Equal(t, "ARIN", resp.Registry)
		assert.Nil(t, resp.Range)
		expCreated, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2000-03-30 00:00:00 +0000 UTC")
		assert.Nil(t, err)
		assert.Equal(t, expCreated, resp.AllocatedAt)
	})

	t.Run("IP", func(t *testing.T) {
		ip := net.ParseIP("8.8.8.8")
		resp, err := c.LookupIP(ip)
		require.Nil(t, err)
		assert.Equal(t, ASN(15169), resp.ASN)
		assert.Equal(t, "Google Inc., US", resp.Name.Long)
		assert.Equal(t, "US", resp.Country)
		assert.Equal(t, "ARIN", resp.Registry)
		_, expRange, err := net.ParseCIDR("8.8.8.0/24")
		assert.Nil(t, err)
		assert.Equal(t, expRange, resp.Range)
		//fmt.Printf("resp: %+v", resp)
	})
}
