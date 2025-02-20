resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[2]s"
  script_name = "%[1]s"
  content = "export default { fetch() { return new Response('Hello world'); }, };"
  main_module = "worker.mjs"
}

resource "cloudflare_workers_route" "%[1]s" {
  zone_id = "%[3]s"
  pattern = "%[4]s/*"
  script = cloudflare_workers_script.%[1]s.script_name
}
