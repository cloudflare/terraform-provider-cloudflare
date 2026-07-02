resource "cloudflare_email_routing_catch_all" "example_email_routing_catch_all" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  actions = [{
    type = "forward"
    value = ["destinationaddress@example.net"]
  }]
  matchers = [{
    type = "all"
  }]
  enabled = true
  name = "Send to user@example.net rule."
  owner_worker_tag = "a7e6fb77503c41d8a7f3113c6918f10c"
  source = "api"
}
