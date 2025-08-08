resource "cloudflare_managed_transforms" "%[1]s" {
  zone_id = "%[2]s"
  managed_request_headers = [{
    id = "add_visitor_location_headers"
    enabled = true
  }]
  
  managed_response_headers = [{
    id = "remove_x_powered_by_headers"
    enabled = true
  }]
}
