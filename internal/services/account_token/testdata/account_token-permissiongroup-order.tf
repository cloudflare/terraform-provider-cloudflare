resource "cloudflare_account_token" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "%[3]s"
    },{
      id = "%[4]s"
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }]
}

data "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  token_id = cloudflare_account_token.%[1]s.id
  depends_on = [cloudflare_account_token.%[1]s]
}