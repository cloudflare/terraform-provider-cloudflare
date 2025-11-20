package main

import (
	"testing"
)

func TestSpectrumApplicationStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "origin_port integer value preserved",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_spectrum_application",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"protocol": "tcp/3306",
									"dns": [
										{
											"type": "CNAME",
											"name": "test.example.com"
										}
									],
									"origin_direct": ["tcp://128.66.0.2:3306"],
									"origin_port": 3306
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
						"type": "cloudflare_spectrum_application",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"protocol": "tcp/3306",
									"dns": {
										"type": "CNAME",
										"name": "test.example.com"
									},
									"origin_direct": ["tcp://128.66.0.2:3306"],
									"origin_port": {
										"value": 3306,
										"type": "number"
									}
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "origin_port_range block converted to origin_port string",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_spectrum_application",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"protocol": "tcp/3306",
									"dns": [
										{
											"type": "CNAME",
											"name": "test.example.com"
										}
									],
									"origin_direct": ["tcp://128.66.0.1:23"],
									"origin_port_range": [
										{
											"start": 3306,
											"end": 3310
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
						"type": "cloudflare_spectrum_application",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"protocol": "tcp/3306",
									"dns": {
										"type": "CNAME",
										"name": "test.example.com"
									},
									"origin_direct": ["tcp://128.66.0.1:23"],
									"origin_port": {
										"value": "3306-3310",
										"type": "string"
									}
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "all nested objects converted from arrays",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_spectrum_application",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"protocol": "tcp/443",
									"dns": [
										{
											"type": "CNAME",
											"name": "test.example.com"
										}
									],
									"edge_ips": [
										{
											"type": "dynamic",
											"connectivity": "all"
										}
									],
									"origin_dns": [
										{
											"name": "origin.example.com",
											"type": "A"
										}
									],
									"origin_direct": ["tcp://128.66.0.3:443"],
									"tls": "flexible",
									"argo_smart_routing": true,
									"proxy_protocol": "v1",
									"ip_firewall": true,
									"traffic_type": "direct"
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
						"type": "cloudflare_spectrum_application",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"protocol": "tcp/443",
									"dns": {
										"type": "CNAME",
										"name": "test.example.com"
									},
									"edge_ips": {
										"type": "dynamic",
										"connectivity": "all"
									},
									"origin_dns": {
										"name": "origin.example.com",
										"type": "A"
									},
									"origin_direct": ["tcp://128.66.0.3:443"],
									"tls": "flexible",
									"argo_smart_routing": true,
									"proxy_protocol": "v1",
									"ip_firewall": true,
									"traffic_type": "direct"
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