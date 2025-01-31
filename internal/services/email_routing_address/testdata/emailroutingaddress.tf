
  resource "cloudflare_email_routing_address" "%[1]s" {
    account_id = "%[2]s"
    email      = "%[1]s@example.com"
  }