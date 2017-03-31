package ipisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhoisClient(t *testing.T) {
	t.Parallel()
	client, err := NewWhoisClient()
	assert.Nil(t, err)
	testClient(client, t)
}
