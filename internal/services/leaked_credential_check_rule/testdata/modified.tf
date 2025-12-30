resource "cloudflare_leaked_credential_check_rule" "%[2]s" {
  password = "lookup_json_string(http.request.body.raw, \"pass_modified\")"
  username = "lookup_json_string(http.request.body.raw, \"username_modified\")"
  zone_id  = "%[1]s"
}