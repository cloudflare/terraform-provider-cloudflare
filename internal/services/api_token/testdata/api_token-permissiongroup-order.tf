resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"
  status = "active"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "%[2]s"
    },{
      id = "%[3]s"
    }]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }]
}

data "cloudflare_api_token" "%[1]s" {
  token_id = cloudflare_api_token.%[1]s.id
  depends_on = [cloudflare_api_token.%[1]s]
}