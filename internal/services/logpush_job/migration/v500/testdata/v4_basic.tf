resource "cloudflare_logpush_job" "%[1]s" {
  account_id       = "%[2]s"
  dataset          = "audit_logs"
  destination_conf = "https://logpush-receiver.sd.cfplat.com"
  enabled          = true
  name             = "%[3]s"
}
