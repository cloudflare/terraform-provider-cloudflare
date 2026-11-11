resource "cloudflare_observatory_scheduled_test" "%[1]s" {
  zone_id   = %[2]q
  url       = urlencode("%[3]s/%[4]s")
  frequency = %[5]q
}
