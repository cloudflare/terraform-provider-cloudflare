resource "cloudflare_access_bookmark" "my_bookmark_app" {
  account_id           = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name                 = "My Bookmark App"
  domain               = "example.com"
  logo_url             = "https://example.com/example.png"
  app_launcher_visible = true
}
