resource "cloudflare_logpush_job" "%s" {
  zone_id          = "%s"
  dataset          = "%s"
  destination_conf = "%s"
  kind             = "%s"
}
