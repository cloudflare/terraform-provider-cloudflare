# Enable Content Scanning before trying to add custom scan expressions
resource "cloudflare_content_scanning" "example" {
    zone_id = "399c6f4950c01a5a141b99ff7fbcbd8b"
    enabled = true
}

resource "cloudflare_content_scanning_expression" "first_example" {
	zone_id = cloudflare_content_scanning.example.zone_id
	payload = "lookup_json_string(http.request.body.raw, \"file\")"
}

resource "cloudflare_content_scanning_expression" "second_example" {
	zone_id = cloudflare_content_scanning.example.zone_id
	payload = "lookup_json_string(http.request.body.raw, \"document\")"
}
