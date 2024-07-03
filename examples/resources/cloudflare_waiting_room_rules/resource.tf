resource "cloudflare_waiting_room_rules" "example" {
  zone_id         = "0da42c8d2132a9ddaf714f9e7c920711"
  waiting_room_id = "d41d8cd98f00b204e9800998ecf8427e"

  rules = [{
    description = "bypass ip list"
    expression  = "src.ip in {192.0.2.0 192.0.2.1}"
    action      = "bypass_waiting_room"
    status      = "enabled"
    },
    {
      description = "bypass query string"
      expression  = "http.request.uri.query contains \"bypass=true\""
      action      = "bypass_waiting_room"
      status      = "enabled"
  }]

}
