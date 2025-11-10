# Look up zone by zone ID
data "cloudflare_zone" "%[1]s" {
  zone_id = "%[2]s"
}