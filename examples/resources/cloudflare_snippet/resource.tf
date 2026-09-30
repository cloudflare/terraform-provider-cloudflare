resource "cloudflare_snippet" "example_snippet" {
  zone_id = "9f1839b6152d298aca64c4e906b6d074"
  snippet_name = "my_snippet"
  metadata = {
    main_module = "main.js"
  }
}
