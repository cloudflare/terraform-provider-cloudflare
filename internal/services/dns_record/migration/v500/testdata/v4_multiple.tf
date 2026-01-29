# A record with content field (some configs use content instead of value)
resource "cloudflare_record" "%[1]s_a" {
  zone_id = "%[2]s"
  name    = "api-%[1]s"
  proxied = true
  tags    = ["tf-applied", "production"]
  ttl     = 1
  type    = "A"
  content = "52.152.96.252"
}

# CNAME with value field
resource "cloudflare_record" "%[1]s_cname" {
  zone_id = "%[2]s"
  name    = "www-%[1]s"
  proxied = true
  ttl     = 1
  type    = "CNAME"
  value   = "api-%[1]s.terraform.cfapi.net"
}

# CAA record with data block
resource "cloudflare_record" "%[1]s_caa" {
  zone_id = "%[2]s"
  name    = "caa-%[1]s"
  proxied = false
  ttl     = 1
  type    = "CAA"

  data {
    flags = 0
    tag   = "issue"
    value = "pki.goog"
  }
}

# TXT record
resource "cloudflare_record" "%[1]s_txt" {
  zone_id = "%[2]s"
  name    = "_dmarc-%[1]s"
  proxied = false
  ttl     = 300
  type    = "TXT"
  content = "v=DMARC1; p=reject; sp=reject;"
}
