data "cloudflare_account_api_token_permission_groups_list" "dns_read" {
  account_id = "%[2]s"
  name       = "DNS Read"
  scope      = "com.cloudflare.api.account.zone"
}

data "cloudflare_account_api_token_permission_groups_list" "dns_write" {
  account_id = "%[2]s"
  name       = "DNS Write"
  scope      = "com.cloudflare.api.account.zone"
}

data "cloudflare_account_api_token_permission_groups_list" "zone_security_center_insights_read" {
  account_id = "%[2]s"
  name       = "Zone Security Center Insights Read"
  scope      = "com.cloudflare.api.account.zone"
}

data "cloudflare_account_api_token_permission_groups_list" "account_api_tokens_write" {
  account_id = "%[2]s"
  name       = "Account API Tokens Write"
  scope      = "com.cloudflare.api.account"
}

data "cloudflare_account_api_token_permission_groups_list" "account_api_tokens_read" {
  account_id = "%[2]s"
  name       = "Account API Tokens Read"
  scope      = "com.cloudflare.api.account"
}

resource "cloudflare_api_token" "test_account_token" {
  name = "%[1]s"
  policies = [
    {
      effect = "allow"
      permission_groups = [
        { id = data.cloudflare_account_api_token_permission_groups_list.account_api_tokens_write.result[0].id },
        { id = data.cloudflare_account_api_token_permission_groups_list.account_api_tokens_read.result[0].id },
      ]
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = "*"
      })
    },
    {
      effect = "allow"
      permission_groups = [
        { id = data.cloudflare_account_api_token_permission_groups_list.dns_write.result[0].id },
        { id = data.cloudflare_account_api_token_permission_groups_list.dns_read.result[0].id },
        // changed this one permission group
        { id = data.cloudflare_account_api_token_permission_groups_list.zone_security_center_insights_read.result[0].id },
      ]
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = {
          "com.cloudflare.api.account.zone.*" = "*"
        }
      })
    },
  ]
  status = "active"
}
