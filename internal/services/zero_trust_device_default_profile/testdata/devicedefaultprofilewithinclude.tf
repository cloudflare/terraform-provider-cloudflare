resource "cloudflare_zero_trust_device_default_profile" "%[1]s" {
  account_id  = "%[2]s"
  auto_connect = 60

  include = [
    {
      address     = "10.0.0.0/8"
      description = "Private network"
    },
    {
      host        = "intranet.example.com"
      description = "Internal services"
    }
  ]

  service_mode_v2 = {
    mode = "warp"
  }
}
