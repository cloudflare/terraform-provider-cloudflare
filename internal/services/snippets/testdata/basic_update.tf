resource "cloudflare_snippet" "%[1]s" {
  zone_id      = "%[2]s"
  snippet_name = "%[1]s"
  files        = "export { async function fetch(request, env) { return new Response('Hello Updated World!'); } }"
  metadata = {
    main_module = "main.js"
  }
}