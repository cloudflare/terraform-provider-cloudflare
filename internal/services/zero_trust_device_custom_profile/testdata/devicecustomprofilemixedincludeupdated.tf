resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  match       = "os.name == \"Windows\""
  precedence  = %[3]d
  enabled     = true
  description = "Profile with mixed include list"

  include = [
    {
      host        = "example.com"
      description = "Example host"
    },
    {
      address     = "172.64.128.0/20"
      description = "Gateway initial resolved IPs"
    }
  ]
}
