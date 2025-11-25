resource "cloudflare_zero_trust_device_default_profile_local_domain_fallback" "%[1]s" {
    account_id = "%[2]s"
    domains = [
        {
            suffix = "home"
        },
        {
            suffix = "corp"
        },
        {
            suffix = "localdomain"
        },
    ]
}