package main

import (
	"testing"
)

func TestArgoTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "smart_routing only",
			Config: `resource "cloudflare_argo" "example" {
  zone_id       = var.zone_id
  smart_routing = "on"
}`,
			Expected: []string{
				`resource "cloudflare_argo_smart_routing" "example" {
  zone_id = var.zone_id
  value   = "on"
}`,
				`moved {
  from = cloudflare_argo.example
  to   = cloudflare_argo_smart_routing.example
}`,
			},
		},
		{
			Name: "smart_routing off",
			Config: `resource "cloudflare_argo" "example" {
  zone_id       = var.api_openai_com_zone_id
  smart_routing = "off"
}`,
			Expected: []string{
				`resource "cloudflare_argo_smart_routing" "example" {
  zone_id = var.api_openai_com_zone_id
  value   = "off"
}`,
				`moved {
  from = cloudflare_argo.example
  to   = cloudflare_argo_smart_routing.example
}`,
			},
		},
		{
			Name: "smart_routing with zone reference",
			Config: `resource "cloudflare_argo" "example" {
  zone_id       = cloudflare_zone.operator_chatgpt_com.id
  smart_routing = "on"
}`,
			Expected: []string{
				`resource "cloudflare_argo_smart_routing" "example" {
  zone_id = cloudflare_zone.operator_chatgpt_com.id
  value   = "on"
}`,
				`moved {
  from = cloudflare_argo.example
  to   = cloudflare_argo_smart_routing.example
}`,
			},
		},
		{
			Name: "both smart_routing and tiered_caching",
			Config: `resource "cloudflare_argo" "example" {
  zone_id         = var.zone_id
  smart_routing   = "on"
  tiered_caching  = "on"
}`,
			Expected: []string{
				`resource "cloudflare_argo_smart_routing" "example" {
  zone_id = var.zone_id
  value   = "on"
}`,
				`moved {
  from = cloudflare_argo.example
  to   = cloudflare_argo_smart_routing.example
}`,
				`resource "cloudflare_argo_tiered_caching" "example_tiered" {
  zone_id = var.zone_id
  value   = "on"
}`,
				`moved {
  from = cloudflare_argo.example
  to   = cloudflare_argo_tiered_caching.example_tiered
}`,
			},
		},
		{
			Name: "tiered_caching only",
			Config: `resource "cloudflare_argo" "example" {
  zone_id         = var.zone_id
  tiered_caching  = "on"
}`,
			Expected: []string{
				`resource "cloudflare_argo_tiered_caching" "example" {
  zone_id = var.zone_id
  value   = "on"
}`,
				`moved {
  from = cloudflare_argo.example
  to   = cloudflare_argo_tiered_caching.example
}`,
			},
		},
		{
			Name: "no attributes defaults to smart_routing off",
			Config: `resource "cloudflare_argo" "example" {
  zone_id = var.zone_id
}`,
			Expected: []string{
				`resource "cloudflare_argo_smart_routing" "example" {
  zone_id = var.zone_id
  value   = "off"
}`,
				`moved {
  from = cloudflare_argo.example
  to   = cloudflare_argo_smart_routing.example
}`,
			},
		},
		{
			Name: "with lifecycle block",
			Config: `resource "cloudflare_argo" "example" {
  zone_id       = var.zone_id
  smart_routing = "on"
  
  lifecycle {
    prevent_destroy = true
  }
}`,
			Expected: []string{
				`resource "cloudflare_argo_smart_routing" "example" {
  zone_id = var.zone_id
  value   = "on"

  lifecycle {
    prevent_destroy = true
  }
}`,
				`moved {
  from = cloudflare_argo.example
  to   = cloudflare_argo_smart_routing.example
}`,
			},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}