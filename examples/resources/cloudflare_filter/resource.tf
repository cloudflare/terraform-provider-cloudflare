resource "cloudflare_filter" "wordpress" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  description = "Wordpress break-in attempts that are outside of the office"
  expression  = "(http.request.uri.path ~ \".*wp-login.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1"
}
