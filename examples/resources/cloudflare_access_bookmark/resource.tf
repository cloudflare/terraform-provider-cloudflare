resource "cloudflare_access_bookmark" "my_bookmark_app" {
  account_id           = "f037e56e89293a057740de681ac9abbe"
  name                 = "My Bookmark App"
  domain               = "example.com"
  logo_url             = "https://example.com/example.png"
  app_launcher_visible = true
}
