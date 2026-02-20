resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-minimal-%[1]s"
  secret     = base64encode("minimal-tunnel-secret-that-is-at-least-32-bytes-long-testing")
  config_src = "local"
}
