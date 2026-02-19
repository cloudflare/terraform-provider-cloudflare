resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  secret     = base64encode("test-secret-that-is-at-least-32-bytes-long-for-testing")
  config_src = "local"
}
