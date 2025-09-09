package main

import (
	"testing"
)

// State transformation tests for zero_trust_access_group
func TestZeroTrustAccessGroupStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "transforms_v4_include_with_email_arrays",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [{
								"email": ["user1@example.com", "user2@example.com"]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"email": {"email": "user1@example.com"}},
								{"email": {"email": "user2@example.com"}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_v4_include_with_email_domain_arrays",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [{
								"email_domain": ["example.com", "test.com"]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"email_domain": {"domain": "example.com"}},
								{"email_domain": {"domain": "test.com"}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_v4_include_with_ip_arrays",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [{
								"ip": ["192.0.2.1/32", "192.0.2.2/32"]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"ip": {"ip": "192.0.2.1/32"}},
								{"ip": {"ip": "192.0.2.2/32"}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_v4_boolean_fields_to_empty_objects",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [{
								"everyone": true,
								"certificate": true,
								"any_valid_service_token": true
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"everyone": {}},
								{"certificate": {}},
								{"any_valid_service_token": {}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_exclude_rules",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"exclude": [{
								"email": ["blocked@example.com"],
								"ip": ["192.168.1.1/32"]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"exclude": [
								{"email": {"email": "blocked@example.com"}},
								{"ip": {"ip": "192.168.1.1/32"}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_require_rules",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"require": [{
								"email_domain": ["company.com"],
								"certificate": true
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"require": [
								{"email_domain": {"domain": "company.com"}},
								{"certificate": {}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_v5_format_unchanged",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"email": {"email": "user@example.com"}},
								{"everyone": {}}
							]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"email": {"email": "user@example.com"}},
								{"everyone": {}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_mixed_rule_types",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [{
								"email": ["user1@example.com", "user2@example.com"],
								"ip": ["10.0.0.1/32"],
								"everyone": true
							}],
							"exclude": [{
								"geo": ["CN", "RU"]
							}],
							"require": [{
								"email_domain": ["company.com"]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"email": {"email": "user1@example.com"}},
								{"email": {"email": "user2@example.com"}},
								{"ip": {"ip": "10.0.0.1/32"}},
								{"everyone": {}}
							],
							"exclude": [
								{"geo": {"country_code": "CN"}},
								{"geo": {"country_code": "RU"}}
							],
							"require": [
								{"email_domain": {"domain": "company.com"}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_empty_rules",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [],
							"exclude": [],
							"require": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [],
							"exclude": [],
							"require": []
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_additional_attributes",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [{
								"group": ["group1", "group2"],
								"service_token": ["token1"],
								"email_list": ["list1"],
								"ip_list": ["iplist1"],
								"login_method": ["method1"],
								"device_posture": ["posture1"]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"group": {"id": "group1"}},
								{"group": {"id": "group2"}},
								{"service_token": {"token_id": "token1"}},
								{"email_list": {"id": "list1"}},
								{"ip_list": {"id": "iplist1"}},
								{"login_method": {"id": "method1"}},
								{"device_posture": {"integration_uid": "posture1"}}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_geo_attribute",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [{
								"geo": ["US", "CA", "GB"]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-group",
							"account_id": "acc-123",
							"include": [
								{"geo": {"country_code": "US"}},
								{"geo": {"country_code": "CA"}},
								{"geo": {"country_code": "GB"}}
							]
						}
					}]
				}]
			}`,
		},
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
