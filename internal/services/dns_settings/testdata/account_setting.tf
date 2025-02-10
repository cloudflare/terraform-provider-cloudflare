resource "cloudflare_dns_settings" "%[1]s" {
  account_id = "%[2]s"

  zone_defaults = {
    flatten_all_cnames = "%[3]t"
    foundation_dns     = "%[4]t"
    multi_provider     = "%[5]t"
    nameservers = {
      type = "cloudflare.standard"
    }
    ns_ttl = 86400
    secondary_overrides = false
    soa = {
      expire  = 604800
      min_ttl = 1800
      mname   = "kristina.ns.cloudflare.com"
      refresh = 10000
      retry   = 2400
      rname   = "admin.example.com"
      ttl     = 3600
    }
    zone_mode = "standard"
  }
}
