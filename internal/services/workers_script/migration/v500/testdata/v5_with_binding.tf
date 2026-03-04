resource "cloudflare_workers_script" "%[1]s" {
  account_id  = "%[2]s"
  script_name = "%[3]s"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World: ' + MY_VAR)); });"

  bindings = [
    {
      type = "plain_text"
      name = "MY_VAR"
      text = "my-value"
    }
  ]
}
