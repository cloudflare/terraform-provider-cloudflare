package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateEphemeralCertAndKey(t *testing.T) {
	cert, key, err := GenerateEphemeralCertAndKey([]string{"example.com"}, time.Now().Add(time.Hour*1))

	assert.NotEqual(t, "", cert)
	assert.NotEqual(t, "", key)
	assert.Equal(t, nil, err)
	assert.Contains(t, cert, "BEGIN CERTIFICATE", "failed to find certificate delimiter in output")
	assert.Contains(t, key, "BEGIN EC PRIVATE KEY", "failed to find private key delimiter in output")
}
