resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  previews_base_config = {
    limits = {
      cpu_ms = %[3]d
    }
  }
}
