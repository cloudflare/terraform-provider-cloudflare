resource "cloudflare_worker_script" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}
