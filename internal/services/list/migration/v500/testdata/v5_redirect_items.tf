resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "redirect"
  description = "%[4]s"

  items = [{
    comment = "Test redirect"
    redirect = {
      include_subdomains    = true
      preserve_path_suffix  = false
      preserve_query_string = true
      source_url            = "https://example.com/old"
      status_code           = 301
      subpath_matching      = false
      target_url            = "https://example.com/new"
    }
  }]
}
