data "cloudflare_account_api_token_permission_groups_list" "dns_read" {
  account_id = "%[2]s"
  name       = "DNS Read"
  scope      = "com.cloudflare.api.account.zone"
}

resource "cloudflare_account_token" "test_account_token" {
  name       = "%[1]s"
  account_id = "%[2]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = data.cloudflare_account_api_token_permission_groups_list.dns_read.result[0].id
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = "*"
    })
  }]

  condition = {
    request_ip = {
      in = ["192.0.2.1/32"]
    }
  }
}
