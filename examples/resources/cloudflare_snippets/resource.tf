resource "cloudflare_snippets" "example_snippets" {
  zone_id      = "9f1839b6152d298aca64c4e906b6d074"
  snippet_name = "my_snippet"
  files = [
    {
      name    = "main.js"
      content = <<-EOT
      export default {
        async fetch(request) {
          return new Response('Hello, World!');
        }
      }
      EOT
    }
  ]
  metadata = {
    main_module = "main.js"
  }
}
