resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Test with custom entries"
	
	entries = [
		{
			enabled = true
			name    = "Credit Card Pattern"
			pattern = {
				regex      = "\\d{4}-\\d{4}-\\d{4}-\\d{4}"
				validation = "luhn"
			}
		},
		{
			enabled = false
			name    = "SSN Pattern"
			pattern = {
				regex = "\\d{3}-\\d{2}-\\d{4}"
			}
		}
	]
	
	context_awareness = {
		enabled = true
		skip = {
			files = false
		}
	}
}