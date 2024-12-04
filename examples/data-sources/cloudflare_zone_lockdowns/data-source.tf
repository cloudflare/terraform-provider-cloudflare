data "cloudflare_zone_lockdowns" "example_zone_lockdowns" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  created_on = "2014-01-01T05:20:00.12345Z"
  description = "endpoints"
  description_search = "endpoints"
  ip = "1.2.3.4"
  ip_range_search = "1.2.3.0/16"
  ip_search = "1.2.3.4"
  modified_on = "2014-01-01T05:20:00.12345Z"
  priority = 5
  uri_search = "/some/path"
}
