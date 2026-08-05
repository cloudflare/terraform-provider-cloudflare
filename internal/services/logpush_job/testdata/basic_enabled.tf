resource "cloudflare_logpush_job" "%s" {
  account_id       = "%s"
  dataset          = "%s"
  destination_conf = "%s"
  name             = "%s"
  enabled          = %t
}
