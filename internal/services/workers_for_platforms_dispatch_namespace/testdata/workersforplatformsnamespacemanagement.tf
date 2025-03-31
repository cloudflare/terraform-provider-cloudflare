
  resource "cloudflare_workers_for_platforms_dispatch_namespace" "%[1]s" {
    account_id = "%[2]s"
    name       = "%[1]s"
  }