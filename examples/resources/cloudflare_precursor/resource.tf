resource "cloudflare_precursor" "example_precursor" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  mode = "monitor"
  rules = [{
    expression = "http.request.uri.path eq \"/login\""
    show_interstitial = true
    description = "Show an interstitial on the login path"
    enabled = true
  }]
  show_interstitial = true
}
