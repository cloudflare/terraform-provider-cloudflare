# Enable the Leaked Credentials Check detection before trying
# to add detections.
resource "cloudflare_leaked_credential_check" "example" {
	zone_id = "399c6f4950c01a5a141b99ff7fbcbd8b"
	enabled = true
}

resource "cloudflare_leaked_credential_check_rule" "example" {
  zone_id = cloudflare_leaked_credential_check.example.zone_id
  username = "lookup_json_string(http.request.body.raw, \"user\")"
	password = "lookup_json_string(http.request.body.raw, \"pass\")"
}
