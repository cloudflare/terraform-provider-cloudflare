resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World: ' + MY_VAR)); });"

  plain_text_binding {
    name = "MY_VAR"
    text = "my-value"
  }
}
