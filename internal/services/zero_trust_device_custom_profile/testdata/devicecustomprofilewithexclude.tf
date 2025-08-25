resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  match       = "os.name == \"Windows\""
  precedence  = %[3]d
  enabled     = true
  description = "Profile with exclude list"
  
  exclude = [
    {
      address     = "10.0.0.0/8"
      description = "Private network range"
    },
    {
      host        = "internal.example.com"
      description = "Internal domain"
    }
  ]
}
