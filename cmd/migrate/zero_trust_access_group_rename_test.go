package main

import (
	"testing"
)

// Test resource type rename from cloudflare_access_group to cloudflare_zero_trust_access_group
func TestZeroTrustAccessGroupResourceRename(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "renames_cloudflare_access_group_to_cloudflare_zero_trust_access_group",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_access_group",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [{
								"email": ["user@example.com"]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"email": {"email": "user@example.com"}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_cloudflare_access_group_with_complex_rules",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_access_group",
					"name": "complex",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [{
								"email": ["user1@example.com", "user2@example.com"],
								"everyone": true,
								"ip": ["192.0.2.0/24"]
							}],
							"exclude": [{
								"email_domain": ["blocked.com"]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"name": "complex",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"email": {"email": "user1@example.com"}},
								{"email": {"email": "user2@example.com"}},
								{"everyone": {}},
								{"ip": {"ip": "192.0.2.0/24"}}
							],
							"exclude": [
								{"email_domain": {"domain": "blocked.com"}}
							]
						}
					}]
				}]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}