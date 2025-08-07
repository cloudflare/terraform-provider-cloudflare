resource "cloudflare_workers_script" "%[1]s" {
  account_id  = "%[2]s"
  script_name = "%[1]s"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello from Worker!')) })"
}

resource "cloudflare_workers_route" "%[1]s" {
  zone_id = "%[3]s"
  pattern = "%[1]s.%[4]s/*"
  script  = cloudflare_workers_script.%[1]s.script_name
}