resource "cloudflare_precursor" "example_precursor" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  default_mode = "min-friction"
  enforcement_rules = [{
    expression = "http.request.uri.path eq \"/login\""
    mode = "max-security"
    description = "Ease friction on the login path"
    enabled = true
  }]
}
