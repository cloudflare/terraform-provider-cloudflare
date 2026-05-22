resource "cloudflare_dls_prefix_binding" "example_dls_prefix_binding" {
  account_id = 0
  cidr = "10.0.1.0/24"
  prefix_id = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
  region_key = "eu"
}
