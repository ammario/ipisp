package ipisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDNSClient(t *testing.T) {
	t.Parallel()
	client, err := NewDNSClient()
	assert.Nil(t, err)
	testClient(client, t)
}
