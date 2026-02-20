resource "cloudflare_zero_trust_device_default_profile_local_domain_fallback" "%[1]s" {
  account_id = "%[2]s"
  domains = [{
    suffix      = "corp.example.com"
    description = "Corporate domain with multiple DNS servers"
    dns_server  = ["10.0.0.1", "10.0.0.2", "10.0.0.3"]
  }]
}
