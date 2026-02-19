resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "redirect"
  description = "%[4]s"
}

resource "cloudflare_list_item" "%[1]s_item" {
  account_id = "%[2]s"
  list_id    = cloudflare_list.%[1]s.id
  redirect {
    source_url            = "https://example.com/old"
    target_url            = "https://example.com/new"
    status_code           = 301
    include_subdomains    = true
    subpath_matching      = false
    preserve_query_string = true
    preserve_path_suffix  = false
  }
  comment = "Test redirect item"
}
