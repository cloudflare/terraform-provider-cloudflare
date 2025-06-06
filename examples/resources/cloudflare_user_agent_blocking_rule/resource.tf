resource "cloudflare_user_agent_blocking_rule" "example_user_agent_blocking_rule" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  configuration = {
    target = "ua"
    value = "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)"
  }
  mode = "challenge"
}
