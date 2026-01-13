resource "cloudflare_leaked_credential_check" "%[2]s" {
  zone_id = "%[1]s"
  enabled = true
}

resource "cloudflare_leaked_credential_check_rule" "%[2]s" {
  password = "lookup_json_string(http.request.body.raw, \"pass\")"
  username = "lookup_json_string(http.request.body.raw, \"username\")"
  zone_id  = "%[1]s"
}