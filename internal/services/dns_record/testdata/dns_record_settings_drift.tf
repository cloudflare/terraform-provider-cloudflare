resource "cloudflare_dns_record" "%[1]s_a_with_settings" {
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

resource "cloudflare_dns_record" "%[1]s_a_empty_settings" {
  zone_id = "%[2]s"
  name    = "tf-acctest-empty-settings.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.31"
  ttl     = 3600
  proxied = false

  settings = {}
}