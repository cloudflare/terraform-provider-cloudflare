resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  proxied = true
  tags    = ["tf-applied"]
  ttl     = 1
  type    = "CNAME"
  content = "abc-browser-external.foo.com"
}
