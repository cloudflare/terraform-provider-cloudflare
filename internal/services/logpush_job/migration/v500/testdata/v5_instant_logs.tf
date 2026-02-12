resource "cloudflare_logpush_job" "%[1]s" {
  zone_id          = "%[2]s"
  dataset          = "http_requests"
  destination_conf = "https://logpush-receiver.sd.cfplat.com"
  enabled          = true
  kind             = ""
}
