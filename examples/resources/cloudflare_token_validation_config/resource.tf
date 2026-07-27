resource "cloudflare_token_validation_config" "example_token_validation_config" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  credentials = {
    keys = [{
      alg = "RS256"
      e = "e"
      kid = "kid"
      kty = "RSA"
      n = "n"
    }]
  }
  description = "Long description for Token Validation Configuration"
  title = "Example Token Validation Configuration"
  token_sources = ["http.request.headers[\"x-auth\"][0]", "http.request.cookies[\"Authorization\"][0]"]
  token_type = "JWT"
}
