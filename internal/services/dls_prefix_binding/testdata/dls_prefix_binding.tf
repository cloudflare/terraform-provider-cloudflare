resource "cloudflare_dls_prefix_binding" "%[1]s" {
  account_id = "%[2]s"
  prefix_id  = "%[3]s"
  cidr       = "%[4]s"
  region_key = "%[5]s"
}

data "cloudflare_dls_prefix_binding" "%[1]s" {
  account_id = "%[2]s"
  binding_id = cloudflare_dls_prefix_binding.%[1]s.id
}

data "cloudflare_dls_prefix_bindings" "%[1]s" {
  account_id = "%[2]s"
  max_items  = 100

  depends_on = [cloudflare_dls_prefix_binding.%[1]s]
}
