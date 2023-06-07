package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateEphemeralCertAndKey(t *testing.T) {
	cert, key, err := GenerateEphemeralCertAndKey([]string{"example.com"})

	assert.NotEqual(t, "", cert)
	assert.NotEqual(t, "", key)
	assert.Equal(t, nil, err)
	assert.Contains(t, cert, "BEGIN CERTIFICATE", "failed to find certificate delimiter in output")
	assert.Contains(t, key, "BEGIN EC PRIVATE KEY", "failed to find private key delimiter in output")
}
