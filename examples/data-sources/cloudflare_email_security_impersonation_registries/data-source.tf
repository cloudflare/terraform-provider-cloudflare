data "cloudflare_email_security_impersonation_registries" "example_email_security_impersonation_registries" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  direction = "asc"
  order = "name"
  provenance = "A1S_INTERNAL"
  search = "search"
}
