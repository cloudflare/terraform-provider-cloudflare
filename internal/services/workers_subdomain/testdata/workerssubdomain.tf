resource "cloudflare_workers_subdomain" "%[1]s" {
  account_id = "%[2]s"
  subdomain  = "%[3]s"
}
