resource "cloudflare_zero_trust_dns_location" "example" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "office"
  client_default = true
  ecs_support    = false

  networks = [{
    network = "203.0.113.1/32"
    },
    {
      network = "203.0.113.2/32"
  }]

}
