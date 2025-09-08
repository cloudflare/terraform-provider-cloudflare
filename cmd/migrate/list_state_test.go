package main

import (
	"testing"
)

func TestCloudflareListStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "IP list state migration",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"account_id": "abc123",
							"name": "ip_list",
							"kind": "ip",
							"item": [
								{
									"comment": "First IP",
									"value": [{
										"ip": "1.1.1.1"
									}]
								},
								{
									"comment": "Second IP",
									"value": [{
										"ip": "1.1.1.2"
									}]
								}
							]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"account_id": "abc123",
							"name": "ip_list",
							"kind": "ip",
							"num_items": 2,
							"items": [
								{
									"comment": "First IP",
									"ip": "1.1.1.1"
								},
								{
									"comment": "Second IP",
									"ip": "1.1.1.2"
								}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "ASN list state migration",
			Input: `{
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"kind": "asn",
							"item": [
								{
									"comment": "Google ASN",
									"value": [{
										"asn": 15169
									}]
								},
								{
									"value": [{
										"asn": 13335
									}]
								}
							]
						}
					}]
				}]
			}`,
			Expected: `{
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"kind": "asn",
							"num_items": 2,
							"items": [
								{
									"comment": "Google ASN",
									"asn": 15169
								},
								{
									"asn": 13335
								}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "Hostname list state migration",
			Input: `{
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"kind": "hostname",
							"item": [
								{
									"comment": "Example hostname",
									"value": [{
										"hostname": [{
											"url_hostname": "example.com"
										}]
									}]
								}
							]
						}
					}]
				}]
			}`,
			Expected: `{
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"kind": "hostname",
							"num_items": 1,
							"items": [
								{
									"comment": "Example hostname",
									"hostname": {
										"url_hostname": "example.com"
									}
								}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "Redirect list state migration with boolean conversions",
			Input: `{
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"kind": "redirect",
							"item": [
								{
									"comment": "Main redirect",
									"value": [{
										"redirect": [{
											"source_url": "example.com/old",
											"target_url": "example.com/new",
											"include_subdomains": "enabled",
											"subpath_matching": "disabled",
											"preserve_query_string": "enabled",
											"preserve_path_suffix": "disabled",
											"status_code": 301
										}]
									}]
								}
							]
						}
					}]
				}]
			}`,
			Expected: `{
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"kind": "redirect",
							"num_items": 1,
							"items": [
								{
									"comment": "Main redirect",
									"redirect": {
										"source_url": "example.com/old",
										"target_url": "example.com/new",
										"include_subdomains": true,
										"subpath_matching": false,
										"preserve_query_string": true,
										"preserve_path_suffix": false,
										"status_code": 301
									}
								}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "Empty item array removal",
			Input: `{
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"kind": "ip",
							"item": []
						}
					}]
				}]
			}`,
			Expected: `{
				"resources": [{
					"type": "cloudflare_list",
					"instances": [{
						"attributes": {
							"kind": "ip",
							"num_items": 0
						}
					}]
				}]
			}`,
		},
	}
	
	RunFullStateTransformationTests(t, tests)
}