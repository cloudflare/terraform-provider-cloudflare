resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  match       = "os.name == \"Windows\""
  precedence  = %[3]d
  enabled     = true
  description = "Profile to test exclude drift"

  exclude = [
    {
      address     = "10.0.0.0/8"
      description = "Private network class A"
    },
    {
      address = "172.16.0.0/12"
    },
    {
      host        = "*.example.com"
      description = "Company domain"
    }
  ]
}
