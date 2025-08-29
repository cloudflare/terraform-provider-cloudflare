package main

import "testing"

func TestListItemTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "",
			Config: `
resource "cloudflare_list" "known_ips" {
  account_id  = local.account_id
  name        = "known_ips"
  description = "dogs are great"
  kind        = "ip"

  dynamic "item" {
    for_each = local.ips_data.ips
    iterator = ip_entry
    content {
      comment = ip_entry.value.comment

      value {
        ip = ip_entry.value.ip
      }
    }
  }
}
`,
			Expected: []string{`
resource "cloudflare_list" "known_ips" {
  account_id  = local.account_id
  description = "dogs are great"
  kind        = "ip"
  name        = "known_ips"
}
`, `
resource "cloudflare_list_item" "known_ips" {
  account_id = cloudflare_list_item.known_ips.account_id
  comment    = each.value.content.comment
  for_each = [
    for ip_entry in
    [
      for value in
      local.ips_data.ips
      : {
        key   = value
        value = value
      }
    ]
    : {
      content = {
        comment = ip_entry.value.comment
        value   = { ip = ip_entry.value.ip }
      }
    }
  ]
  ip      = each.value.content.ip.value
  list_id = cloudflare_list_item.known_ips.id
}`,
			},
		},
		{

			Name: "",
			Config: `
resource "cloudflare_list" "persona_webhook_ips" {
  account_id  = local.account_id
  name        = "persona_webhook_ips"
  description = "dogs are great"
  kind        = "ip"

  dynamic "item" {
    for_each = local.persona_webhook_ips
    content {
      value {
        ip = item.value
      }
    }
  }
}
`,
			Expected: []string{`
resource "cloudflare_list" "persona_webhook_ips" {
  account_id  = local.account_id
  description = "dogs are great"
  kind        = "ip"
  name        = "persona_webhook_ips"
}
`, `
resource "cloudflare_list_item" "persona_webhook_ips" {
  account_id = cloudflare_list_item.persona_webhook_ips.account_id
  comment    = each.value.content.comment
  for_each = [
    for item in
    [
      for value in
      local.persona_webhook_ips
      : {
        key   = value
        value = value
      }
    ]
    : { content = { value = { ip = item.value } } }
  ]
  ip      = each.value.content.ip.value
  list_id = cloudflare_list_item.persona_webhook_ips.id
}
`,
			},
		},
	}

	RunTransformationTests(t, tests, transformListItem)
}
