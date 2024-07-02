
resource "cloudflare_teams_proxy_endpoint" "%[1]s" {
  name        = "%[1]s"
  account_id  = "%[2]s"
  ips  = ["104.16.132.229/32"]
}
