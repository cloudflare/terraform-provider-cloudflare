resource "cloudflare_workers_script" "%[1]s" {
  account_id  = "%[2]s"
  script_name = "%[1]s"
  content     = "export default { fetch() { return new Response('Hello world'); } };"
  main_module = "worker.mjs"
  bindings = [
    {
      name         = "MY_RATE_LIMITER"
      type         = "ratelimit"
      namespace_id = "1234"
      simple = {
        limit  = 100
        period = 60
      }
    }
  ]
}
