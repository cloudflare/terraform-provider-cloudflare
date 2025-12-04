resource "cloudflare_zero_trust_device_default_profile" "%[1]s" {
  account_id  = "%[2]s"
  auto_connect = 60

  exclude = [
    {
      address     = "192.168.1.0/24"
      description = "Local network"
    },
    {
      host        = "example.com"
      description = "Corporate domain"
    }
  ]

  service_mode_v2 = {
    mode = "warp"
  }
}
