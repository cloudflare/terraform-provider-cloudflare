resource "cloudflare_zero_trust_access_application" "%[1]s" {
  name       = "%[1]s"
  account_id = "%[3]s"
  domain     = "%[1]s.%[2]s"
  type       = "self_hosted"
}

resource "cloudflare_zero_trust_access_group" "%[1]s" {
    account_id     = "%[3]s"
    name           = "%[1]s"
    include =[{
      ip = {
        ip = "127.0.0.1/32"
      }
    }]
}

resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  name           = "%[1]s"
  account_id     = "%[3]s"
  decision       = "non_identity"
  include = [{
    group = {
      id = cloudflare_zero_trust_access_group.%[1]s.id
    }
  }]
}
