resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-cloudflare-config-%[1]s"
  secret     = base64encode("cloudflare-config-tunnel-secret-at-least-32-bytes-long-test")
  config_src = "cloudflare"
}
