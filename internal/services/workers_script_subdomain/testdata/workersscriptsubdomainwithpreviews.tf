resource "cloudflare_workers_script" "%[1]s" {
  account_id  = "%[2]s"
  script_name = "%[1]s"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}

resource "cloudflare_workers_script_subdomain" "%[1]s" {
  account_id      = "%[2]s"
  script_name     = cloudflare_workers_script.%[1]s.script_name
  enabled         = true
  previews_enabled = true
}

