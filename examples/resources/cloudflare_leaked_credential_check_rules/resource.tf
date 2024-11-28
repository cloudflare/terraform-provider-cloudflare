# Enable the Leaked Credentials Check detection
resource "cloudflare_leaked_credential_check" "example" {
	zone_id = "399c6f4950c01a5a141b99ff7fbcbd8b"
	enabled = true
}

# Create user-defined detection patterns for Leaked Credentials Check
resource "cloudflare_leaked_credential_check_rules" "example" {
	# The following line ensures that rules are created only for a zone
	# where the Leaked Credentials Check module is enabled, and that the 
	# module is activated before user-defined detection patterns are created.
    zone_id = cloudflare_leaked_credential_check.example.zone_id

	rule {
		username = "lookup_json_string(http.request.body.raw, \"user\")"
		password = "lookup_json_string(http.request.body.raw, \"pass\")"
	}
    
	rule {
		username = "lookup_json_string(http.request.body.raw, \"id\")"
		password = "lookup_json_string(http.request.body.raw, \"secret\")"
	}
}