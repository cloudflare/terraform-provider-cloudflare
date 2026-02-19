resource "cloudflare_zero_trust_device_default_profile" "%[1]s" {
  account_id = "%[2]s"
  exclude    = []

  service_mode_v2 = {
    mode = "warp"
  }
}
