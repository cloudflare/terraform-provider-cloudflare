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

resource "cloudflare_snippet_rules" "%[1]s" {
  zone_id = "%[2]s"
  rules = [
    {
      expression   = "ip.src eq 2.2.2.2"
      snippet_name = "%[1]s"
      description  = "Execute my_snippet when IP address is 2.2.2.2."
      enabled      = true
    }
  ]
  depends_on = [cloudflare_snippet.%[1]s]
}
