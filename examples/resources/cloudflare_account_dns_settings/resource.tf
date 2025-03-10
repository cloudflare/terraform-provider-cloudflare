resource "cloudflare_account_dns_settings" "example_account_dns_settings" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  zone_defaults = {
    flatten_all_cnames = false
    foundation_dns = false
    internal_dns = {
      reference_zone_id = "reference_zone_id"
    }
    multi_provider = false
    nameservers = {
      type = "cloudflare.standard"
    }
    ns_ttl = 86400
    secondary_overrides = false
    soa = {
      expire = 604800
      min_ttl = 1800
      mname = "kristina.ns.cloudflare.com"
      refresh = 10000
      retry = 2400
      rname = "admin.example.com"
      ttl = 3600
    }
    zone_mode = "standard"
  }
}
