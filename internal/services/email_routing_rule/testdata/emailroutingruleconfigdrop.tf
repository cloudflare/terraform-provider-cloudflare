resource "cloudflare_email_routing_rule" "%[1]s" {
  zone_id = "%[2]s"
  enabled = "%[3]t"
  priority = "%[4]d"
  name = "%[1]s"
  matchers = [{
    field  = "to"
    type = "literal"
    value = "test@example.com"
  }]
  actions = [{
    type = "drop"
    value = []
  }]
}
