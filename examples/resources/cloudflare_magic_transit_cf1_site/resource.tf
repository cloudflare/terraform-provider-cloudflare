resource "cloudflare_magic_transit_cf1_site" "example_magic_transit_cf1_site" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  body = [{
    name = "Pad 34"
    description = "Launch Pad 34"
    location = {
      lat = 28.521339842093845
      long = -80.56092644815843
      name = "Cape Canaveral"
    }
  }]
}
