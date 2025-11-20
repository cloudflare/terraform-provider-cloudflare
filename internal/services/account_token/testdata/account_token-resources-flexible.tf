data "cloudflare_account_api_token_permission_groups_list" "dns_read" {
  account_id = "%[2]s"
  name       = "DNS Read"
  scope      = "com.cloudflare.api.account.zone"
}

resource "cloudflare_account_token" "test_account_token" {
  account_id = "%[2]s"
  name       = "%[1]s"
  status     = "active"

  policies = [
    {
      effect            = "allow"
      permission_groups = [{ id = data.cloudflare_account_api_token_permission_groups_list.dns_read.result[0].id }]
      resources         = jsonencode({ "com.cloudflare.api.account.%[2]s" = "*" })
    },
    {
      effect            = "allow"
      permission_groups = [{ id = data.cloudflare_account_api_token_permission_groups_list.dns_read.result[0].id }]
      resources = jsonencode(
        {
          "com.cloudflare.api.account.%[2]s" = {
            "com.cloudflare.api.account.zone.*" : "*"
          }
        }
      )
    }
  ]
}
