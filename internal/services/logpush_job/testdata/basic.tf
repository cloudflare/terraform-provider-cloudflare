resource "cloudflare_logpush_job" "%[1]s" {
  enabled = false
  account_id = "%[2]s"
  destination_conf = "%[3]s"
  dataset = "%[4]s"
  name = "%[5]s"
}
