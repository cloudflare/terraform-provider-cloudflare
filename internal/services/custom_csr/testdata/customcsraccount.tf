resource "cloudflare_custom_csr" "%[1]s" {
  account_id   = "%[2]s"
  country      = "US"
  state        = "California"
  locality     = "San Francisco"
  organization = "Terraform Test"
  common_name  = "%[3]s"
  sans         = ["%[3]s", "www.%[3]s"]
  name         = "test-csr-%[1]s"
  description  = "Acceptance test CSR"
}
