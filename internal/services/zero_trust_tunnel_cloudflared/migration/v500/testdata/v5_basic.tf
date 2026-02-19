resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s" {
  account_id    = "%[2]s"
  name          = "tf-acc-test-%[1]s"
  tunnel_secret = base64encode("test-secret-that-is-at-least-32-bytes-long-for-testing")
  config_src    = "local"
}
