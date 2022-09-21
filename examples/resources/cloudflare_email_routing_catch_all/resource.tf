resource "cloudflare_email_routing_catch_all" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "example catch all"
  enabled = true

  matcher {
    type = "all"
  }

  action {
    type  = "forward"
    value = ["destinationaddress@example.net"]
  }
}
