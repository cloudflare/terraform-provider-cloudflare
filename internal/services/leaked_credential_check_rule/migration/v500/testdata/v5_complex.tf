resource "cloudflare_leaked_credential_check_rule" "%s" {
  zone_id  = "%s"
  username = "lookup_json_string(lookup_json_string(http.request.body.raw, \"payload\"), \"username\")"
  password = "http.request.headers[\"x-password\"][0]"
}
