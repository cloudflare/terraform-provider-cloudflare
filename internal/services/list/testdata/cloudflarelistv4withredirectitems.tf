resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "redirect"
  description = "%[4]s"
  
  item {
    value {
      redirect {
        source_url           = "https://example.com/old"
        target_url           = "https://example.com/new"
        status_code          = 301
        include_subdomains   = "enabled"
        subpath_matching     = "disabled"
        preserve_query_string = "enabled"
        preserve_path_suffix = "disabled"
      }
    }
    comment = "Test redirect"
  }
}