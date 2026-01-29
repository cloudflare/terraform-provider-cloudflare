resource "cloudflare_record" "%[1]s" {
  zone_id  = "%[2]s"
  name     = "%[3]s"
  type     = "MX"
  content  = "mail.example.com"
  priority = 10
}
