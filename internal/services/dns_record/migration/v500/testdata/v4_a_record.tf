resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  proxied = true
  tags    = ["tf-applied"]
  ttl     = 1
  type    = "A"
  content = "52.152.96.252"
}
