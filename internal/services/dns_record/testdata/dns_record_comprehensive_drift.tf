# Basic A record
resource "cloudflare_dns_record" "%[1]s_basic_a" {
  zone_id = "%[2]s"
  name    = "tf-acctest-basic.%[1]s.%[3]s"
  type    = "A"
  content = "%[4]s"
  ttl     = 3600
  proxied = false
}

# CAA record with data field
resource "cloudflare_dns_record" "%[1]s_caa" {
  zone_id = "%[2]s"
  name    = "tf-acctest-caa.%[1]s.%[3]s"
  type    = "CAA"
  ttl     = 3600

  data = {
    flags = "0"
    tag   = "issue"
    value = "letsencrypt.org"
  }
}

# CNAME record (case-sensitive test)
resource "cloudflare_dns_record" "%[1]s_cname" {
  zone_id = "%[2]s"
  name    = "tf-acctest-cname.%[1]s.%[3]s"
  type    = "CNAME"
  content = "Target.%[3]s"
  ttl     = 3600
  proxied = false
}

# Record with tags
resource "cloudflare_dns_record" "%[1]s_with_tags" {
  zone_id = "%[2]s"
  name    = "tf-acctest-tags.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.50"
  ttl     = 3600
  proxied = false
  tags    = ["tag1", "tag2"]
}

# Record with settings
resource "cloudflare_dns_record" "%[1]s_with_settings" {
  zone_id = "%[2]s"
  name    = "tf-acctest-settings.%[1]s.%[3]s"
  type    = "CNAME"
  content = "target.%[3]s"
  ttl     = 3600
  proxied = false

  settings = {
    flatten_cname = false
  }
}