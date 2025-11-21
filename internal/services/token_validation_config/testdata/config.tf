resource "cloudflare_token_validation_config" "%[1]s" {
  zone_id = "%[2]s"
  token_type = "JWT"
  title = "%[3]s"
  description = "%[4]s"
  token_sources = [%[5]s]
  credentials = {
    keys = [
%[6]s
    ]
  }
}