resource "cloudflare_leaked_credential_check_rule" "%s" {
  zone_id  = "%s"
  username = "lookup_json_string(http.request.body.raw, \"user\")"
  password = "lookup_json_string(http.request.body.raw, \"pass\")"
}
