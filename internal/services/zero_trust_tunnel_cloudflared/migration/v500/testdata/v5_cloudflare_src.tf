resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s" {
  account_id    = "%[2]s"
  name          = "tf-acc-test-cloudflare-config-%[1]s"
  tunnel_secret = base64encode("cloudflare-config-tunnel-secret-at-least-32-bytes-long-test")
  config_src    = "cloudflare"
}
