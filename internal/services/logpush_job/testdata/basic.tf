resource "cloudflare_logpush_job" "%[1]s" {
  account_id       = "%[2]s"
  dataset          = "%[3]s"
  destination_conf = "%[4]s"
}
