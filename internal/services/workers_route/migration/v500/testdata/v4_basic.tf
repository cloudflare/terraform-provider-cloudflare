resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[3]s"
  name       = "%[4]s"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}

resource "cloudflare_workers_route" "%[1]s" {
  zone_id     = "%[2]s"
  pattern     = "%[1]s.cfapi.net/*"
  script_name = "%[4]s"
  depends_on  = [cloudflare_workers_script.%[1]s]
}
