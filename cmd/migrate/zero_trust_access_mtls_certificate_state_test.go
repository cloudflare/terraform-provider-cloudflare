package main

import (
	"testing"
)

func TestZeroTrustAccessMTLSCertificateStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "basic certificate state transformation from v4",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_access_mutual_tls_certificate",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "test-account-id",
									"name": "test-cert",
									"certificate": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
									"associated_hostnames": ["example.com", "test.example.com"],
									"id": "cert-id-123",
									"fingerprint": "abc123def456"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zero_trust_access_mtls_certificate",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "test-account-id",
									"name": "test-cert",
									"certificate": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
									"associated_hostnames": ["example.com", "test.example.com"],
									"id": "cert-id-123",
									"fingerprint": "abc123def456"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "certificate with zone_id state transformation",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_access_mutual_tls_certificate",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"name": "test-cert",
									"certificate": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
									"associated_hostnames": [],
									"id": "cert-id-456",
									"fingerprint": "def789ghi012"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zero_trust_access_mtls_certificate",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"name": "test-cert",
									"certificate": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
									"associated_hostnames": [],
									"id": "cert-id-456",
									"fingerprint": "def789ghi012"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "v5 certificate state remains unchanged",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zero_trust_access_mtls_certificate",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "test-account-id",
									"name": "test-cert",
									"certificate": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
									"associated_hostnames": ["example.com"],
									"id": "cert-id-789",
									"fingerprint": "xyz456abc789",
									"expires_on": "2026-08-21T12:34:53Z"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zero_trust_access_mtls_certificate",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "test-account-id",
									"name": "test-cert",
									"certificate": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
									"associated_hostnames": ["example.com"],
									"id": "cert-id-789",
									"fingerprint": "xyz456abc789",
									"expires_on": "2026-08-21T12:34:53Z"
								}
							}
						]
					}
				]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}