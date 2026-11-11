resource "cloudflare_ai_search_namespace" "%[1]s_a" {
  account_id = "%[2]s"
  name       = "%[3]s"
}

resource "cloudflare_ai_search_namespace" "%[1]s_b" {
  account_id = "%[2]s"
  name       = "%[4]s"
}

resource "cloudflare_ai_search_instance" "%[1]s" {
  account_id      = "%[2]s"
  id              = "%[1]s"
  namespace       = cloudflare_ai_search_namespace.%[1]s_%[5]s.name
  max_num_results = %[6]d
}
