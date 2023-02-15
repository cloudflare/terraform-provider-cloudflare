# Predefined profile
resource "cloudflare_dlp_profile" "example_predefined" {
  account_id          = "0da42c8d2132a9ddaf714f9e7c920711"
  name                = "Example Predefined Profile"
  type                = "predefined"
  allowed_match_count = 0

  entry {
	name = "Mastercard Card Number"
	enabled = true
  }

  entry {
	name = "Union Pay Card Number"
	enabled = false
  }
}

# Custom profile
resource "cloudflare_dlp_profile" "example_custom" {
  account_id          = "0da42c8d2132a9ddaf714f9e7c920711"
  name                = "Example Custom Profile"
  description         = "A profile with example entries"
  type                = "custom"
  allowed_match_count = 0

  entry {
	name = "Matches visa credit cards"
	enabled = true
	pattern {
		regex = "4\d{3}([-\. ])?\d{4}([-\. ])?\d{4}([-\. ])?\d{4}"
		validation = "luhn"
	}
  }

  entry {
	name = "Matches diners club card"
	enabled = true
	pattern {
		regex = "(?:0[0-5]|[68][0-9])[0-9]{11}"
		validation = "luhn"
	}
  }
}