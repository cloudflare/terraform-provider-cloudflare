resource "cloudflare_account_subscription" "%[1]s" {
  account_id = "%[2]s"
  rate_plan = {
    id                 = "teams_free"
    sets               = []
  }
}