resource "cloudflare_email_routing_rule" "example_email_routing_rule" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  actions = [{
    type = "forward"
    value = ["destinationaddress@example.net"]
  }]
  matchers = [{
    type = "literal"
    field = "to"
    value = "test@example.com"
  }]
  enabled = true
  name = "Send to user@example.net rule."
  priority = 0
}
