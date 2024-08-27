# Account level CA certificate import.
$ terraform import cloudflare_zero_trust_access_short_lived_certificate.example account/<account_id>/<application_id>

# Zone level CA certificate import.
$ terraform import cloudflare_zero_trust_access_short_lived_certificate.example account/<zone_id>/<application_id>
