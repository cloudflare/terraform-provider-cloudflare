resource "cloudflare_filter" "wordpress" {
  zone_id     = "d41d8cd98f00b204e9800998ecf8427e"
  description = "Wordpress break-in attempts that are outside of the office"
  expression  = "(http.request.uri.path ~ \".*wp-login.php\" or http.request.uri.path ~ \".*xmlrpc.php\") and ip.src ne 192.0.2.1"
}
