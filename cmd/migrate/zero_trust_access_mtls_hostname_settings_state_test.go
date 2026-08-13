package main

import (
	"testing"
)

func TestZeroTrustAccessMTLSHostnameSettingsStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "basic settings transformation with boolean defaults",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_access_mutual_tls_hostname_settings",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "test-account-id",
									"settings": [
										{
											"hostname": "example.com"
										}
									]
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
						"type": "cloudflare_zero_trust_access_mtls_hostname_settings",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "test-account-id",
									"settings": [
										{
											"hostname": "example.com",
											"china_network": false,
											"client_certificate_forwarding": false
										}
									]
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "settings with existing boolean values",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zero_trust_access_mtls_hostname_settings",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "test-account-id",
									"settings": [
										{
											"hostname": "example.com",
											"china_network": true,
											"client_certificate_forwarding": true
										}
									]
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
						"type": "cloudflare_zero_trust_access_mtls_hostname_settings",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "test-account-id",
									"settings": [
										{
											"hostname": "example.com",
											"china_network": true,
											"client_certificate_forwarding": true
										}
									]
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "multiple settings entries",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_access_mutual_tls_hostname_settings",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"settings": [
										{
											"hostname": "app1.example.com",
											"client_certificate_forwarding": true
										},
										{
											"hostname": "app2.example.com",
											"china_network": false
										}
									]
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
						"type": "cloudflare_zero_trust_access_mtls_hostname_settings",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"settings": [
										{
											"hostname": "app1.example.com",
											"china_network": false,
											"client_certificate_forwarding": true
										},
										{
											"hostname": "app2.example.com",
											"china_network": false,
											"client_certificate_forwarding": false
										}
									]
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