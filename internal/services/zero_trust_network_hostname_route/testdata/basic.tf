resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s" {
  account_id    = "%[2]s"
  name          = "%[1]s"
  tunnel_secret = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_zero_trust_network_hostname_route" "%[1]s" {
  account_id = "%[2]s"
  hostname   = "%[3]s.test.example.com"
  tunnel_id  = cloudflare_zero_trust_tunnel_cloudflared.%[1]s.id
  comment    = "Test hostname route for tf-acctest-%[1]s"
}

data "cloudflare_zero_trust_network_hostname_route" "%[1]s" {
  account_id = "%[2]s"
  hostname_route_id = cloudflare_zero_trust_network_hostname_route.%[1]s.id
}
