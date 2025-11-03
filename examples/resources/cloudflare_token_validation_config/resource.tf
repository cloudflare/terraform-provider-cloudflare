resource "cloudflare_token_validation_config" "example_token_validation_config" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  credentials = {
    keys = [{
      alg = "ES256"
      crv = "P-256"
      kid = "38013f13-c266-4eec-a72a-92ec92779f21"
      kty = "EC"
      x = "KN53JRwN3wCjm2o39bvZUX2VdrsHzS8pxOAGjm8m7EQ"
      y = "lnkkzIxaveggz-HFhcMWW15nxvOj0Z_uQsXbpK0GFcY"
    }]
  }
  description = "Long description for Token Validation Configuration"
  title = "Example Token Validation Configuration"
  token_sources = ["http.request.headers[\"x-auth\"][0]", "http.request.cookies[\"Authorization\"][0]"]
  token_type = "JWT"
}
