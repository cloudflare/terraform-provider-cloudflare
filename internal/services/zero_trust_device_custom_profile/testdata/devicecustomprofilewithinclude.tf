resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  match       = "os.name == \"Windows\""
  precedence  = 100
  enabled     = true
  description = "Profile with include list"
  
  include = [
    {
      address     = "192.168.1.0/24"
      description = "Corporate network"
    },
    {
      host        = "app.example.com"
      description = "Corporate application"
    }
  ]
}