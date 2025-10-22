resource "cloudflare_r2_bucket" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  jurisdiction = "eu"
}

resource "cloudflare_r2_custom_domain" "%[1]s" {
  account_id   = "%[2]s"
  bucket_name  = cloudflare_r2_bucket.%[1]s.name
  domain       = "%[4]s"
  zone_id      = "%[3]s"
  jurisdiction = "eu"
  enabled      = true
  min_tls      = "1.3"
}
