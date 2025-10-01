resource "cloudflare_account_subscription" "%[1]s" {
  account_id = "%[2]s"
  rate_plan = {
    id                 = "teams_free"
    currency           = "USD"
    externally_managed = false
    is_contract        = false
    public_name        = "Teams Free Base"
    scope              = "account"
  }
}