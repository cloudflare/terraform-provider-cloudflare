resource "cloudflare_email_routing_rule" "%[1]s" {
  zone_id = "%[2]s"
  enabled = "%[3]t"
  priority = "%[4]d"
  name = "terraform rule"
  matchers = [ {
    field  = "to"
    type = "literal"
    value = "test@example.com"
  }]
  actions = [ {
    type = "forward"
    value = ["destinationaddress@example.net"]
  }]
}
