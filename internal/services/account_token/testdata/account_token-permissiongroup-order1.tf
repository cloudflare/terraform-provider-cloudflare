data "cloudflare_account_api_token_permission_groups_list" "dns_read" {
  account_id = "%[2]s"
  name       = "DNS Read"
  scope      = "com.cloudflare.api.account.zone"
}

data "cloudflare_account_api_token_permission_groups_list" "disable_esc_read" {
  account_id = "%[2]s"
  name       = "Disable ESC Read"
  scope      = "com.cloudflare.api.account.zone"
}

resource "cloudflare_account_token" "test_account_token" {
  name       = "%[1]s"
  account_id = "%[2]s"
  status     = "active"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = data.cloudflare_account_api_token_permission_groups_list.disable_esc_read.result[0].id
      }, {
      id = data.cloudflare_account_api_token_permission_groups_list.dns_read.result[0].id
    }]
    resources = jsonencode(
      {
        "com.cloudflare.api.account.%[2]s" = {
          "com.cloudflare.api.account.zone.*" : "*"
        }
      }
    )
  }]
}

data "cloudflare_account_token" "test_account_token" {
  account_id = "%[2]s"
  token_id   = cloudflare_account_token.test_account_token.id
  depends_on = [cloudflare_account_token.test_account_token]
}
