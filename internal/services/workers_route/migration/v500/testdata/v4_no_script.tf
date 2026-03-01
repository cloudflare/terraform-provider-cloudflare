resource "cloudflare_workers_route" "%[1]s" {
  zone_id = "%[2]s"
  pattern = "%[1]s.cfapi.net/api/*"
}
