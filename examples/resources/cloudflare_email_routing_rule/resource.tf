resource "cloudflare_email_routing_rule" "main" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "terraform rule"
  enabled = true
  matchers = [
    {
      type  = "literal",
      field = "to",
      value = "test@example.com"
    }
  ]

  actions = [
    {
      type  = "forward"
      value = ["destinationaddress@example.net"]
    }
  ]
}
