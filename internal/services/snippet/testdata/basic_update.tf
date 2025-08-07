resource "cloudflare_snippet" "%[1]s" {
  zone_id      = "%[2]s"
  snippet_name = "%[1]s"
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