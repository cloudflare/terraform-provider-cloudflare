data "cloudflare_email_security_domains" "example_email_security_domains" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  active_delivery_mode = "DIRECT"
  allowed_delivery_mode = "DIRECT"
  direction = "asc"
  domain = ["string"]
  integration_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  order = "domain"
  search = "search"
  status = "PENDING"
}
